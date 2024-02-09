package utils

import (
	"errors"

	"github.com/aaantiii/goclash"

	"github.com/aaantiii/lostapp/backend/api/types"
)

func ComparableStatisticByName(name string) (*types.ComparableStatistic, error) {
	for _, stat := range types.ComparableStats {
		if stat.Name == name {
			return stat, nil
		}
	}
	return nil, types.ErrInvalidStatName
}

func StatisticValueFromPlayers(players goclash.Players, stat *types.ComparableStatistic) ([]int, error) {
	if len(players) == 0 {
		return nil, errors.New("no tags provided")
	}

	values := make([]int, len(players))
	if achievements, err := players.GetAchievement(&goclash.Achievement{Name: stat.Name}); err == nil {
		for i := range players {
			values[i] = achievements[i].Value
		}
	} else {
		for i, player := range players {
			values[i] = nonAchievementStatisticValue(player, stat)
		}
	}

	return values, nil
}

// for stats that are not listed in coc.Achievements, e.g. Season Wins
func nonAchievementStatisticValue(player *goclash.Player, stat *types.ComparableStatistic) int {
	switch stat.Name {
	case types.StatSeasonWins.Name:
		return player.AttackWins
	default:
		return 0
	}
}
