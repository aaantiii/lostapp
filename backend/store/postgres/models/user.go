package models

type User struct {
	DiscordID string `gorm:"size:18;primaryKey;not null" json:"discordId"`
	Name      string `gorm:"size:32;not null" json:"name"`
	AvatarURL string `gorm:"size:1024;not null" json:"avatarUrl"`
	IsAdmin   bool   `gorm:"default:false;not null" json:"-"`

	Players *[]Player `gorm:"foreignKey:DiscordID;references:DiscordID;onUpdate:CASCADE;onDelete:RESTRICT" json:"players,omitempty"`
}
