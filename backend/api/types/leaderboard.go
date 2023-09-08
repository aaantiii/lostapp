package types

type LeaderboardPlayersParams struct {
	PaginationParams
	StatsID int    `binding:"-"`
	ClanTag string `form:"clanTag"`
}
