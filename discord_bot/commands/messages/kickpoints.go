package messages

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"

	"bot/commands/util"
	"bot/store/postgres/models"
	"bot/types"
)

func SendClanKickpoints(s *discordgo.Session, i *discordgo.InteractionCreate, clanName string, members []*types.ClanMemberKickpoints) {
	desc := "" // send in description because of embed field limit
	for _, m := range members {
		desc += fmt.Sprintf("%s (%s): %d Kickpunkte\n\n", m.Name, m.Tag, m.Amount)
	}

	SendEmbed(s, i, NewEmbed(
		fmt.Sprintf("Kickpunkte von %s", clanName),
		desc,
		ColorAqua,
	))
}

func SendMemberKickpoints(s *discordgo.Session, i *discordgo.InteractionCreate, clanName string, kickpoints []*models.Kickpoint) {
	fields := make([]*discordgo.MessageEmbedField, len(kickpoints))
	for index, k := range kickpoints {
		field := &discordgo.MessageEmbedField{
			Name: fmt.Sprintf("Kickpunkt #%d", k.ID),
		}

		for _, f := range DetailedKickpointFields(k) {
			field.Value += fmt.Sprintf("%s: %s\n", f.Name, f.Value)
		}

		fields[index] = field
	}

	player := kickpoints[0].Player
	SendEmbed(s, i, NewFieldEmbed(
		"Aktive Kickpunkte",
		fmt.Sprintf("Aktive Kickpunkte von %s (%s) in %s", player.Name, player.CocTag, clanName),
		ColorAqua,
		fields,
	))
}

func DetailedKickpointFields(kickpoint *models.Kickpoint) []*discordgo.MessageEmbedField {
	activeTime := time.Until(kickpoint.Date)
	activeLabel := "Aktiv in"
	if activeTime < 0 {
		activeTime = -activeTime
		activeLabel = "Aktiv seit"
	}

	fields := []*discordgo.MessageEmbedField{
		{Name: "Grund", Value: kickpoint.Description},
		{Name: "Anzahl Kickpunkte", Value: strconv.Itoa(kickpoint.Amount)},
		{Name: "Erhalten am", Value: util.FormatDate(kickpoint.Date)},
		{Name: activeLabel, Value: util.FormatDuration(activeTime)},
	}

	createdAtField := &discordgo.MessageEmbedField{Name: "Erstellt"}
	if kickpoint.CreatedByUser != nil {
		createdAtField.Value = fmt.Sprintf("von %s am ", kickpoint.CreatedByUser.Name)
	}
	createdAtField.Value += util.FormatDateTime(kickpoint.CreatedAt)

	updatedAtField := &discordgo.MessageEmbedField{Name: "Zuletzt bearbeitet"}
	if kickpoint.UpdatedByUser != nil {
		updatedAtField.Value = fmt.Sprintf("von %s am ", kickpoint.UpdatedByUser.Name)
	}
	updatedAtField.Value += util.FormatDateTime(kickpoint.UpdatedAt)

	fields = append(fields, createdAtField, updatedAtField)
	return fields
}
