package util

import (
	"github.com/amaanq/coc.go"

	"bot/store/postgres/models"
)

const (
	IndexRoleLeader = iota
	IndexRoleCoLeader
	IndexRoleElder
	IndexRoleMember
)

func SortMembersByRole(members models.Members) []models.Members {
	sortedMembers := make([]models.Members, 4)
	for _, member := range members {
		switch member.ClanRole {
		case coc.Leader:
			sortedMembers[IndexRoleLeader] = append(sortedMembers[IndexRoleLeader], member)
		case coc.CoLeader:
			sortedMembers[IndexRoleCoLeader] = append(sortedMembers[IndexRoleCoLeader], member)
		case coc.Elder:
			sortedMembers[IndexRoleElder] = append(sortedMembers[IndexRoleElder], member)
		case coc.Member:
			sortedMembers[IndexRoleMember] = append(sortedMembers[IndexRoleMember], member)
		}
	}

	return sortedMembers
}
