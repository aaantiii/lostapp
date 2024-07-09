package repos

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
	"github.com/aaantiii/lostapp/backend/store/postgres"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IKickpointsRepo interface {
	KickpointByID(id uint) (*models.Kickpoint, error)
	ActiveClanKickpoints(settings *models.ClanSettings) ([]*types.ClanKickpointsEntry, error)
	ActiveMemberKickpoints(params types.KickpointParams, settings *models.ClanSettings) (*types.PaginatedResponse[*models.Kickpoint], error)
	CountMemberKickpoints(params types.KickpointParams, minDate time.Time) (int64, error)
	ActiveMemberKickpointsSum(memberTag string, settings *models.ClanSettings) (int, error)
	CreateKickpoint(kickpoint *models.Kickpoint) error
	UpdateKickpoint(kickpoint *models.Kickpoint) error
	DeleteKickpoint(id uint) error
}

type KickpointsRepo struct {
	db *gorm.DB
}

func NewKickpointsRepo(db *gorm.DB) IKickpointsRepo {
	return &KickpointsRepo{db: db}
}

func (r *KickpointsRepo) KickpointByID(id uint) (*models.Kickpoint, error) {
	var kickpoint *models.Kickpoint
	err := r.db.Preload(clause.Associations).First(&kickpoint, id).Error
	return kickpoint, err
}

func (r *KickpointsRepo) ActiveClanKickpoints(settings *models.ClanSettings) ([]*types.ClanKickpointsEntry, error) {
	minDate := utils.KickpointMinDate(settings.KickpointsExpireAfterDays)

	var memberKickpoints []*types.ClanKickpointsEntry
	if err := r.db.
		Raw("SELECT p.name AS name, p.coc_tag as tag, SUM(k.amount) AS amount FROM kickpoints k INNER JOIN players p ON k.player_tag = p.coc_tag WHERE k.clan_tag = ? AND k.date BETWEEN ? AND NOW() GROUP BY p.name, p.coc_tag ORDER BY amount DESC", settings.ClanTag, minDate).
		Scan(&memberKickpoints).Error; err != nil {
		return nil, err
	}

	if len(memberKickpoints) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return memberKickpoints, nil
}

func (r *KickpointsRepo) ActiveMemberKickpoints(params types.KickpointParams, settings *models.ClanSettings) (*types.PaginatedResponse[*models.Kickpoint], error) {
	minDate := utils.KickpointMinDate(settings.KickpointsExpireAfterDays)

	count, err := r.CountMemberKickpoints(params, minDate)
	if err != nil {
		return nil, err
	}
	if err = utils.ValidatePagination(params.PaginationParams, count); err != nil {
		return nil, err
	}

	var kickpoints []*models.Kickpoint
	if err = r.db.
		Scopes(postgres.WithPagination(params.PaginationParams)).
		Preload(clause.Associations).
		Order("created_at").
		Find(&kickpoints, "player_tag = ? AND clan_tag = ? AND date BETWEEN ? AND NOW()", params.PlayerTag, params.ClanTag, minDate).Error; err != nil {
		return nil, err
	}

	if len(kickpoints) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return types.NewPaginatedResponse(kickpoints, params.PaginationParams, count), nil
}

func (r *KickpointsRepo) CountMemberKickpoints(params types.KickpointParams, minDate time.Time) (int64, error) {
	var count int64
	if err := r.db.
		Model(&models.Kickpoint{}).
		Where("player_tag = ? AND clan_tag = ? AND date BETWEEN ? AND NOW()", params.PlayerTag, params.ClanTag, minDate).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *KickpointsRepo) ActiveMemberKickpointsSum(memberTag string, settings *models.ClanSettings) (int, error) {
	minDate := utils.KickpointMinDate(settings.KickpointsExpireAfterDays)

	var v struct{ Sum int }
	if err := r.db.
		Model(&models.Kickpoint{}).
		Where("player_tag = ? AND clan_tag = ? AND date BETWEEN ? AND NOW()", memberTag, settings.ClanTag, minDate).
		Select("SUM(amount) as sum").
		Scan(&v).Error; err != nil {
		return 0, err
	}

	return v.Sum, nil
}

func (r *KickpointsRepo) CreateKickpoint(kickpoint *models.Kickpoint) error {
	return r.db.Create(&kickpoint).Error
}

func (r *KickpointsRepo) UpdateKickpoint(kickpoint *models.Kickpoint) error {
	return r.db.Save(kickpoint).Error
}

func (r *KickpointsRepo) DeleteKickpoint(id uint) error {
	return r.db.Delete(&models.Kickpoint{}, id).Error
}
