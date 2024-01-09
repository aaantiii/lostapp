package messages

import (
	"fmt"

	"github.com/amaanq/coc.go"
	"github.com/bwmarrin/discordgo"

	"bot/commands/util"
	"bot/store/postgres/models"
)

func SendClanMembers(s *discordgo.Session, i *discordgo.InteractionCreate, clan *models.Clan) {
	membersByRole := util.SortMembersByRole(clan.ClanMembers)

	var fields []*discordgo.MessageEmbedField
	for _, members := range membersByRole {
		if len(members) == 0 {
			continue
		}

		field := &discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s (%d)", members[0].ClanRole.Format(), len(members)),
		}
		for _, member := range members {
			field.Value += fmt.Sprintf("%s\n", member.Player.Name)
		}
		fields = append(fields, field)
	}

	SendEmbed(s, i, NewFieldEmbed(
		fmt.Sprintf("Mitglieder von %s", clan.Name),
		fmt.Sprintf("%s hat momentan %d Mitglieder.", clan.Name, len(clan.ClanMembers)),
		ColorAqua,
		fields,
	))
}

func SendClansMembersStatus(s *discordgo.Session, i *discordgo.InteractionCreate, dbMembers models.ClanMembers, clan *coc.Clan) {
	if clan.MemberList == nil || len(*clan.MemberList) == 0 {
		SendCocApiError(s, i)
		return
	}

	SendEmbed(s, i, NewFieldEmbed(
		fmt.Sprintf("Mitgliederstatus von %s", clan.Name),
		"Übersicht des Verifizerungsstatus aller Mitglieder.",
		ColorAqua,
		[]*discordgo.MessageEmbedField{getUnverifiedMembers(dbMembers, *clan.MemberList), getMembersNotInClan(dbMembers, *clan.MemberList)},
	))
}

// getUnverifiedMembers returns all members that are currently in the clan but not in the database.
func getUnverifiedMembers(dbMembers models.ClanMembers, currentMembers []coc.ClanMember) *discordgo.MessageEmbedField {
	dbMemberByTag := make(map[string]*models.ClanMember, len(dbMembers))
	for _, member := range dbMembers {
		dbMemberByTag[member.PlayerTag] = member
	}

	field := &discordgo.MessageEmbedField{Name: "Kein Mitglied, ingame im Clan"}
	for _, member := range currentMembers {
		if _, ok := dbMemberByTag[member.Tag]; !ok {
			field.Value += fmt.Sprintf("%s (%s)\n", member.Name, member.Tag)
		}
	}

	if field.Value == "" {
		field.Value = "Alle Personen im Clan sind als Mitglied hinzugefügt."
	}

	return field
}

// getMembersNotInClan returns all members that are in the database but currently not in the clan.
func getMembersNotInClan(members models.ClanMembers, currentMembers []coc.ClanMember) *discordgo.MessageEmbedField {
	currentMembersByTag := make(map[string]coc.ClanMember, len(currentMembers))
	for _, m := range currentMembers {
		currentMembersByTag[m.Tag] = m
	}

	field := &discordgo.MessageEmbedField{Name: "Mitglied, gerade nicht im Clan"}
	for _, member := range members {
		if _, ok := currentMembersByTag[member.PlayerTag]; !ok {
			field.Value += fmt.Sprintf("%s (%s)\n", member.Player.Name, member.PlayerTag)
		}
	}

	if field.Value == "" {
		field.Value = "Alle Mitglieder befinden sich momentan im Clan."
	}

	return field
}
