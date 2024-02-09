package repos

import (
	"gorm.io/gorm"

	"bot/store/postgres"
	"bot/store/postgres/models"
)

type IKickpointReasonsRepo interface {
	KickpointReasons(clanTag string) ([]*models.KickpointReason, error)
	KickpointReason(name, clanTag string) (*models.KickpointReason, error)
	FindKickpointReasons(clanTag, query string) ([]*models.KickpointReason, error)
	CreateKickpointReason(reason *models.KickpointReason) error
	UpdateKickpointReason(reason *models.KickpointReason) error
	DeleteKickpointReason(name, clanTag string) error
}

type KickpointReasonsRepo struct {
	db *gorm.DB
}

func NewKickpointReasonsRepo(db *gorm.DB) IKickpointReasonsRepo {
	return &KickpointReasonsRepo{db: db}
}

func (repo *KickpointReasonsRepo) KickpointReasons(clanTag string) ([]*models.KickpointReason, error) {
	var reasons []*models.KickpointReason
	err := repo.db.Find(&reasons, "clan_tag = ?", clanTag).Error
	return reasons, err
}

func (repo *KickpointReasonsRepo) KickpointReason(name, clanTag string) (*models.KickpointReason, error) {
	var reason models.KickpointReason
	err := repo.db.First(&reason, "name = ? AND clan_tag = ?", name, clanTag).Error
	return &reason, err
}

func (repo *KickpointReasonsRepo) FindKickpointReasons(clanTag, query string) ([]*models.KickpointReason, error) {
	var reasons []*models.KickpointReason
	err := repo.db.
		Scopes(postgres.WithSearchQuery(query, "name")).
		Limit(25).
		Find(&reasons, "clan_tag = ?", clanTag).Error
	return reasons, err
}

func (repo *KickpointReasonsRepo) CreateKickpointReason(reason *models.KickpointReason) error {
	return repo.db.Create(reason).Error
}

func (repo *KickpointReasonsRepo) UpdateKickpointReason(reason *models.KickpointReason) error {
	return repo.db.Save(reason).Error
}

func (repo *KickpointReasonsRepo) DeleteKickpointReason(name, clanTag string) error {
	return repo.db.Delete(&models.KickpointReason{}, "name = ? AND clan_tag = ?", name, clanTag).Error
}
