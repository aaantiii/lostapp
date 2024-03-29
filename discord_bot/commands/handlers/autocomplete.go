package handlers

import (
	"github.com/bwmarrin/discordgo"

	"bot/commands/messages"
	"bot/commands/repos"
)

func autocompleteClans(i *discordgo.InteractionCreate, repo repos.IClansRepo, query string) {
	clans, err := repo.Clans(query)
	if err != nil {
		messages.SendAutoCompletion(i, nil)
		return
	}

	messages.SendAutoCompletion(i, clans.Choices())
}

func autocompleteMembers(i *discordgo.InteractionCreate, repo repos.IPlayersRepo, query, clanTag string) {
	if clanTag == "" {
		messages.SendAutoCompletion(i, []*discordgo.ApplicationCommandOptionChoice{{
			Name:  "Gib zuerst einen Clan an",
			Value: "Clan auswählen",
		}})
		return
	}

	players, err := repo.MembersPlayersByClan(clanTag, query)
	if err != nil {
		messages.SendAutoCompletion(i, nil)
		return
	}

	messages.SendAutoCompletion(i, players.Choices())
}

func autocompletePlayers(i *discordgo.InteractionCreate, repo repos.IPlayersRepo, query string) {
	players, err := repo.Players(query)
	if err != nil {
		messages.SendAutoCompletion(i, nil)
		return
	}

	messages.SendAutoCompletion(i, players.Choices())
}
