package commands

import (
	"github.com/aaantiii/goclash"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/handlers"
	"bot/commands/middleware"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/types"
)

func clanInteractionCommands(db *gorm.DB, clashClient *goclash.Client) types.Commands[types.InteractionHandler] {
	handler := handlers.NewClanHandler(
		repos.NewClansRepo(db),
		repos.NewMembersRepo(db),
		repos.NewClanEventsRepo(db),
		middleware.NewAuthMiddleware(repos.NewGuildsRepo(db), repos.NewClansRepo(db), repos.NewUsersRepo(db)),
		clashClient,
	)
	return types.Commands[types.InteractionHandler]{{
		Handler: types.InteractionHandler{
			Main:         handler.ClanStats,
			Autocomplete: handler.HandleAutocomplete,
		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:         "clanstats",
			Description:  "Statistiken der Mitglieder eines Clans anzeigen.",
			Type:         discordgo.ChatApplicationCommand,
			DMPermission: util.BoolPtr(false),
			Options: []*discordgo.ApplicationCommandOption{
				optionClanTag("Clan, dessen Statistiken angezeigt werden sollen."),
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        handlers.StatisticOptionName,
					Description: "Statistik, die angezeigt werden soll.",
					Required:    true,
					Choices:     types.ComparableStatisticChoices(types.ComparableStats),
				},
			},
		},
	},
		{
			Handler: types.InteractionHandler{
				Main:         handler.CWDonator,
				Autocomplete: handler.HandleAutocomplete,
			},
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:         "cwdonator",
				Description:  "Teilt automatisch zufällige Spender für den aktuellen Clan War ein. Einteilung basiert auf Clan War größe.",
				Type:         discordgo.ChatApplicationCommand,
				DMPermission: util.BoolPtr(false),
				Options: []*discordgo.ApplicationCommandOption{
					optionClanTag("Clan, für den die Spender eingeteilt werden sollen."),
				},
			},
		},
		{
			Handler: types.InteractionHandler{
				Main:         handler.RaidPing,
				Autocomplete: handler.HandleAutocomplete,
			},
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:         "raidping",
				Description:  "Pingt alle Mitglieder, die noch fehlende Raid Angriffe haben.",
				Type:         discordgo.ChatApplicationCommand,
				DMPermission: util.BoolPtr(false),
				Options: []*discordgo.ApplicationCommandOption{
					optionClanTag("Clan, dessen Mitglieder einen Ping erhalten sollen."),
				},
			},
		}, {
			Handler: types.InteractionHandler{
				Main: handler.EventInfo,
			},
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:         "eventinfo",
				Description:  "Ruft die aktuellen Informationen zu einem Event ab.",
				Type:         discordgo.ChatApplicationCommand,
				DMPermission: util.BoolPtr(false),
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        handlers.IDOptionName,
						Description: "ID des Events, dessen Informationen angezeigt werden sollen.",
						Required:    true,
						MinValue:    util.FloatPtr(1),
					},
				},
			},
		}, {
			Handler: types.InteractionHandler{
				Main:         handler.CreateEvent,
				Autocomplete: handler.HandleAutocomplete,
			},
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:         "eventcreate",
				Description:  "Erstellt ein Clan Event, bspw. für das Farmen von dunklem Elixir in einem bestimmten Zeitraum.",
				Type:         discordgo.ChatApplicationCommand,
				DMPermission: util.BoolPtr(false),
				Options: []*discordgo.ApplicationCommandOption{
					optionClanTag("Clan, für den das Event erstellt werden soll."),
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        handlers.StatisticOptionName,
						Description: "Aufgabe, für die ein Event erstellt werden soll.",
						Required:    true,
						Choices:     types.ComparableStatisticTaskChoices(types.ComparableAchievements),
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        handlers.StartsAtOptionName,
						Description: "Startzeitpunkt des Events. Format: DD.MM.YYYY HH:MM",
						Required:    true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        handlers.EndsAtOptionName,
						Description: "Endzeitpunkt des Events. Format: DD.MM.YYYY HH:MM",
						Required:    true,
					},
				},
			},
		}, {
			Handler: types.InteractionHandler{
				Main: handler.DeleteEvent,
			},
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:         "eventremove",
				Description:  "Löscht ein bereits erstelltes Event, welches noch nicht vorbei ist.",
				Type:         discordgo.ChatApplicationCommand,
				DMPermission: util.BoolPtr(false),
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionInteger,
						Name:        handlers.IDOptionName,
						Description: "ID des Events, das gelöscht werden soll.",
						Required:    true,
						MinValue:    util.FloatPtr(1),
					},
				},
			},
		}}
}
