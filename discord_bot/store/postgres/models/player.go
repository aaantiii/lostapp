package models

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Player links a Clash of Clans player tag with a Discord ID.
type Player struct {
	Name      string    `gorm:"not null"`
	CocTag    string    `gorm:"primaryKey;not null;size:12"`
	DiscordID string    `gorm:"size:18"`
	UpdatedAt time.Time `gorm:"column:last_updated;not null"`
	Members   Members   `gorm:"foreignKey:PlayerTag;references:CocTag"`
}

func (*Player) TableName() string {
	return "player"
}

type Players []*Player

func (players Players) Choices() []*discordgo.ApplicationCommandOptionChoice {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(players))
	for i, player := range players {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  fmt.Sprintf("%s | %s", player.Name, player.CocTag),
			Value: player.CocTag,
		}
	}

	return choices
}
