package commands

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"

	"bot/commands/util"
	"bot/types"
)

func interactionCommandMap(commands types.Commands[types.InteractionHandler]) map[string]*types.Command[types.InteractionHandler] {
	interactionsMap := make(map[string]*types.Command[types.InteractionHandler])
	for _, c := range commands {
		interactionsMap[c.Name] = c
	}

	return interactionsMap
}

func interactionHandler(interactions types.Commands[types.InteractionHandler]) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commands := interactionCommandMap(interactions)
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		now := time.Now()
		defer func() {
			log.Printf("Interaction called by %s took %s.", i.Member.User.Username, time.Now().Sub(now).Round(time.Millisecond))
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
			commandName, _ := util.ParseCustomID(i.ModalSubmitData().CustomID)
			if command, ok := commands[commandName]; ok {
				if command.Handler.ModalSubmit == nil {
					return
				}
				command.Handler.ModalSubmit(s, i)
				return
			}
		}
		commandNotFound(s, i)
	}
}
