package models

import "github.com/amaanq/coc.go"

type PlayerClan struct {
	PlayerTag string   `gorm:"primaryKey;not null;size:12" json:"-"`
	ClanTag   string   `gorm:"primaryKey;not null;size:12" json:"tag"`
	Name      string   `gorm:"not null;size:32" json:"name"`
	Role      coc.Role `gorm:"not null;size:16" json:"role"`

	Member Member `gorm:"foreignKey:PlayerTag,ClanTag;references:PlayerTag,ClanTag;onUpdate:CASCADE;onDelete:CASCADE" json:"-"`
}

func (*PlayerClan) TableName() string {
	return "lostapp_player_clans_live"
}

func NewPlayerClan(member *Member, name string) *PlayerClan {
	return &PlayerClan{
		PlayerTag: member.PlayerTag,
		ClanTag:   member.ClanTag,
		Name:      name,
		Role:      member.ClanRole,
	}
}
