package repos

import (
	"gorm.io/gorm"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/store/postgres"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IClansRepo interface {
	Clans(params types.ClansParams, preload ...string) (*types.PaginatedResponse[*models.Clan], error)
	ClanByTag(tag string, preload ...string) (*models.Clan, error)
	Count(params types.ClansParams) (int64, error)
}

type ClansRepo struct {
	db *gorm.DB
}

var clansQueryFields = []string{"name", "tag"}

func NewClansRepo(db *gorm.DB) IClansRepo {
	return &ClansRepo{db: db}
}

func (r *ClansRepo) Clans(params types.ClansParams, preload ...string) (*types.PaginatedResponse[*models.Clan], error) {
	count, err := r.Count(params)
	if err != nil {
		return nil, err
	}

	var clans models.Clans
	if err = r.db.
		Scopes(
			postgres.WithContains(params.Query, clansQueryFields...),
			postgres.WithPagination(params.PaginationParams),
			postgres.WithPreloading(preload...),
		).
		Find(&clans).Error; err != nil {
		return nil, err
	}

	return types.NewPaginatedResponse(clans, params.PaginationParams, count), nil
}

func (r *ClansRepo) ClanByTag(tag string, preload ...string) (*models.Clan, error) {
	var clan *models.Clan
	err := r.db.
		Scopes(postgres.WithPreloading(preload...)).
		First(&clan, "tag = ?", tag).Error
	return clan, err
}

func (r *ClansRepo) Count(params types.ClansParams) (int64, error) {
	var count int64
	err := r.db.
		Model(&models.Clan{}).
		Scopes(postgres.WithContains(params.Query, clansQueryFields...)).
		Count(&count).Error
	return count, err
}
