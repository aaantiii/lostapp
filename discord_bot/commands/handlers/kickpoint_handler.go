package handlers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/amaanq/coc.go"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/components"
	"bot/commands/messages"
	"bot/commands/middleware"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/store/postgres/models"
)

type IKickpointHandler interface {
	ClanKickpoints(s *discordgo.Session, i *discordgo.InteractionCreate)
	MemberKickpoints(s *discordgo.Session, i *discordgo.InteractionCreate)
	KickpointInfo(s *discordgo.Session, i *discordgo.InteractionCreate)
	CreateKickpointModal(s *discordgo.Session, i *discordgo.InteractionCreate)
	CreateKickpointModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate)
	EditPenalty(s *discordgo.Session, i *discordgo.InteractionCreate)
	DeletePenalty(s *discordgo.Session, i *discordgo.InteractionCreate)
	ClansAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate)
	MembersAndClansAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type KickpointHandler struct {
	kickpoints     repos.IKickpointsRepo
	clans          repos.IClansRepo
	players        repos.IPlayersRepo
	clanSettings   repos.IClanSettingsRepo
	authMiddleware middleware.AuthMiddleware
}

func NewKickpointHandler(kickpoints repos.IKickpointsRepo, clans repos.IClansRepo, players repos.IPlayersRepo, guilds repos.IGuildsRepo, users repos.IUsersRepo, clanSettings repos.IClanSettingsRepo) IKickpointHandler {
	return &KickpointHandler{
		kickpoints:     kickpoints,
		clans:          clans,
		players:        players,
		clanSettings:   clanSettings,
		authMiddleware: middleware.NewAuthMiddleware(guilds, clans, users),
	}
}

func (h *KickpointHandler) ClanKickpoints(s *discordgo.Session, i *discordgo.InteractionCreate) {
	clanTag := i.ApplicationCommandData().Options[0].StringValue()
	if err := h.authMiddleware.NewHandler(clanTag, coc.Member)(s, i); err != nil {
		return
	}

	settings, err := h.clanSettings.ClanSettingsPreload(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	kickpoints, err := h.kickpoints.ActiveClanKickpoints(settings)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			messages.SendEmbed(s, i, messages.NewEmbed(
				"Keine Kickpunkte gefunden",
				fmt.Sprintf("In %s gibt es keine Mitglieder mit aktiven Kickpunkten.", settings.Clan.Name),
				messages.ColorRed,
			))
			return
		}

		messages.SendUnknownError(s, i)
		return
	}

	messages.SendClanKickpoints(s, i, settings.Clan.Name, kickpoints)
}

func (h *KickpointHandler) MemberKickpoints(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	if len(opts) != 2 {
		messages.SendInvalidInputError(s, i, "Du musst einen Spieler und einen Clan angeben.")
		return
	}

	clanTag := opts[0].StringValue()
	if err := h.authMiddleware.NewHandler(clanTag, coc.Member)(s, i); err != nil {
		return
	}

	settings, err := h.clanSettings.ClanSettingsPreload(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	playerTag := opts[1].StringValue()
	kickpoints, err := h.kickpoints.ActiveMemberKickpoints(playerTag, settings)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			messages.SendEmbed(s, i, messages.NewEmbed(
				"Keine aktiven Kickpunkte gefunden",
				"Dieses Mitglied hat keine aktiven Kickpunkte.",
				messages.ColorRed,
			))
			return
		}

		messages.SendMemberNotFound(s, i, playerTag)
		return
	}

	messages.SendMemberKickpoints(s, i, settings.Clan.Name, kickpoints)
}

func (h *KickpointHandler) KickpointInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Print("sent")
}

func (h *KickpointHandler) CreateKickpointModal(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	if len(opts) != 3 {
		messages.SendInvalidInputError(s, i, "Du musst einen Spieler und einen Clan angeben.")
		return
	}

	clanTag := opts[0].StringValue()
	if err := h.authMiddleware.NewHandler(clanTag, coc.CoLeader)(s, i); err != nil {
		return
	}

	clanSettings, err := h.clanSettings.ClanSettings(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	kickpointAmount, err := clanSettings.KickpointAmountFromName(opts[2].StringValue())
	if err != nil {
		messages.SendInvalidInputError(s, i, "Es wurde ein ungültiger Grund angegeben.")
		return
	}

	playerTag := opts[1].StringValue()
	now := time.Now()
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: util.BuildCustomID(i.ApplicationCommandData().Name, i.Interaction.Member.User.ID),
			Title:    "Kickpunkt hinzufügen",
			Components: components.GenModalComponents(
				components.KickpointReason(""),
				components.KickpointDate(util.FormatDate(now)),
				components.Tag("Spieler Tag", playerTag, components.PlayerTagID),
				components.Tag("Clan Tag", clanTag, components.ClanTagID),
				components.KickpointAmount(kickpointAmount),
			),
		},
	})

	if err != nil {
		log.Println("Error while handling CreateKickpoint", "err", err)
	}
}

func (h *KickpointHandler) CreateKickpointModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	_, userID := util.ParseCustomID(data.CustomID)

	if len(data.Components) != 5 {
		messages.SendEmbed(s, i, messages.NewEmbed("Fehler", "Es wurden nicht alle Felder ausgefüllt.", messages.ColorRed))
		return
	}

	date, err := util.ParseDateInput(data.Components[1])
	if err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Ungültiges Datum",
			"Beim Erstellen des Kickpunkts wurde ein ungültiges Datum angegeben. Das Datum muss im Format 'ClanTag.Monat.Jahr' angegeben werden.",
			messages.ColorRed,
		))
		return
	}

	kickpoint := &models.Kickpoint{
		Description:        util.ParseStringModalInput(data.Components[0]),
		PlayerTag:          util.ParseStringModalInput(data.Components[2]),
		ClanTag:            util.ParseStringModalInput(data.Components[3]),
		Amount:             util.ParseIntModalInput(data.Components[4]),
		Date:               date,
		CreatedByDiscordID: userID,
		UpdatedByDiscordID: userID,
	}

	if err = h.kickpoints.CreateKickpoint(kickpoint); err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Datenbankfehler",
			"Beim Speichern der Daten in der Datenbank ist ein Fehler aufgetreten. Dies liegt wahrscheinlich an der Eingabe ungültiger Daten.",
			messages.ColorRed,
		))
		return
	}

	playerName, err := h.players.NameByTag(kickpoint.PlayerTag)
	if err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Ungültige Eingaben",
			"Der Spieler konnte nicht gefunden werden.",
			messages.ColorRed,
		))
		return
	}

	clanName, err := h.clans.ClanNameByTag(kickpoint.ClanTag)
	if err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Ungültige Eingaben",
			"Der Clan konnte nicht gefunden werden.",
			messages.ColorRed,
		))
		return
	}

	messages.SendEmbed(s, i, messages.NewFieldEmbed(
		"Kickpunkt erstellt",
		"Der Kickpunkt wurde erstellt und gespeichert!",
		messages.ColorGreen,
		append([]*discordgo.MessageEmbedField{
			{Name: "Mitglied", Value: fmt.Sprintf("%s in %s", playerName, clanName)}},
			messages.DetailedKickpointFields(kickpoint)...,
		),
	))
}

func (h *KickpointHandler) EditPenalty(s *discordgo.Session, i *discordgo.InteractionCreate) {

}

func (h *KickpointHandler) DeletePenalty(s *discordgo.Session, i *discordgo.InteractionCreate) {
	idOpt := i.ApplicationCommandData().Options[0].UintValue()
	if idOpt == 0 {
		messages.SendInvalidInputError(s, i, "Du musst eine gültige Kickpunkt ID angeben.")
		return
	}

	id := uint(idOpt)
	kickpoint, err := h.kickpoints.KickpointByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			messages.SendEmbed(s, i, messages.NewEmbed(
				"Ungültige ID",
				"Es wurde kein Kickpunkt mit dieser ID gefunden.",
				messages.ColorRed,
			))
			return
		}

		messages.SendUnknownError(s, i)
		return
	}

	if err = h.authMiddleware.NewHandler(kickpoint.ClanTag, coc.CoLeader)(s, i); err != nil {
		return
	}

	if err = h.kickpoints.DeleteKickpoint(id); err != nil {
		messages.SendUnknownError(s, i)
		return
	}

	messages.SendEmbed(s, i, messages.NewFieldEmbed(
		fmt.Sprintf("Kickpunkt #%d gelöscht", kickpoint.ID),
		fmt.Sprintf("Der Kickpunkt von %s in %s wurde gelöscht!", kickpoint.Player.Name, kickpoint.Clan.Name),
		messages.ColorGreen,
		messages.DetailedKickpointFields(kickpoint),
	))
}

func (h *KickpointHandler) ClansAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	autocompleteClans(h.clans, i.ApplicationCommandData().Options[0].StringValue())(s, i)
}

func (h *KickpointHandler) MembersAndClansAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	for _, opt := range data.Options {
		if opt.Focused {
			if opt.Name == "clan" {
				autocompleteClans(h.clans, opt.StringValue())(s, i)
			} else {
				autocompleteMembers(h.players, opt.StringValue(), data.Options[0].StringValue())(s, i)
			}
			continue // sometimes both options are focused...
		}
	}
}
