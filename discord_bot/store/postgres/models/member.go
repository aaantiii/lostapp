package models

import (
	"github.com/amaanq/coc.go"
	"github.com/bwmarrin/discordgo"
)

type Member struct {
	PlayerTag        string   `gorm:"primaryKey;not null;size:12"`
	ClanTag          string   `gorm:"primaryKey;not null;size:12"`
	AddedByDiscordID string   `gorm:"size:18;not null"`
	ClanRole         coc.Role `gorm:"not null"`

	LostClan  LostClan   `gorm:"foreignKey:ClanTag;references:Tag;onUpdate:CASCADE;onDelete:CASCADE"`
	Player    Player     `gorm:"foreignKey:PlayerTag;references:CocTag"`
	Penalties *[]Penalty `gorm:"foreignKey:PlayerTag,ClanTag;references:PlayerTag,ClanTag;onUpdate:CASCADE;onDelete:CASCADE"`
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
	if members == nil || len(members) == 0 {
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

func (members Members) Choices() []*discordgo.ApplicationCommandOptionChoice {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(members))
	for i, member := range members {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  member.Player.Name,
			Value: member.PlayerTag,
		}
	}

	return choices
}
