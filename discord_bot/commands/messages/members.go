package messages

import (
	"fmt"
	"sort"
	"strings"

	"github.com/aaantiii/goclash"
	"github.com/bwmarrin/discordgo"

	"bot/commands/util"
	"bot/store/postgres/models"
)

func SendClanMembers(i *discordgo.InteractionCreate, clan *models.Clan) {
	sort.SliceStable(clan.ClanMembers, func(i, j int) bool {
		return strings.ToLower(clan.ClanMembers[i].Player.Name) < strings.ToLower(clan.ClanMembers[j].Player.Name)
	})
	membersByRole := util.GroupMembersByRole(clan.ClanMembers)

	var fields []*discordgo.MessageEmbedField
	for _, members := range membersByRole {
		if len(members) == 0 {
			continue
		}

		field := &discordgo.MessageEmbedField{Name: fmt.Sprintf("%s (%d)", members[0].ClanRole.Format(), len(members))}
		for _, member := range members {
			field.Value += fmt.Sprintf("%s\n", member.Player.Name)
		}
		field.Value += " "
		fields = append(fields, field)
	}

	SendEmbedResponse(i, NewFieldEmbed(
		fmt.Sprintf("Mitglieder von %s", clan.Name),
		fmt.Sprintf("%s hat momentan %d Mitglieder.", clan.Name, len(clan.ClanMembers)),
		ColorAqua,
		fields,
	))
}

func SendClansMembersStatus(i *discordgo.InteractionCreate, dbMembers models.ClanMembers, clan *goclash.Clan) {
	if len(clan.MemberList) == 0 {
		SendCocApiErr(i, nil)
		return
	}

	SendEmbedResponse(i, NewFieldEmbed(
		fmt.Sprintf("Mitgliederstatus von %s", clan.Name),
		"Übersicht aller Mitglieder, welche sich gerade nicht im Clan befinden, sowie nicht hinzugefügte Mitglieder, welche gerade im Clan sind.",
		ColorAqua,
		[]*discordgo.MessageEmbedField{getUnverifiedMembers(dbMembers, clan.MemberList), getMembersNotInClan(dbMembers, clan.MemberList)},
	))
}

// getUnverifiedMembers returns all members that are currently in the clan but not in the database.
func getUnverifiedMembers(dbMembers models.ClanMembers, currentMembers []goclash.ClanMember) *discordgo.MessageEmbedField {
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
func getMembersNotInClan(members models.ClanMembers, currentMembers []goclash.ClanMember) *discordgo.MessageEmbedField {
	currentMembersByTag := make(map[string]goclash.ClanMember, len(currentMembers))
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
