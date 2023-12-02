package util

import "github.com/bwmarrin/discordgo"

func StringOptionByName(name string, options []*discordgo.ApplicationCommandInteractionDataOption) string {
	for _, o := range options {
		if o.Name == name {
			return o.StringValue()
		}
	}

	return ""
}

func IntOptionByName(name string, options []*discordgo.ApplicationCommandInteractionDataOption) int {
	for _, o := range options {
		if o.Name == name {
			return int(o.IntValue())
		}
	}

	return 0
}

func UintOptionByName(name string, options []*discordgo.ApplicationCommandInteractionDataOption) uint {
	for _, o := range options {
		if o.Name == name {
			return uint(o.IntValue())
		}
	}

	return 0
}
