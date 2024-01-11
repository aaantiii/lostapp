package messages

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/message"

	"bot/store/postgres/models"
	"bot/types"
)

func SendClanStats(s *discordgo.Session, i *discordgo.InteractionCreate, clan *models.Clan, playerStats []*types.PlayerStatistic, compStat *types.ComparableStatistic) {
	printer := message.NewPrinter(message.MatchLanguage("de"))
	desc := fmt.Sprintf("Statistik: **%s**\n\n", compStat.DisplayName)
	for _, player := range playerStats {
		desc += fmt.Sprintf("\n%d. %s (%v)", player.Placement, player.Name, printer.Sprint(player.Value))
	}

	SendEmbed(s, i, NewEmbed(
		fmt.Sprintf("Clan Statistik von %s", clan.Name),
		desc,
		ColorAqua,
	))
}
