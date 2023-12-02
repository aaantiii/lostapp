package commands

import (
	"github.com/amaanq/coc.go"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/handlers"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/types"
)

func memberInteractionCommands(db *gorm.DB) types.Commands[types.InteractionHandler] {
	handler := handlers.NewMemberHandler(
		repos.NewMembersRepo(db),
		repos.NewClansRepo(db),
		repos.NewPlayersRepo(db),
	)

	return types.Commands[types.InteractionHandler]{{
		Handler: types.InteractionHandler{
			Main:         handler.ListMembers,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "listmembers",
			Description:  "Listet alle Mitglieder eines Clans auf.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, dessen Mitglieder aufgelistet werden sollen."),
			},
		},
	}, {
		Handler: types.InteractionHandler{
			Main:         handler.AddMember,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "addmember",
			Description:  "Fügt ein Mitglied zu einem Clan hinzu.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, zu dem das Mitglied hinzugefügt werden soll."),
				optionPlayerTag("Mitglied, das zum Clan hinzugefügt werden soll."),
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "role",
					Description: "Rolle, die das Mitglied im Clan haben soll.",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: coc.Leader.String(), Value: coc.Leader},
						{Name: coc.CoLeader.String(), Value: coc.CoLeader},
						{Name: coc.Elder.String(), Value: coc.Elder},
						{Name: coc.Member.String(), Value: coc.Member},
					},
				},
			},
		},
	}, {
		Handler: types.InteractionHandler{
			Main:         handler.EditMember,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "editmember",
			Description:  "Ändere die Rolle eines Mitglieds in einem Clan.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, zu dem das Mitglied gehört."),
				optionPlayerTag("Mitglied, dessen Rolle bearbeitet werden soll."),
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "role",
					Description: "Rolle, die das Mitglied im Clan haben soll.",
					Required:    true,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{Name: coc.Leader.String(), Value: coc.Leader},
						{Name: coc.CoLeader.String(), Value: coc.CoLeader},
						{Name: coc.Elder.String(), Value: coc.Elder},
						{Name: coc.Member.String(), Value: coc.Member},
					},
				},
			},
		},
	}, {
		Handler: types.InteractionHandler{
			Main:         handler.RemoveMember,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "removemember",
			Description:  "Mitglied von einem Clan entfernen.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.OptionalBool(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, von dem das Mitglied entfernt werden soll."),
				optionPlayerTag("Mitglied, das vom Clan entfernt werden soll."),
			},
		},
	}}
}
