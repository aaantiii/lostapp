package repos

import (
	"gorm.io/gorm"

	"bot/store/postgres"
	"bot/store/postgres/models"
	"bot/types"
)

type ILostClansRepo interface {
	LostClans(query string) (models.LostClans, error)
	LostClan(tag string) (*models.LostClan, error)
	NameByTag(tag string) (string, error)
}

type LostClansRepo struct {
	db *gorm.DB
}

func NewLostClansRepo(db *gorm.DB) ILostClansRepo {
	return &LostClansRepo{db: db}
}

func (repo *LostClansRepo) LostClans(query string) (models.LostClans, error) {
	var lostClans models.LostClans
	err := repo.db.
		Scopes(
			postgres.ScopeLimit(types.MaxCommandChoices),
			postgres.ScopeContains(query, "name", "tag"),
		).
		Find(&lostClans).Error
	return lostClans, err
}

func (repo *LostClansRepo) LostClan(tag string) (*models.LostClan, error) {
	var clan *models.LostClan
	err := repo.db.First(&clan, "tag = ?", tag).Error
	return clan, err
}

func (repo *LostClansRepo) NameByTag(tag string) (string, error) {
	var name string
	err := repo.db.
		Model(&models.LostClan{}).
		Select("name").
		First(&name, "tag = ?", tag).Error
	return name, err
}
