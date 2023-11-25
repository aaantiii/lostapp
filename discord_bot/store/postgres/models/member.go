package models

import "github.com/amaanq/coc.go"

type Member struct {
	PlayerTag        string
	ClanTag          string
	AddedByDiscordID string
	ClanRole         coc.Role

	Player     *Player     `gorm:"foreignKey:CocTag;references:PlayerTag"`
	Clan       *Clan       `gorm:"foreignKey:Tag;references:ClanTag"`
	Kickpoints []Kickpoint `gorm:"foreignKey:PlayerTag,ClanTag;references:PlayerTag,ClanTag;onUpdate:CASCADE;onDelete:CASCADE"`
}

func (*Member) TableName() string {
	return "clan_member"
}

type Members []Member

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
