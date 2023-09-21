package types

import (
	"strings"

	"github.com/amaanq/coc.go"
)

type Player struct {
	ComparableStatsByName map[string]int `json:"-"` // maps a stat name to value
	DiscordID             string         `json:"discordId"`
	Clans                 []PlayerClan   `json:"clans"`
	Name                  string         `json:"name"`
	Tag                   string         `json:"tag"`
	WarPreference         string         `json:"warPreference"`
	Trophies              int            `json:"trophies"`
	ExpLevel              int            `json:"expLevel"`
	TownHallLevel         int            `json:"townHallLevel"`
	AttackWins            int            `json:"attackWins"`
	DefenseWins           int            `json:"defenseWins"`
}

func NewPlayer(player *coc.Player) *Player {
	if player == nil {
		return nil
	}

	return &Player{
		Name:          player.Name,
		Tag:           player.Tag,
		WarPreference: player.WarPreference,
		Trophies:      player.Trophies,
		ExpLevel:      player.ExpLevel,
		TownHallLevel: player.TownHallLevel,
		AttackWins:    player.AttackWins,
		DefenseWins:   player.DefenseWins,
	}
}

type PlayerClan struct {
	Name string   `json:"name"`
	Tag  string   `json:"tag"`
	Role coc.Role `json:"role"`
}

type Players []*Player

type PlayerStatistic struct {
	PlayerName string `json:"playerName"`
	PlayerTag  string `json:"playerTag"`
	ClanNames  string `json:"clanNames"`
	Value      int    `json:"value"`
}

func NewPlayerStatistic(player *Player, value int) *PlayerStatistic {
	statistic := &PlayerStatistic{
		PlayerName: player.Name,
		PlayerTag:  player.Tag,
		Value:      value,
	}

	clanNames := make([]string, len(player.Clans))
	for i, clan := range player.Clans {
		clanNames[i] = clan.Name
	}
	statistic.ClanNames = strings.Join(clanNames, ", ")

	return statistic
}

type PlayersParams struct {
	PaginationParams

	Name      string `form:"name" binding:"omitempty,min=3,max=30"`
	Tag       string `form:"tag" binding:"omitempty,min=3,max=9"`
	ClanName  string `form:"clanName" binding:"omitempty,min=3,max=30"`
	ClanTag   string `form:"clanTag" binding:"omitempty,min=3,max=9"`
	DiscordID string `form:"discordID" binding:"omitempty,len=18"`
}
