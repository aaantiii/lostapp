package types

import (
	"time"

	"github.com/amaanq/coc.go"
)

type ClanMemberKickpoints struct {
	Tag    string   `json:"tag"`
	Name   string   `json:"name"`
	Role   coc.Role `json:"role"`
	Amount int      `json:"amount"`
}

type CreateKickpointPayload struct {
	Date             time.Time `json:"date" binding:"required"`
	Amount           int       `json:"amount" binding:"required,min=1,max=10"`
	Description      string    `json:"description" binding:"required"`
	PlayerTag        string    `binding:"-"`
	ClanTag          string    `binding:"-"`
	AddedByDiscordID string    `binding:"-"`
}

type UpdateKickpointPayload struct {
	Date               time.Time `json:"date" binding:"required"`
	Amount             int       `json:"amount" binding:"required,min=1,max=10"`
	Description        string    `json:"description" binding:"required"`
	UpdatedByDiscordID string    `binding:"-"`
}
