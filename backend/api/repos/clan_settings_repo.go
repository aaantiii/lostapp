package repos

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/aaantiii/lostapp/backend/store/postgres"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IClanSettingsRepo interface {
	ClanSettings(clanTag string, preload ...string) (*models.ClanSettings, error)
	UpdateClanSettings(settings *models.ClanSettings) error
}

type ClanSettingsRepo struct {
	db *gorm.DB
}

func NewClanSettingsRepo(db *gorm.DB) IClanSettingsRepo {
	return &ClanSettingsRepo{db: db}
}

func (r *ClanSettingsRepo) ClanSettings(clanTag string, preload ...string) (*models.ClanSettings, error) {
	clanSettings := &models.ClanSettings{ClanTag: clanTag}
	err := r.db.
		Clauses(clause.Returning{}).
		Scopes(postgres.WithPreloading(preload...)).
		FirstOrCreate(&clanSettings).Error
	return clanSettings, err
}

func (r *ClanSettingsRepo) UpdateClanSettings(settings *models.ClanSettings) error {
	return r.db.Save(settings).Error
}
