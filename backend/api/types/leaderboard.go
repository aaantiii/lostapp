package types

type LeaderboardParams struct {
	PaginationParams
	StatName string `form:"statName" binding:"required"`
	ClanTag  string `form:"clanTag" binding:"omitempty,min=3,max=12"`
}
