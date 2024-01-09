package messages

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"

	"bot/env"
)

func SendAutoCompletion(s *discordgo.Session, i *discordgo.InteractionCreate, choices []*discordgo.ApplicationCommandOptionChoice) {
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{Choices: choices},
	}); err != nil {
		log.Print(err)
	}
}

func SendEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, embed *discordgo.MessageEmbed) {
	embed.Footer = &discordgo.MessageEmbedFooter{
		Text: fmt.Sprintf("%s | %s | v%s", s.State.User.Username, "Made by Anti", env.VERSION.Value()),
	}

	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Embeds: []*discordgo.MessageEmbed{embed}},
	}); err != nil {
		log.Print(err)
	}
}
