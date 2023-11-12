package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

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

func ParseMonthYearInput(c discordgo.MessageComponent) (month, year int) {
	input, err := parseTextInputComponent(c)
	if err != nil {
		return 0, 0
	}

	date := strings.Split(input.Value, "-")
	if len(date) != 2 {
		return 0, 0
	}

	month, err = strconv.Atoi(date[0])
	if err != nil {
		return 0, 0
	}

	year, err = strconv.Atoi(date[1])
	if err != nil {
		return 0, 0
	}

	return month, year
}

func parseTextInputComponent(c discordgo.MessageComponent) (*discordgo.TextInput, error) {
	input, ok := c.(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)
	if !ok {
		return nil, fmt.Errorf("could not parse modal input")
	}

	return input, nil
}
