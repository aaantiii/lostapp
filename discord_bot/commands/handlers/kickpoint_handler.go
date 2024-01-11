package handlers

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/components"
	"bot/commands/messages"
	"bot/commands/middleware"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/commands/validation"
	"bot/store/postgres/models"
	"bot/types"
)

type IKickpointHandler interface {
	ClanKickpoints(s *discordgo.Session, i *discordgo.InteractionCreate)
	MemberKickpoints(s *discordgo.Session, i *discordgo.InteractionCreate)
	KickpointInfo(s *discordgo.Session, i *discordgo.InteractionCreate)
	KickpointConfig(s *discordgo.Session, i *discordgo.InteractionCreate)
	KickpointHelp(s *discordgo.Session, i *discordgo.InteractionCreate)
	CreateKickpointModal(s *discordgo.Session, i *discordgo.InteractionCreate)
	CreateKickpointModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate)
	EditKickpoint(s *discordgo.Session, i *discordgo.InteractionCreate)
	EditKickpointModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate)
	DeleteKickpoint(s *discordgo.Session, i *discordgo.InteractionCreate)
	NewKickpointLockHandler(lock bool) func(s *discordgo.Session, i *discordgo.InteractionCreate)
	HandleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type KickpointHandler struct {
	kickpoints     repos.IKickpointsRepo
	clans          repos.IClansRepo
	players        repos.IPlayersRepo
	clanSettings   repos.IClanSettingsRepo
	memberStates   repos.IMemberStatesRepo
	authMiddleware middleware.AuthMiddleware
}

func NewKickpointHandler(kickpoints repos.IKickpointsRepo, clans repos.IClansRepo, players repos.IPlayersRepo, guilds repos.IGuildsRepo, users repos.IUsersRepo, clanSettings repos.IClanSettingsRepo, memberStates repos.IMemberStatesRepo) IKickpointHandler {
	return &KickpointHandler{
		kickpoints:     kickpoints,
		clans:          clans,
		players:        players,
		clanSettings:   clanSettings,
		memberStates:   memberStates,
		authMiddleware: middleware.NewAuthMiddleware(guilds, clans, users),
	}
}

func (h *KickpointHandler) ClanKickpoints(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	if len(opts) != 1 {
		messages.SendInvalidInputError(s, i, "Du musst einen Clan angeben.")
		return
	}

	clanTag := opts[0].StringValue()
	if err := h.authMiddleware.NewClanHandler(clanTag, types.AuthRoleMember)(s, i); err != nil {
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
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	memberTag := util.StringOptionByName(MemberTagOptionName, opts)
	if clanTag == "" || memberTag == "" {
		messages.SendInvalidInputError(s, i, "Du musst einen Clan und ein Mitglied angeben.")
		return
	}

	if err := h.authMiddleware.NewClanHandler(clanTag, types.AuthRoleMember)(s, i); err != nil {
		return
	}

	settings, err := h.clanSettings.ClanSettingsPreload(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	kickpoints, err := h.kickpoints.ActiveMemberKickpoints(memberTag, settings)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			messages.SendEmbed(s, i, messages.NewEmbed(
				"Keine aktiven Kickpunkte gefunden",
				"Dieses Mitglied hat keine aktiven Kickpunkte.",
				messages.ColorRed,
			))
			return
		}

		messages.SendMemberNotFound(s, i, memberTag, clanTag)
		return
	}

	messages.SendMemberKickpoints(s, i, settings, kickpoints)
}

func (h *KickpointHandler) KickpointInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	if clanTag == "" {
		messages.SendInvalidInputError(s, i, "Du musst einen Clan angeben.")
		return
	}

	settings, err := h.clanSettings.ClanSettingsPreload(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	messages.SendClanSettings(s, i, settings)
}

func (h *KickpointHandler) KickpointConfig(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	settingName := models.KickpointSetting(util.StringOptionByName(SettingOptionName, opts))
	settingValue := util.IntOptionByName(AmountOptionName, opts)

	if clanTag == "" || settingName == "" || settingValue == nil {
		messages.SendInvalidInputError(s, i, "Du musst einen Clan, eine Einstellung und einen Wert angeben.")
		return
	}

	if err := h.authMiddleware.NewClanHandler(clanTag, types.AuthRoleCoLeader)(s, i); err != nil {
		return
	}

	if _, err := h.clanSettings.ClanSettings(clanTag); err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	if msg, ok := validation.ValidateKickpointSettings(settingName, *settingValue); !ok {
		messages.SendInvalidInputError(s, i, msg)
		return
	}

	if err := h.clanSettings.UpdateKickpointSetting(clanTag, settingName, *settingValue); err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed("Fehler", "Beim Speichern der Einstellung ist ein Fehler aufgetreten.", messages.ColorRed))
		return
	}

	messages.SendEmbed(s, i, messages.NewEmbed("Einstellung geändert!", "Die Einstellung wurde erfolgreich geändert.", messages.ColorGreen))
}

func (h *KickpointHandler) KickpointHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	messages.SendKickpointHelp(s, i)
}

func (h *KickpointHandler) CreateKickpointModal(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	memberTag := util.StringOptionByName(MemberTagOptionName, opts)
	settingName := models.KickpointSetting(util.StringOptionByName(SettingOptionName, opts))

	log.Print(clanTag, memberTag, settingName)

	if clanTag == "" || memberTag == "" || settingName == "" {
		messages.SendInvalidInputError(s, i, "Du musst einen Clan, ein Mitglied und einen Grund angeben.")
		return
	}

	if err := h.authMiddleware.NewClanHandler(clanTag, types.AuthRoleCoLeader)(s, i); err != nil {
		return
	}

	isLocked, err := h.memberStates.IsKickpointLocked(memberTag, clanTag)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		messages.SendUnknownError(s, i)
		return
	}
	if isLocked {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Mitglied gesperrt",
			"Dieses Mitglied kann momentan keine Kickpunkte erhalten, da es abgemeldet ist.",
			messages.ColorRed,
		))
		return
	}

	clanSettings, err := h.clanSettings.ClanSettings(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	kickpointAmount, err := clanSettings.KickpointAmountFromSetting(settingName)
	if err != nil {
		messages.SendInvalidInputError(s, i, "Es wurde ein ungültiger Grund angegeben.")
		return
	}

	if kickpointAmount == 0 {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Kein Kickpunkt erforderlich",
			fmt.Sprintf("Dieser Regelbruch gibt in %s keine Kickpunkte. Du kannst die Anzahl der Kickpunkte mit ```kpconfig``` ändern.", clanSettings.Clan.Name),
			messages.ColorRed,
		))
		return
	}

	settings, err := h.clanSettings.ClanSettings(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	totalKickpoints, err := h.kickpoints.ActiveMemberKickpointsSum(memberTag, settings)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		messages.SendUnknownError(s, i)
		return
	}

	if totalKickpoints >= settings.MaxKickpoints {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Maximale Kickpunkte erreicht",
			fmt.Sprintf("Dieses Mitglied hat bereits %d/%d Kickpunkte und kann keine weiteren erhalten.", totalKickpoints, settings.MaxKickpoints),
			messages.ColorRed,
		))
		return
	}

	if err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: util.BuildCustomID(i.ApplicationCommandData().Name, i.Interaction.Member.User.ID, ""),
			Title:    "Kickpunkt hinzufügen",
			Components: components.GenModalComponents(
				components.KickpointReason(settingName.DisplayStringShort()),
				components.KickpointDate(""),
				components.KickpointAmount(kickpointAmount),
				components.Tag("Spieler Tag", memberTag, components.PlayerTagID),
				components.Tag("Clan Tag", clanTag, components.ClanTagID),
			),
		},
	}); err != nil {
		log.Printf("Error while responding to CreateKickpoint: %v", err)
	}
}

func (h *KickpointHandler) CreateKickpointModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	if len(data.Components) != 5 {
		messages.SendEmbed(s, i, messages.NewEmbed("Fehler", "Es wurden nicht alle Felder ausgefüllt.", messages.ColorRed))
		return
	}

	clanTag := util.ParseStringModalInput(data.Components[4])
	if err := h.authMiddleware.NewClanHandler(clanTag, types.AuthRoleCoLeader)(s, i); err != nil {
		return
	}

	date, err := util.ParseDateInput(data.Components[1])
	if err != nil {
		messages.SendInvalidInputError(s, i, "Das eingegebene Datum ist ungültig. Es muss im DisplayString `DD.MM.YYYY` angegeben werden.")
		return
	}

	if date.After(time.Now()) {
		messages.SendInvalidInputError(s, i, "Das eingegebene Datum liegt in der Zukunft.")
		return
	}

	settings, err := h.clanSettings.ClanSettingsPreload(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	if minDate := util.KickpointMinDate(settings.KickpointsExpireAfterDays); minDate.After(date) {
		messages.SendInvalidInputError(s, i, fmt.Sprintf("Es können keine Kickpunkte vor %s vergeben werden, da diese schon abgelaufen wären.", util.FormatDate(minDate)))
		return
	}

	amount := util.ParseIntModalInput(data.Components[2])
	if amount < validation.MinKickpointAmount || amount > validation.MaxKickpointAmount {
		messages.SendInvalidInputError(s, i, fmt.Sprintf(
			"Die Anzahl der Kickpunkte muss zwischen %d und %d liegen.", validation.MinKickpointAmount, validation.MaxKickpointAmount),
		)
		return
	}

	playerTag := util.ParseStringModalInput(data.Components[3])

	_, userID, _ := util.ParseCustomID(data.CustomID)
	kickpoint := &models.Kickpoint{
		Description:        util.ParseStringModalInput(data.Components[0]),
		Date:               date,
		Amount:             amount,
		PlayerTag:          playerTag,
		ClanTag:            clanTag,
		CreatedByDiscordID: userID,
		UpdatedByDiscordID: userID,
	}

	playerName, err := h.players.NameByTag(kickpoint.PlayerTag)
	if err != nil {
		messages.SendMemberNotFound(s, i, kickpoint.PlayerTag, kickpoint.ClanTag)
		return
	}

	if err = h.kickpoints.CreateKickpoint(kickpoint); err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Datenbankfehler",
			"Beim Speichern der Daten in der Datenbank ist ein Fehler aufgetreten. Dies liegt wahrscheinlich an der Eingabe ungültiger Daten.",
			messages.ColorRed,
		))
		return
	}

	messages.SendEmbed(s, i, messages.NewFieldEmbed(
		fmt.Sprintf("Kickpunkt #%d erstellt", kickpoint.ID),
		"Der Kickpunkt wurde erstellt und gespeichert!",
		messages.ColorGreen,
		append([]*discordgo.MessageEmbedField{
			{Name: "Mitglied", Value: fmt.Sprintf("%s in %s", playerName, settings.Clan.Name)}},
			messages.DetailedKickpointFields(kickpoint, settings.KickpointsExpireAfterDays)...,
		),
	))

	totalKickpoints, err := h.kickpoints.ActiveMemberKickpointsSum(kickpoint.PlayerTag, settings)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if err != nil {
		messages.SendChannelWarning(s, i.ChannelID, fmt.Sprintf("Bei der Überprüfung ob %s die maximale Anzahl an Kickpunkten erreicht hat, ist ein Fehler aufgetreten. Bitte überprüfe dies manuell.", playerName))
		return
	}

	if totalKickpoints >= settings.MaxKickpoints {
		messages.SendChannelWarning(s, i.ChannelID, fmt.Sprintf("%s hat die maximale Anzahl an Kickpunkten erreicht.", playerName))
	}
}

func (h *KickpointHandler) EditKickpoint(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	id := util.UintOptionByName(IDOptionName, opts)
	if id == nil {
		messages.SendInvalidInputError(s, i, "Du musst eine gültige Kickpunkt ID angeben.")
		return
	}

	kickpoint, err := h.kickpoints.KickpointByID(*id)
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

	if err = h.authMiddleware.NewClanHandler(kickpoint.ClanTag, types.AuthRoleCoLeader)(s, i); err != nil {
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: util.BuildCustomID(i.ApplicationCommandData().Name, i.Interaction.Member.User.ID, strconv.Itoa(int(*id))),
			Title:    "Kickpunkt bearbeiten",
			Components: components.GenModalComponents(
				components.KickpointReason(kickpoint.Description),
				components.KickpointDate(util.FormatDate(kickpoint.Date)),
				components.KickpointAmount(kickpoint.Amount),
			),
		},
	})

	if err != nil {
		log.Printf("Error while handling EditKickpoint: %v", err)
	}
}

func (h *KickpointHandler) EditKickpointModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	if len(data.Components) != 3 {
		messages.SendEmbed(s, i, messages.NewEmbed("Fehler", "Es wurden nicht alle Felder ausgefüllt.", messages.ColorRed))
		return
	}

	date, err := util.ParseDateInput(data.Components[1])
	if err != nil {
		messages.SendInvalidInputError(s, i, "Das eingegebene Datum ist ungültig. Es muss im DisplayString `DD.MM.YYYY` angegeben werden.")
		return
	}

	amount := util.ParseIntModalInput(data.Components[2])
	if amount < validation.MinKickpointAmount || amount > validation.MaxKickpointAmount {
		messages.SendInvalidInputError(s, i, fmt.Sprintf(
			"Die Anzahl der Kickpunkte muss zwischen %d und %d liegen.", validation.MinKickpointAmount, validation.MaxKickpointAmount),
		)
		return
	}

	_, userID, idStr := util.ParseCustomID(data.CustomID)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		messages.SendInvalidInputError(s, i, "Es wurde eine ungültige ID angegeben.")
		return
	}

	prevKickpoint, err := h.kickpoints.KickpointByID(uint(id))
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

	if err = h.authMiddleware.NewClanHandler(prevKickpoint.ClanTag, types.AuthRoleCoLeader)(s, i); err != nil {
		return
	}

	settings, err := h.clanSettings.ClanSettings(prevKickpoint.ClanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, prevKickpoint.ClanTag)
		return
	}

	updatedKickpoint := &models.Kickpoint{
		ID:                 prevKickpoint.ID,
		Description:        util.ParseStringModalInput(data.Components[0]),
		Date:               date,
		Amount:             amount,
		UpdatedByDiscordID: userID,
	}

	updatedKickpoint, err = h.kickpoints.UpdateKickpoint(updatedKickpoint)
	if err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Datenbankfehler",
			"Beim Speichern der Daten in der Datenbank ist ein Fehler aufgetreten. Dies liegt wahrscheinlich an der Eingabe ungültiger Daten.",
			messages.ColorRed,
		))
		return
	}

	messages.SendEmbed(s, i, messages.NewFieldEmbed(
		fmt.Sprintf("Kickpunkt #%d bearbeitet", updatedKickpoint.ID),
		"Der Kickpunkt wurde bearbeitet und gespeichert!",
		messages.ColorGreen,
		append([]*discordgo.MessageEmbedField{
			{Name: "Mitglied", Value: fmt.Sprintf("%s in %s", updatedKickpoint.Player.Name, updatedKickpoint.Clan.Name)}},
			messages.DetailedKickpointFields(updatedKickpoint, settings.KickpointsExpireAfterDays)...,
		),
	))
}

func (h *KickpointHandler) DeleteKickpoint(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	id := util.UintOptionByName(IDOptionName, opts)
	if id == nil {
		messages.SendInvalidInputError(s, i, "Du musst eine gültige Kickpunkt ID angeben.")
		return
	}

	kickpoint, err := h.kickpoints.KickpointByID(*id)
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

	if err = h.authMiddleware.NewClanHandler(kickpoint.ClanTag, types.AuthRoleCoLeader)(s, i); err != nil {
		return
	}

	if err = h.kickpoints.DeleteKickpoint(*id); err != nil {
		messages.SendUnknownError(s, i)
		return
	}

	messages.SendEmbed(s, i, messages.NewFieldEmbed(
		fmt.Sprintf("Kickpunkt #%d gelöscht", kickpoint.ID),
		fmt.Sprintf("Der Kickpunkt von %s in %s wurde gelöscht!", kickpoint.Player.Name, kickpoint.Clan.Name),
		messages.ColorGreen,
		messages.DetailedKickpointFields(kickpoint, 0),
	))
}

func (h *KickpointHandler) NewKickpointLockHandler(lock bool) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		opts := i.ApplicationCommandData().Options
		clanTag := util.StringOptionByName(ClanTagOptionName, opts)
		memberTag := util.StringOptionByName(MemberTagOptionName, opts)
		if clanTag == "" || memberTag == "" {
			messages.SendInvalidInputError(s, i, "Du musst einen Clan und ein Mitglied angeben.")
			return
		}

		if err := h.authMiddleware.NewClanHandler(clanTag, types.AuthRoleCoLeader)(s, i); err != nil {
			return
		}

		if err := h.memberStates.UpdateKickpointLockStatus(memberTag, clanTag, lock); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				messages.SendEmbed(s, i, messages.NewEmbed(
					"Ungültiger Spieler Tag",
					"Es wurde kein Mitglied mit diesem Spieler Tag gefunden.",
					messages.ColorRed,
				))
				return
			}

			messages.SendUnknownError(s, i)
			return
		}

		title := "Mitglied angemeldet"
		desc := "Das Mitglied kann ab sofort wieder Kickpunkte erhalten."
		if lock {
			title = "Mitglied abgemeldet"
			desc = "Das Mitglied kann ab sofort keine Kickpunkte mehr erhalten."
		}

		messages.SendEmbed(s, i, messages.NewEmbed(
			title,
			desc,
			messages.ColorGreen,
		))
	}
}

func (h *KickpointHandler) HandleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	for _, opt := range opts {
		if !opt.Focused {
			continue
		}

		switch opt.Name {
		case ClanTagOptionName:
			autocompleteClans(h.clans, opt.StringValue())(s, i)
		case MemberTagOptionName:
			autocompleteMembers(h.players, opt.StringValue(), util.StringOptionByName(ClanTagOptionName, opts))(s, i)
		}
	}
}
