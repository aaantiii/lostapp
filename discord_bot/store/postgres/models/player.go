package models

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// Player links a Clash of Clans player tag with a Discord ID.
type Player struct {
	Name      string
	CocTag    string
	DiscordID *string

	UpdatedAt time.Time `gorm:"column:last_updated"`

	Members Members `gorm:"foreignKey:PlayerTag;references:CocTag"`
}

func (*Player) TableName() string {
	return "player"
}

type Players []*Player

func (players Players) Choices() []*discordgo.ApplicationCommandOptionChoice {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(players))
	for i, player := range players {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  player.Name,
			Value: player.CocTag,
		}
	}

	return choices
}
