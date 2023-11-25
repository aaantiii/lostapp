package repos

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"backend/api/types"
	"backend/api/util"
	"backend/store/cache"
	"backend/store/postgres/models"
)

type IKickpointsRepo interface {
	Kickpoint(id uint) (*models.Kickpoint, error)
	ActiveClanMemberKickpoints(clanTag string, settings *models.LostClanSettings) ([]*types.ClanMemberKickpoints, error)
	ActivePlayerKickpoints(playerTag, clanTag string, settings *models.LostClanSettings) ([]*models.Kickpoint, error)
	FuturePlayerKickpoints(playerTag, clanTag string, settings *models.LostClanSettings) ([]*models.Kickpoint, error)
	CreateKickpoint(kickpoint *models.Kickpoint) error
	UpdateKickpoint(kickpoint *models.Kickpoint) error
	DeleteKickpoint(id uint) error
}

type KickpointsRepo struct {
	db    *gorm.DB
	cache *cache.CocCache
}

func NewKickpointsRepo(db *gorm.DB, cache *cache.CocCache) IKickpointsRepo {
	return &KickpointsRepo{db: db, cache: cache}
}

func (repo *KickpointsRepo) Kickpoint(id uint) (*models.Kickpoint, error) {
	var kickpoint *models.Kickpoint
	err := repo.db.First(&kickpoint, id).Error
	return kickpoint, err
}

func (repo *KickpointsRepo) ActiveClanMemberKickpoints(clanTag string, settings *models.LostClanSettings) ([]*types.ClanMemberKickpoints, error) {
	minDate := time.Now().AddDate(0, 0, -int(settings.KickpointsExpireAfterDays))

	var kickpoints []*models.Kickpoint
	err := repo.db.
		Select("SUM(amount) as amount, player_tag").
		Group("player_tag").
		Find(&kickpoints, "clan_tag = ? AND date > ? AND date <= ?", clanTag, minDate, time.Now()).Error
	if err != nil {
		return nil, err
	}

	clan, found := repo.cache.ClanByTag.Get(clanTag)
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
	err := repo.db.
		//Preload("CreatedByUser").
		//Preload("UpdatedByUser").
		Preload(clause.Associations).
		Order("id").
		Find(&kickpoints, "player_tag = ? AND clan_tag = ? AND date > ? AND date <= ?", playerTag, clanTag, minDate, time.Now()).Error

	return kickpoints, err
}

func (repo *KickpointsRepo) FuturePlayerKickpoints(playerTag, clanTag string, settings *models.LostClanSettings) ([]*models.Kickpoint, error) {
	timestamp := util.TruncateToDay(time.Now()).AddDate(0, 0, 1)
	var kickpoints []*models.Kickpoint
	err := repo.db.Find(&kickpoints, "player_tag = ? AND clan_tag = ? AND date > ?", playerTag, clanTag, timestamp).Error
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
