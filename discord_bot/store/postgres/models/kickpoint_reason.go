package models

type KickpointReason struct {
	Name    string `gorm:"primaryKey;not null"`
	ClanTag string `gorm:"primaryKey;not null"`
	Amount  int    `gorm:"not null"`
}
