package components

import (
	"github.com/bwmarrin/discordgo"
)

func GenModalComponents(components ...discordgo.MessageComponent) []discordgo.MessageComponent {
	c := make([]discordgo.MessageComponent, len(components))
	for i, component := range components {
		c[i] = &discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{component},
		}
	}

	return c
}

func autofillLabel(label string, defaultValue any) string {
	isAutofilled := false
	switch defaultValue.(type) {
	case string:
		if defaultValue != "" {
			isAutofilled = true
		}
	case int:
		if defaultValue != 0 {
			isAutofilled = true
		}
	}

	if isAutofilled {
		label += " (automatisch ausgefüllt)"
	}

	return label
}
