package commands

import (
	"log"
	"slices"

	"github.com/aaantiii/goclash"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/messages"
	"bot/commands/util"
	"bot/env"
	"bot/types"
)

func createInteractions(db *gorm.DB, clashClient *goclash.Client) types.Commands[types.InteractionHandler] {
	interactions := []types.Commands[types.InteractionHandler]{
		kickpointInteractionCommands(db),
		playerInteractionCommands(db, clashClient),
		memberInteractionCommands(db, clashClient),
		adminInteractionCommands(db),
		clanInteractionCommands(db, clashClient),
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
		if i.GuildID == "" {
			if i.User != nil {
				log.Printf("Aborted interaction called by %s in DMs (not supported).", i.User.Username)
			}
			sendDMNotSupported(i)
			return
		}
		defer handleRecovery()

		if env.MODE.Value() == "TEST" || env.MODE.Value() == "DEBUG" {
			if !slices.Contains(i.Member.Roles, "1192095424094408724") {
				messages.SendEmbedResponse(i, messages.NewEmbed(
					"Entwicklung",
					"Dies ist der Test Bot. Bitte verwende den richtigen Bot.",
					messages.ColorYellow,
				))
				return
			}
		}

		switch i.Type {
		case discordgo.InteractionApplicationCommandAutocomplete:
			if command, ok := commands[i.ApplicationCommandData().Name]; ok {
				if command.Handler.Autocomplete == nil {
					return
				}
				command.Handler.Autocomplete(s, i)
				return
			}

		case discordgo.InteractionApplicationCommand:
			if command, ok := commands[i.ApplicationCommandData().Name]; ok {
				command.Handler.Main(s, i)
				return
			}

		case discordgo.InteractionModalSubmit:
			commandName, _, _ := util.ParseCustomID(i.ModalSubmitData().CustomID)
			if command, ok := commands[commandName]; ok {
				if command.Handler.ModalSubmit == nil {
					log.Printf("Tried to run modal submit handler for command '%s', but it is nil.", commandName)
					return
				}
				command.Handler.ModalSubmit(s, i)
				log.Printf("Ran modal submit handler for command '%s', called by %s.", commandName, i.Interaction.Member.User.Username)
				return
			}
		}
		sendCommandNotFound(i)
	}
}

func handleRecovery() {
	if err := recover(); err != nil {
		log.Printf("Interaction panicked: %v", err)
	}
}
