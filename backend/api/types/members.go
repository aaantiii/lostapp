package types

type MembersParams struct {
	PlayerTagParam
	ClanTagParam
	QueryParam
	PaginationParams
}

func (p MembersParams) Conds() map[string]any {
	conds := make(map[string]any)
	if p.PlayerTag != "" {
		conds["player_tag"] = p.PlayerTag
	}
	if p.ClanTag != "" {
		conds["clan_tag"] = p.ClanTag
	}
	return conds
}

type CreateMemberPayload struct {
	PlayerTag string `binding:"-"`
	ClanTag   string `binding:"-"`
	Role      string `form:"role" binding:"required,oneof=member admin coLeader leader"`
}

type UpdateMemberPayload struct {
	CreateMemberPayload
}
