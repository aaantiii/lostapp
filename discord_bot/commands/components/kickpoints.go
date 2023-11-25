package components

import (
	"strconv"

	"github.com/bwmarrin/discordgo"

	"bot/commands/validation"
)

const (
	PlayerTagID       = "player"
	ClanTagID         = "clan"
	KickpointAmountID = "kickpoint_amount"
	KickpointReasonID = "kickpoint_reason"
	KickpointDateID   = "kickpoint_date"
)

func Tag(label, defaultValue, customId string) discordgo.TextInput {
	if defaultValue != "" {
		label += " (automatisch ausgefüllt)"
	}

	return discordgo.TextInput{
		CustomID:    customId,
		Label:       label,
		Placeholder: "z.B. #18742069",
		Style:       discordgo.TextInputShort,
		Value:       defaultValue,
		Required:    true,
		MinLength:   validation.TagMinLength,
		MaxLength:   validation.TagMaxLength,
	}
}

func KickpointAmount(defaultValue int) discordgo.TextInput {
	label := "Anzahl Kickpunkte"
	if defaultValue != 0 {
		label += " (automatisch ausgefüllt)"
	}

	return discordgo.TextInput{
		CustomID:    KickpointAmountID,
		Label:       label,
		Placeholder: "Zahl von 1 bis 10",
		Style:       discordgo.TextInputShort,
		Value:       strconv.Itoa(defaultValue),
		Required:    true,
		MinLength:   1,
		MaxLength:   1,
	}
}

func KickpointReason(defaultValue string) discordgo.TextInput {
	return discordgo.TextInput{
		CustomID:    KickpointReasonID,
		Label:       "Grund",
		Placeholder: "z.B. CWL Feb 2. Tag nicht angegriffen",
		Style:       discordgo.TextInputShort,
		Value:       defaultValue,
		Required:    true,
		MinLength:   10,
		MaxLength:   100,
	}
}

func KickpointDate(defaultValue string) discordgo.TextInput {
	return discordgo.TextInput{
		CustomID:    KickpointDateID,
		Label:       "Datum",
		Placeholder: "z.B. 31.01.2023",
		Style:       discordgo.TextInputShort,
		Value:       defaultValue,
		Required:    true,
		MinLength:   8,
		MaxLength:   10,
	}
}
