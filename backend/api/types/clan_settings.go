package types

type UpdateClanSettingsPayload struct {
	MaxKickpoints             int `binding:"required,min=3,max=20"`
	MinSeasonWins             int `binding:"required,min=0,max=200"`
	KickpointsExpireAfterDays int `binding:"required,min=0,max=9"`
	KickpointsSeasonWins      int `binding:"required,min=0,max=9"`
	KickpointsCWMissed        int `binding:"required,min=0,max=9"`
	KickpointsCWFail          int `binding:"required,min=0,max=9"`
	KickpointsCWLMissed       int `binding:"required,min=0,max=9"`
	KickpointsCWLZeroStars    int `binding:"required,min=0,max=9"`
	KickpointsCWLOneStar      int `binding:"required,min=0,max=9"`
	KickpointsRaidMissed      int `binding:"required,min=0,max=9"`
	KickpointsRaidFail        int `binding:"required,min=0,max=9"`
	KickpointsClanGames       int `binding:"required,min=0,max=9"`
}
