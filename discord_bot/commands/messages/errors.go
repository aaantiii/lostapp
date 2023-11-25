package messages

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SendUnknownError(s *discordgo.Session, i *discordgo.InteractionCreate) {
	SendEmbed(s, i, NewEmbed(
		"Unbekannter Fehler",
		"Es ist ein unbekannter Fehler aufgetreten.",
		ColorRed,
	))
}

func SendClanNotFound(s *discordgo.Session, i *discordgo.InteractionCreate, clanTag string) {
	SendEmbed(s, i, NewEmbed(
		"Clan nicht gefunden",
		fmt.Sprintf("Der Clan mit dem ClanTag '%s' konnte nicht gefunden werden.", clanTag),
		ColorRed,
	))
}

func SendMemberNotFound(s *discordgo.Session, i *discordgo.InteractionCreate, memberTag string) {
	SendEmbed(s, i, NewEmbed(
		"Mitglied nicht gefunden",
		fmt.Sprintf("Das Mitglied mit dem SpielerTag '%s' konnte nicht gefunden werden.", memberTag),
		ColorRed,
	))
}

func SendInvalidInputError(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	SendEmbed(s, i, NewEmbed(
		"Ungültige Eingaben",
		msg,
		ColorRed,
	))
}
