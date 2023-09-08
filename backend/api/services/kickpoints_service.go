package services

import (
	"errors"

	"backend/api/repos"
	"backend/api/types"
	"backend/store/postgres/models"
)

type IKickpointsService interface {
	ActiveClanKickpoints(clanTag string) ([]*types.ClanMemberKickpoints, error)
	ActivePlayerKickpoints(playerTag, clanTag string) ([]*types.Kickpoint, error)
	CreateKickpoint(addedByDiscordID string, kickpoint *types.CreateKickpointPayload) error
	UpdateKickpoint(updatedByDiscordID string, kickpoint *types.UpdateKickpointPayload) error
	DeleteKickpoint(id uint) error
}

type KickpointsService struct {
	kickpointsRepo   repos.IKickpointsRepo
	playersRepo      repos.IPlayersRepo
	clanSettingsRepo repos.IClanSettingsRepo
}

func NewKickpointsService(kickpointsRepo repos.IKickpointsRepo, playersRepo repos.IPlayersRepo, clanSettingsRepo repos.IClanSettingsRepo) *KickpointsService {
	return &KickpointsService{
		kickpointsRepo:   kickpointsRepo,
		playersRepo:      playersRepo,
		clanSettingsRepo: clanSettingsRepo,
	}
}

func (service *KickpointsService) ActiveClanKickpoints(clanTag string) ([]*types.ClanMemberKickpoints, error) {
	clanSettings, err := service.clanSettingsRepo.ClanSettings(clanTag)
	if err != nil {
		return nil, err
	}

	return service.kickpointsRepo.ActiveClanKickpoints(clanTag, clanSettings)
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

func (service *KickpointsService) CreateKickpoint(addedByDiscordID string, payload *types.CreateKickpointPayload) error {
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
		PlayerTag:              payload.PlayerTag,
		ClanTag:                payload.ClanTag,
		Date:                   payload.Date,
		Amount:                 payload.Amount,
		Reason:                 payload.Reason,
		Description:            payload.Description,
		AddedByDiscordID:       addedByDiscordID,
		LastUpdatedByDiscordID: addedByDiscordID,
	}

	return service.kickpointsRepo.CreateKickpoint(newKickpoint)
}

func (service *KickpointsService) UpdateKickpoint(updatedByDiscordID string, payload *types.UpdateKickpointPayload) error {
	updatedKickpoint := &models.Kickpoint{
		PlayerTag:              payload.PlayerTag,
		ClanTag:                payload.ClanTag,
		Date:                   payload.Date,
		Amount:                 payload.Amount,
		Reason:                 payload.Reason,
		Description:            payload.Description,
		LastUpdatedByDiscordID: updatedByDiscordID,
	}

	return service.kickpointsRepo.UpdateKickpoint(updatedKickpoint)
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
