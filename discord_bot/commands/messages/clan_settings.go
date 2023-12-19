package messages

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bot/store/postgres/models"
)

func SendClanSettings(s *discordgo.Session, i *discordgo.InteractionCreate, settings *models.ClanSettings) {
	fields := []*discordgo.MessageEmbedField{
		kickpointAmountField("Season Wins", settings.KickpointsSeasonWins),
		kickpointAmountField("CK nicht angegriffen", settings.KickpointsCWMissed),
		kickpointAmountField("CK Fails (z.B. 0 Sterne)", settings.KickpointsCWFail),
		kickpointAmountField("CWL nicht angegriffen", settings.KickpointsCWLMissed),
		kickpointAmountField("CWL 0 Sterne Angriff", settings.KickpointsCWLZeroStars),
		kickpointAmountField("CWL 1 Sterne Angriff", settings.KickpointsCWLOneStar),
		kickpointAmountField("Raid nicht angegriffen", settings.KickpointsRaidMissed),
		kickpointAmountField("Raid Fails (z.B. nur 1 Angriff)", settings.KickpointsRaidFail),
		kickpointAmountField("Clanspiele nicht gemacht", settings.KickpointsClanGames),
	}

	SendEmbed(s, i, NewFieldEmbed(
		fmt.Sprintf("Einstellungen von %s", settings.Clan.Name),
		fmt.Sprintf("Kickpunkte in %s laufen nach **%d Tagen** ab.\nBei **%d Kickpunkten** erfolgt ein Kick.\nPro Season sind **mindestens %d Wins** erforderlich.\n\nFür folgende Regelbrüche wird die jeweilige Anzahl an Kickpunkten vergeben:",
			settings.Clan.Name, settings.KickpointsExpireAfterDays, settings.MaxKickpoints, settings.MinSeasonWins,
		),
		ColorAqua,
		fields,
	))
}

func kickpointAmountField(name string, value int) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:  name,
		Value: fmt.Sprintf("%d Kickpunkte", value),
	}
}
