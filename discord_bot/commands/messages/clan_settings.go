package messages

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/bwmarrin/discordgo"

	"bot/store/postgres/models"
)

func SendKickpointInfo(i *discordgo.InteractionCreate, settings *models.ClanSettings, reasons []*models.KickpointReason) {
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "Grund"},
			{Align: simpletable.AlignRight, Text: "Kickpunḱte"},
		},
	}

	for _, reason := range reasons {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: reason.Name},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", reason.Amount)},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	SendEmbedResponse(i, NewFieldEmbed(
		fmt.Sprintf("Einstellungen von %s", settings.Clan.Name),
		"```\n"+table.String()+"\n```",
		ColorAqua,
		[]*discordgo.MessageEmbedField{
			{
				Name:   "Gültigkeitsdauer von Kickpunkten",
				Value:  fmt.Sprintf("%d Tage", settings.KickpointsExpireAfterDays),
				Inline: true,
			},
			{
				Name:   "Maximale Anzahl an Kickpunkten",
				Value:  fmt.Sprintf("%d Kickpunkte", settings.MaxKickpoints),
				Inline: true,
			},
			{
				Name:   "Mindestanzahl an Season-Wins",
				Value:  fmt.Sprintf("%d Wins", settings.MinSeasonWins),
				Inline: true,
			},
		},
	))
}

func kickpointAmountField(name string, value int) *discordgo.MessageEmbedField {
	return &discordgo.MessageEmbedField{
		Name:  name,
		Value: fmt.Sprintf("%d Kickpunkte", value),
	}
}
