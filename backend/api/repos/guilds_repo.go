package repos

import (
	"gorm.io/gorm"

	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IGuildsRepo interface {
	GuildsByGuildID(id string) ([]*models.Guild, error)
	GuildByClanTag(id, clanTag string) (*models.Guild, error)
}

type GuildsRepo struct {
	db *gorm.DB
}

func NewGuildsRepo(db *gorm.DB) IGuildsRepo {
	return &GuildsRepo{db: db}
}

func (r *GuildsRepo) GuildsByGuildID(id string) ([]*models.Guild, error) {
	var guilds []*models.Guild
	err := r.db.Find(&guilds, "guild_id = ?", id).Error
	return guilds, err
}

func (r *GuildsRepo) GuildByClanTag(id, clanTag string) (*models.Guild, error) {
	var guild *models.Guild
	err := r.db.First(&guild, "guild_id = ? AND clan_tag = ?", id, clanTag).Error
	return guild, err
}
