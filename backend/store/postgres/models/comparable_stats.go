package models

import (
	"errors"

	"github.com/amaanq/coc.go"
	"gorm.io/gorm"
)

// ComparableStats is a coc.Player statistic that can be compared between players.
type ComparableStats struct {
	ID            uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name          string `gorm:"not null;size:64" json:"name"`
	DisplayName   string `gorm:"not null;size:64" json:"displayName"`
	IsAchievement bool   `gorm:"not null" json:"-"`
}

func (*ComparableStats) TableName() string {
	return "lostapp_comparable_stats"
}

func SeedComparableStats(db *gorm.DB) error {
	if db.Migrator().HasTable(&ComparableStats{}) {
		if err := db.First(&ComparableStats{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
	}

	return db.Create(comparableStatsSeedData()).Error
}

func comparableStatsSeedData() []*ComparableStats {
	return []*ComparableStats{{
		Name:          "Season Wins",
		DisplayName:   "Angriffe gewonnen (Saison)",
		IsAchievement: false,
	}, {
		Name:          coc.HeroicHeist.Name,
		DisplayName:   "Dunkles Elixier gesammelt",
		IsAchievement: true,
	}, {
		Name:          coc.WellSeasoned.Name,
		DisplayName:   "Season Challenge Punkte",
		IsAchievement: true,
	}, {
		Name:          coc.NiceAndTidy.Name,
		DisplayName:   "Hindernisse entfernt",
		IsAchievement: true,
	}, {
		Name:          coc.FriendInNeed.Name,
		DisplayName:   "Truppen gespendet",
		IsAchievement: true,
	}, {
		Name:          coc.SharingIsCaring.Name,
		DisplayName:   "Zauber gespendet",
		IsAchievement: true,
	}, {
		Name:          coc.SiegeSharer.Name,
		DisplayName:   "Siege Machines gespendet",
		IsAchievement: true,
	}, {
		Name:          coc.WarHero.Name,
		DisplayName:   "CW Sterne",
		IsAchievement: true,
	}, {
		Name:          coc.WarLeagueLegend.Name,
		DisplayName:   "CWL Sterne",
		IsAchievement: true,
	}, {
		Name:          coc.GamesChampion.Name,
		DisplayName:   "Clan Games Punkte",
		IsAchievement: true,
	}, {
		Name:          coc.Unbreakable.Name,
		DisplayName:   "Erfolgreiche Verteidigungen",
		IsAchievement: true,
	}, {
		Name:          coc.SweetVictory.Name,
		DisplayName:   "Meiste Trophäen",
		IsAchievement: true,
	}, {
		Name:          coc.Conqueror.Name,
		DisplayName:   "Angriffe gewonnen",
		IsAchievement: true,
	}, {
		Name:          coc.Humiliator.Name,
		DisplayName:   "Rathäuser zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.NotSoEasyThisTime.Name,
		DisplayName:   "Bewaffnete Rathäuser zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.UnionBuster.Name,
		DisplayName:   "Bauhütten zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.BustThis.Name,
		DisplayName:   "Bewaffnete Bauhütten zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.WallBuster.Name,
		DisplayName:   "Mauern zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.MortarMauler.Name,
		DisplayName:   "Minenwerfer zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.XBowExterminator.Name,
		DisplayName:   "X-Bögen zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.Firefighter.Name,
		DisplayName:   "Infernotürme zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.AntiArtillery.Name,
		DisplayName:   "Adlerartillerien zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.ShatteredAndScattered.Name,
		DisplayName:   "Scattershots zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.SuperbWork.Name,
		DisplayName:   "Supertruppen geboosted",
		IsAchievement: true,
	}, {
		Name:          coc.ChampionBuilder.Name,
		DisplayName:   "Meiste Trophäen (BB)",
		IsAchievement: true,
	}, {
		Name:          coc.Counterspell.Name,
		DisplayName:   "Zaubertürme zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.MonolithMasher.Name,
		DisplayName:   "Monolithen zerstört",
		IsAchievement: true,
	}, {
		Name:          coc.AggressiveCapitalism.Name,
		DisplayName:   "Clan Gold aus Raids",
		IsAchievement: true,
	}, {
		Name:          coc.MostValuableClanmate.Name,
		DisplayName:   "Clan Gold ausgegeben",
		IsAchievement: true,
	}}
}
