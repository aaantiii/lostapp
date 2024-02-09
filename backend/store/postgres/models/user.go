package models

import (
	"gorm.io/gorm"
)

type User struct {
	DiscordID string `gorm:"primaryKey;not null" json:"discordId"`
	Name      string `gorm:"not null" json:"name"`
	AvatarURL string `json:"avatarUrl"`
	IsAdmin   bool   `gorm:"default:false;not null" json:"isAdmin"`
}

func (*User) BeforeCreate(tx *gorm.DB) error {
	tx.Omit("is_admin")
	return nil
}

func (*User) BeforeUpdate(tx *gorm.DB) error {
	tx.Omit("is_admin")
	return nil
}
