package commands

import (
	"github.com/bwmarrin/discordgo"

	"bot/commands/handlers"
	"bot/commands/util"
	"bot/commands/validation"
)

func optionClanTag(desc string) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:         handlers.ClanTagOptionName,
		Description:  desc,
		Type:         discordgo.ApplicationCommandOptionString,
		Required:     true,
		Autocomplete: true,
		MinLength:    util.IntPtr(validation.TagMinLength),
		MaxLength:    validation.TagMaxLength,
	}
}

func optionMemberTag(desc string) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:         handlers.MemberTagOptionName,
		Description:  desc,
		Type:         discordgo.ApplicationCommandOptionString,
		Required:     true,
		Autocomplete: true,
		MinLength:    util.IntPtr(validation.TagMinLength),
		MaxLength:    validation.TagMaxLength,
	}
}

func optionPlayerTag(desc string) *discordgo.ApplicationCommandOption {
	return &discordgo.ApplicationCommandOption{
		Name:         handlers.PlayerTagOptionName,
		Description:  desc,
		Type:         discordgo.ApplicationCommandOptionString,
		Required:     true,
		Autocomplete: true,
		MinLength:    util.IntPtr(validation.TagMinLength),
		MaxLength:    validation.TagMaxLength,
	}
}
