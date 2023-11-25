package types

type PlayersParams struct {
	Name      string `form:"name" binding:"omitempty,min=3,max=30"`
	Tag       string `form:"tag" binding:"omitempty,min=3,max=11"`
	ClanName  string `form:"clanName" binding:"omitempty,min=3,max=30"`
	ClanTag   string `form:"clanTag" binding:"omitempty,min=3,max=11"`
	DiscordID string `form:"discordID" binding:"omitempty,len=18"`
}
