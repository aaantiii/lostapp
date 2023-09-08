package repos

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"backend/api/types"
	"backend/store/cache"
	"backend/store/postgres/models"
)

type IKickpointsRepo interface {
	ActiveClanKickpoints(clanTag string, settings *models.LostClanSettings) ([]*types.ClanMemberKickpoints, error)
	ActivePlayerKickpoints(playerTag, clanTag string, settings *models.LostClanSettings) ([]*models.Kickpoint, error)
	CreateKickpoint(kickpoint *models.Kickpoint) error
	UpdateKickpoint(kickpoint *models.Kickpoint) error
	DeleteKickpoint(id uint) error
}

type KickpointsRepo struct {
	db    *gorm.DB
	cache *cache.CocCache
}

func NewKickpointsRepo(db *gorm.DB, cache *cache.CocCache) *KickpointsRepo {
	return &KickpointsRepo{db: db, cache: cache}
}

func (repo *KickpointsRepo) ActiveClanKickpoints(clanTag string, settings *models.LostClanSettings) ([]*types.ClanMemberKickpoints, error) {
	minDate := time.Now().AddDate(0, 0, -int(settings.KickpointsExpireAfterDays))

	var kickpoints []*models.Kickpoint
	err := repo.db.Select("SUM(amount) as amount, player_tag").Group("player_tag").Find(&kickpoints, "clan_tag = ? AND date > ?", clanTag, minDate).Error
	if err != nil {
		return nil, err
	}

	for _, kickpoint := range kickpoints {
		log.Println(kickpoint.PlayerTag, kickpoint.Amount)
	}

	clan, found := repo.cache.ClanByTag[clanTag]
	if !found {
		return nil, errors.New("clan not found")
	}

	memberKickpoints := make([]*types.ClanMemberKickpoints, clan.Members)
	for i, member := range clan.MemberList {
		memberKickpoints[i] = &types.ClanMemberKickpoints{
			Tag:  member.Tag,
			Name: member.Name,
			Role: member.Role,
		}
		for _, kickpoint := range kickpoints {
			if kickpoint.PlayerTag == member.Tag {
				memberKickpoints[i].Amount = kickpoint.Amount
			}
		}
	}

	return memberKickpoints, nil
}

func (repo *KickpointsRepo) ActivePlayerKickpoints(playerTag, clanTag string, settings *models.LostClanSettings) ([]*models.Kickpoint, error) {
	minDate := time.Now().AddDate(0, 0, -int(settings.KickpointsExpireAfterDays))
	var kickpoints []*models.Kickpoint
	err := repo.db.Find(&kickpoints, "player_tag = ? AND clan_tag = ? AND date > ?", playerTag, clanTag, minDate).Error
	return kickpoints, err
}

func (repo *KickpointsRepo) CreateKickpoint(kickpoint *models.Kickpoint) error {
	return repo.db.Create(kickpoint).Error
}

func (repo *KickpointsRepo) UpdateKickpoint(kickpoint *models.Kickpoint) error {
	return repo.db.Save(kickpoint).Error
}

func (repo *KickpointsRepo) DeleteKickpoint(id uint) error {
	return repo.db.Delete(&models.Kickpoint{}, id).Error
}
