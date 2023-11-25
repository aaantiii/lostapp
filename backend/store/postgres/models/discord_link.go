package models

import "time"

// DiscordLink links a Clash of Clans player tag with a Discord ID.
type DiscordLink struct {
	Name      string    `gorm:"not null"`
	CocTag    string    `gorm:"primaryKey;not null;size:12"`
	DiscordID string    `gorm:"size:18"`
	UpdatedAt time.Time `gorm:"column:last_updated;not null"`
}

func (*DiscordLink) TableName() string {
	return "player"
}
