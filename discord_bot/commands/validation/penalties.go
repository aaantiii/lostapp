package validation

import (
	"errors"
	"strings"

	"bot/store/postgres/models"
)

func PenaltyType(t string) (models.PenaltyType, error) {
	tLower := strings.ToLower(t)

	if strings.HasPrefix(tLower, "k") {
		return models.PenaltyTypeKickpoint, nil
	}

	if strings.HasPrefix(tLower, "v") {
		return models.PenaltyTypeWarning, nil
	}

	return "", errors.New("invalid penalty type")
}
