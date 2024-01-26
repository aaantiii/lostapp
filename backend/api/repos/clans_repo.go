package repos

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/aaantiii/lostapp/backend/store/postgres"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IClansRepo interface {
	Clans(params types.ClansParams) (models.Clans, error)
	ClanByTag(tag string) (*models.Clan, error)
	ClanByTagPreload(tag string) (*models.Clan, error)
	Count(params types.ClansParams) (int64, error)
	ClanNameByTag(tag string) (string, error)
}

type ClansRepo struct {
	db *gorm.DB
}

func NewClansRepo(db *gorm.DB) IClansRepo {
	return &ClansRepo{db: db}
}

func (repo *ClansRepo) Clans(query string) (models.Clans, error) {
	var clans models.Clans
	err := repo.db.
		Scopes(postgres.ScopeContains(query, "name", "tag")).
		Find(&clans).Error
	return clans, err
}

func (repo *ClansRepo) ClanByTag(tag string) (*models.Clan, error) {
	var clan *models.Clan
	err := repo.db.
		Preload("ClanMembers").
		First(&clan, "tag = ?", tag).Error
	return clan, err
}

func (repo *ClansRepo) ClanByTagPreload(tag string) (*models.Clan, error) {
	var clan *models.Clan
	err := repo.db.
		Preload(clause.Associations).
		Preload("ClanMembers.Player").
		First(&clan, "tag = ?", tag).Error
	return clan, err
}

func (repo *ClansRepo) ClanNameByTag(tag string) (string, error) {
	var c struct{ Name string }
	err := repo.db.
		Model(&models.Clan{}).
		Select("name").
		First(&c, "tag = ?", tag).Error
	return c.Name, err
}
