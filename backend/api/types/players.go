package types

import (
	"github.com/aaantiii/goclash"

	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type PlayersParams struct {
	PaginationParams
	QueryParam
	DiscordID string `form:"discordId" binding:"omitempty,len=18"`
}

type Player struct {
	DiscordID           string             `json:"discordId"`
	ClanMembers         models.ClanMembers `json:"clanMembers,omitempty"`
	*goclash.PlayerBase `json:",inline"`
}

func (p PlayersParams) Conds() map[string]interface{} {
	conds := make(map[string]interface{})
	if p.DiscordID != "" {
		conds["discord_id"] = p.DiscordID
	}
	return conds
}

type MembersPlayersParams struct {
	ClanTagParam
}
