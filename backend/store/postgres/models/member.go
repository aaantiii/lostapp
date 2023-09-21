package models

import "github.com/amaanq/coc.go"

type Member struct {
	PlayerTag        string   `gorm:"primaryKey;not null;size:10"`
	ClanTag          string   `gorm:"primaryKey;not null;size:10"`
	AddedByDiscordID string   `gorm:"size:18;not null"`
	ClanRole         coc.Role `gorm:"not null"`
	IsAdmin          bool     `gorm:"not null;default:false"`

	LostClan    LostClan     `gorm:"foreignKey:ClanTag;references:Tag;onUpdate:CASCADE;onDelete:CASCADE"`
	DiscordLink DiscordLink  `gorm:"foreignKey:PlayerTag;references:CocTag"`
	Kickpoints  *[]Kickpoint `gorm:"foreignKey:PlayerTag,ClanTag;references:PlayerTag,ClanTag;onUpdate:CASCADE;onDelete:CASCADE"`
}

func (*Member) TableName() string {
	return "clan_member"
}

type Members []*Member

func (members Members) Tags() []string {
	if members == nil {
		return nil
	}

	tags := make([]string, len(members))
	for i, member := range members {
		tags[i] = member.PlayerTag
	}

	return tags
}

func (members Members) TagsDistinct() []string {
	if members == nil {
		return nil
	}

	tags := make([]string, 0)
	seen := make(map[string]bool)
	for _, member := range members {
		if seen[member.PlayerTag] {
			continue
		}

		tags = append(tags, member.PlayerTag)
		seen[member.PlayerTag] = true
	}

	return tags
}
