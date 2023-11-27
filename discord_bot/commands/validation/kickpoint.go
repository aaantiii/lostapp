package validation

import (
	"fmt"

	"bot/store/postgres/models"
)

const (
	MinTotalKickpoints        = 3
	MaxTotalKickpoints        = 20
	MinExpirationDays         = 30
	MaxExpirationDays         = 90
	MinSeasonWins             = 0
	MaxSeasonWins             = 100
	MinKickpointAmount        = 1
	MinKickpointSettingAmount = 0
	MaxKickpointAmount        = 9
)

func ValidateKickpointSettings(name string, value int) (string, bool) {
	switch name {
	case models.ClanSettingsMaxKickpoints:
		return fmt.Sprintf("Die Maximale Anzahl an Kickpunkte muss zwischen %d und %d liegen", MinTotalKickpoints, MaxTotalKickpoints), value >= MinTotalKickpoints && value <= MaxTotalKickpoints
	case models.ClanSettingsMinSeasonWins:
		return fmt.Sprintf("Die Anzahl an Season Wins muss zwischen %d und %d liegen.", MinSeasonWins, MaxSeasonWins), value >= MinSeasonWins && value <= MaxSeasonWins
	case models.ClanSettingsExpireAfterDays:
		return fmt.Sprintf("Die Anzahl an Tagen muss zwischen %d und %d liegen.", MinExpirationDays, MaxExpirationDays), value >= MinExpirationDays && value <= MaxExpirationDays
	case models.ClanSettingsSeasonWins, models.ClanSettingsCWMissed, models.ClanSettingsCWFail, models.ClanSettingsCWLMissed, models.ClanSettingsCWLZero, models.ClanSettingsCWLOne, models.ClanSettingsRaidMissed, models.ClanSettingsRaidFail, models.ClanSettingsClanGames:
		return fmt.Sprintf("Die Anzahl an Kickpunkten muss zwischen %d und %d liegen.", MinKickpointAmount, MaxKickpointAmount), value >= MinKickpointSettingAmount && value <= MaxKickpointAmount
	default:
		return "Es wurde eine ungültige Einstellung angegeben.", false
	}
}
