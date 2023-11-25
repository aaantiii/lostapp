package models

import "time"

type Kickpoint struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement;not null"`
	PlayerTag          string    `gorm:"size:12;not null"`
	ClanTag            string    `gorm:"size:12;not null"`
	Date               time.Time `gorm:"not null"`
	Amount             int       `gorm:"not null"`
	Description        string    `gorm:"size:100"`
	CreatedByDiscordID string    `gorm:"size:18;not null"`
	CreatedByUser      User      `gorm:"foreignKey:CreatedByDiscordID;references:DiscordID"`
	UpdatedByDiscordID *string   `gorm:"size:18"`
	UpdatedByUser      *User     `gorm:"foreignKey:UpdatedByDiscordID;references:DiscordID"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
