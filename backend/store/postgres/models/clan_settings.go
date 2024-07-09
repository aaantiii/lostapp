package models

import (
	"time"
)

type ClanSettings struct {
	ClanTag                   string    `gorm:"size:12;primaryKey;not null" json:"clanTag"`
	MaxKickpoints             int       `gorm:"not null;default:6" json:"maxKickpoints"`
	MinSeasonWins             int       `gorm:"not null;default:80" json:"minSeasonWins"`
	KickpointsExpireAfterDays int       `gorm:"not null;default:45" json:"kickpointsExpireAfterDays"`
	UpdatedAt                 time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedByDiscordID        *string   `gorm:"size:18" json:"-"`

	Clan          *Clan `gorm:"foreignKey:Tag;references:ClanTag;onUpdate:CASCADE;onDelete:CASCADE" json:"clan,omitempty"`
	UpdatedByUser *User `gorm:"foreignKey:DiscordID;references:UpdatedByDiscordID" json:"updatedByUser,omitempty"`
}
