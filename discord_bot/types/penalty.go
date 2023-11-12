package types

import "bot/store/postgres/models"

type MemberPenaltiesList struct {
	Member          *models.Member
	ActivePenalties []*models.Penalty
}

type ClanPenaltiesList struct {
	Clan    *models.LostClans
	Members []*MemberPenaltiesList
}
