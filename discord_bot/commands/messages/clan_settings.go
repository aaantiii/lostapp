package messages

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bot/store/postgres/models"
)

func SendClanSettings(i *discordgo.InteractionCreate, settings *models.ClanSettings) {
	SendEmbedResponse(i, NewEmbed(
		fmt.Sprintf("Einstellungen von %s", settings.Clan.Name),
		fmt.Sprintf("Kickpunkte in %s laufen nach **%d Tagen** ab.\nBei **%d Kickpunkten** erfolgt ein Kick.\nPro Season sind **mindestens %d Wins** erforderlich.\n\nFür folgende Regelbrüche wird die jeweilige Anzahl an Kickpunkten vergeben:",
			settings.Clan.Name, settings.KickpointsExpireAfterDays, settings.MaxKickpoints, settings.MinSeasonWins,
		),
		ColorAqua,
	))
}

func kickpointAmountField(name string, value int) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:  name,
		Value: fmt.Sprintf("%d Kickpunkte", value),
	}
}
