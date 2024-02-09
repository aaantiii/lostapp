package models

import (
	"errors"
	"time"
)

type ClanSettings struct {
	ClanTag                   string    `gorm:"size:12;primaryKey;not null" json:"clanTag"`
	MaxKickpoints             int       `gorm:"not null;default:6" json:"maxKickpoints"`
	MinSeasonWins             int       `gorm:"not null;default:80" json:"minSeasonWins"`
	KickpointsExpireAfterDays int       `gorm:"not null;default:45" json:"kickpointsExpireAfterDays"`
	KickpointsSeasonWins      int       `gorm:"not null;default:2" json:"kickpointsSeasonWins"`
	KickpointsCWMissed        int       `gorm:"not null;default:1" json:"kickpointsCWMissed"`
	KickpointsCWFail          int       `gorm:"not null;default:0" json:"kickpointsCWFail"`
	KickpointsCWLMissed       int       `gorm:"not null;default:2" json:"kickpointsCWLMissed"`
	KickpointsCWLZeroStars    int       `gorm:"not null;default:2" json:"kickpointsCWLZeroStars"`
	KickpointsCWLOneStar      int       `gorm:"not null;default:2" json:"kickpointsCWLOneStar"`
	KickpointsRaidMissed      int       `gorm:"not null;default:2" json:"kickpointsRaidMissed"`
	KickpointsRaidFail        int       `gorm:"not null;default:2" json:"kickpointsRaidFail"`
	KickpointsClanGames       int       `gorm:"not null;default:2" json:"kickpointsClanGames"`
	UpdatedAt                 time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	UpdatedByDiscordID        *string   `gorm:"size:18" json:"-"`

	Clan          *Clan `gorm:"foreignKey:Tag;references:ClanTag;onUpdate:CASCADE;onDelete:CASCADE" json:"clan,omitempty"`
	UpdatedByUser *User `gorm:"foreignKey:DiscordID;references:UpdatedByDiscordID" json:"updatedByUser,omitempty"`
}

type KickpointSetting string

const (
	KickpointSettingMaxKickpoints   KickpointSetting = "max_kickpoints"
	KickpointSettingMinSeasonWins   KickpointSetting = "min_season_wins"
	KickpointSettingExpireAfterDays KickpointSetting = "kickpoints_expire_after_days"
	KickpointSettingSeasonWins      KickpointSetting = "kickpoints_season_wins"
	KickpointSettingCWMissed        KickpointSetting = "kickpoints_cw_missed"
	KickpointSettingCWFail          KickpointSetting = "kickpoints_cw_fail"
	KickpointSettingCWLMissed       KickpointSetting = "kickpoints_cwl_missed"
	KickpointSettingCWLZeroStars    KickpointSetting = "kickpoints_cwl_zero_stars"
	KickpointSettingCWLOneStar      KickpointSetting = "kickpoints_cwl_one_star"
	KickpointSettingRaidMissed      KickpointSetting = "kickpoints_raid_missed"
	KickpointSettingRaidFail        KickpointSetting = "kickpoints_raid_fail"
	KickpointSettingClanGames       KickpointSetting = "kickpoints_clan_games"
	KickpointSettingOther           KickpointSetting = "kickpoints_other"
)

func (s *ClanSettings) KickpointAmountFromSetting(setting KickpointSetting) (int, error) {
	switch setting {
	case KickpointSettingSeasonWins:
		return s.KickpointsSeasonWins, nil
	case KickpointSettingCWMissed:
		return s.KickpointsCWMissed, nil
	case KickpointSettingCWFail:
		return s.KickpointsCWFail, nil
	case KickpointSettingCWLMissed:
		return s.KickpointsCWLMissed, nil
	case KickpointSettingCWLZeroStars:
		return s.KickpointsCWLZeroStars, nil
	case KickpointSettingCWLOneStar:
		return s.KickpointsCWLOneStar, nil
	case KickpointSettingRaidMissed:
		return s.KickpointsRaidMissed, nil
	case KickpointSettingRaidFail:
		return s.KickpointsRaidFail, nil
	case KickpointSettingClanGames:
		return s.KickpointsClanGames, nil
	case KickpointSettingOther:
		return 1, nil
	default:
		return 0, errors.New("invalid name")
	}
}
