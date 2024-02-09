package messages

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aaantiii/goclash"
	"github.com/bwmarrin/discordgo"

	"bot/commands/util"
)

func SendErr(i *discordgo.InteractionCreate, msg string) {
	SendEmbedResponse(i, NewEmbed(
		"Fehler",
		msg,
		ColorRed,
	))
}

func SendWarning(i *discordgo.InteractionCreate, msg string) {
	SendEmbedResponse(i, NewEmbed(
		"Warnung",
		msg,
		ColorYellow,
	))
}

func SendChannelWarning(channelID, msg string) {
	if _, err := util.Session.ChannelMessageSendEmbed(channelID, NewEmbed(":warning: Warnung :warning:", msg, ColorYellow)); err != nil {
		log.Printf("Error while sending warning message: %v", err)
	}
}

func SendUnknownErr(i *discordgo.InteractionCreate) {
	SendEmbedResponse(i, NewEmbed(
		"Unbekannter Fehler",
		"Es ist ein unbekannter Fehler aufgetreten.",
		ColorRed,
	))
}

func SendClanNotFound(i *discordgo.InteractionCreate, clanTag string) {
	SendEmbedResponse(i, NewEmbed(
		"Clan nicht gefunden",
		fmt.Sprintf("Der Clan mit dem Tag `%s` konnte nicht gefunden werden. Stelle sicher, dass du den Clan aus der Liste ausgewählt hast, oder direkt einen Clan Tag eingegeben hast.", clanTag),
		ColorRed,
	))
}

func SendMemberNotFound(i *discordgo.InteractionCreate, memberTag, clanTag string) {
	SendEmbedResponse(i, NewEmbed(
		"Mitglied nicht gefunden",
		fmt.Sprintf("Das Mitglied mit dem Tag '%s' konnte im Clan '%s' nicht gefunden werden.", memberTag, clanTag),
		ColorRed,
	))
}

func SendClanHasNoMembers(i *discordgo.InteractionCreate, clanName string) {
	SendEmbedResponse(i, NewEmbed(
		"Clan hat keine Mitglieder",
		fmt.Sprintf("Der Clan '%s' hat keine Mitglieder.", clanName),
		ColorRed,
	))
}

func SendInvalidInputErr(i *discordgo.InteractionCreate, msg string) {
	SendEmbedResponse(i, NewEmbed(
		"Ungültige Eingaben",
		msg,
		ColorRed,
	))
}

func SendCocApiErr(i *discordgo.InteractionCreate, err error) {
	var cErr *goclash.ClientError
	if errors.As(err, &cErr) {
		switch cErr.Status {
		case http.StatusNotFound:
			SendEmbedResponse(i, NewEmbed(
				"API Fehler",
				"Die angeforderten Daten konnten nicht gefunden werden.",
				ColorRed,
			))
			return
		case http.StatusTooManyRequests:
			SendEmbedResponse(i, NewEmbed(
				"API Fehler",
				"Die Anfrage an die Clash of Clans API wurde zu oft gestellt. Bitte versuche es später erneut.",
				ColorRed,
			))
			return
		case http.StatusServiceUnavailable:
			SendEmbedResponse(i, NewEmbed(
				"API Fehler",
				"Die Clash of Clans API ist momentan unter Wartungsarbeiten. Bitte versuche es später erneut.",
				ColorRed,
			))
			return
		}
	}

	SendEmbedResponse(i, NewEmbed(
		"API Fehler",
		"Beim Abrufen der Daten von der Clash of Clan API ist ein Fehler aufgetreten.",
		ColorRed,
	))
}

func SendInvalidDateTimeFormat(i *discordgo.InteractionCreate, field string) {
	SendInvalidInputErr(i, fmt.Sprintf("Das Datumsformat vom Feld %s ist ungültig. Bitte gib ein Datum im Format `DD.MM.YYYY HH:MM` an.", field))
}
