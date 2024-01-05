package util

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func StringOptionByName(name string, options []*discordgo.ApplicationCommandInteractionDataOption) string {
	for _, o := range options {
		if o.Name == name {
			return strings.Trim(o.StringValue(), " ")
		}
	}

	return ""
}

func IntOptionByName(name string, options []*discordgo.ApplicationCommandInteractionDataOption) *int {
	for _, o := range options {
		if o.Name == name {
			value := int(o.IntValue())
			return &value
		}
	}

	return nil
}

func UintOptionByName(name string, options []*discordgo.ApplicationCommandInteractionDataOption) *uint {
	for _, o := range options {
		if o.Name == name {
			value := uint(o.IntValue())
			return &value
		}
	}

	return nil
}
