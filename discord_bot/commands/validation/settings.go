package validation

import (
	"fmt"

	"bot/store/postgres/models"
)

const (
	MinTotalKickpoints = 3
	MaxTotalKickpoints = 20
	MinExpirationDays  = 30
	MaxExpirationDays  = 90
	MinSeasonWins      = 0
	MaxSeasonWins      = 100
)

func ValidateClanSettings(setting *models.ClanSettings) (string, bool) {
	if setting.MaxKickpoints < MinTotalKickpoints || setting.MaxKickpoints > MaxTotalKickpoints {
		return fmt.Sprintf("Die Maximale Anzahl an Kickpunkte muss zwischen %d und %d liegen", MinTotalKickpoints, MaxTotalKickpoints), false
	}
	if setting.MinSeasonWins < MinSeasonWins || setting.MinSeasonWins > MaxSeasonWins {
		return fmt.Sprintf("Die Anzahl an Season Wins muss zwischen %d und %d liegen.", MinSeasonWins, MaxSeasonWins), false
	}
	if setting.KickpointsExpireAfterDays < MinExpirationDays || setting.KickpointsExpireAfterDays > MaxExpirationDays {
		return fmt.Sprintf("Die Anzahl an Tagen muss zwischen %d und %d liegen.", MinExpirationDays, MaxExpirationDays), false
	}
	return "", true
}
