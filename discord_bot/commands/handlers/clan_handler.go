package handlers

import (
	"sort"

	"github.com/amaanq/coc.go"
	"github.com/bwmarrin/discordgo"

	"bot/client"
	"bot/commands/messages"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/types"
)

type IClanHandler interface {
	ClanStats(s *discordgo.Session, i *discordgo.InteractionCreate)
	HandleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type ClanHandler struct {
	clans     repos.IClansRepo
	cocClient *client.CocClient
}

func NewClanHandler(clans repos.IClansRepo, cocClient *client.CocClient) IClanHandler {
	return &ClanHandler{
		clans:     clans,
		cocClient: cocClient,
	}
}

func (h *ClanHandler) ClanStats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	statistic := util.StringOptionByName(StatisticOptionName, opts)

	if clanTag == "" || statistic == "" {
		messages.SendInvalidInputError(s, i, "Bitte gib einen Clan und eine Statistik an.")
		return
	}

	var compStat *types.ComparableStatistic
	for _, stat := range types.ComparableStats {
		if stat.Name == statistic {
			compStat = &stat
			break
		}
	}
	if compStat == nil {
		messages.SendInvalidInputError(s, i, "Es wurde keine gültige Statistik angegeben.")
		return
	}

	clan, err := h.clans.ClanByTagPreload(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	players := h.cocClient.GetPlayers(clan.ClanMembers.Tags())
	if len(players) == 0 {
		messages.SendCocApiError(s, i)
		return
	}

	var playerStats []*types.PlayerStatistic
	for _, player := range players {
		if player == nil {
			continue
		}

		achv, err := coc.GetAchievement(player.Achievements, coc.Achievement{Name: statistic})
		if err != nil {
			switch statistic {
			case types.StatSeasonWins.Name:
				achv.Value = player.AttackWins
			default:
				messages.SendInvalidInputError(s, i, "Es wurde keine gültige Statistik angegeben.")
				return
			}
		}

		stats := &types.PlayerStatistic{
			Tag:   player.Tag,
			Name:  player.Name,
			Value: achv.Value,
		}
		playerStats = append(playerStats, stats)
	}

	sort.SliceStable(playerStats, func(i, j int) bool {
		return playerStats[i].Value > playerStats[j].Value
	})

	for index, player := range playerStats {
		player.Placement = index + 1
	}

	messages.SendClanStats(s, i, clan, playerStats, compStat)
}

func (h *ClanHandler) HandleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	for _, opt := range opts {
		if !opt.Focused {
			continue
		}

		switch opt.Name {
		case ClanTagOptionName:
			autocompleteClans(h.clans, opt.StringValue())(s, i)
		}
	}
}
