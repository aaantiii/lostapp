package commands

import (
	"github.com/aaantiii/goclash"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/handlers"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/commands/validation"
	"bot/types"
)

func playerInteractionCommands(db *gorm.DB, client *goclash.Client) types.Commands[types.InteractionHandler] {
	handler := handlers.NewPlayerHandler(repos.NewPlayersRepo(db), client)

	return types.Commands[types.InteractionHandler]{{
		Handler: types.InteractionHandler{
			Main: handler.VerifyPlayer,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "verify",
			Description:  "Verknüpfe deinen Discord Account mit deinem COC-Account.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        handlers.PlayerTagOptionName,
					Description: "Spieler-Tag deines Clash of Clans Accounts.",
					Required:    true,
					MinLength:   util.OptionalInt(validation.TagMinLength),
					MaxLength:   validation.TagMaxLength,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        handlers.ApiTokenOptionName,
					Description: "API Token deines Clash of Clans Accounts.",
					Required:    true,
					MinLength:   util.OptionalInt(8),
					MaxLength:   8,
				},
			},
		},
	}, {
		Handler: types.InteractionHandler{
			Main:         handler.PingPlayer,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "pingplayer",
			Description:  "Pingt einen Spieler auf Discord durch seinen Namen oder Spieler Tag.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:         discordgo.ApplicationCommandOptionString,
					Name:         handlers.PlayerTagOptionName,
					Description:  "Spieler-Tag des Spielers, der gepingt werden soll.",
					Required:     true,
					MinLength:    util.OptionalInt(validation.TagMinLength),
					MaxLength:    validation.TagMaxLength,
					Autocomplete: true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        handlers.MessageOptionName,
					Description: "Nachricht, die an den Spieler gesendet werden soll.",
					Required:    true,
					MinLength:   util.OptionalInt(1),
					MaxLength:   200,
				},
			},
		},
	}}
}
