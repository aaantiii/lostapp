package services

import (
	"github.com/aaantiii/lostapp/backend/api/repos"
	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IClansService interface {
	Clans(params types.ClansParams) (*types.PaginatedResponse[*models.Clan], error)
	ClanByTag(tag string) (*models.Clan, error)
	ClanSettings(tag string) (*models.ClanSettings, error)
	UpdateClanSettings(tag, updatedByID string, payload *types.UpdateClanSettingsPayload) error
	ActiveClanKickpoints(clanTag string) ([]*types.ClanKickpointsEntry, error)
	ActiveClanMemberKickpoints(memberTag string, clanTag string) ([]*models.Kickpoint, error)
	CreateKickpoint(payload *types.CreateKickpointPayload) error
	UpdateKickpoint(id uint, payload *types.UpdateKickpointPayload) error
	DeleteKickpoint(id uint) error
}

type ClansService struct {
	clans      repos.IClansRepo
	kickpoints repos.IKickpointsRepo
	settings   repos.IClanSettingsRepo
}

func NewClansService(clans repos.IClansRepo, kickpoints repos.IKickpointsRepo, settings repos.IClanSettingsRepo) IClansService {
	return &ClansService{
		clans:      clans,
		kickpoints: kickpoints,
		settings:   settings,
	}
}

func (s *ClansService) Clans(params types.ClansParams) (*types.PaginatedResponse[*models.Clan], error) {
	return s.clans.Clans(params)
}

func (s *ClansService) ClanByTag(tag string) (*models.Clan, error) {
	return s.clans.ClanByTag(tag)
}

func (s *ClansService) ClanSettings(tag string) (*models.ClanSettings, error) {
	return s.settings.ClanSettings(tag)
}

func (s *ClansService) UpdateClanSettings(tag, updatedByID string, payload *types.UpdateClanSettingsPayload) error {
	settings, err := s.settings.ClanSettings(tag)
	if err != nil {
		return err
	}

	settings.MaxKickpoints = payload.MaxKickpoints
	settings.MinSeasonWins = payload.MinSeasonWins
	settings.KickpointsExpireAfterDays = payload.KickpointsExpireAfterDays
	settings.KickpointsSeasonWins = payload.KickpointsSeasonWins
	settings.KickpointsCWMissed = payload.KickpointsCWMissed
	settings.KickpointsCWFail = payload.KickpointsCWFail
	settings.KickpointsCWLMissed = payload.KickpointsCWLMissed
	settings.KickpointsCWLZeroStars = payload.KickpointsCWLZeroStars
	settings.KickpointsCWLOneStar = payload.KickpointsCWLOneStar
	settings.KickpointsRaidMissed = payload.KickpointsRaidMissed
	settings.KickpointsRaidFail = payload.KickpointsRaidFail
	settings.KickpointsClanGames = payload.KickpointsClanGames
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

func (s *ClansService) ActiveClanMemberKickpoints(memberTag string, clanTag string) ([]*models.Kickpoint, error) {
	settings, err := s.settings.ClanSettings(clanTag)
	if err != nil {
		return nil, err
	}
	return s.kickpoints.ActiveMemberKickpoints(memberTag, settings)
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
