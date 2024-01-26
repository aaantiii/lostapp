package models

import "time"

type ClanEventMember struct {
	ClanEventID uint      `gorm:"primaryKey;not null"`
	PlayerTag   string    `gorm:"primaryKey;not null"`
	ClanTag     string    `gorm:"primaryKey;not null"`
	Timestamp   time.Time `gorm:"primaryKey;not null"`
	Name        string    `gorm:"not null"`
	Value       int       `gorm:"not null"`
}
