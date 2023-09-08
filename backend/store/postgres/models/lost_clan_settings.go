package models

type LostClanSettings struct {
	ClanTag                   string `gorm:"size:10;primaryKey;not null" json:"-"`
	MaxKickpoints             uint8  `gorm:"default:10;not null" json:"maxKickpoints"`
	MinSeasonWins             uint8  `gorm:"default:80;not null" json:"minSeasonWins"`
	KickpointsExpireAfterDays uint8  `gorm:"default:60;not null" json:"kickpointsExpireAfterDays"`
	KickpointsSeasonWins      uint8  `gorm:"default:3;not null"  json:"kickpointsSeasonWins"`
	KickpointsCWMissed        uint8  `gorm:"default:3;not null" json:"kickpointsCWMissed"`
	KickpointsCWFail          uint8  `gorm:"default:1;not null" json:"kickpointsCWFail"`
	KickpointsCWLMissed       uint8  `gorm:"default:3;not null" json:"kickpointsCWLMissed"`
	KickpointsCWLZeroStars    uint8  `gorm:"default:3;not null" json:"kickpointsCWLZeroStars"`
	KickpointsCWLOneStar      uint8  `gorm:"default:2;not null" json:"kickpointsCWLOneStar"`
	KickpointsRaidMissed      uint8  `gorm:"default:3;not null" json:"kickpointsRaidMissed"`
	KickpointsRaidFail        uint8  `gorm:"default:2;not null" json:"kickpointsRaidFail"`
	KickpointsClanGames       uint8  `gorm:"default:3;not null" json:"kickpointsClanGames"`
}

func (*LostClanSettings) TableName() string {
	return "lost_clan_settings"
}
