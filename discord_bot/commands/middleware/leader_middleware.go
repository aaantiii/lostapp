package middleware

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bot/commands/messages"
	"bot/commands/repos"
)

// NewClanLeaderMiddleware returns a middleware that checks if the user is a leader.
//
// If anyLeader is true, the user can be leader of any clan. Otherwise, the user must be leader of the clan.
func ClanLeaderMiddleware(guilds repos.IGuildsRepo, clanTag string, allowCoLeader bool) func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
		guild, err := guilds.GuildByClanTag(i.GuildID, clanTag)
		if err != nil {
			messages.SendEmbed(s, i, messages.NewMessageEmbed(
				"Fehler beim Abrufen des Clans",
				fmt.Sprintf("Der Clan mit dem Tag '%s' konnte nicht gefunden werden.", clanTag),
				messages.ColorRed,
			))
			return err
		}

		if guild.IsLeader(i.Member.Roles) || (allowCoLeader && guild.IsCoLeader(i.Member.Roles)) {
			return nil
		}

		messages.SendEmbed(s, i, messages.NewMessageEmbed(
			"Keine Berechtigung",
			fmt.Sprintf("Der Clan mit dem Tag '%s' konnte nicht gefunden werden.", clanTag),
			messages.ColorRed,
		))
		return errors.New("member is not a leader")
	}
}
