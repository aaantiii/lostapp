package types

import (
	"time"

	"github.com/amaanq/coc.go"

	"backend/store/postgres/models"
)

type Kickpoint struct {
	ID            uint        `json:"id"`
	Amount        int         `json:"amount"`
	Date          time.Time   `json:"date"`
	Description   string      `json:"description"`
	CreatedByUser models.User `json:"createdByUser"`
}

func NewKickpoint(kickpoint *models.Kickpoint) *Kickpoint {
	return &Kickpoint{
		ID:            kickpoint.ID,
		Date:          kickpoint.Date,
		Amount:        kickpoint.Amount,
		Description:   kickpoint.Description,
		CreatedByUser: kickpoint.CreatedByUser,
	}
}

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
