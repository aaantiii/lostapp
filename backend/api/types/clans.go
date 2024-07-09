package types

import (
	"github.com/aaantiii/goclash"

	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type ClansParams struct {
	PaginationParams
	QueryParam
}

type Clan struct {
	*goclash.Clan
	LostMembers models.ClanMembers `json:"lostMembers"`
}
