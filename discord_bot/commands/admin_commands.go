package commands

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/handlers"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/types"
)

func adminInteractionCommands(db *gorm.DB) types.Commands[types.InteractionHandler] {
	handler := handlers.NewAdminHandler(repos.NewUsersRepo(db))
	return types.Commands[types.InteractionHandler]{
		{
			Handler: types.InteractionHandler{
				Main: handler.DeleteMessages,
			},
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:         "deletemessages",
				Description:  "Löscht eine bestimmte Anzahl an Nachrichten im aktuellen Channel.",
				Type:         discordgo.ChatApplicationCommand,
				DMPermission: util.OptionalBool(false),
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        "limit",
						Description: "Anzahl an Nachrichten, die gelöscht werden sollen.",
						Required:    true,
						MinValue:    util.OptionalFloat(1),
						MaxValue:    100,
					},
				},
			},
		},
	}
}
