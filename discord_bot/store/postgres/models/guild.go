package models

import "slices"

const LostFamilyGuildID = "733857906117574717"

type Guild struct {
	GuildID        string
	ClanTag        string
	LeaderRoleID   string
	CoLeaderRoleID string
	ElderRoleID    string
	MemberRoleID   string
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
