package messages

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bot/commands/util"
	"bot/store/postgres/models"
)

func SendClanMembers(s *discordgo.Session, i *discordgo.InteractionCreate, clan *models.Clan) {
	membersByRole := util.SortMembersByRole(clan.Members)

	var fields []*discordgo.MessageEmbedField
	for _, members := range membersByRole {
		if len(members) == 0 {
			continue
		}

		field := &discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s (%d)", members[0].ClanRole.String(), len(members)),
		}
		for _, member := range members {
			field.Value += fmt.Sprintf("%s\n", member.Player.Name)
		}
		fields = append(fields, field)
	}

	SendEmbed(s, i, NewFieldEmbed(
		fmt.Sprintf("Mitglieder von %s", clan.Name),
		fmt.Sprintf("%s hat momentan %d Mitglieder.", clan.Name, len(clan.Members)),
		ColorAqua,
		fields,
	))
}
