package commands

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/handlers"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/commands/validation"
	"bot/types"
)

func penaltyInteractionCommands(db *gorm.DB) types.Commands[types.InteractionHandler] {
	penaltiesRepo := repos.NewPenaltiesRepo(db)
	clansRepo := repos.NewLostClansRepo(db)
	playersRepo := repos.NewPlayersRepo(db)
	membersRepo := repos.NewMembersRepo(db)
	handler := handlers.NewPenaltyHandler(penaltiesRepo, clansRepo, playersRepo, membersRepo)

	return types.Commands[types.InteractionHandler]{{
		Handler: types.InteractionHandler{
			Main:                handler.ClanPenalties,
			AutocompleteHandler: handler.ClansAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "clankicks",
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
			Main:                handler.MemberPenalties,
			AutocompleteHandler: handler.MembersAndClansAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "memberkicks",
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
			Main:                handler.CreatePenaltyModal,
			ModalSubmitHandler:  handler.CreatePenaltyModalSubmit,
			AutocompleteHandler: handler.MembersAndClansAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "addkick",
			Description:  "Neuen Kickpunkt hinzufügen",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{{
				Name:         "clan",
				Description:  "Clan in dem dieses Mitglied ist",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
				MinLength:    util.OptionalInt(validation.TagMinLength),
				MaxLength:    validation.TagMaxLength,
			}, {
				Name:         "player",
				Description:  "Mitglied, dessen Kickpunkte angezeigt werden sollen",
				Type:         discordgo.ApplicationCommandOptionString,
				Required:     true,
				Autocomplete: true,
				MinLength:    util.OptionalInt(validation.TagMinLength),
				MaxLength:    validation.TagMaxLength,
			}},
		},
	}, {
		Handler: types.InteractionHandler{Main: handler.EditPenalty},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Options: []*discordgo.ApplicationCommandOption{{
				Name:        "id",
				Description: "ID des Kickpunktes, der verändert werden soll.",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    true,
				MinValue:    util.OptionalFloat(1),
			}},
			Name:         "editkick",
			Description:  "Bestehenden Kickpunkt verändern",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
		},
	}, {
		Handler: types.InteractionHandler{Main: handler.DeletePenalty},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "removekick",
			Description:  "Bestehenden Kickpunkt verändern",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
		},
	}}
}
