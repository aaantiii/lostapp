package messages

import (
	"github.com/bwmarrin/discordgo"
)

const (
	ColorAqua  = 0x08f8fc
	ColorGreen = 0x13E708
	ColorRed   = 0xe30000
)

func NewMessageEmbed(title, desc string, color int) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       title,
		Description: desc,
		Color:       color,
	}
}

func NewMessageEmbedWithFields(title, desc string, color int, fields []*discordgo.MessageEmbedField) *discordgo.MessageEmbed {
	for _, field := range fields {
		field.Name += ":"
	}

	return &discordgo.MessageEmbed{
		Title:       title,
		Description: desc,
		Fields:      fields,
		Color:       color,
	}
}
