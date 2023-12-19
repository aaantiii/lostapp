package handlers

import (
	"errors"
	"fmt"

	"github.com/amaanq/coc.go"
	"github.com/bwmarrin/discordgo"

	"bot/client"
	"bot/commands/messages"
	"bot/commands/repos"
	"bot/store/postgres/models"
)

type IPlayerHandler interface {
	VerifyPlayer(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type PlayerHandler struct {
	players   repos.IPlayersRepo
	cocClient *client.CocClient
}

const cocVerificationStatusOK = "ok"

func NewPlayerHandler(players repos.IPlayersRepo, cocClient *client.CocClient) IPlayerHandler {
	return &PlayerHandler{
		players:   players,
		cocClient: cocClient,
	}
}

func (h *PlayerHandler) VerifyPlayer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options

	if len(opts) != 2 {
		messages.SendInvalidInputError(s, i, "Bitte gib einen Spieler-Tag und einen API-Token an.")
		return
	}

	playerTag := opts[0].StringValue()
	apiToken := opts[1].StringValue()

	verification, err := h.cocClient.VerifyPlayerToken(playerTag, apiToken)
	if err != nil {
		var apiErr *coc.APIError
		if !errors.As(err, &apiErr) {
			messages.SendUnknownError(s, i)
			return
		}

		messages.SendInvalidInputError(s, i, fmt.Sprintf("Bei der Anfrage zur Clash of Clans-API ist ein Fehler aufgetreten (Fehlercode %d).", apiErr.StatusCode))
		return
	}

	if verification.Status != cocVerificationStatusOK {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Ungültiger API-Token",
			"Dein API-Token ist ungültig. Bitte versuche es erneut mit einem gültigen Token.",
			messages.ColorRed,
		))
		return
	}

	cocPlayer, err := h.cocClient.GetPlayer(playerTag)
	if err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Fehler",
			"Der API-Token konnte erfolgreich verifiziert werden, jedoch konnte der Spieler nicht abgerufen werden.",
			messages.ColorRed,
		))
		return
	}

	if _, err = h.players.PlayerByTagAndDiscordID(playerTag, i.Member.User.ID); err == nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Account bereits verifiziert",
			fmt.Sprintf("%s, dein Clash of Clans Account %s (%s) ist bereits verifiziert und mit deinem Discord Account verknüpft!", i.Member.Mention(), cocPlayer.Name, cocPlayer.Tag),
			messages.ColorRed,
		))
		return
	}

	if existingPlayers, _ := h.players.PlayersByDiscordID(i.Member.User.ID); len(existingPlayers) == 0 {
		if _, err = s.GuildMemberEdit(i.GuildID, i.Member.User.ID, &discordgo.GuildMemberParams{
			Nick: cocPlayer.Name,
		}); err != nil {
			messages.SendEmbed(s, i, messages.NewEmbed(
				"Fehler",
				"Der Bot konnte deinen Discord Namen nicht zu deinem Clash of Clans Namen ändern.",
				messages.ColorGreen,
			))
			return
		}
	}

	if err = h.players.CreateOrUpdatePlayer(&models.Player{
		CocTag:    cocPlayer.Tag,
		Name:      cocPlayer.Name,
		DiscordID: i.Member.User.ID,
	}); err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Datenbankfehler",
			"Beim Speichern deines Accounts in der Datenbank ist ein Fehler aufgetreten. Bitte versuche es erneut.",
			messages.ColorRed,
		))
		return
	}

	messages.SendEmbed(s, i, messages.NewEmbed(
		"Erfolgreich verifiziert",
		fmt.Sprintf("%s, dein Clash of Clans Account %s (%s) wurde erfolgreich verifiziert und mit deinem Discord Account verknüpft!", i.Member.Mention(), cocPlayer.Name, cocPlayer.Tag),
		messages.ColorGreen,
	))
}
