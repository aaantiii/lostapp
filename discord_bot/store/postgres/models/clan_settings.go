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

func (s KickpointSetting) DisplayString() string {
	switch s {
	case KickpointSettingMaxKickpoints:
		return "Maximale Kickpunkte"
	case KickpointSettingMinSeasonWins:
		return "Minimum Season Wins"
	case KickpointSettingExpireAfterDays:
		return "Gültigkeitsdauer in Tagen"
	case KickpointSettingSeasonWins:
		return "Kickpunkte: Season Wins"
	case KickpointSettingCWMissed:
		return "Kickpunkte: CW nicht angegriffen"
	case KickpointSettingCWFail:
		return "Kickpunkte: CW 0 Sterne"
	case KickpointSettingCWLMissed:
		return "Kickpunte: CKL nicht angegriffen"
	case KickpointSettingCWLZeroStars:
		return "Kickpunte: CKL 0 Sterne"
	case KickpointSettingCWLOneStar:
		return "Kickpunte: CKL 1 Stern"
	case KickpointSettingRaidMissed:
		return "Kickpunte: Raid nicht angegriffen"
	case KickpointSettingRaidFail:
		return "Kickpunte: Raid Fail"
	case KickpointSettingClanGames:
		return "Kickpunte: Clan Spiele nicht gemacht"
	case KickpointSettingOther:
		return "Kickpunkte: Sonstiges"
	default:
		return ""
	}
}

func (s KickpointSetting) DisplayStringShort() string {
	switch s {
	case KickpointSettingSeasonWins:
		return "Season Wins"
	case KickpointSettingCWMissed:
		return "CW nicht angegriffen"
	case KickpointSettingCWFail:
		return "CW 0 Sterne"
	case KickpointSettingCWLMissed:
		return "CKL nicht angegriffen"
	case KickpointSettingCWLZeroStars:
		return "CKL 0 Sterne"
	case KickpointSettingCWLOneStar:
		return "CKL 1 Stern"
	case KickpointSettingRaidMissed:
		return "Raid nicht angegriffen"
	case KickpointSettingRaidFail:
		return "Raid Fail"
	case KickpointSettingClanGames:
		return "Clan Spiele nicht gemacht"
	case KickpointSettingOther:
		return "Sonstiges"
	default:
		return ""
	}
}

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
