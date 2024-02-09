package repos

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"bot/store/postgres"
	"bot/store/postgres/models"
	"bot/types"
)

type IClansRepo interface {
	Clans(query string) (models.Clans, error)
	ClanByTag(tag string) (*models.Clan, error)
	ClanByTagPreload(tag string) (*models.Clan, error)
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
		Scopes(postgres.WithSearchQuery(query, "name", "tag")).
		Limit(types.MaxCommandChoices).
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
