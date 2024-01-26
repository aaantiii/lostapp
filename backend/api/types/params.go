package types

type ClanTagParam struct {
	ClanTag string `form:"clanTag" binding:"omitempty,min=3,max=12"`
}

type PlayerTagParam struct {
	PlayerTag string `form:"playerTag" binding:"omitempty,min=3,max=12"`
}

type QueryParam struct {
	Query string `form:"q" binding:"omitempty,min=1,max=100"`
}

type CreatedAtParam struct {
	CreatedAt string `form:"createdAt" binding:"omitempty,oneof=asc desc"`
}

func (p CreatedAtParam) OrderCreatedAt() string {
	if p.CreatedAt == "asc" {
		return "created_at asc"
	}

	return "created_at desc"
}
