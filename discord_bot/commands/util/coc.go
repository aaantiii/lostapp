package util

import (
	"bot/store/postgres/models"
)

const (
	IndexRoleLeader = iota
	IndexRoleCoLeader
	IndexRoleElder
	IndexRoleMember
)

func SortMembersByRole(members models.ClanMembers) []models.ClanMembers {
	sortedMembers := make([]models.ClanMembers, 4)
	for _, member := range members {
		switch member.ClanRole {
		case models.RoleLeader:
			sortedMembers[IndexRoleLeader] = append(sortedMembers[IndexRoleLeader], member)
		case models.RoleCoLeader:
			sortedMembers[IndexRoleCoLeader] = append(sortedMembers[IndexRoleCoLeader], member)
		case models.RoleElder:
			sortedMembers[IndexRoleElder] = append(sortedMembers[IndexRoleElder], member)
		case models.RoleMember:
			sortedMembers[IndexRoleMember] = append(sortedMembers[IndexRoleMember], member)
		}
	}

	return sortedMembers
}
