package repos

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
	"github.com/aaantiii/lostapp/backend/store/postgres"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IPlayersRepo interface {
	// Players returns a paginated list of players matching given params.
	Players(params types.PlayersParams) (*types.PaginatedResponse[*models.Player], error)
	// Count returns the number of players matching given params.
	Count(params types.PlayersParams) (int64, error)
	// PlayerByTag returns a player by it's tag.
	PlayerByTag(tag string) (*models.Player, error)
	// PlayerByTagAndDiscordID returns a player by tag and discord id.
	PlayerByTagAndDiscordID(tag, discordID string) (*models.Player, error)
}

var playersQueryFields = []string{"name", "coc_tag"}

type PlayersRepo struct {
	db *gorm.DB
}

func NewPlayersRepo(db *gorm.DB) IPlayersRepo {
	return &PlayersRepo{db: db}
}

func (r *PlayersRepo) Players(params types.PlayersParams) (*types.PaginatedResponse[*models.Player], error) {
	count, err := r.Count(params)
	if err != nil {
		return nil, err
	}

	if err = utils.ValidatePagination(params.PaginationParams, count); err != nil {
		return nil, err
	}

	var players models.Players
	if err = r.db.
		Scopes(
			postgres.WithContains(params.Query, playersQueryFields...),
			postgres.WithPagination(params.PaginationParams),
		).
		Where(params.Conds()).
		Preload(clause.Associations).
		Preload("Members.Clan").
		Order("name").
		Find(&players).Error; err != nil {
		return nil, err
	}

	return types.NewPaginatedResponse(players, params.PaginationParams, count), nil
}

func (r *PlayersRepo) Count(params types.PlayersParams) (int64, error) {
	var count int64
	err := r.db.
		Model(&models.Player{}).
		Where(params.Conds()).
		Scopes(postgres.WithContains(params.Query, playersQueryFields...)).
		Count(&count).Error
	return count, err
}

func (r *PlayersRepo) PlayerByTag(tag string) (*models.Player, error) {
	var clan *models.Player
	err := r.db.First(&clan, "coc_tag = ?", tag).Error
	return clan, err
}

func (r *PlayersRepo) PlayerByTagAndDiscordID(tag, discordID string) (*models.Player, error) {
	var clan *models.Player
	err := r.db.First(&clan, "coc_tag = ? AND discord_id = ?", tag, discordID).Error
	return clan, err
}
