package repos

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"bot/store/postgres/models"
)

type IClanSettingsRepo interface {
	ClanSettings(clanTag string) (*models.ClanSettings, error)
	ClanSettingsPreload(clanTag string) (*models.ClanSettings, error)
	UpdateClanSettings(settings *models.ClanSettings) error
}

type ClanSettingsRepo struct {
	db *gorm.DB
}

func NewClanSettingsRepo(db *gorm.DB) IClanSettingsRepo {
	return &ClanSettingsRepo{db: db}
}

func (repo *ClanSettingsRepo) ClanSettings(clanTag string) (*models.ClanSettings, error) {
	clanSettings := &models.ClanSettings{ClanTag: clanTag}
	err := repo.db.Clauses(clause.Returning{}).FirstOrCreate(&clanSettings).Error
	return clanSettings, err
}

func (repo *ClanSettingsRepo) ClanSettingsPreload(clanTag string) (*models.ClanSettings, error) {
	clanSettings := &models.ClanSettings{ClanTag: clanTag}
	err := repo.db.
		Preload(clause.Associations).
		Clauses(clause.Returning{}).
		FirstOrCreate(&clanSettings).Error
	return clanSettings, err
}

func (repo *ClanSettingsRepo) UpdateClanSettings(settings *models.ClanSettings) error {
	return repo.db.Save(settings).Error
}
