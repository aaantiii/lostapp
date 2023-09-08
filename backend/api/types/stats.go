package types

import (
	"github.com/amaanq/coc.go"
)

// ComparableStatistic is a coc.Player statistic that can be compared between players.
type ComparableStatistic struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

var (
	StatSeasonWins = ComparableStatistic{
		ID:          1,
		Name:        "Season Wins",
		DisplayName: "Angriffe gewonnen (Saison)",
	}
	StatDarkElixirLooted = ComparableStatistic{
		ID:          2,
		Name:        coc.HeroicHeist.Name,
		DisplayName: "Dunkles Elixier gesammelt",
	}
	StatSeasonPoints = ComparableStatistic{
		ID:          3,
		Name:        coc.WellSeasoned.Name,
		DisplayName: "Season Challenge Punkte",
	}
	StatObstaclesRemoved = ComparableStatistic{
		ID:          4,
		Name:        coc.NiceAndTidy.Name,
		DisplayName: "Hindernisse entfernt",
	}
	StatTroopsDonated = ComparableStatistic{
		ID:          5,
		Name:        coc.FriendInNeed.Name,
		DisplayName: "Truppen gespendet",
	}
	StatSpellsDonated = ComparableStatistic{
		ID:          6,
		Name:        coc.SharingIsCaring.Name,
		DisplayName: "Zauber gespendet",
	}
	StatSiegeMachinesDonated = ComparableStatistic{
		ID:          7,
		Name:        coc.SiegeSharer.Name,
		DisplayName: "Siege Machines gespendet",
	}
	StatWarStars = ComparableStatistic{
		ID:          8,
		Name:        coc.WarHero.Name,
		DisplayName: "CW Sterne",
	}
	StatCWLStars = ComparableStatistic{
		ID:          9,
		Name:        coc.WarLeagueLegend.Name,
		DisplayName: "CWL Sterne",
	}
	StatClanGamesPoints = ComparableStatistic{
		ID:          10,
		Name:        coc.GamesChampion.Name,
		DisplayName: "Clan Games Punkte",
	}
	StatSuccessfulDefenses = ComparableStatistic{
		ID:          11,
		Name:        coc.Unbreakable.Name,
		DisplayName: "Erfolgreiche Verteidigungen",
	}
	StatHighestTrophies = ComparableStatistic{
		ID:          12,
		Name:        coc.SweetVictory.Name,
		DisplayName: "Meiste Trophäen",
	}
	StatAttackWins = ComparableStatistic{
		ID:          13,
		Name:        coc.Conqueror.Name,
		DisplayName: "Angriffe gewonnen",
	}
	StatTownHallsDestroyed = ComparableStatistic{
		ID:          14,
		Name:        coc.Humiliator.Name,
		DisplayName: "Rathäuser zerstört",
	}
	StatWeaponizedTownHallsDestroyed = ComparableStatistic{
		ID:          15,
		Name:        coc.NotSoEasyThisTime.Name,
		DisplayName: "Bewaffnete Rathäuser zerstört",
	}
	StatBuilderHutsDestroyed = ComparableStatistic{
		ID:          16,
		Name:        coc.UnionBuster.Name,
		DisplayName: "Bauhütten zerstört",
	}
	StatWeaponizedBuilderHutsDestroyed = ComparableStatistic{
		ID:          17,
		Name:        coc.BustThis.Name,
		DisplayName: "Bewaffnete Bauhütten zerstört",
	}
	StatWallsDestroyed = ComparableStatistic{
		ID:          18,
		Name:        coc.WallBuster.Name,
		DisplayName: "Mauern zerstört",
	}
	StatMortarsDestroyed = ComparableStatistic{
		ID:          19,
		Name:        coc.MortarMauler.Name,
		DisplayName: "Minenwerfer zerstört",
	}
	StatXBowsDestroyed = ComparableStatistic{
		ID:          20,
		Name:        coc.XBowExterminator.Name,
		DisplayName: "X-Bögen zerstört",
	}
	StatInfernoTowersDestroyed = ComparableStatistic{
		ID:          21,
		Name:        coc.Firefighter.Name,
		DisplayName: "Infernotürme zerstört",
	}
	StatEagleArtilleriesDestroyed = ComparableStatistic{
		ID:          22,
		Name:        coc.AntiArtillery.Name,
		DisplayName: "Adlerartillerien zerstört",
	}
	StatScattershotsDestroyed = ComparableStatistic{
		ID:          23,
		Name:        coc.ShatteredAndScattered.Name,
		DisplayName: "Scattershots zerstört",
	}
	StatSuperTroopsBoosted = ComparableStatistic{
		ID:          24,
		Name:        coc.SuperbWork.Name,
		DisplayName: "Supertruppen geboosted",
	}
	StatHighestBuilderBaseTrophies = ComparableStatistic{
		ID:          25,
		Name:        coc.ChampionBuilder.Name,
		DisplayName: "Meiste Trophäen (BB)",
	}
	StatSpellTowersDestroyed = ComparableStatistic{
		ID:          26,
		Name:        coc.Counterspell.Name,
		DisplayName: "Zaubertürme zerstört",
	}
	StatMonolithsDestroyed = ComparableStatistic{
		ID:          27,
		Name:        coc.MonolithMasher.Name,
		DisplayName: "Monolithen zerstört",
	}
	StatClanGoldLooted = ComparableStatistic{
		ID:          28,
		Name:        coc.AggressiveCapitalism.Name,
		DisplayName: "Clan Gold aus Raids",
	}
	StatClanGoldContributed = ComparableStatistic{
		ID:          29,
		Name:        coc.MostValuableClanmate.Name,
		DisplayName: "Clan Gold ausgegeben",
	}
)

// ComparableStats is a list of all comparable stats a coc.Player has. This includes achievements and other stats.
var ComparableStats = []ComparableStatistic{
	StatSeasonWins, // player.AttackWins
	StatDarkElixirLooted,
	StatSeasonPoints,
	StatObstaclesRemoved,
	StatTroopsDonated,
	StatSpellsDonated,
	StatSiegeMachinesDonated,
	StatWarStars,
	StatCWLStars,
	StatClanGamesPoints,
	StatSuccessfulDefenses,
	StatHighestTrophies,
	StatAttackWins,
	StatTownHallsDestroyed,
	StatWeaponizedTownHallsDestroyed,
	StatBuilderHutsDestroyed,
	StatWeaponizedBuilderHutsDestroyed,
	StatWallsDestroyed,
	StatMortarsDestroyed,
	StatXBowsDestroyed,
	StatInfernoTowersDestroyed,
	StatEagleArtilleriesDestroyed,
	StatScattershotsDestroyed,
	StatSuperTroopsBoosted,
	StatHighestBuilderBaseTrophies,
	StatSpellTowersDestroyed,
	StatMonolithsDestroyed,
	StatClanGoldLooted,
	StatClanGoldContributed,
}

// ComparableAchievements are the achievements that can be compared between players.
var ComparableAchievements = []ComparableStatistic{
	StatDarkElixirLooted,
	StatSeasonPoints,
	StatObstaclesRemoved,
	StatTroopsDonated,
	StatSpellsDonated,
	StatSiegeMachinesDonated,
	StatWarStars,
	StatCWLStars,
	StatClanGamesPoints,
	StatSuccessfulDefenses,
	StatHighestTrophies,
	StatAttackWins,
	StatTownHallsDestroyed,
	StatWeaponizedTownHallsDestroyed,
	StatBuilderHutsDestroyed,
	StatWeaponizedBuilderHutsDestroyed,
	StatWallsDestroyed,
	StatMortarsDestroyed,
	StatXBowsDestroyed,
	StatInfernoTowersDestroyed,
	StatEagleArtilleriesDestroyed,
	StatScattershotsDestroyed,
	StatSuperTroopsBoosted,
	StatHighestBuilderBaseTrophies,
	StatSpellTowersDestroyed,
	StatMonolithsDestroyed,
	StatClanGoldLooted,
	StatClanGoldContributed,
}
