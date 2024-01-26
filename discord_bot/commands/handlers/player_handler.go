package handlers

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aaantiii/goclash"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/messages"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/env"
	"bot/store/postgres/models"
)

type IPlayerHandler interface {
	VerifyPlayer(s *discordgo.Session, i *discordgo.InteractionCreate)
	PingPlayer(s *discordgo.Session, i *discordgo.InteractionCreate)
	HandleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type PlayerHandler struct {
	players     repos.IPlayersRepo
	clashClient *goclash.Client
}

const cocVerificationStatusOK = "ok"

func NewPlayerHandler(players repos.IPlayersRepo, clashClient *goclash.Client) IPlayerHandler {
	return &PlayerHandler{
		players:     players,
		clashClient: clashClient,
	}
}

func (h *PlayerHandler) VerifyPlayer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	go util.DeleteInteractionResponseWithTimeout(s, i.Interaction, time.Minute*2)

	playerTag := util.StringOptionByName(PlayerTagOptionName, opts)
	apiToken := util.StringOptionByName(ApiTokenOptionName, opts)
	if playerTag == "" || apiToken == "" {
		messages.SendInvalidInputErr(i, "Bitte gib einen Spieler-Tag und einen API-Token an.")
		return
	}

	if !strings.HasPrefix(playerTag, "#") {
		messages.SendInvalidInputErr(i, "Bitte gib einen gültigen Spieler-Tag an.")
		return
	}

	verification, err := h.clashClient.VerifyPlayer(playerTag, apiToken)
	if err != nil {
		var apiErr *goclash.ClientError
		if !errors.As(err, &apiErr) {
			messages.SendUnknownErr(i)
			return
		}

		messages.SendInvalidInputErr(i, fmt.Sprintf("Bei der Anfrage zur Clash of Clans-API ist ein Fehler aufgetreten (Fehlercode %d).", apiErr.Status))
		return
	}

	if verification.Status != cocVerificationStatusOK {
		messages.SendInvalidInputErr(i, "Dein API-Token oder Spieler Tag ist ungültig. Bitte versuche es erneut mit gültigen Eingaben.")
		return
	}

	player, err := h.clashClient.GetPlayer(playerTag)
	if err != nil {
		messages.SendCocApiErr(i)
		return
	}

	if _, err = h.players.PlayerByTagAndDiscordID(playerTag, i.Member.User.ID); err == nil {
		messages.SendEmbedResponse(i, messages.NewEmbed(
			"Account bereits verifiziert",
			fmt.Sprintf("%s, dein Clash of Clans Account %s (%s) ist bereits verifiziert und mit deinem Discord Account verknüpft!", i.Member.Mention(), player.Name, player.Tag),
			messages.ColorRed,
		))
		return
	}

	existingPlayers, err := h.players.PlayersByDiscordID(i.Member.User.ID)
	if (err == nil || errors.Is(err, gorm.ErrRecordNotFound)) && len(existingPlayers) == 0 {
		if _, err = s.GuildMemberEdit(i.GuildID, i.Member.User.ID, &discordgo.GuildMemberParams{
			Nick: player.Name,
		}); err != nil {
			messages.SendEmbedResponse(i, messages.NewEmbed(
				"Fehler",
				"Der Bot konnte deinen Discord Namen nicht zu deinem Clash of Clans Namen ändern. Dies liegt wahrscheinlich an fehlenden Berechtigungen des Bots.",
				messages.ColorRed,
			))
			return
		}
	}

	if err = h.players.CreateOrUpdatePlayer(&models.Player{
		CocTag:    player.Tag,
		Name:      player.Name,
		DiscordID: i.Member.User.ID,
	}); err != nil {
		messages.SendEmbedResponse(i, messages.NewEmbed(
			"Datenbankfehler",
			"Beim Speichern deines Accounts in der Datenbank ist ein Fehler aufgetreten. Bitte versuche es erneut.",
			messages.ColorRed,
		))
		return
	}

	if err = s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, env.DISCORD_VERIFIED_ROLE_ID.Value()); err != nil {
		messages.SendEmbedResponse(i, messages.NewEmbed(
			"Erfolgreich verifiziert",
			fmt.Sprintf(
				"%s, dein Clash of Clans Account %s (%s) wurde erfolgreich verifiziert und mit deinem Discord Account verknüpft!\n**Achtung**: Der Bot konnte dir die Verified-Rolle nicht automatisch zuweisen. Bitte frage einen Vize, ob er sie dir manuell geben kann.",
				i.Member.Mention(),
				player.Name,
				player.Tag),
			messages.ColorGreen,
		))
		return
	}

	messages.SendEmbedResponse(i, messages.NewEmbed(
		"Erfolgreich verifiziert",
		fmt.Sprintf("%s, dein Clash of Clans Account %s (%s) wurde erfolgreich verifiziert und mit deinem Discord Account verknüpft!", i.Member.Mention(), player.Name, player.Tag),
		messages.ColorGreen,
	))
}

func (h *PlayerHandler) PingPlayer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	playerTag := util.StringOptionByName(PlayerTagOptionName, opts)
	msg := util.StringOptionByName(MessageOptionName, opts)

	if playerTag == "" || msg == "" {
		messages.SendInvalidInputErr(i, "Bitte gib einen Spieler-Tag und eine Nachricht ein.")
		return
	}

	if !strings.HasPrefix(playerTag, "#") {
		messages.SendInvalidInputErr(i, "Bitte gib einen gültigen Spieler-Tag an.")
		return
	}

	player, err := h.players.PlayerByTag(playerTag)
	if err != nil {
		messages.SendErr(i, "Es wurde kein Spieler mit dem angegebenen Spieler-Tag gefunden.")
		return
	}

	if _, err = s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("%s, du wurdest von %s gepingt:\n%s", util.MentionUserID(player.DiscordID), i.Member.Mention(), msg)); err != nil {
		log.Printf("PingPlayer failed to send message: %v", err)
		messages.SendErr(i, "Beim Senden des Pings ist ein Fehler aufgetreten.")
		return
	}

	messages.SendEmptyResponse(i)
}

func (h *PlayerHandler) HandleAutocomplete(_ *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	for _, opt := range opts {
		if !opt.Focused {
			continue
		}

		switch opt.Name {
		case PlayerTagOptionName:
			autocompletePlayers(i, h.players, opt.StringValue())
		}
	}
}
