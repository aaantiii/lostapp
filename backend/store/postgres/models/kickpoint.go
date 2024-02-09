package models

import "time"

type Kickpoint struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	PlayerTag          string    `gorm:"not null" json:"playerTag"`
	ClanTag            string    `gorm:"not null" json:"clanTag"`
	Date               time.Time `gorm:"not null" json:"date"`
	Amount             int       `gorm:"not null" json:"amount"`
	Description        string    `gorm:"not null" json:"description"`
	CreatedByDiscordID string    `gorm:"not null" json:"-"`
	UpdatedByDiscordID string    `json:"-"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`

	Member        *ClanMember `gorm:"foreignKey:PlayerTag,ClanTag;references:PlayerTag,ClanTag" json:"member,omitempty"`
	Clan          *Clan       `gorm:"foreignKey:Tag;references:ClanTag" json:"clan,omitempty"`
	Player        *Player     `gorm:"foreignKey:CocTag;references:PlayerTag" json:"player,omitempty"`
	CreatedByUser *User       `gorm:"foreignKey:DiscordID;references:CreatedByDiscordID" json:"createdByUser,omitempty"`
	UpdatedByUser *User       `gorm:"foreignKey:DiscordID;references:UpdatedByDiscordID" json:"updatedByUser,omitempty"`
}
