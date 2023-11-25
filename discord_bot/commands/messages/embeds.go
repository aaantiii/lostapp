package messages

import (
	"github.com/bwmarrin/discordgo"
)

const (
	ColorAqua  = 0x64abe1
	ColorGreen = 0x13E708
	ColorRed   = 0xe30000
)

func NewEmbed(title, desc string, color int) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       title,
		Description: desc,
		Color:       color,
		Type:        discordgo.EmbedTypeRich,
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
	}
}
