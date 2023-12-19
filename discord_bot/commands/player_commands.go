package commands

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/client"
	"bot/commands/handlers"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/commands/validation"
	"bot/types"
)

func playerInteractionCommands(db *gorm.DB, cocClient *client.CocClient) types.Commands[types.InteractionHandler] {
	handler := handlers.NewPlayerHandler(repos.NewPlayersRepo(db), cocClient)

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
					Name:        "api_token",
					Description: "API Token deines Clash of Clans Accounts.",
					Required:    true,
					MinLength:   util.OptionalInt(8),
					MaxLength:   8,
				},
			},
		},
	}}
}
