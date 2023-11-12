package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"

	"bot/commands/components"
	"bot/commands/messages"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/commands/validation"
	"bot/store/postgres/models"
)

type IPenaltyHandler interface {
	ClanPenalties(s *discordgo.Session, i *discordgo.InteractionCreate)
	MemberPenalties(s *discordgo.Session, i *discordgo.InteractionCreate)
	CreatePenaltyModal(s *discordgo.Session, i *discordgo.InteractionCreate)
	CreatePenaltyModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate)
	EditPenalty(s *discordgo.Session, i *discordgo.InteractionCreate)
	DeletePenalty(s *discordgo.Session, i *discordgo.InteractionCreate)
	ClansAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate)
	MembersAndClansAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type PenaltyHandler struct {
	penalties repos.IPenaltiesRepo
	clans     repos.ILostClansRepo
	players   repos.IPlayersRepo
	members   repos.IMembersRepo
}

func NewPenaltyHandler(penalties repos.IPenaltiesRepo, clans repos.ILostClansRepo, players repos.IPlayersRepo, members repos.IMembersRepo) IPenaltyHandler {
	return &PenaltyHandler{
		penalties: penalties,
		clans:     clans,
		players:   players,
	}
}

func (h *PenaltyHandler) ClanPenalties(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Print("ClanPenalties Main")
}

func (h *PenaltyHandler) MemberPenalties(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Print("sent")
}

func (h *PenaltyHandler) CreatePenaltyModal(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	if len(opts) != 2 {
		messages.SendEmbed(s, i, messages.NewMessageEmbed(
			"Ungültige Eingaben",
			"Du musst einen Spieler und einen Clan angeben.",
			messages.ColorRed,
		))
		return
	}

	clanTag := opts[0].StringValue()
	playerTag := opts[1].StringValue()
	now := time.Now()
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: util.BuildCustomID(i.ApplicationCommandData().Name, i.Interaction.Member.User.ID),
			Title:    "Kickpunkt/Verwarnung hinzufügen",
			Components: components.GenModalComponents(
				components.Tag("Spieler Tag", playerTag, components.PlayerTagID),
				components.Tag("Clan Tag", clanTag, components.ClanTagID),
				components.PenaltyReason(""),
				components.PenaltyType(models.PenaltyTypeNone),
				components.PenaltyDate(util.FormatMonthYear(int(now.Month()), now.Year())),
			),
		},
	})

	if err != nil {
		log.Print(err.Error())
	}
}

func (h *PenaltyHandler) CreatePenaltyModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	_, userID := util.ParseCustomID(data.CustomID)

	if len(data.Components) != 5 {
		messages.SendEmbed(s, i, messages.NewMessageEmbed("Fehler", "Es wurden nicht alle Felder ausgefüllt.", messages.ColorRed))
		return
	}

	penalty := &models.Penalty{
		PlayerTag:          util.ParseStringModalInput(data.Components[0]),
		ClanTag:            util.ParseStringModalInput(data.Components[1]),
		Reason:             util.ParseStringModalInput(data.Components[2]),
		CreatedByDiscordID: userID,
	}

	penaltyType, err := validation.PenaltyType(util.ParseStringModalInput(data.Components[3]))
	if err != nil {
		messages.SendEmbed(s, i, messages.NewMessageEmbed(
			"Ungültige Eingaben",
			"Beim Erstellen des Kickpunkts oder der Verwarnung wurde ein ungültiger Typ angegeben. Erlaubt sind 'K' für Kickpunkt und 'V' für Verwarnung.",
			messages.ColorRed,
		))
		return
	}
	penalty.Type = penaltyType

	month, year := util.ParseMonthYearInput(data.Components[4])
	penalty.Month = month
	penalty.Year = year
	if err = h.penalties.Create(penalty); err != nil {
		messages.SendEmbed(s, i, messages.NewMessageEmbed(
			"Datenbankfehler",
			"Beim Speichern der Daten in der Datenbank ist ein Fehler aufgetreten. Dies liegt wahrscheinlich an der Eingabe ungültiger Daten.",
			messages.ColorRed,
		))
		return
	}

	playerName, err := h.players.NameByTag(penalty.PlayerTag)
	if err != nil {
		messages.SendEmbed(s, i, messages.NewMessageEmbed(
			"Ungültige Eingaben",
			"Der Spieler konnte nicht gefunden werden.",
			messages.ColorRed,
		))
		return
	}

	clanName, err := h.clans.NameByTag(penalty.ClanTag)
	if err != nil {
		messages.SendEmbed(s, i, messages.NewMessageEmbed(
			"Ungültige Eingaben",
			"Der Clan konnte nicht gefunden werden.",
			messages.ColorRed,
		))
		return
	}

	typeStr := penalty.Type.DisplayString()
	messages.SendEmbed(s, i, messages.NewMessageEmbedWithFields(
		fmt.Sprintf("%s erstellt", typeStr),
		fmt.Sprintf("%s wurde erfolgreich erstellt und gespeichert!", typeStr),
		messages.ColorGreen,
		[]*discordgo.MessageEmbedField{
			{Name: "Mitglied", Value: fmt.Sprintf("%s in %s", playerName, clanName)},
			{Name: "Typ", Value: typeStr},
			{Name: "Erhalten am (Monat-Jahr)", Value: util.FormatMonthYear(penalty.Month, penalty.Year)},
			{Name: "Erstellt am", Value: util.FormatDateString(time.Now())},
			{Name: "Grund", Value: penalty.Reason},
		},
	))
}

func (h *PenaltyHandler) EditPenalty(s *discordgo.Session, i *discordgo.InteractionCreate) {

}

func (h *PenaltyHandler) DeletePenalty(s *discordgo.Session, i *discordgo.InteractionCreate) {

}

func (h *PenaltyHandler) ClansAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	autocompleteClans(h.clans, i.ApplicationCommandData().Options[0].StringValue())(s, i)
}

func (h *PenaltyHandler) MembersAndClansAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	for _, opt := range data.Options {
		if opt.Focused {
			if opt.Name == "clan" {
				autocompleteClans(h.clans, opt.StringValue())(s, i)
			} else {
				autocompleteMembers(h.players, opt.StringValue(), data.Options[0].StringValue())(s, i)
			}
			break
		}
	}
}
