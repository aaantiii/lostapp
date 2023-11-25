package services

import (
	"github.com/amaanq/coc.go"

	"backend/api/repos"
	"backend/store/postgres/models"
)

type IMembersService interface {
	MembersByDiscordID(discordID string) (models.Members, error)
	MemberIsLeadingClan(discordID, clanTag string) bool
}

type MembersService struct {
	membersRepo repos.IMembersRepo
}

func NewMembersService(memberRepo repos.IMembersRepo) IMembersService {
	return &MembersService{membersRepo: memberRepo}
}

func (service *MembersService) MembersByDiscordID(discordID string) (models.Members, error) {
	return service.membersRepo.MembersByDiscordID(discordID)
}

func (service *MembersService) MemberIsLeadingClan(discordID, clanTag string) bool {
	members, err := service.MembersByDiscordID(discordID)
	if err != nil {
		return false
	}

	var member *models.Member
	for _, m := range members {
		if m.ClanTag == clanTag {
			member = m
		}
	}

	if member == nil {
		return false
	}

	return member.ClanRole == coc.Leader || member.ClanRole == coc.CoLeader
}
