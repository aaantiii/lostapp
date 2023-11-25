package models

import (
	"github.com/amaanq/coc.go"
)

type PlayerStats struct {
	Tag                   string `gorm:"primaryKey;not null;size:12"`
	SeasonWins            int    `gorm:"not null"`
	WellSeasoned          int    `gorm:"not null"`
	NiceAndTidy           int    `gorm:"not null"`
	FriendInNeed          int    `gorm:"not null"`
	SharingIsCaring       int    `gorm:"not null"`
	SiegeSharer           int    `gorm:"not null"`
	WarHero               int    `gorm:"not null"`
	WarLeagueLegend       int    `gorm:"not null"`
	GamesChampion         int    `gorm:"not null"`
	Unbreakable           int    `gorm:"not null"`
	SweetVictory          int    `gorm:"not null"`
	Conqueror             int    `gorm:"not null"`
	Humiliator            int    `gorm:"not null"`
	NotSoEasyThisTime     int    `gorm:"not null"`
	UnionBuster           int    `gorm:"not null"`
	BustThis              int    `gorm:"not null"`
	WallBuster            int    `gorm:"not null"`
	MortarMauler          int    `gorm:"not null"`
	XBowExterminator      int    `gorm:"not null"`
	Firefighter           int    `gorm:"not null"`
	AntiArtillery         int    `gorm:"not null"`
	ShatteredAndScattered int    `gorm:"not null"`
	SuperbWork            int    `gorm:"not null"`
	ChampionBuilder       int    `gorm:"not null"`
	Counterspell          int    `gorm:"not null"`
	MonolithMasher        int    `gorm:"not null"`
	AggressiveCapitalism  int    `gorm:"not null"`
	MostValuableClanmate  int    `gorm:"not null"`
}

func (*PlayerStats) TableName() string {
	return "lostapp_player_stats_live"
}

func NewPlayerStats(player *coc.Player) *PlayerStats {
	achievements := make(map[string]int)
	for _, achievement := range player.Achievements {
		achievements[achievement.Name] = achievement.Value
	}

	return &PlayerStats{
		Tag:                   player.Tag,
		SeasonWins:            player.AttackWins,
		WellSeasoned:          achievements[coc.WellSeasoned.Name],
		NiceAndTidy:           achievements[coc.NiceAndTidy.Name],
		FriendInNeed:          achievements[coc.FriendInNeed.Name],
		SharingIsCaring:       achievements[coc.SharingIsCaring.Name],
		SiegeSharer:           achievements[coc.SiegeSharer.Name],
		WarHero:               achievements[coc.WarHero.Name],
		WarLeagueLegend:       achievements[coc.WarLeagueLegend.Name],
		GamesChampion:         achievements[coc.GamesChampion.Name],
		Unbreakable:           achievements[coc.Unbreakable.Name],
		SweetVictory:          achievements[coc.SweetVictory.Name],
		Conqueror:             achievements[coc.Conqueror.Name],
		Humiliator:            achievements[coc.Humiliator.Name],
		NotSoEasyThisTime:     achievements[coc.NotSoEasyThisTime.Name],
		UnionBuster:           achievements[coc.UnionBuster.Name],
		BustThis:              achievements[coc.BustThis.Name],
		WallBuster:            achievements[coc.WallBuster.Name],
		MortarMauler:          achievements[coc.MortarMauler.Name],
		XBowExterminator:      achievements[coc.XBowExterminator.Name],
		Firefighter:           achievements[coc.Firefighter.Name],
		AntiArtillery:         achievements[coc.AntiArtillery.Name],
		ShatteredAndScattered: achievements[coc.ShatteredAndScattered.Name],
		SuperbWork:            achievements[coc.SuperbWork.Name],
		ChampionBuilder:       achievements[coc.ChampionBuilder.Name],
		Counterspell:          achievements[coc.Counterspell.Name],
		MonolithMasher:        achievements[coc.MonolithMasher.Name],
		AggressiveCapitalism:  achievements[coc.AggressiveCapitalism.Name],
		MostValuableClanmate:  achievements[coc.MostValuableClanmate.Name],
	}
}
