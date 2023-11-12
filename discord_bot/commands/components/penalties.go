package components

import (
	"github.com/bwmarrin/discordgo"

	"bot/commands/validation"
	"bot/store/postgres/models"
)

const (
	PlayerTagID     = "player"
	ClanTagID       = "clan"
	PenaltyReasonID = "penalty_reason"
	PenaltyTypeID   = "penalty_type"
	PenaltyDateID   = "penalty_month"
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

func PenaltyReason(defaultValue string) discordgo.TextInput {
	return discordgo.TextInput{
		CustomID:    PenaltyReasonID,
		Label:       "Grund",
		Placeholder: "z.B. CWL Feb Tag 2 nicht angegriffen",
		Style:       discordgo.TextInputShort,
		Value:       defaultValue,
		Required:    true,
		MinLength:   10,
		MaxLength:   100,
	}
}

func PenaltyType(defaultValue models.PenaltyType) discordgo.TextInput {
	return discordgo.TextInput{
		CustomID:    PenaltyTypeID,
		Label:       "Typ",
		Placeholder: "K = Kickpunkt, V = Verwarnung",
		Style:       discordgo.TextInputShort,
		Value:       string(defaultValue),
		Required:    true,
		MinLength:   1,
		MaxLength:   1,
	}
}

func PenaltyDate(defaultValue string) discordgo.TextInput {
	return discordgo.TextInput{
		CustomID:    PenaltyDateID,
		Label:       "Datum (Monat und Jahr)",
		Placeholder: "z.B. 01-2023",
		Style:       discordgo.TextInputShort,
		Value:       defaultValue,
		Required:    true,
		MinLength:   6,
		MaxLength:   7,
	}
}
