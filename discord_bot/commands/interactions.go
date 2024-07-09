package commands

import (
	"log/slog"
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
		defer handleRecovery()
		if i.GuildID == "" {
			if i.User != nil {
				slog.Info("Aborted interaction executed per DM.", slog.String("username", i.User.Username))
			}
			sendDMNotSupported(i)
			return
		}

		if env.MODE.Value() != "PROD" {
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
					slog.Error("Tried to run autocomplete handler, but it is nil.", slog.String("command", command.Name), slog.String("username", i.Member.User.Username))
					return
				}
				command.Handler.Autocomplete(s, i)
				return
			}

		case discordgo.InteractionApplicationCommand:
			if command, ok := commands[i.ApplicationCommandData().Name]; ok {
				command.Handler.Main(s, i)
				slog.Info("Interaction handler was executed.", slog.String("command", command.Name), slog.String("username", i.Member.User.Username))
				return
			}

		case discordgo.InteractionModalSubmit:
			commandName, _, _ := util.ParseCustomID(i.ModalSubmitData().CustomID)
			if command, ok := commands[commandName]; ok {
				if command.Handler.ModalSubmit == nil {
					slog.Error("Tried to run modal submit handler but it is nil.", slog.String("command", commandName), slog.String("username", i.Member.User.Username))
					return
				}
				command.Handler.ModalSubmit(s, i)
				slog.Info("Modal submit handler was executed.", slog.String("command", commandName), slog.String("username", i.Member.User.Username))
				return
			}
		}
		sendCommandNotFound(i)
	}
}

func handleRecovery() {
	if err := recover(); err != nil {
		slog.Error("Recovered from panic after error.", slog.Any("err", err), slog.String("func", "handleRecovery"))
	}
}
