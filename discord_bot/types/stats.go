package types

import (
	"github.com/amaanq/coc.go"
	"github.com/bwmarrin/discordgo"
)

type PlayerStatistic struct {
	Tag       string
	Name      string
	Placement int
	Value     int
}

// ComparableStatistic is a coc.Player statistic that can be compared between players.
type ComparableStatistic struct {
	Name        string
	DisplayName string
}

func ComparableStatisticChoices(stats []ComparableStatistic) []*discordgo.ApplicationCommandOptionChoice {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, len(stats))
	for i, stat := range stats {
		choices[i] = &discordgo.ApplicationCommandOptionChoice{
			Name:  stat.DisplayName,
			Value: stat.Name,
		}
	}

	return choices
}

var (
	StatSeasonWins = ComparableStatistic{
		Name:        "Season Wins",
		DisplayName: "Angriffe gewonnen (Saison)",
	}
	StatDarkElixirLooted = ComparableStatistic{
		Name:        coc.HeroicHeist.Name,
		DisplayName: "Dunkles Elixier gesammelt",
	}
	StatSeasonPoints = ComparableStatistic{
		Name:        coc.WellSeasoned.Name,
		DisplayName: "Season Challenge Punkte",
	}
	StatObstaclesRemoved = ComparableStatistic{
		Name:        coc.NiceAndTidy.Name,
		DisplayName: "Hindernisse entfernt",
	}
	StatTroopsDonated = ComparableStatistic{
		Name:        coc.FriendInNeed.Name,
		DisplayName: "Truppen gespendet",
	}
	StatSpellsDonated = ComparableStatistic{
		Name:        coc.SharingIsCaring.Name,
		DisplayName: "Zauber gespendet",
	}
	StatSiegeMachinesDonated = ComparableStatistic{
		Name:        coc.SiegeSharer.Name,
		DisplayName: "Siege Machines gespendet",
	}
	StatWarStars = ComparableStatistic{
		Name:        coc.WarHero.Name,
		DisplayName: "CW Sterne",
	}
	StatCWLStars = ComparableStatistic{
		Name:        coc.WarLeagueLegend.Name,
		DisplayName: "CWL Sterne",
	}
	StatClanGamesPoints = ComparableStatistic{
		Name:        coc.GamesChampion.Name,
		DisplayName: "Clan Games Punkte",
	}
	StatSuccessfulDefenses = ComparableStatistic{
		Name:        coc.Unbreakable.Name,
		DisplayName: "Erfolgreiche Verteidigungen",
	}
	StatHighestTrophies = ComparableStatistic{
		Name:        coc.SweetVictory.Name,
		DisplayName: "Meiste Trophäen",
	}
	StatAttackWins = ComparableStatistic{
		Name:        coc.Conqueror.Name,
		DisplayName: "Angriffe gewonnen",
	}
	StatTownHallsDestroyed = ComparableStatistic{
		Name:        coc.Humiliator.Name,
		DisplayName: "Rathäuser zerstört",
	}
	StatWeaponizedTownHallsDestroyed = ComparableStatistic{
		Name:        coc.NotSoEasyThisTime.Name,
		DisplayName: "Bewaffnete Rathäuser zerstört",
	}
	StatWallsDestroyed = ComparableStatistic{
		Name:        coc.WallBuster.Name,
		DisplayName: "Mauern zerstört",
	}
	StatXBowsDestroyed = ComparableStatistic{
		Name:        coc.XBowExterminator.Name,
		DisplayName: "X-Bögen zerstört",
	}
	StatInfernoTowersDestroyed = ComparableStatistic{
		Name:        coc.Firefighter.Name,
		DisplayName: "Infernotürme zerstört",
	}
	StatEagleArtilleriesDestroyed = ComparableStatistic{
		Name:        coc.AntiArtillery.Name,
		DisplayName: "Adlerartillerien zerstört",
	}
	StatScattershotsDestroyed = ComparableStatistic{
		Name:        coc.ShatteredAndScattered.Name,
		DisplayName: "Scattershots zerstört",
	}
	StatHighestBuilderBaseTrophies = ComparableStatistic{
		Name:        coc.ChampionBuilder.Name,
		DisplayName: "Meiste Trophäen (BB)",
	}
	StatSpellTowersDestroyed = ComparableStatistic{
		Name:        coc.Counterspell.Name,
		DisplayName: "Zaubertürme zerstört",
	}
	StatMonolithsDestroyed = ComparableStatistic{
		Name:        coc.MonolithMasher.Name,
		DisplayName: "Monolithen zerstört",
	}
	StatClanGoldLooted = ComparableStatistic{
		Name:        coc.AggressiveCapitalism.Name,
		DisplayName: "Clan Gold aus Raids",
	}
	StatClanGoldContributed = ComparableStatistic{
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
