package types

import "github.com/amaanq/coc.go"

type Clan struct {
	MemberList     []ClanMember     `json:"memberList,omitempty"`
	WarTies        *int             `json:"warTies,omitempty"`
	Description    *string          `json:"description,omitempty"`
	WarLosses      *int             `json:"warLosses,omitempty"`
	ClanCapital    *coc.ClanCapital `json:"clanCapital,omitempty"`
	WarLeague      WarLeague        `json:"warLeague"`
	BadgeURL       string           `json:"badgeUrl"`
	Tag            string           `json:"tag"`
	Name           string           `json:"name"`
	Level          int              `json:"level"`
	WarWins        int              `json:"warWins"`
	Members        int              `json:"members"`
	IsWarLogPublic bool             `json:"isWarLogPublic"`
}

func NewClan(clan *coc.Clan, memberList []ClanMember) *Clan {
	return &Clan{
		MemberList:  memberList,
		WarTies:     clan.WarTies,
		Description: clan.Description,
		WarLosses:   clan.WarLosses,
		ClanCapital: clan.ClanCapital,
		WarLeague: WarLeague{
			Name:    clan.WarLeague.Name,
			ID:      clan.WarLeague.ID,
			IconURL: clan.WarLeague.IconUrls.Medium,
		},
		BadgeURL:       clan.BadgeURLs.Medium,
		Tag:            clan.Tag,
		Name:           clan.Name,
		Level:          clan.ClanLevel,
		WarWins:        clan.WarWins,
		Members:        len(memberList),
		IsWarLogPublic: clan.IsWarLogPublic,
	}
}

type ClanMember struct {
	Tag      string   `json:"tag"`
	Name     string   `json:"name"`
	Role     coc.Role `json:"role"`
	ExpLevel int      `json:"expLevel"`
}

func NewClanMember(player *coc.Player, role coc.Role) ClanMember {
	return ClanMember{
		Tag:      player.Tag,
		Name:     player.Name,
		Role:     role,
		ExpLevel: player.ExpLevel,
	}
}

type ClanListItem struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type ClanList []ClanListItem

func NewClanList(clans []*Clan) ClanList {
	clanList := make(ClanList, len(clans))
	for i, clan := range clans {
		clanList[i] = ClanListItem{
			Name: clan.Name,
			Tag:  clan.Tag,
		}
	}

	return clanList
}

type WarLeague struct {
	Name    string `json:"name,omitempty"`
	ID      int    `json:"id,omitempty"`
	IconURL string `json:"iconUrl,omitempty"`
}
