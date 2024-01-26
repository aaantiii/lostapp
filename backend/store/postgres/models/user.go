package models

import (
	"gorm.io/gorm"
)

type User struct {
	DiscordID string `gorm:"primaryKey;not null"`
	Name      string `gorm:"not null"`
	AvatarURL string
	IsAdmin   bool `gorm:"default:false;not null"`
}

func (*User) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Omit("is_admin")
	return nil
}

func (*User) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Omit("is_admin")
	return nil
}
