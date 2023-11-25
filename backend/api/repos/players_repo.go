package repos

import (
	"gorm.io/gorm"

	"backend/api/types"
	"backend/store/postgres"
	"backend/store/postgres/models"
)

type IPlayersRepo interface {
	Players(params *types.PaginationParams) ([]*models.Player, error)
	PlayerByTag(tag string) (*models.Player, error)
	PlayersByTags(tags []string) ([]*models.Player, error)
	PlayersByDiscordID(discordID string) ([]*models.Player, error)
	PlayersByClanTag(tag string) ([]*models.Player, error)
}

type PlayersRepo struct {
	db *gorm.DB
}

func NewPlayersRepo(db *gorm.DB) IPlayersRepo {
	return &PlayersRepo{db: db}
}

func (repo *PlayersRepo) Players(params *types.PaginationParams) ([]*models.Player, error) {
	var players []*models.Player
	if err := repo.db.Scopes(postgres.Paginate(params)).Find(players).Error; err != nil {
		panic(err)
	}

	return players, nil
}

func (repo *PlayersRepo) PlayerByTag(tag string) (*models.Player, error) {
	var player *models.Player
	err := repo.db.First(&player, "tag = ?", tag).Error
	return player, err
}

func (repo *PlayersRepo) PlayersByTags(tags []string) ([]*models.Player, error) {
	var players []*models.Player
	err := repo.db.Find(&players, "tag IN ?", tags).Error
	return players, err
}

func (repo *PlayersRepo) PlayersByDiscordID(discordID string) ([]*models.Player, error) {
	var players []*models.Player
	err := repo.db.Find(&players, "discord_id = ?", discordID).Error
	return players, err
}

func (repo *PlayersRepo) PlayersByClanTag(tag string) ([]*models.Player, error) {
	var playerTags []string
	if err := repo.db.
		Model(&models.Member{}).
		Where("clan_tag = ?", tag).
		Pluck("player_tag", &playerTags).Error; err != nil {
		return nil, err
	}

	var players []*models.Player
	err := repo.db.Find(&players, "tag IN ?", playerTags).Error
	return players, err
}
