package types

import (
	"github.com/amaanq/coc.go"
)

type AddMemberPayload struct {
	Tag              string   `json:"tag" binding:"required"`
	ClanTag          string   `json:"clanTag" binding:"required"`
	Role             coc.Role `json:"role" binding:"required"`
	AddedByDiscordID string   `binding:"-"`
}

type UpdateMemberPayload AddMemberPayload
