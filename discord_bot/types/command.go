package types

import "github.com/bwmarrin/discordgo"

type Command[T CommandHandler] struct {
	*discordgo.ApplicationCommand
	Handler T
}

type Commands[T CommandHandler] []*Command[T]

type CommandHandler interface {
	InteractionHandler | EventHandler
}

type InteractionHandler struct {
	Main         func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Autocomplete func(s *discordgo.Session, i *discordgo.InteractionCreate)
	ModalSubmit  func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type EventHandler = func(s *discordgo.Session, e *discordgo.Event)

func (commands Commands[T]) ApplicationCommands() []*discordgo.ApplicationCommand {
	acs := make([]*discordgo.ApplicationCommand, len(commands))
	for i, c := range commands {
		acs[i] = c.ApplicationCommand
	}

	return acs
}
