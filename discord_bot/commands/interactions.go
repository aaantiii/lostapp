package commands

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/client"
	"bot/commands/util"
	"bot/types"
)

func createInteractions(db *gorm.DB, cocClient *client.CocClient) types.Commands[types.InteractionHandler] {
	interactions := []types.Commands[types.InteractionHandler]{
		kickpointInteractionCommands(db),
		playerInteractionCommands(db, cocClient),
		memberInteractionCommands(db),
	}

	var flat types.Commands[types.InteractionHandler]
	for _, i := range interactions {
		flat = append(flat, i...)
	}

	return flat
}

func interactionCommandMap(commands types.Commands[types.InteractionHandler]) map[string]*types.Command[types.InteractionHandler] {
	interactionsMap := make(map[string]*types.Command[types.InteractionHandler])
	for _, c := range commands {
		interactionsMap[c.Name] = c
	}

	return interactionsMap
}

// InteractionHandler is the handler for all interactions. It handles every interaction in its own goroutine.
func interactionHandler(interactions types.Commands[types.InteractionHandler]) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commands := interactionCommandMap(interactions)
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		go func() {
			if i.GuildID == "" {
				if i.User != nil {
					log.Printf("Interaction called by %s in DMs.", i.User.Username)
				}
				sendDMNotSupported(s, i)
				return
			}

			now := time.Now()
			defer func() {
				if err := recover(); err != nil {
					log.Printf("Interaction called by %s panicked: %v", i.Member.User.Username, err)
				} else {
					log.Printf("Interaction called by %s took %s.", i.Member.User.Username, time.Now().Sub(now).Round(time.Millisecond))
				}
			}()

			switch i.Type {
			case discordgo.InteractionApplicationCommand:
				if command, ok := commands[i.ApplicationCommandData().Name]; ok {
					command.Handler.Main(s, i)
					return
				}

			case discordgo.InteractionApplicationCommandAutocomplete:
				if command, ok := commands[i.ApplicationCommandData().Name]; ok {
					if command.Handler.Autocomplete == nil {
						return
					}
					command.Handler.Autocomplete(s, i)
					return
				}

			case discordgo.InteractionModalSubmit:
				commandName, _, _ := util.ParseCustomID(i.ModalSubmitData().CustomID)
				if command, ok := commands[commandName]; ok {
					if command.Handler.ModalSubmit == nil {
						return
					}
					command.Handler.ModalSubmit(s, i)
					return
				}
			}
			sendCommandNotFound(s, i)
		}()
	}
}
