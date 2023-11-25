package commands

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/handlers"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/commands/validation"
	"bot/store/postgres/models"
	"bot/types"
)

func kickpointInteractionCommands(db *gorm.DB) types.Commands[types.InteractionHandler] {
	kickpointsRepo := repos.NewKickpointsRepo(db)
	clansRepo := repos.NewClansRepo(db)
	clanSettingsRepo := repos.NewClanSettingsRepo(db)
	playersRepo := repos.NewPlayersRepo(db)
	guildsRepo := repos.NewGuildsRepo(db)
	usersRepo := repos.NewUsersRepo(db)

	handler := handlers.NewKickpointHandler(kickpointsRepo, clansRepo, playersRepo, guildsRepo, usersRepo, clanSettingsRepo)
	return types.Commands[types.InteractionHandler]{{
		Handler: types.InteractionHandler{
			Main:         handler.ClanKickpoints,
			Autocomplete: handler.ClansAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpclans",
			Description:  "Alle Kickpunkte eines Clans anzeigen",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{{
				Name:         "clan",
				Description:  "Clan, dessen Kickpunkte angezeigt werden sollen",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
				MinLength:    util.OptionalInt(validation.TagMinLength),
				MaxLength:    validation.TagMaxLength,
			}},
		},
	}, {
		Handler: types.InteractionHandler{
			Main:         handler.MemberKickpoints,
			Autocomplete: handler.MembersAndClansAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpmember",
			Description:  "Alle Kickpunkte eines Mitglieds anzeigen",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:         "clan",
					Description:  "Clan, aus dem das Mitglied stammt",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
					MinLength:    util.OptionalInt(validation.TagMinLength),
					MaxLength:    validation.TagMaxLength,
				}, {
					Name:         "player",
					Description:  "Member, dessen Kickpunkte angezeigt werden sollen",
					Type:         discordgo.ApplicationCommandOptionString,
					Required:     true,
					Autocomplete: true,
					MinLength:    util.OptionalInt(validation.TagMinLength),
					MaxLength:    validation.TagMaxLength,
				},
			},
		},
	}, {
		Handler: types.InteractionHandler{
			Main:         handler.KickpointInfo,
			Autocomplete: handler.ClansAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpinfo",
			Description:  "Übersicht wie viele Kickpunkte verschiedene Regelbrüche geben.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{{
				Name:         "clan",
				Description:  "Clan, dessen Übersicht angezeigt werden soll",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
				MinLength:    util.OptionalInt(validation.TagMinLength),
				MaxLength:    validation.TagMaxLength,
			}},
		},
	}, {
		Handler: types.InteractionHandler{
			Main:         handler.CreateKickpointModal,
			ModalSubmit:  handler.CreateKickpointModalSubmit,
			Autocomplete: handler.MembersAndClansAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "kpadd",
			Description:  "Neuen Kickpunkt hinzufügen",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{{
				Name:         "clan",
				Description:  "Clan in dem das Mitglied ist",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
				MinLength:    util.OptionalInt(validation.TagMinLength),
				MaxLength:    validation.TagMaxLength,
			}, {
				Name:         "player",
				Description:  "Mitglied das einen Kickpunkt erhält",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
				MinLength:    util.OptionalInt(validation.TagMinLength),
				MaxLength:    validation.TagMaxLength,
			}, {
				Name:        "reason",
				Description: "Grund des Kickpunktes (wird für die Anzahl der Kickpunkte benötigt)",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{Name: "Season Wins", Value: models.ClanSettingsSeasonWins},
					{Name: "CW nicht angegriffen", Value: models.ClanSettingsCWMissed},
					{Name: "CW 0 Sterne", Value: models.ClanSettingsCWFail},
					{Name: "CWL nicht angegriffen", Value: models.ClanSettingsCWLMissed},
					{Name: "CWL 0 Sterne", Value: models.ClanSettingsCWLZero},
					{Name: "CWL 1 Stern", Value: models.ClanSettingsCWLOne},
					{Name: "Raid nicht angegriffen", Value: models.ClanSettingsRaidMissed},
					{Name: "Raid Fail", Value: models.ClanSettingsRaidFail},
					{Name: "Clan Spiele nicht gemacht", Value: models.ClanSettingsClanGames},
					{Name: "Anderer Grund", Value: models.ClanSettingsOther}},
				MinLength: util.OptionalInt(10),
				MaxLength: 100,
			}},
		},
	}, {
		Handler: types.InteractionHandler{Main: handler.EditPenalty},
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
		Handler: types.InteractionHandler{Main: handler.DeletePenalty},
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
