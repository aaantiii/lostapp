package messages

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"

	"bot/commands/util"
)

func SendAutoCompletion(i *discordgo.InteractionCreate, choices []*discordgo.ApplicationCommandOptionChoice) {
	if err := util.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{Choices: choices},
	}); err != nil {
		slog.Error("Error autocompleting interaction.", slog.Any("err", err))
	}
}

func SendEmbedResponse(i *discordgo.InteractionCreate, embed *discordgo.MessageEmbed) {
	if err := util.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{embed}},
	}); err != nil {
		slog.Error("Error responding to interaction.", slog.Any("err", err))
	}
}

func SendChannelEmbed(channelID string, embed *discordgo.MessageEmbed) {
	if _, err := util.Session.ChannelMessageSendEmbed(channelID, embed); err != nil {
		slog.Error("Error sending embed.", slog.Any("err", err))
	}
}

func SendMessageResponse(i *discordgo.InteractionCreate, title, message string) {
	if err := util.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: fmt.Sprintf("## %s\n%s", title, message)},
	}); err != nil {
		slog.Error("Error responding to interaction.", slog.Any("err", err))
	}
}

func SendEmptyResponse(i *discordgo.InteractionCreate) {
	if err := util.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: ".",
		},
	}); err != nil {
		slog.Error("Error responding to interaction.", slog.Any("err", err))
		return
	}

	if err := util.Session.InteractionResponseDelete(i.Interaction); err != nil {
		slog.Error("Error deleting interaction response.", slog.Any("err", err))
	}
}

func SendChannelMessage(channelID, message string) {
	if _, err := util.Session.ChannelMessageSend(channelID, message); err != nil {
		slog.Error("Error sending message.", slog.Any("err", err))
	}
}
