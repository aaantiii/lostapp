package components

import "github.com/bwmarrin/discordgo"

func GenModalComponents(components ...discordgo.MessageComponent) []discordgo.MessageComponent {
	c := make([]discordgo.MessageComponent, len(components))
	for i, component := range components {
		c[i] = &discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{component},
		}
	}

	return c
}
