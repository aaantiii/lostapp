package types

import "github.com/bwmarrin/discordgo"

type CommandHandler interface {
	InteractionHandler
}

type Command[T CommandHandler] struct {
	*discordgo.ApplicationCommand
	Handler T
}

type Commands[T CommandHandler] []*Command[T]

type InteractionHandler struct {
	Main         func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Autocomplete func(s *discordgo.Session, i *discordgo.InteractionCreate)
	ModalSubmit  func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func (commands Commands[T]) ApplicationCommands() []*discordgo.ApplicationCommand {
	appCmds := make([]*discordgo.ApplicationCommand, len(commands))
	for i, c := range commands {
		appCmds[i] = c.ApplicationCommand
	}

	return appCmds
}
