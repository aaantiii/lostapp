package messages

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"

	"bot/commands/util"
)

func SendAutoCompletion(i *discordgo.InteractionCreate, choices []*discordgo.ApplicationCommandOptionChoice) {
	if err := util.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{Choices: choices},
	}); err != nil {
		log.Printf("Error auto completing interaction: %v", err)
	}
}

func SendEmbedResponse(i *discordgo.InteractionCreate, embed *discordgo.MessageEmbed) {
	if err := util.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{embed}},
	}); err != nil {
		log.Printf("Error responding to interaction: %v", err)
	}
}

func SendChannelEmbed(channelID string, embed *discordgo.MessageEmbed) {
	if _, err := util.Session.ChannelMessageSendEmbed(channelID, embed); err != nil {
		log.Printf("Error sending embed: %v", err)
	}
}

func SendMessageResponse(i *discordgo.InteractionCreate, title, message string) {
	if err := util.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: fmt.Sprintf("## %s\n%s", title, message)},
	}); err != nil {
		log.Printf("Error responding to interaction: %v", err)
	}
}

func SendEmptyResponse(i *discordgo.InteractionCreate) {
	if err := util.Session.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: ".",
		},
	}); err != nil {
		log.Printf("Error sending empty interaction response: %v", err)
		return
	}

	if err := util.Session.InteractionResponseDelete(i.Interaction); err != nil {
		log.Printf("Error deleting empty interaction response: %v", err)
	}
}
