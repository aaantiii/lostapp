package commands

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/client"
	"bot/commands/messages"
	"bot/commands/util"
	"bot/env"
	"bot/types"
)

func createInteractions(db *gorm.DB, cocClient *client.CocClient) types.Commands[types.InteractionHandler] {
	interactions := []types.Commands[types.InteractionHandler]{
		kickpointInteractionCommands(db),
		playerInteractionCommands(db, cocClient),
		memberInteractionCommands(db, cocClient),
		adminInteractionCommands(db),
		clanInteractionCommands(db, cocClient),
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
					log.Printf("Aborted interaction called by %s in DMs (not supported).", i.User.Username)
				}
				sendDMNotSupported(s, i)
				return
			}
			defer handleRecovery(i.Interaction.Member.User.Username, time.Now())

			if env.MODE.Value() == "TEST" {
				if i.Member.User.ID != "289735223313301504" {
					messages.SendEmbed(s, i, messages.NewEmbed(
						"Wartungsarbeiten",
						"Der Bot wird gerade gewartet. Bitte versuche es später erneut.",
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
					return
				}
			}
			sendCommandNotFound(s, i)
		}()
	}
}

func handleRecovery(caller string, start time.Time) {
	took := time.Since(start).Round(time.Millisecond)
	if err := recover(); err != nil {
		log.Printf("Interaction panicked after %s: %v", took, err)
	} else {
		log.Printf("Interaction called by %s took %s.", caller, took)
	}
}
