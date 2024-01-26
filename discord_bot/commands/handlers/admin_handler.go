package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bot/commands/messages"
	"bot/commands/middleware"
	"bot/commands/util"
)

type IAdminHandler interface {
	DeleteMessages(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type AdminHandler struct {
	auth middleware.AuthMiddleware
}

func NewAdminHandler(auth middleware.AuthMiddleware) IAdminHandler {
	return &AdminHandler{
		auth: auth,
	}
}

func (h *AdminHandler) DeleteMessages(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if err := h.auth.AuthorizeAdminInteraction(i); err != nil {
		return
	}

	opts := i.ApplicationCommandData().Options
	limit := util.UintOptionByName(LimitOptionName, opts)
	if len(opts) != 1 || limit == nil {
		messages.SendInvalidInputErr(i, "Bitte gib eine gültige Anzahl an Nachrichten an.")
		return
	}

	msgs, err := s.ChannelMessages(i.ChannelID, int(*limit), "", "", "")
	if err != nil {
		messages.SendUnknownErr(i)
		return
	}

	if len(msgs) == 0 {
		messages.SendErr(i, "Es gibt keine Nachrichten, die gelöscht werden können.")
		return
	}

	msgIDs := make([]string, len(msgs))
	for index, msg := range msgs {
		msgIDs[index] = msg.ID
	}

	if err = s.ChannelMessagesBulkDelete(i.ChannelID, msgIDs); err != nil {
		messages.SendUnknownErr(i)
		return
	}

	messages.SendEmbedResponse(i, messages.NewEmbed(
		"Nachrichten gelöscht",
		fmt.Sprintf("%d Nachrichten wurden gelöscht.", len(msgs)),
		messages.ColorGreen,
	))
}
