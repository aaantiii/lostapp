package types

type UpdateClanSettingsPayload struct {
	MaxKickpoints             uint8 `json:"maxKickpoints" binding:"min=1,max=100"`
	MinSeasonWins             uint8 `json:"minSeasonWins" binding:"min=0,max=200"`
	KickpointsExpireAfterDays uint8 `json:"kickpointsExpireAfterDays" binding:"min=7,max=100"`
	KickpointsSeasonWins      uint8 `json:"kickpointsSeasonWins" binding:"min=0,max=10"`
	KickpointsCWMissed        uint8 `json:"kickpointsCWMissed" binding:"min=0,max=10"`
	KickpointsCWFail          uint8 `json:"kickpointsCWFail" binding:"min=0,max=10"`
	KickpointsCWLMissed       uint8 `json:"kickpointsCWLMissed" binding:"min=0,max=10"`
	KickpointsCWLZeroStars    uint8 `json:"kickpointsCWLZeroStars" binding:"min=0,max=10"`
	KickpointsCWLOneStar      uint8 `json:"kickpointsCWLOneStar" binding:"min=0,max=10"`
	KickpointsRaidMissed      uint8 `json:"kickpointsRaidMissed" binding:"min=0,max=10"`
	KickpointsRaidFail        uint8 `json:"kickpointsRaidFail" binding:"min=0,max=10"`
	KickpointsClanGames       uint8 `json:"kickpointsClanGames" binding:"min=0,max=10"`
}
