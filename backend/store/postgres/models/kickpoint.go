package models

import "time"

type Kickpoint struct {
	ID                     uint            `gorm:"primaryKey;autoIncrement;not null"`
	PlayerTag              string          `gorm:"size:10;not null"`
	ClanTag                string          `gorm:"size:10;not null"`
	Date                   time.Time       `gorm:"not null"`
	Amount                 int             `gorm:"not null"`
	Reason                 KickpointReason `gorm:"not null;size:60"`
	Description            string          `gorm:"size:255"`
	AddedByDiscordID       string          `gorm:"size:18;not null"`
	LastUpdatedByDiscordID string          `gorm:"size:18;not null"`
}

type KickpointReason string

const (
	KickPointReasonSeasonWins   KickpointReason = "Saison Siege"
	KickPointReasonCWMissed     KickpointReason = "CK Angriff vergessen"
	KickPointReasonCWFail       KickpointReason = "CK Fail"
	KickPointReasonCWLMissed    KickpointReason = "CWL Angriff vergessen"
	KickPointReasonCWLZeroStars KickpointReason = "CWL Fail (0 Sterne)"
	KickPointReasonCWLOneStar   KickpointReason = "CWL Fail (1 Stern)"
	KickPointReasonRaidMissed   KickpointReason = "Raid vergessen"
	KickPointReasonRaidFail     KickpointReason = "Raid Fail"
	KickPointReasonClanGames    KickpointReason = "Clan Spiele"
)
