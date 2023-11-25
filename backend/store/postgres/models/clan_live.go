package models

import "github.com/amaanq/coc.go"

type Clan struct {
	Tag            string        `gorm:"primaryKey;not null" json:"tag"`
	Name           string        `gorm:"not null" json:"name"`
	WarTies        *int          `json:"warTies,omitempty"`
	Description    *string       `json:"description,omitempty"`
	WarLosses      *int          `json:"warLosses,omitempty"`
	BadgeURL       string        `gorm:"not null" json:"badgeUrl"`
	Level          int           `gorm:"not null" json:"level"`
	WarWins        int           `gorm:"not null" json:"warWins"`
	MemberCount    int           `gorm:"not null" json:"members"`
	IsWarLogPublic bool          `gorm:"not null" json:"isWarLogPublic"`
	Members        *[]Members    `gorm:"foreignKey:Tag;references:Tag;onUpdate:CASCADE;onDelete:CASCADE" json:"-"`
	MemberList     *[]ClanMember `gorm:"foreignKey:ClanTag;references:Tag;onUpdate:CASCADE;onDelete:CASCADE" json:"memberList,omitempty"`
}

func (*Clan) TableName() string {
	return "lostapp_clans_live"
}

func NewClan(clan *coc.Clan) *Clan {
	return &Clan{
		Tag:            clan.Tag,
		Name:           clan.Name,
		WarTies:        clan.WarTies,
		Description:    clan.Description,
		WarLosses:      clan.WarLosses,
		BadgeURL:       clan.BadgeURLs.Medium,
		Level:          clan.ClanLevel,
		WarWins:        clan.WarWins,
		MemberCount:    clan.Members,
		IsWarLogPublic: clan.IsWarLogPublic,
	}
}
