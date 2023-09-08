package repos

import (
	"fmt"

	"gorm.io/gorm"

	"backend/api/types"
	"backend/store/cache"
)

type IPlayersRepo interface {
	Players() types.Players
	PlayerByTag(tag string) (*types.Player, error)
	PlayersByTags(tags []string) (types.Players, error)
	PlayersByDiscordID(discordID string) (types.Players, error)
	PlayersByClanTag(tag string) (types.Players, error)
}

type PlayersRepo struct {
	db    *gorm.DB
	cache *cache.CocCache
}

func NewPlayersRepo(db *gorm.DB, cache *cache.CocCache) *PlayersRepo {
	return &PlayersRepo{db: db, cache: cache}
}

func (repo *PlayersRepo) Players() types.Players {
	return repo.cache.Players
}

func (repo *PlayersRepo) PlayerByTag(tag string) (*types.Player, error) {
	player, found := repo.cache.PlayerByTag[tag]
	if !found {
		return nil, fmt.Errorf("no player found with tag=%s", tag)
	}

	return player, nil
}

func (repo *PlayersRepo) PlayersByTags(tags []string) (types.Players, error) {
	var players types.Players
	for _, tag := range tags {
		if player, err := repo.PlayerByTag(tag); err == nil {
			players = append(players, player)
		} else {
			return nil, err
		}
	}

	return players, nil
}

func (repo *PlayersRepo) PlayersByDiscordID(discordID string) (types.Players, error) {
	players, found := repo.cache.PlayersByDiscordID[discordID]
	if !found {
		return nil, fmt.Errorf("no players found with discord id=%s", discordID)
	}

	return players, nil
}

func (repo *PlayersRepo) PlayersByClanTag(tag string) (types.Players, error) {
	players, found := repo.cache.PlayersByClanTag[tag]
	if !found {
		return nil, fmt.Errorf("no players found with clan tag=%s", tag)
	}

	return players, nil
}
