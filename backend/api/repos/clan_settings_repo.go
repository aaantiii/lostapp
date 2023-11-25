package repos

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"backend/store/postgres/models"
)

type IClanSettingsRepo interface {
	ClanSettings(clanTag string) (*models.LostClanSettings, error)
	UpdateClanSettings(updatedSettings *models.LostClanSettings) error
}

type ClanSettingsRepo struct {
	db *gorm.DB
}

func NewClanSettingsRepo(db *gorm.DB) IClanSettingsRepo {
	return &ClanSettingsRepo{db: db}
}

func (repo *ClanSettingsRepo) ClanSettings(tag string) (*models.LostClanSettings, error) {
	log.Print(tag)
	clanSettings := &models.LostClanSettings{ClanTag: tag}
	if err := repo.db.Preload(clause.Associations).FirstOrCreate(&clanSettings).Error; err != nil {
		return nil, err
	}

	return clanSettings, nil
}

func (repo *ClanSettingsRepo) UpdateClanSettings(updatedSettings *models.LostClanSettings) error {
	return repo.db.Save(updatedSettings).Error
}
