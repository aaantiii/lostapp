package repos

import (
	"github.com/amaanq/coc.go"
	"gorm.io/gorm"

	"backend/store/postgres/models"
)

type IClansRepo interface {
	Clans() ([]*models.Clan, error)
	Clan(clanTag string) (*models.Clan, error)
}

type ClansRepo struct {
	db        *gorm.DB
	cocClient *coc.Client
}

func NewClansRepo(db *gorm.DB, cocClient *coc.Client) IClansRepo {
	return &ClansRepo{db: db, cocClient: cocClient}
}

func (repo *ClansRepo) Clans() ([]*models.Clan, error) {
	var clans []*models.Clan
	err := repo.db.Order("id").Find(&clans).Error
	return clans, err
}

func (repo *ClansRepo) Clan(tag string) (*models.Clan, error) {
	var clan *models.Clan
	if err := repo.db.Preload("MemberList").First(&clan, "tag = ?", tag).Error; err != nil {
		return nil, err
	}

	return clan, nil
}
