package types

import (
	"time"

	"github.com/amaanq/coc.go"

	"backend/store/postgres/models"
)

type Kickpoint struct {
	ID          uint                   `json:"id"`
	Amount      int                    `json:"amount"`
	Date        time.Time              `json:"date"`
	Reason      models.KickpointReason `json:"reason"`
	Description string                 `json:"description"`
}

func NewKickpoint(kickpoint *models.Kickpoint) *Kickpoint {
	return &Kickpoint{
		ID:          kickpoint.ID,
		Amount:      kickpoint.Amount,
		Date:        kickpoint.Date,
		Reason:      kickpoint.Reason,
		Description: kickpoint.Description,
	}
}

type ClanMemberKickpoints struct {
	Tag    string   `json:"tag"`
	Name   string   `json:"name"`
	Role   coc.Role `json:"role"`
	Amount int      `json:"amount"`
}

type CreateKickpointPayload struct {
	PlayerTag   string                 `json:"playerTag" binding:"-"`
	ClanTag     string                 `json:"clanTag" binding:"-"`
	Date        time.Time              `json:"date" binding:"required"`
	Amount      int                    `json:"amount" binding:"required,min=1"`
	Reason      models.KickpointReason `json:"reason" binding:"required"`
	Description string                 `json:"description" binding:"required"`
}

type UpdateKickpointPayload CreateKickpointPayload
