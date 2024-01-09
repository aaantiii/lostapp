package messages

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SendError(s *discordgo.Session, i *discordgo.InteractionCreate, msg string) {
	SendEmbed(s, i, NewEmbed(
		"Fehler",
		msg,
		ColorRed,
	))
}

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
		fmt.Sprintf("Der Clan mit dem Tag '%s' konnte nicht gefunden werden.", clanTag),
		ColorRed,
	))
}

func SendMemberNotFound(s *discordgo.Session, i *discordgo.InteractionCreate, memberTag, clanTag string) {
	SendEmbed(s, i, NewEmbed(
		"Mitglied nicht gefunden",
		fmt.Sprintf("Das Mitglied mit dem Tag '%s' konnte im Clan '%s' nicht gefunden werden.", memberTag, clanTag),
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

func SendCocApiError(s *discordgo.Session, i *discordgo.InteractionCreate) {
	SendEmbed(s, i, NewEmbed(
		"API Fehler",
		"Beim Abrufen der Daten von der Clash of Clan API ist ein Fehler aufgetreten.",
		ColorRed,
	))
}
