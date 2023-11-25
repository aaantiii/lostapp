package models

import "slices"

const LostFamilyGuildID = "733857906117574717"

type Guild struct {
	GuildID        string `gorm:"primaryKey;not null;size:20"`
	ClanTag        string `gorm:"primaryKey;not null;size:12"`
	LeaderRoleID   string `gorm:"size:20"`
	CoLeaderRoleID string `gorm:"size:20"`
	ElderRoleID    string `gorm:"size:20"`
	MemberRoleID   string `gorm:"size:20"`
}

func (*Guild) TableName() string {
	return "guild"
}

func (g *Guild) IsLeader(roles []string) bool {
	return slices.Contains(roles, g.LeaderRoleID)
}

func (g *Guild) IsCoLeader(roles []string) bool {
	return slices.Contains(roles, g.CoLeaderRoleID)
}

func (g *Guild) IsElder(roles []string) bool {
	return slices.Contains(roles, g.ElderRoleID)
}

func (g *Guild) IsMember(roles []string) bool {
	return slices.Contains(roles, g.MemberRoleID)
}
