package messages

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func SendAutoCompletion(s *discordgo.Session, i *discordgo.InteractionCreate, choices []*discordgo.ApplicationCommandOptionChoice) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{Choices: choices},
	}); err != nil {
		log.Printf("Error auto completing interaction: %v", err)
	}
}

func SendEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, embed *discordgo.MessageEmbed) {
	embed.Footer = newDefaultFooter(s)
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{embed}},
	}); err != nil {
		log.Printf("Error responding to interaction: %v", err)
	}
}

func SendEmptyResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: ".",
		},
	}); err != nil {
		log.Printf("Error sending empty interaction response: %v", err)
		return
	}

	if err := s.InteractionResponseDelete(i.Interaction); err != nil {
		log.Printf("Error deleting empty interaction response: %v", err)
	}
}
