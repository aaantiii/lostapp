package messages

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bot/commands/util"
	"bot/env"
)

const (
	ColorAqua   = 0x64abe1
	ColorGreen  = 0x13E708
	ColorRed    = 0xe30000
	ColorYellow = 0xe3e300
)

func NewEmbed(title, desc string, color int) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       title,
		Description: desc,
		Color:       color,
		Type:        discordgo.EmbedTypeRich,
		Footer:      newDefaultFooter(),
	}
}

func NewFieldEmbed(title, desc string, color int, fields []*discordgo.MessageEmbedField) *discordgo.MessageEmbed {
	for _, field := range fields {
		field.Name += ":"
	}

	return &discordgo.MessageEmbed{
		Title:       title,
		Description: desc,
		Fields:      fields,
		Color:       color,
		Type:        discordgo.EmbedTypeRich,
		Footer:      newDefaultFooter(),
	}
}

func CreateAndEditEmbed(s *discordgo.Session, i *discordgo.InteractionCreate, title, description string, color int) error {
	embed := NewEmbed(
		title,
		description,
		color,
	)
	_, err := s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{embed},
	})
	return err
}

func newDefaultFooter() *discordgo.MessageEmbedFooter {
	return &discordgo.MessageEmbedFooter{
		Text: fmt.Sprintf("%s | %s | v%s", util.Session.State.User.Username, "Made by LOST Dev-Team", env.VERSION.Value()),
	}
}
