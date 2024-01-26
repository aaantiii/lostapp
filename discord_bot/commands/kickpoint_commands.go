package commands

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/handlers"
	"bot/commands/middleware"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/store/postgres/models"
	"bot/types"
)

func kickpointInteractionCommands(db *gorm.DB) types.Commands[types.InteractionHandler] {
	handler := handlers.NewKickpointHandler(
		repos.NewKickpointsRepo(db),
		repos.NewClansRepo(db),
		repos.NewPlayersRepo(db),
		repos.NewClanSettingsRepo(db),
		repos.NewMemberStatesRepo(db),
		middleware.NewAuthMiddleware(repos.NewGuildsRepo(db), repos.NewClansRepo(db), repos.NewUsersRepo(db)),
	)

	return types.Commands[types.InteractionHandler]{{
		Handler: types.InteractionHandler{
			Main:         handler.ClanKickpoints,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpclan",
			Description:  "Alle Kickpunkte eines Clans anzeigen.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, dessen Kickpunkte angezeigt werden sollen."),
			},
		}}, {
		Handler: types.InteractionHandler{
			Main:         handler.MemberKickpoints,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpmember",
			Description:  "Alle Kickpunkte eines Mitglieds anzeigen.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, aus dem das Mitglied stammt."),
				optionMemberTag("ClanMember, dessen Kickpunkte angezeigt werden sollen."),
			},
		}}, {
		Handler: types.InteractionHandler{
			Main:         handler.KickpointInfo,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpinfo",
			Description:  "Übersicht wie viele Kickpunkte verschiedene Regelbrüche geben.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, dessen Übersicht angezeigt werden soll."),
			},
		}}, {
		Handler: types.InteractionHandler{
			Main:         handler.KickpointConfig,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpconfig",
			Description:  "Kickpunkte Einstellungen eines Clans ändern.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, dessen Konfiguration geändert werden soll."),
				{
					Name:        "setting",
					Description: "Einstellung, die geändert werden soll.",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: models.KickpointSettingMaxKickpoints.DisplayString(), Value: models.KickpointSettingMaxKickpoints},
						{Name: models.KickpointSettingExpireAfterDays.DisplayString(), Value: models.KickpointSettingExpireAfterDays},
						{Name: models.KickpointSettingMinSeasonWins.DisplayString(), Value: models.KickpointSettingMinSeasonWins},
						{Name: models.KickpointSettingSeasonWins.DisplayString(), Value: models.KickpointSettingSeasonWins},
						{Name: models.KickpointSettingCWMissed.DisplayString(), Value: models.KickpointSettingCWMissed},
						{Name: models.KickpointSettingCWFail.DisplayString(), Value: models.KickpointSettingCWFail},
						{Name: models.KickpointSettingCWLMissed.DisplayString(), Value: models.KickpointSettingCWLMissed},
						{Name: models.KickpointSettingCWLZeroStars.DisplayString(), Value: models.KickpointSettingCWLZeroStars},
						{Name: models.KickpointSettingCWLOneStar.DisplayString(), Value: models.KickpointSettingCWLOneStar},
						{Name: models.KickpointSettingRaidMissed.DisplayString(), Value: models.KickpointSettingRaidMissed},
						{Name: models.KickpointSettingRaidFail.DisplayString(), Value: models.KickpointSettingRaidFail},
						{Name: models.KickpointSettingClanGames.DisplayString(), Value: models.KickpointSettingClanGames}},
				}, {
					Name:        "amount",
					Description: "Wert, auf den die Einstellung geändert werden soll.",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Required:    true,
					MinValue:    util.OptionalFloat(0),
					MaxValue:    100,
				}},
		}}, {
		Handler: types.InteractionHandler{Main: handler.KickpointHelp},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kphelp",
			Description:  "Erklärung des Kickpunkte Systems vom Bot sowie den wichtigsten Befehlen.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
		}}, {
		Handler: types.InteractionHandler{
			Main:         handler.CreateKickpointModal,
			ModalSubmit:  handler.CreateKickpointModalSubmit,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpadd",
			Description:  "Neuen Kickpunkt hinzufügen",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan aus dem das Mitglied stammt."),
				optionMemberTag("Mitglied, welches einen Kickpunkt erhält."),
				{
					Name:        handlers.SettingOptionName,
					Description: "Grund des Kickpunktes (wird für die Anzahl der Kickpunkte benötigt).",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: models.KickpointSettingSeasonWins.DisplayStringShort(), Value: models.KickpointSettingSeasonWins},
						{Name: models.KickpointSettingCWMissed.DisplayStringShort(), Value: models.KickpointSettingCWMissed},
						{Name: models.KickpointSettingCWFail.DisplayStringShort(), Value: models.KickpointSettingCWFail},
						{Name: models.KickpointSettingCWLMissed.DisplayStringShort(), Value: models.KickpointSettingCWLMissed},
						{Name: models.KickpointSettingCWLZeroStars.DisplayStringShort(), Value: models.KickpointSettingCWLZeroStars},
						{Name: models.KickpointSettingCWLOneStar.DisplayStringShort(), Value: models.KickpointSettingCWLOneStar},
						{Name: models.KickpointSettingRaidMissed.DisplayStringShort(), Value: models.KickpointSettingRaidMissed},
						{Name: models.KickpointSettingRaidFail.DisplayStringShort(), Value: models.KickpointSettingRaidFail},
						{Name: models.KickpointSettingClanGames.DisplayStringShort(), Value: models.KickpointSettingClanGames}},
				}},
		}}, {
		Handler: types.InteractionHandler{
			Main:        handler.EditKickpoint,
			ModalSubmit: handler.EditKickpointModalSubmit,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpedit",
			Description:  "Bestehenden Kickpunkt bearbeiten",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{{
				Name:        "id",
				Description: "ID des Kickpunktes, den du bearbeiten möchtest.",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
				MinValue:    util.OptionalFloat(1),
			}},
		}}, {
		Handler: types.InteractionHandler{Main: handler.DeleteKickpoint},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpremove",
			Description:  "Bestehenden Kickpunkt löschen (kann nicht rückgängig gemacht werden)",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{{
				Name:        "id",
				Description: "ID des Kickpunktes, den du löschen möchtest.",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
				MinValue:    util.OptionalFloat(1),
			}},
		}}, {
		Handler: types.InteractionHandler{
			Main:         handler.NewKickpointLockHandler(true),
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kplock",
			Description:  "Mitglied abmelden, sodass es keine Kickpunkte mehr erhalten kann.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan aus dem das Mitglied stammt."),
				optionMemberTag("Mitglied, welches abgemeldet werden soll."),
			},
		}}, {
		Handler: types.InteractionHandler{
			Main:         handler.NewKickpointLockHandler(false),
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpunlock",
			Description:  "Mitglied anmelden, sodass es wieder Kickpunkte erhalten kann.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan aus dem das Mitglied stammt."),
				optionMemberTag("Mitglied, welches wieder angemeldets werden soll."),
			},
		}},
	}
}
