package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"bot/commands/messages"
	"bot/commands/repos"
)

func autocompleteClans(repo repos.ILostClansRepo, query string) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		clans, err := repo.LostClans(query)
		if err != nil {
			log.Print("err clans")
			messages.SendAutoCompletion(s, i, nil)
			return
		}

		messages.SendAutoCompletion(s, i, clans.Choices())
	}
}

func autocompleteMembers(repo repos.IPlayersRepo, query, clanTag string) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if clanTag == "" {
			messages.SendAutoCompletion(s, i, []*discordgo.ApplicationCommandOptionChoice{{
				Name:  "Gib zuerst einen Clan an",
				Value: "Clan auswählen",
			}})
			return
		}

		players, err := repo.MembersPlayersByClan(clanTag, query)
		if err != nil {
			messages.SendAutoCompletion(s, i, nil)
			return
		}

		messages.SendAutoCompletion(s, i, players.Choices())
	}
}
