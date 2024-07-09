package types

import "github.com/aaantiii/goclash"

type PlayerStatistic struct {
	Tag       string `json:"tag"`
	Name      string `json:"name"`
	ClanName  string `json:"clanName"`
	Placement int    `json:"placement"`
	Value     int    `json:"value"`
}

type PlayerStatistics []*PlayerStatistic

// ComparableStatistic is a goclash.Player statistic that can be compared between players.
type ComparableStatistic struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Task        string `json:"task"`
}

var (
	StatSeasonWins = &ComparableStatistic{
		Name:        "Season Wins",
		DisplayName: "Angriffe gewonnen (Saison)",
	}
	StatDarkElixirLooted = &ComparableStatistic{
		Name:        goclash.AchievementHeroicHeist.Name,
		DisplayName: "Dunkles Elixier gesammelt",
		Task:        "Dunkles Elixier sammeln",
	}
	StatSeasonPoints = &ComparableStatistic{
		Name:        goclash.AchievementWellSeasoned.Name,
		DisplayName: "Season Challenge Punkte",
		Task:        "Season Challenge Punkte sammeln",
	}
	StatObstaclesRemoved = &ComparableStatistic{
		Name:        goclash.AchievementNiceAndTidy.Name,
		DisplayName: "Hindernisse entfernt",
		Task:        "Hindernisse entfernen",
	}
	StatTroopsDonated = &ComparableStatistic{
		Name:        goclash.AchievementFriendInNeed.Name,
		DisplayName: "Truppen gespendet",
		Task:        "Truppen spenden",
	}
	StatSpellsDonated = &ComparableStatistic{
		Name:        goclash.AchievementSharingIsCaring.Name,
		DisplayName: "Zauber gespendet",
		Task:        "Zauber spenden",
	}
	StatSiegeMachinesDonated = &ComparableStatistic{
		Name:        goclash.AchievementSiegeSharer.Name,
		DisplayName: "Siege Machines gespendet",
		Task:        "Siege Machines spenden",
	}
	StatWarStars = &ComparableStatistic{
		Name:        goclash.AchievementWarHero.Name,
		DisplayName: "CW Sterne",
		Task:        "CW Sterne sammeln",
	}
	StatCWLStars = &ComparableStatistic{
		Name:        goclash.AchievementWarLeagueLegend.Name,
		DisplayName: "CWL Sterne",
		Task:        "CWL Sterne sammeln",
	}
	StatClanGamesPoints = &ComparableStatistic{
		Name:        goclash.AchievementGamesChampion.Name,
		DisplayName: "Clan Games Punkte",
		Task:        "Clan Games Punkte sammeln",
	}
	StatSuccessfulDefenses = &ComparableStatistic{
		Name:        goclash.AchievementUnbreakable.Name,
		DisplayName: "Erfolgreiche Verteidigungen",
		Task:        "Erfolgreiche Verteidigungen sammeln",
	}
	StatHighestTrophies = &ComparableStatistic{
		Name:        goclash.AchievementSweetVictory.Name,
		DisplayName: "Höchste Trophäen",
		Task:        "Höchste Trophäen",
	}
	StatAttackWins = &ComparableStatistic{
		Name:        goclash.AchievementConqueror.Name,
		DisplayName: "Angriffe gewonnen",
		Task:        "Angriffe gewinnen",
	}
	StatTownHallsDestroyed = &ComparableStatistic{
		Name:        goclash.AchievementHumiliator.Name,
		DisplayName: "Rathäuser zerstört",
		Task:        "Rathäuser zerstören",
	}
	StatWeaponizedTownHallsDestroyed = &ComparableStatistic{
		Name:        goclash.AchievementNotSoEasyThisTime.Name,
		DisplayName: "Bewaffnete Rathäuser zerstört",
		Task:        "Bewaffnete Rathäuser zerstören",
	}
	StatWallsDestroyed = &ComparableStatistic{
		Name:        goclash.AchievementWallBuster.Name,
		DisplayName: "Mauern zerstört",
		Task:        "Mauern zerstören",
	}
	StatXBowsDestroyed = &ComparableStatistic{
		Name:        goclash.AchievementXBowExterminator.Name,
		DisplayName: "X-Bögen zerstört",
		Task:        "X-Bögen zerstören",
	}
	StatInfernoTowersDestroyed = &ComparableStatistic{
		Name:        goclash.AchievementFirefighter.Name,
		DisplayName: "Infernotürme zerstört",
		Task:        "Infernotürme zerstören",
	}
	StatEagleArtilleriesDestroyed = &ComparableStatistic{
		Name:        goclash.AchievementAntiArtillery.Name,
		DisplayName: "Adlerartillerien zerstört",
		Task:        "Adlerartillerien zerstören",
	}
	StatScattershotsDestroyed = &ComparableStatistic{
		Name:        goclash.AchievementShatteredAndScattered.Name,
		DisplayName: "Scattershots zerstört",
		Task:        "Scattershots zerstören",
	}
	StatHighestBuilderBaseTrophies = &ComparableStatistic{
		Name:        goclash.AchievementChampionBuilder.Name,
		DisplayName: "Höchste Trophäen (BB)",
		Task:        "Höchste Trophäen (BB)",
	}
	StatSpellTowersDestroyed = &ComparableStatistic{
		Name:        goclash.AchievementCounterspell.Name,
		DisplayName: "Zaubertürme zerstört",
		Task:        "Zaubertürme zerstören",
	}
	StatMonolithsDestroyed = &ComparableStatistic{
		Name:        goclash.AchievementMonolithMasher.Name,
		DisplayName: "Monolithen zerstört",
		Task:        "Monolithen zerstören",
	}
	StatClanGoldLooted = &ComparableStatistic{
		Name:        goclash.AchievementAggressiveCapitalism.Name,
		DisplayName: "Clan Gold aus Raids",
		Task:        "Clan Gold aus Raids sammeln",
	}
	StatClanGoldContributed = &ComparableStatistic{
		Name:        goclash.AchievementMostValuableClanmate.Name,
		DisplayName: "Clan Gold ausgegeben",
		Task:        "Clan Gold ausgeben",
	}
)

// ComparableStats is a list of all comparable stats a coc.Player has. This includes achievements and other stats.
var ComparableStats = []*ComparableStatistic{
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
	StatWallsDestroyed,
	StatXBowsDestroyed,
	StatInfernoTowersDestroyed,
	StatEagleArtilleriesDestroyed,
	StatScattershotsDestroyed,
	StatHighestBuilderBaseTrophies,
	StatSpellTowersDestroyed,
	StatMonolithsDestroyed,
	StatClanGoldLooted,
	StatClanGoldContributed,
}

// ComparableAchievements are the achievements that can be compared between players.
var ComparableAchievements = []*ComparableStatistic{
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
	StatWallsDestroyed,
	StatXBowsDestroyed,
	StatInfernoTowersDestroyed,
	StatEagleArtilleriesDestroyed,
	StatScattershotsDestroyed,
	StatHighestBuilderBaseTrophies,
	StatSpellTowersDestroyed,
	StatMonolithsDestroyed,
	StatClanGoldLooted,
	StatClanGoldContributed,
}
