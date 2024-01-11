package commands

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/client"
	"bot/commands/handlers"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/types"
)

func clanInteractionCommands(db *gorm.DB, cocClient *client.CocClient) types.Commands[types.InteractionHandler] {
	handler := handlers.NewClanHandler(repos.NewClansRepo(db), cocClient)
	return types.Commands[types.InteractionHandler]{{
		Handler: types.InteractionHandler{
			Main:         handler.ClanStats,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "clanstats",
			Description:  "Statistiken der Mitglieder eines Clans anzeigen.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, dessen Statistiken angezeigt werden sollen."),
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "statistic",
					Description: "Statistik, die angezeigt werden soll.",
					Required:    true,
					Choices:     types.ComparableStatisticChoices(types.ComparableStats),
				},
			},
		},
	}}
}
