package commands

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/handlers"
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
		repos.NewGuildsRepo(db),
		repos.NewUsersRepo(db),
		repos.NewClanSettingsRepo(db),
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
		},
	}, {
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
				optionPlayerTag("Member, dessen Kickpunkte angezeigt werden sollen."),
			},
		},
	}, {
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
		},
	}, {
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
						{Name: "Maximale Kickpunkte", Value: models.ClanSettingsMaxKickpoints},
						{Name: "Gültigkeitsdauer in Tagen", Value: models.ClanSettingsExpireAfterDays},
						{Name: "Minimum Season Wins", Value: models.ClanSettingsMinSeasonWins},
						{Name: "Kickpunkte: Season Wins", Value: models.ClanSettingsSeasonWins},
						{Name: "Kickpunkte: CW nicht angegriffen", Value: models.ClanSettingsCWMissed},
						{Name: "Kickpunkte: CW 0 Sterne", Value: models.ClanSettingsCWFail},
						{Name: "Kickpunte: CKL nicht angegriffen", Value: models.ClanSettingsCWLMissed},
						{Name: "Kickpunte: CKL 0 Sterne", Value: models.ClanSettingsCWLZero},
						{Name: "Kickpunte: CKL 1 Stern", Value: models.ClanSettingsCWLOne},
						{Name: "Kickpunte: Raid nicht angegriffen", Value: models.ClanSettingsRaidMissed},
						{Name: "Kickpunte: Raid Fail", Value: models.ClanSettingsRaidFail},
						{Name: "Kickpunte: Clan Spiele nicht gemacht", Value: models.ClanSettingsClanGames}},
				}, {
					Name:        "amount",
					Description: "Wert, auf den die Einstellung geändert werden soll.",
					Type:        discordgo.ApplicationCommandOptionInteger,
					Required:    true,
					MinValue:    util.OptionalFloat(0),
					MaxValue:    100,
				}},
		},
	}, {
		Handler: types.InteractionHandler{
			Main: handler.KickpointHelp,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kphelp",
			Description:  "Erklärung des Kickpunkte Systems vom Bot sowie den wichtigsten Befehlen.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
		},
	}, {
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
				optionPlayerTag("Mitglied, welches einen Kickpunkt erhält."),
				{
					Name:        "reason",
					Description: "Grund des Kickpunktes (wird für die Anzahl der Kickpunkte benötigt).",
					Type:        discordgo.ApplicationCommandOptionString,
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: "Season Wins", Value: models.ClanSettingsSeasonWins},
						{Name: "CW nicht angegriffen", Value: models.ClanSettingsCWMissed},
						{Name: "CW 0 Sterne", Value: models.ClanSettingsCWFail},
						{Name: "CKL nicht angegriffen", Value: models.ClanSettingsCWLMissed},
						{Name: "CKL 0 Sterne", Value: models.ClanSettingsCWLZero},
						{Name: "CKL 1 Stern", Value: models.ClanSettingsCWLOne},
						{Name: "Raid nicht angegriffen", Value: models.ClanSettingsRaidMissed},
						{Name: "Raid Fail", Value: models.ClanSettingsRaidFail},
						{Name: "Clan Spiele nicht gemacht", Value: models.ClanSettingsClanGames},
						{Name: "Anderer Grund", Value: models.ClanSettingsOther}},
				}},
		},
	}, {
		Handler: types.InteractionHandler{Main: handler.EditKickpoint, ModalSubmit: handler.EditKickpointModalSubmit},
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
		},
	}, {
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
		},
	}}
}
