package models

import "time"

type Penalty struct {
	ID                 uint        `gorm:"primaryKey;autoIncrement;not null"`
	PlayerTag          string      `gorm:"size:12;not null"`
	ClanTag            string      `gorm:"size:12;not null"`
	Type               PenaltyType `gorm:"not null;check:type IN ('kickpoint', 'warning')"`
	Month              int         `gorm:"not null;check:month between 1 and 12"`
	Year               int         `gorm:"not null;check:year >= 2023"`
	Reason             string      `gorm:"size:100"`
	CreatedAt          time.Time
	CreatedByDiscordID string `gorm:"size:18;not null"`
	UpdatedAt          *time.Time
	UpdatedByDiscordID *string `gorm:"size:18"`
}

func (*Penalty) TableName() string {
	return "member_penalties"
}

type PenaltyType string

func (p PenaltyType) DisplayString() string {
	switch p {
	case PenaltyTypeKickpoint:
		return "Kickpunkt"
	case PenaltyTypeWarning:
		return "Verwarnung"
	default:
		return ""
	}
}

const (
	PenaltyTypeKickpoint PenaltyType = "kickpoint"
	PenaltyTypeWarning   PenaltyType = "warning"
	PenaltyTypeNone      PenaltyType = ""
)
