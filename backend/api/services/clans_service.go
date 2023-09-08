package services

import (
	"errors"

	"backend/api/repos"
	"backend/api/types"
	"backend/store/postgres/models"
)

type IClansService interface {
	IsMaintenance() bool
	Clans() ([]*types.Clan, error)
	ClanByTag(tag string) (*types.Clan, error)
	ClanMembers(tag string) ([]types.ClanMember, error)
	ClanSettings(tag string) (*models.LostClanSettings, error)
	UpdateClanSettings(tag string, payload *types.UpdateClanSettingsPayload) error
	ClansWhereMemberIsLeader(discordID string) ([]*types.Clan, error)
}

type ClansService struct {
	clansRepo        repos.IClansRepo
	playersRepo      repos.IPlayersRepo
	clanSettingsRepo repos.IClanSettingsRepo
}

func NewClansService(clansRepo repos.IClansRepo, playersRepo repos.IPlayersRepo, clanSettingsRepo repos.IClanSettingsRepo) *ClansService {
	return &ClansService{
		clansRepo:        clansRepo,
		playersRepo:      playersRepo,
		clanSettingsRepo: clanSettingsRepo,
	}
}

func (service *ClansService) IsMaintenance() bool {
	return service.clansRepo.IsMaintenance()
}

func (service *ClansService) Clans() ([]*types.Clan, error) {
	clans := service.clansRepo.Clans()
	if clans == nil {
		return nil, errors.New("clans is nil")
	}

	return clans, nil
}

func (service *ClansService) ClanByTag(tag string) (*types.Clan, error) {
	clan, err := service.clansRepo.ClanByTag(tag)
	if err != nil {
		return nil, err
	}

	return clan, nil
}

func (service *ClansService) ClanMembers(tag string) ([]types.ClanMember, error) {
	clan, err := service.clansRepo.ClanByTag(tag)
	if err != nil {
		return nil, err
	}

	return clan.MemberList, nil
}

func (service *ClansService) ClanSettings(tag string) (*models.LostClanSettings, error) {
	if _, err := service.clansRepo.ClanByTag(tag); err != nil {
		return nil, err
	}

	return service.clanSettingsRepo.ClanSettings(tag)
}

func (service *ClansService) UpdateClanSettings(tag string, payload *types.UpdateClanSettingsPayload) error {
	_, err := service.clanSettingsRepo.ClanSettings(tag)
	if err != nil {
		return err
	}

	return service.clanSettingsRepo.UpdateClanSettings(&models.LostClanSettings{
		ClanTag:                   tag,
		MaxKickpoints:             payload.MaxKickpoints,
		MinSeasonWins:             payload.MinSeasonWins,
		KickpointsExpireAfterDays: payload.KickpointsExpireAfterDays,
		KickpointsSeasonWins:      payload.KickpointsSeasonWins,
		KickpointsCWMissed:        payload.KickpointsCWMissed,
		KickpointsCWFail:          payload.KickpointsCWFail,
		KickpointsCWLMissed:       payload.KickpointsCWLMissed,
		KickpointsCWLZeroStars:    payload.KickpointsCWLZeroStars,
		KickpointsCWLOneStar:      payload.KickpointsCWLOneStar,
		KickpointsRaidMissed:      payload.KickpointsRaidMissed,
		KickpointsRaidFail:        payload.KickpointsRaidFail,
		KickpointsClanGames:       payload.KickpointsClanGames,
	})
}

func (service *ClansService) ClansWhereMemberIsLeader(discordID string) ([]*types.Clan, error) {
	players, err := service.playersRepo.PlayersByDiscordID(discordID)
	if err != nil {
		return nil, err
	}

	var clans []*types.Clan
	for _, player := range players {
		for _, playerClan := range player.Clans {
			if playerClan.Role.IsLeader() || playerClan.Role.IsCoLeader() {
				clan, err := service.ClanByTag(playerClan.Tag)
				if err != nil {
					continue
				}
				clans = append(clans, clan)
			}
		}
	}

	return clans, nil
}
