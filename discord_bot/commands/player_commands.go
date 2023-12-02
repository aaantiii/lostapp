package commands

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/client"
	"bot/commands/handlers"
	"bot/commands/repos"
	"bot/commands/util"
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
			Description:  "Verifiziert einen Spieler.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        handlers.PlayerTagOptionName,
					Description: "Der Spieler-Tag des zu verifizierenden Spielers.",
					Required:    true,
					MinLength:   util.OptionalInt(4),
					MaxLength:   12,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "api_token",
					Description: "Der API Token des zu verifizierenden Spielers.",
					Required:    true,
					MinLength:   util.OptionalInt(8),
					MaxLength:   8,
				},
			},
		},
	}}
}
