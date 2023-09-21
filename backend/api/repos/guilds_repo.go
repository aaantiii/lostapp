package repos

import (
	"gorm.io/gorm"

	"backend/store/postgres/models"
)

type IGuildsRepo interface {
	Guilds() ([]*models.Guild, error)
	Guild(clanTag string) (*models.Guild, error)
}

type GuildsRepo struct {
	db *gorm.DB
}

func NewGuildsRepo(db *gorm.DB) *GuildsRepo {
	return &GuildsRepo{db: db}
}

func (r *GuildsRepo) Guilds() ([]*models.Guild, error) {
	var guilds []*models.Guild
	err := r.db.Find(&guilds, "guild_id = ?", models.LostFamilyGuildID).Error
	return guilds, err
}

func (r *GuildsRepo) Guild(clanTag string) (*models.Guild, error) {
	var guild *models.Guild
	err := r.db.First(&guild, "guild_id = ? AND clan_tag = ?", models.LostFamilyGuildID, clanTag).Error
	return guild, err
}
