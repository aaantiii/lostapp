package models

import (
	"errors"
	"time"
)

type ClanSettings struct {
	ClanTag                   string `gorm:"size:12;primaryKey;not null"`
	MaxKickpoints             int    `gorm:"not null;default:10"`
	MinSeasonWins             int    `gorm:"not null;default:60"`
	KickpointsExpireAfterDays int    `gorm:"not null;default:45"`
	KickpointsSeasonWins      int    `gorm:"not null;default:3"`
	KickpointsCWMissed        int    `gorm:"not null;default:1"`
	KickpointsCWFail          int    `gorm:"not null;default:0"`
	KickpointsCWLMissed       int    `gorm:"not null;default:4"`
	KickpointsCWLZeroStars    int    `gorm:"not null;default:3"`
	KickpointsCWLOneStar      int    `gorm:"not null;default:1"`
	KickpointsRaidMissed      int    `gorm:"not null;default:3"`
	KickpointsRaidFail        int    `gorm:"not null;default:2"`
	KickpointsClanGames       int    `gorm:"not null;default:3"`
	UpdatedAt                 time.Time
	UpdatedByDiscordID        *string

	Clan          *Clan `gorm:"foreignKey:Tag;references:ClanTag;onUpdate:CASCADE;onDelete:CASCADE"`
	UpdatedByUser *User `gorm:"foreignKey:DiscordID;references:UpdatedByDiscordID"`
}

const (
	ClanSettingsMaxKickpoints   = "max_kickpoints"
	ClanSettingsMinSeasonWins   = "min_season_wins"
	ClanSettingsExpireAfterDays = "kickpoints_expire_after_days"
	ClanSettingsSeasonWins      = "kickpoints_season_wins"
	ClanSettingsCWMissed        = "kickpoints_cw_missed"
	ClanSettingsCWFail          = "kickpoints_cw_fail"
	ClanSettingsCWLMissed       = "kickpoints_cwl_missed"
	ClanSettingsCWLZero         = "kickpoints_cwl_zero"
	ClanSettingsCWLOne          = "kickpoints_cwl_one"
	ClanSettingsRaidMissed      = "kickpoints_raid_missed"
	ClanSettingsRaidFail        = "kickpoints_raid_fail"
	ClanSettingsClanGames       = "kickpoints_clan_games"
	ClanSettingsOther           = "kickpoints_other"
)

func (s *ClanSettings) KickpointAmountFromName(name string) (int, error) {
	switch name {
	case ClanSettingsSeasonWins:
		return s.KickpointsSeasonWins, nil
	case ClanSettingsCWMissed:
		return s.KickpointsCWMissed, nil
	case ClanSettingsCWFail:
		return s.KickpointsCWFail, nil
	case ClanSettingsCWLMissed:
		return s.KickpointsCWLMissed, nil
	case ClanSettingsCWLZero:
		return s.KickpointsCWLZeroStars, nil
	case ClanSettingsCWLOne:
		return s.KickpointsCWLOneStar, nil
	case ClanSettingsRaidMissed:
		return s.KickpointsRaidMissed, nil
	case ClanSettingsRaidFail:
		return s.KickpointsRaidFail, nil
	case ClanSettingsClanGames:
		return s.KickpointsClanGames, nil
	case ClanSettingsOther:
		return 1, nil
	default:
		return 0, errors.New("invalid name")
	}
}
