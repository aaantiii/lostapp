package models

import "github.com/amaanq/coc.go"

type Player struct {
	Tag           string `gorm:"primaryKey;not null;size:12" json:"tag"`
	DiscordID     string `gorm:"size:18;not null" json:"discordId"`
	Name          string `gorm:"size:32;not null" json:"name"`
	WarPreference string `gorm:"size:16;not null" json:"warPreference"`
	Trophies      int    `gorm:"not null" json:"trophies"`
	ExpLevel      int    `gorm:"not null" json:"expLevel"`
	TownHallLevel int    `gorm:"not null" json:"townHallLevel"`
	AttackWins    int    `gorm:"not null" json:"attackWins"`
	DefenseWins   int    `gorm:"not null" json:"defenseWins"`

	Clans []PlayerClan `gorm:"foreignKey:Tag;references:Tag;onUpdate:CASCADE;onDelete:CASCADE" json:"clans,omitempty"`
	Stats *PlayerStats `gorm:"foreignKey:Tag;references:Tag;onUpdate:CASCADE;onDelete:CASCADE" json:"stats,omitempty"`
}

func (*Player) TableName() string {
	return "lostapp_players_live"
}

func NewPlayer(player *coc.Player, discordID string) *Player {
	return &Player{
		DiscordID:     discordID,
		Name:          player.Name,
		Tag:           player.Tag,
		WarPreference: player.WarPreference,
		Trophies:      player.Trophies,
		ExpLevel:      player.ExpLevel,
		TownHallLevel: player.TownHallLevel,
		AttackWins:    player.AttackWins,
		DefenseWins:   player.DefenseWins,
	}
}
