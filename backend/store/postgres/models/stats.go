package models

import "time"

type Stats struct {
	CocTag                              string
	Timestamp                           time.Time
	AttackWins                          int
	BestTrophies                        int
	BestVersusTrophies                  int
	BuilderHallLevel                    int
	ClanTag                             string
	TotalClanCapitalContribution        int
	DefenseWins                         int
	TroopCapacityDonated                int
	ExpLevel                            int
	BarbarianKingLevel                  int
	ArcherQueenLevel                    int
	GrandWardenLevel                    int
	RoyalChampionLevel                  int
	BattleMachineLevel                  int
	Name                                string
	TroopCapacityReceived               int
	Role                                string
	TownHallLevel                       int
	TownHallWeaponLevel                 int
	Trophies                            int
	VersusBattleWins                    int
	VersusTrophies                      int
	WarOptedIn                          bool
	TotalWarStars                       int
	TotalGoldLooted                     int
	TotalElixirLooted                   int
	TotalDarkElixirLooted               int
	TotalSeasonChallengePoints          int
	ObstaclesCleared                    int
	ClanCastleLevel                     int
	ClanCastleGoldCollected             int
	TotalTroopCapacityDonated           int
	TotalSpellCapacityDonated           int
	TotalSiegeMachinesDonated           int
	TotalCwlStars                       int
	TotalClanGamesPoints                int
	TotalDefenseWins                    int
	TotalAttackWins                     int
	TotalTownHallsDestroyed             int
	TotalWeaponizedTownHallsDestroyed   int
	TotalBuilderHutsDestroyed           int
	TotalWeaponizedBuilderHutsDestroyed int
	TotalWallsDestroyed                 int
	TotalMortarsDestroyed               int
	TotalXBowsDestroyed                 int
	TotalInfernosDestroyed              int
	TotalEagleArtilleriesDestroyed      int
	TotalScattershotsDestroyed          int
	CampaignStars                       int
	TotalSuperTroopBoosts               int
	TotalBuilderHallsDestroyed          int
	TotalClanCapitalLooted              int
}

func (*Stats) TableName() string {
	return "stats"
}
