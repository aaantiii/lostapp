package types

type PlayersParams struct {
	Name      string `form:"name" binding:"omitempty,min=3,max=50"`
	Tag       string `form:"tag" binding:"omitempty,min=3,max=12"`
	ClanName  string `form:"clanName" binding:"omitempty,min=3,max=50"`
	ClanTag   string `form:"clanTag" binding:"omitempty,min=3,max=12"`
	DiscordID string `form:"discordID" binding:"omitempty,len=18"`
}
