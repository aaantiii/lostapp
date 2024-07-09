package services

import (
	"errors"
	"log/slog"
	"sort"

	"github.com/aaantiii/goclash"

	"github.com/aaantiii/lostapp/backend/api/repos"
	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IPlayersService interface {
	Players(params types.PlayersParams) (*types.PaginatedResponse[*models.Player], error)
	PlayersLive(params types.PlayersParams) (*types.PaginatedResponse[*types.Player], error)
	PlayerByTag(tag string) (*models.Player, error)
	Leaderboard(params types.LeaderboardParams, stat *types.ComparableStatistic) (*types.PaginatedResponse[*types.PlayerStatistic], error)
}

type PlayersService struct {
	players     repos.IPlayersRepo
	members     repos.IMembersRepo
	clashClient *goclash.Client
}

func NewPlayersService(player repos.IPlayersRepo, member repos.IMembersRepo, clashClient *goclash.Client) IPlayersService {
	return &PlayersService{
		players:     player,
		members:     member,
		clashClient: clashClient,
	}
}

func (s *PlayersService) Players(params types.PlayersParams) (*types.PaginatedResponse[*models.Player], error) {
	return s.players.Players(params)
}

func (s *PlayersService) PlayersLive(params types.PlayersParams) (*types.PaginatedResponse[*types.Player], error) {
	players, err := s.players.Players(params)
	if err != nil {
		return nil, err
	}
	slog.Debug("players: ", slog.Any("players", players))
	tags := make([]string, len(players.Items))
	for i, player := range players.Items {
		tags[i] = player.CocTag
	}

	livePlayers, err := s.clashClient.GetPlayers(tags...)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(livePlayers, func(i, j int) bool {
		return livePlayers[i].ExpLevel > livePlayers[j].ExpLevel
	})

	res := make([]*types.Player, len(livePlayers))
	for i, player := range livePlayers {
		res[i] = &types.Player{
			PlayerBase:  player.PlayerBase,
			DiscordID:   players.Items[i].DiscordID,
			ClanMembers: players.Items[i].Members,
		}
	}

	return &types.PaginatedResponse[*types.Player]{Items: res, Pagination: players.Pagination}, nil
}

func (s *PlayersService) PlayerByTag(tag string) (*models.Player, error) {
	return s.players.PlayerByTag(tag)
}

func (s *PlayersService) Leaderboard(params types.LeaderboardParams, stat *types.ComparableStatistic) (*types.PaginatedResponse[*types.PlayerStatistic], error) {
	mp := types.MembersParams{}
	mp.ClanTag = params.ClanTag
	tags, err := s.members.MemberTagsDistinct(mp)
	if err = utils.ValidatePagination(params.PaginationParams, int64(len(tags))); err != nil {
		return nil, err
	}

	players, err := s.clashClient.GetPlayers(tags...)
	if err != nil {
		return nil, err
	}
	values, err := utils.StatisticValueFromPlayers(players, stat)
	if err != nil {
		return nil, err
	}

	playerStats := make([]*types.PlayerStatistic, len(players))
	for i, player := range players {
		playerStats[i] = &types.PlayerStatistic{
			Tag:      player.Tag,
			Name:     player.Name,
			ClanName: player.Clan.Name,
			Value:    values[i],
		}
	}
	sort.SliceStable(playerStats, func(i, j int) bool {
		return playerStats[i].Value > playerStats[j].Value
	})

	for i, player := range playerStats {
		player.Placement = i + 1
	}

	start := params.Limit * (params.Page - 1)
	if start > len(playerStats) {
		return nil, errors.New("page out of bounds")
	}

	end := params.Limit * params.Page
	if end > len(playerStats) {
		end = len(playerStats)
	}

	return types.NewPaginatedResponse(playerStats[start:end], params.PaginationParams, int64(len(playerStats))), nil
}
