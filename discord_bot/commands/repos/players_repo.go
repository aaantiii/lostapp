package repos

import (
	"sort"
	"strings"

	"gorm.io/gorm"

	"bot/store/postgres"
	"bot/store/postgres/models"
	"bot/types"
)

type IPlayersRepo interface {
	Players(query string) (models.Players, error)
	PlayerByTag(tag string) (*models.Player, error)
	NameByTag(tag string) (string, error)
	MembersPlayersByClan(clanTag, query string) (models.Players, error)
}

type PlayersRepo struct {
	db *gorm.DB
}

func NewPlayersRepo(db *gorm.DB) IPlayersRepo {
	return &PlayersRepo{db: db}
}

func (repo *PlayersRepo) Players(query string) (models.Players, error) {
	var players models.Players
	err := repo.db.
		Scopes(
			postgres.ScopeLimit(types.MaxCommandChoices),
			postgres.ScopeContains(query, "name", "coc_tag"),
		).
		Find(&players).Error
	return players, err
}

func (repo *PlayersRepo) PlayerByTag(tag string) (*models.Player, error) {
	var clan *models.Player
	err := repo.db.First(&clan, "coc_tag = ?", tag).Error
	return clan, err
}

func (repo *PlayersRepo) NameByTag(tag string) (string, error) {
	var name string
	err := repo.db.
		Model(&models.Player{}).
		Select("name").
		First(&name, "coc_tag = ?", tag).Error
	return name, err
}

func (repo *PlayersRepo) MembersPlayersByClan(clanTag, query string) (models.Players, error) {
	var players models.Players
	if err := repo.db.
		Scopes(postgres.ScopeContains(query, "coc_tag", "name")).
		Where("coc_tag IN (?)", repo.db.
			Model(&models.Member{}).
			Select("player_tag").
			Where("clan_tag = ?", clanTag),
		).
		Limit(types.MaxCommandChoices).
		Find(&players).Error; err != nil {
		return nil, err
	}

	sort.Slice(players, func(i, _ int) bool {
		return strings.HasPrefix(strings.ToLower(players[i].Name), query)
	})

	return players, nil
}
