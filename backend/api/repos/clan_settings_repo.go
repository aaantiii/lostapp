package repos

import (
	"gorm.io/gorm"

	"backend/store/postgres/models"
)

type IClanSettingsRepo interface {
	ClanSettings(clanTag string) (*models.LostClanSettings, error)
	ClansSettings(clanTags []string) ([]*models.LostClanSettings, error)
	UpdateClanSettings(updatedSettings *models.LostClanSettings) error
}

type ClanSettingsRepo struct {
	db *gorm.DB
}

func NewClanSettingsRepo(db *gorm.DB) *ClanSettingsRepo {
	return &ClanSettingsRepo{db: db}
}

func (repo *ClanSettingsRepo) ClanSettings(tag string) (*models.LostClanSettings, error) {
	clanSettings := &models.LostClanSettings{ClanTag: tag}
	if err := repo.db.FirstOrCreate(&clanSettings).Error; err != nil {
		return nil, err
	}

	return clanSettings, nil
}

func (repo *ClanSettingsRepo) ClansSettings(clanTags []string) ([]*models.LostClanSettings, error) {
	var clanSettings []*models.LostClanSettings
	if err := repo.db.Find(&clanSettings, "clan_tag IN ?", clanTags).Error; err != nil {
		return nil, err
	}

	orderedClanSettings := make([]*models.LostClanSettings, len(clanSettings))
	for i, tag := range clanTags {
		for _, clanSetting := range clanSettings {
			if clanSetting.ClanTag == tag {
				orderedClanSettings[i] = clanSetting
			}
		}
	}

	return orderedClanSettings, nil
}

func (repo *ClanSettingsRepo) UpdateClanSettings(updatedSettings *models.LostClanSettings) error {
	return repo.db.Save(updatedSettings).Error
}
