package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var timeNil = time.Time{}

func ParseStringModalInput(c discordgo.MessageComponent) string {
	input, err := parseTextInputComponent(c)
	if err != nil {
		return ""
	}

	return input.Value
}

func ParseIntModalInput(c discordgo.MessageComponent) int {
	input, err := parseTextInputComponent(c)
	if err != nil {
		return 0
	}

	if v, err := strconv.Atoi(input.Value); err == nil {
		return v
	}

	return 0
}

func ParseDateInput(c discordgo.MessageComponent) (time.Time, error) {
	input, err := parseTextInputComponent(c)
	if err != nil {
		return timeNil, err
	}

	return ParseDateString(input.Value)
}

func parseTextInputComponent(c discordgo.MessageComponent) (*discordgo.TextInput, error) {
	input, ok := c.(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
	if !ok {
		return nil, fmt.Errorf("could not parse modal input")
	}

	return input, nil
}
