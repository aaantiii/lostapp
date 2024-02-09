package models

import "time"

type ClanEventMember struct {
	ClanEventID uint      `gorm:"primaryKey;not null" json:"clanEventId"`
	PlayerTag   string    `gorm:"primaryKey;not null" json:"playerTag"`
	ClanTag     string    `gorm:"primaryKey;not null" json:"clanTag"`
	Timestamp   time.Time `gorm:"primaryKey;not null" json:"timestamp"`
	Name        string    `gorm:"not null" json:"name"`
	Value       int       `gorm:"not null" json:"value"`
}
