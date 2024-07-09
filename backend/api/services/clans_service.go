package services

import (
	"github.com/aaantiii/goclash"

	"github.com/aaantiii/lostapp/backend/api/repos"
	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
	"github.com/aaantiii/lostapp/backend/store/postgres"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IClansService interface {
	Clans(params types.ClansParams) (*types.PaginatedResponse[*models.Clan], error)
	ClansList(params types.ClansParams) (models.Clans, error)
	LiveClans(params types.ClansParams) (*types.PaginatedResponse[*types.Clan], error)
	ClanByTag(tag string) (*models.Clan, error)
	ClanSettings(tag string) (*models.ClanSettings, error)
	ClanEvents(tag string, params types.PaginationParams) (*types.PaginatedResponse[*models.ClanEvent], error)
	ClanMembers(params types.MembersParams) (*types.PaginatedResponse[*models.ClanMember], error)
	UpdateClanSettings(tag, updatedByID string, payload *types.UpdateClanSettingsPayload) error
	ActiveClanKickpoints(clanTag string) ([]*types.ClanKickpointsEntry, error)
	ActiveClanMemberKickpoints(params types.KickpointParams) (*types.PaginatedResponse[*models.Kickpoint], error)
	CreateMember(payload *types.CreateMemberPayload, addedByDiscordID string) error
	UpdateMember(payload *types.UpdateMemberPayload) error
	DeleteMember(playerTag, clanTag string) error
	CreateKickpoint(payload *types.CreateKickpointPayload) error
	UpdateKickpoint(id uint, payload *types.UpdateKickpointPayload) error
	DeleteKickpoint(id uint) error
}

type ClansService struct {
	clans       repos.IClansRepo
	kickpoints  repos.IKickpointsRepo
	settings    repos.IClanSettingsRepo
	events      repos.IClanEventsRepo
	members     repos.IMembersRepo
	clashClient *goclash.Client
}

func NewClansService(clans repos.IClansRepo, kickpoints repos.IKickpointsRepo, settings repos.IClanSettingsRepo, events repos.IClanEventsRepo, members repos.IMembersRepo, clashClient *goclash.Client) IClansService {
	return &ClansService{
		clans:       clans,
		kickpoints:  kickpoints,
		settings:    settings,
		members:     members,
		clashClient: clashClient,
	}
}

func (s *ClansService) Clans(params types.ClansParams) (*types.PaginatedResponse[*models.Clan], error) {
	return s.clans.Clans(params)
}

func (s *ClansService) ClansList(params types.ClansParams) (models.Clans, error) {
	return s.clans.ClansList(params)
}

func (s *ClansService) LiveClans(params types.ClansParams) (*types.PaginatedResponse[*types.Clan], error) {
	clans, err := s.clans.Clans(params, postgres.Preloader{Field: "ClanMembers"})
	if err != nil {
		return nil, err
	}

	tags := make([]string, len(clans.Items))
	for i, clan := range clans.Items {
		tags[i] = clan.Tag
	}

	liveClans, err := s.clashClient.GetClans(tags...)
	if err != nil {
		return nil, err
	}

	res := make([]*types.Clan, len(liveClans))
	for i, clan := range liveClans {
		res[i] = &types.Clan{
			Clan:        clan,
			LostMembers: clans.Items[i].ClanMembers,
		}
	}

	return &types.PaginatedResponse[*types.Clan]{Items: res, Pagination: clans.Pagination}, nil
}

func (s *ClansService) ClanByTag(tag string) (*models.Clan, error) {
	settings, err := s.settings.ClanSettings(tag)
	if err != nil {
		return nil, err
	}

	kickpointMinDate := utils.KickpointMinDate(settings.KickpointsExpireAfterDays)
	clan, err := s.clans.ClanByTag(
		tag,
		postgres.Preloader{Field: "ClanMembers.Player"},
		postgres.Preloader{Field: "ClanMembers.Kickpoints", Args: []any{"date BETWEEN ? AND NOW()", kickpointMinDate}},
	)
	if err != nil {
		return nil, err
	}
	clan.Settings = settings
	return clan, nil
}

func (s *ClansService) CreateMember(payload *types.CreateMemberPayload, addedByDiscordID string) error {
	return s.members.CreateMember(&models.ClanMember{
		ClanTag:          payload.ClanTag,
		PlayerTag:        payload.PlayerTag,
		AddedByDiscordID: addedByDiscordID,
		ClanRole:         models.ClanRole(payload.Role),
	})
}

func (s *ClansService) UpdateMember(payload *types.UpdateMemberPayload) error {
	return s.members.UpdateMemberRole(payload.PlayerTag, payload.ClanTag, models.ClanRole(payload.Role))
}

func (s *ClansService) DeleteMember(playerTag, clanTag string) error {
	return s.members.DeleteMember(playerTag, clanTag)
}

func (s *ClansService) ClanMembers(params types.MembersParams) (*types.PaginatedResponse[*models.ClanMember], error) {
	return s.members.MembersPaginated(params)
}

func (s *ClansService) ClanSettings(tag string) (*models.ClanSettings, error) {
	return s.settings.ClanSettings(tag)
}

func (s *ClansService) ClanEvents(tag string, params types.PaginationParams) (*types.PaginatedResponse[*models.ClanEvent], error) {
	return s.events.ClanEvents(tag, params)
}

func (s *ClansService) UpdateClanSettings(tag, updatedByID string, payload *types.UpdateClanSettingsPayload) error {
	settings, err := s.settings.ClanSettings(tag)
	if err != nil {
		return err
	}

	settings.MaxKickpoints = payload.MaxKickpoints
	settings.MinSeasonWins = payload.MinSeasonWins
	settings.KickpointsExpireAfterDays = payload.KickpointsExpireAfterDays
	settings.UpdatedByDiscordID = &updatedByID
	return s.settings.UpdateClanSettings(settings)
}

func (s *ClansService) ActiveClanKickpoints(clanTag string) ([]*types.ClanKickpointsEntry, error) {
	settings, err := s.settings.ClanSettings(clanTag)
	if err != nil {
		return nil, err
	}
	return s.kickpoints.ActiveClanKickpoints(settings)
}

func (s *ClansService) ActiveClanMemberKickpoints(params types.KickpointParams) (*types.PaginatedResponse[*models.Kickpoint], error) {
	settings, err := s.settings.ClanSettings(params.ClanTag)
	if err != nil {
		return nil, err
	}
	return s.kickpoints.ActiveMemberKickpoints(params, settings)
}

func (s *ClansService) CreateKickpoint(payload *types.CreateKickpointPayload) error {
	return s.kickpoints.CreateKickpoint(&models.Kickpoint{
		Date:               payload.Date,
		Amount:             payload.Amount,
		ClanTag:            payload.ClanTag,
		PlayerTag:          payload.PlayerTag,
		CreatedByDiscordID: payload.CreatedByDiscordID,
		UpdatedByDiscordID: payload.CreatedByDiscordID,
		Description:        payload.Description,
	})
}

func (s *ClansService) UpdateKickpoint(id uint, payload *types.UpdateKickpointPayload) error {
	kickpoint, err := s.kickpoints.KickpointByID(id)
	if err != nil {
		return err
	}

	kickpoint.Date = payload.Date
	kickpoint.Amount = payload.Amount
	kickpoint.UpdatedByDiscordID = payload.UpdatedByDiscordID
	kickpoint.Description = payload.Description
	return s.kickpoints.UpdateKickpoint(kickpoint)
}

func (s *ClansService) DeleteKickpoint(id uint) error {
	return s.kickpoints.DeleteKickpoint(id)
}
