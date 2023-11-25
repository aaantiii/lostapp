package services

import (
	"errors"
	"sort"
	"strings"

	"backend/api/repos"
	"backend/api/types"
)

type IPlayersService interface {
	Players(params types.PlayersParams) (types.Players, error)
	PlayerByTag(tag string) (*types.Player, error)
	PlayersByTags(tags []string) (types.Players, error)
	PlayersByDiscordID(discordID string) (types.Players, error)
	PlayersByClanTag(tag string) (types.Players, error)
	PlayersLeaderboard(params types.LeaderboardPlayersParams) ([]*types.PlayerStatistic, error)
}

type PlayersService struct {
	playersRepo repos.IPlayersRepo
	clansRepo   repos.IClansRepo
}

func NewPlayersService(playersRepo repos.IPlayersRepo, clansRepo repos.IClansRepo) IPlayersService {
	return &PlayersService{playersRepo: playersRepo, clansRepo: clansRepo}
}

func (service *PlayersService) Players(params types.PlayersParams) (types.Players, error) {
	return service.runPlayersFilter(service.playersRepo.Players(), params)
}

func (service *PlayersService) PlayerByTag(tag string) (*types.Player, error) {
	return service.playersRepo.PlayerByTag(tag)
}

func (service *PlayersService) PlayersByTags(tags []string) (types.Players, error) {
	return service.playersRepo.PlayersByTags(tags)
}

func (service *PlayersService) PlayersByDiscordID(discordID string) (types.Players, error) {
	return service.playersRepo.PlayersByDiscordID(discordID)
}

func (service *PlayersService) PlayersByClanTag(tag string) (types.Players, error) {
	return service.playersRepo.PlayersByClanTag(tag)
}

func (service *PlayersService) PlayersLeaderboard(params types.LeaderboardPlayersParams) ([]*types.PlayerStatistic, error) {
	players, err := service.Players(types.PlayersParams{ClanTag: params.ClanTag})
	if err != nil || len(players) == 0 {
		return nil, err
	}

	wantedStatName := ""
	for _, stat := range types.ComparableStats {
		if stat.ID == params.StatsID {
			wantedStatName = stat.Name
		}
	}

	stats := make([]*types.PlayerStatistic, len(players))
	for i, player := range players {
		stats[i] = types.NewPlayerStatistic(player, player.ComparableStatsByName[wantedStatName])
	}

	sort.SliceStable(stats, func(i, j int) bool {
		return stats[i].Value > stats[j].Value
	})

	return stats, nil
}

func (service *PlayersService) runPlayersFilter(players []*types.Player, params types.PlayersParams) ([]*types.Player, error) {
	if params.DiscordID != "" {
		return service.PlayersByDiscordID(params.DiscordID)
	}

	if params.Name != "" {
		params.Name = strings.ToLower(params.Name)
		var filteredPlayers []*types.Player
		for _, player := range players {
			if strings.Contains(strings.ToLower(player.Name), params.Name) {
				filteredPlayers = append(filteredPlayers, player)
			}
		}

		if filteredPlayers == nil {
			return nil, errors.New("no players found with the given name")
		}

		sort.SliceStable(filteredPlayers, func(i, j int) bool {
			return strings.HasPrefix(filteredPlayers[i].Name, params.Name)
		})
		players = filteredPlayers
	}

	if params.Tag != "" {
		params.Tag = strings.ToUpper(params.Tag)
		var filteredPlayers []*types.Player
		for _, player := range players {
			if strings.Contains(player.Tag, params.Tag) {
				filteredPlayers = append(filteredPlayers, player)
			}
		}

		if filteredPlayers == nil {
			return nil, errors.New("no players found with the given tag")
		}
		players = filteredPlayers
	}

	if params.ClanName != "" {
		params.ClanName = strings.ToLower(params.ClanName)
		var filteredPlayers []*types.Player
		for _, player := range players {
			for _, clan := range player.Clans {
				if strings.Contains(strings.ToLower(clan.Name), params.ClanName) {
					filteredPlayers = append(filteredPlayers, player)
					break
				}
			}
		}

		if filteredPlayers == nil {
			return nil, errors.New("no players found with the given clan name")
		}

		players = filteredPlayers
	}

	if params.ClanTag != "" {
		params.ClanTag = strings.ToUpper(params.ClanTag)
		var filteredPlayers []*types.Player
		var matchingClanTags []string
		for _, clan := range service.clansRepo.Clans() {
			if strings.Contains(clan.Tag, params.ClanTag) {
				matchingClanTags = append(matchingClanTags, clan.Tag)
			}
		}

		for _, tag := range matchingClanTags {
			clanMembers, err := service.PlayersByClanTag(tag)
			if err != nil {
				continue
			}
			filteredPlayers = append(filteredPlayers, clanMembers...)
		}

		if filteredPlayers == nil {
			return nil, errors.New("no players found with the given clan tag")
		}
		players = filteredPlayers
	}

	return players, nil
}
