package models

import "time"

type ClanEvent struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	ClanTag         string    `gorm:"not null" json:"clanTag"`
	StatName        string    `gorm:"not null" json:"statName"`
	StartsAt        time.Time `gorm:"not null" json:"startsAt"`
	EndsAt          time.Time `gorm:"not null" json:"endsAt"`
	ChannelID       string    `gorm:"not null" json:"channelID"`
	WinnerPlayerTag *string   `gorm:"default:null" json:"winnerPlayerTag,omitempty"`

	Clan *Clan `gorm:"foreignKey:Tag;references:ClanTag" json:"clan,omitempty"`
}
