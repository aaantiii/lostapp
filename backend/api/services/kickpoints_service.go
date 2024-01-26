package services

import (
	"errors"

	"github.com/aaantiii/lostapp/backend/api/repos"
	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/util"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IKickpointsService interface {
	ActiveClanMemberKickpoints(clanTag string) ([]*types.ClanMemberKickpoints, error)
	ActivePlayerKickpoints(playerTag, clanTag string) ([]*types.Kickpoint, error)
	FuturePlayerKickpoints(playerTag, clanTag string) ([]*types.Kickpoint, error)
	CreateKickpoint(payload *types.CreateKickpointPayload) error
	UpdateKickpoint(id uint, payload *types.UpdateKickpointPayload) error
	DeleteKickpoint(id uint) error
}

type KickpointsService struct {
	kickpointsRepo   repos.IKickpointsRepo
	playersRepo      repos.IPlayersRepo
	clanSettingsRepo repos.IClanSettingsRepo
}

func NewKickpointsService(kickpointsRepo repos.IKickpointsRepo, playersRepo repos.IPlayersRepo, clanSettingsRepo repos.IClanSettingsRepo) IKickpointsService {
	return &KickpointsService{
		kickpointsRepo:   kickpointsRepo,
		playersRepo:      playersRepo,
		clanSettingsRepo: clanSettingsRepo,
	}
}

func (service *KickpointsService) ActiveClanMemberKickpoints(clanTag string) ([]*types.ClanMemberKickpoints, error) {
	clanSettings, err := service.clanSettingsRepo.ClanSettings(clanTag)
	if err != nil {
		return nil, err
	}

	return service.kickpointsRepo.ActiveClanMemberKickpoints(clanTag, clanSettings)
}

func (service *KickpointsService) ActivePlayerKickpoints(playerTag, clanTag string) ([]*types.Kickpoint, error) {
	clanSettings, err := service.clanSettingsRepo.ClanSettings(clanTag)
	if err != nil {
		return nil, err
	}

	playerKickpoints, err := service.kickpointsRepo.ActivePlayerKickpoints(playerTag, clanTag, clanSettings)
	if err != nil {
		return nil, err
	}

	kickpoints := make([]*types.Kickpoint, len(playerKickpoints))
	for i, playerKickpoint := range playerKickpoints {
		kickpoints[i] = types.NewKickpoint(playerKickpoint)
	}

	return kickpoints, nil
}

func (service *KickpointsService) FuturePlayerKickpoints(playerTag, clanTag string) ([]*types.Kickpoint, error) {
	clanSettings, err := service.clanSettingsRepo.ClanSettings(clanTag)
	if err != nil {
		return nil, err
	}

	playerKickpoints, err := service.kickpointsRepo.FuturePlayerKickpoints(playerTag, clanTag, clanSettings)
	if err != nil {
		return nil, err
	}

	kickpoints := make([]*types.Kickpoint, len(playerKickpoints))
	for i, playerKickpoint := range playerKickpoints {
		kickpoints[i] = types.NewKickpoint(playerKickpoint)
	}

	return kickpoints, nil
}

func (service *KickpointsService) CreateKickpoint(payload *types.CreateKickpointPayload) error {
	player, err := service.playersRepo.PlayerByTag(payload.PlayerTag)
	if err != nil {
		return err
	}

	playerIsMemberOfClan := false
	for _, clan := range player.Clans {
		if clan.Tag == payload.ClanTag {
			playerIsMemberOfClan = true
			break
		}
	}
	if !playerIsMemberOfClan {
		return errors.New("player is not a member of the clan")
	}

	newKickpoint := &models.Kickpoint{
		PlayerTag:          payload.PlayerTag,
		ClanTag:            payload.ClanTag,
		Date:               util.TruncateToDay(payload.Date),
		Amount:             payload.Amount,
		Description:        payload.Description,
		CreatedByDiscordID: payload.AddedByDiscordID,
	}

	return service.kickpointsRepo.CreateKickpoint(newKickpoint)
}

func (service *KickpointsService) UpdateKickpoint(id uint, payload *types.UpdateKickpointPayload) error {
	kickpoint, err := service.kickpointsRepo.Kickpoint(id)
	if err != nil {
		return err
	}

	kickpoint.Date = util.TruncateToDay(payload.Date)
	kickpoint.Amount = payload.Amount
	kickpoint.Description = payload.Description
	kickpoint.UpdatedByDiscordID = &payload.UpdatedByDiscordID

	return service.kickpointsRepo.UpdateKickpoint(kickpoint)
}

func (service *KickpointsService) DeleteKickpoint(id uint) error {
	return service.kickpointsRepo.DeleteKickpoint(id)
}

func (service *KickpointsService) kickpointDTOs(kickpoints []*models.Kickpoint) []*types.Kickpoint {
	dtos := make([]*types.Kickpoint, len(kickpoints))
	for i, kickpoint := range kickpoints {
		dtos[i] = types.NewKickpoint(kickpoint)
	}

	return dtos
}
