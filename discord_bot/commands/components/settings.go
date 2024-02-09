package components

import (
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"bot/commands/validation"
)

func ClanSettingMaxKickpoints(defaultValue int) *discordgo.TextInput {
	return &discordgo.TextInput{
		CustomID:    "max_kickpoints",
		Label:       "Maximale Kickpunkte",
		Placeholder: fmt.Sprintf("Zahl zwischen %d und %d", validation.MinTotalKickpoints, validation.MaxTotalKickpoints),
		Style:       discordgo.TextInputShort,
		Value:       strconv.Itoa(defaultValue),
		Required:    true,
		MinLength:   1,
		MaxLength:   2,
	}
}

func ClanSettingSeasonWins(defaultValue int) *discordgo.TextInput {
	return &discordgo.TextInput{
		CustomID:    "min_season_wins",
		Label:       "Season Wins (Minimum)",
		Placeholder: fmt.Sprintf("Zahl zwischen %d und %d", validation.MinSeasonWins, validation.MaxSeasonWins),
		Style:       discordgo.TextInputShort,
		Value:       strconv.Itoa(defaultValue),
		Required:    true,
		MinLength:   1,
		MaxLength:   3,
	}
}

func ClanSettingExpiration(defaultValue int) *discordgo.TextInput {
	return &discordgo.TextInput{
		CustomID:    "kickpoints_expire_after_days",
		Label:       "Kickpunkte Ablaufdatum",
		Placeholder: "Anzahl Tage",
		Style:       discordgo.TextInputShort,
		Value:       strconv.Itoa(defaultValue),
		Required:    true,
		MinLength:   2,
		MaxLength:   2,
	}
}
