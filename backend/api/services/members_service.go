package services

import (
	"errors"
	"fmt"

	"github.com/amaanq/coc.go"
	"gorm.io/gorm"

	"backend/api/repos"
	"backend/api/types"
	"backend/store/postgres/models"
)

type IMembersService interface {
	MembersByDiscordID(discordID string) (models.Members, error)
	CreateMember(payload types.AddMemberPayload) error
	UpdateMember(payload types.UpdateMemberPayload) error
	MemberIsLeadingClan(discordID, clanTag string) bool
}

type MembersService struct {
	membersRepo repos.IMembersRepo
}

func NewMembersService(memberRepo repos.IMembersRepo) *MembersService {
	return &MembersService{membersRepo: memberRepo}
}

func (service *MembersService) MembersByDiscordID(discordID string) (models.Members, error) {
	return service.membersRepo.MembersByDiscordID(discordID)
}

func (service *MembersService) CreateMember(payload types.AddMemberPayload) error {
	_, err := service.membersRepo.MembersByID(payload.Tag, payload.ClanTag)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("member already exists")
	}

	return service.membersRepo.CreateMember(&models.Member{
		PlayerTag:        payload.Tag,
		ClanTag:          payload.ClanTag,
		AddedByDiscordID: payload.AddedByDiscordID,
		ClanRole:         payload.Role,
	})
}

func (service *MembersService) UpdateMember(payload types.UpdateMemberPayload) error {
	_, err := service.membersRepo.MembersByID(payload.Tag, payload.ClanTag)
	if err != nil {
		return err
	}

	return service.membersRepo.UpdateMember(&models.Member{
		PlayerTag: payload.Tag,
		ClanTag:   payload.ClanTag,
		ClanRole:  payload.Role,
	})
}

func (service *MembersService) DeleteMember(tag, clanTag string) error {
	return service.membersRepo.DeleteMember(tag, clanTag)
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
