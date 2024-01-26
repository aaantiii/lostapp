package types

type LeaderboardParams struct {
	PaginationParams
	ClanTag string `form:"clanTag" binding:"omitempty,min=3,max=12"`
}
