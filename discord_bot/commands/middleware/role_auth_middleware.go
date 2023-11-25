package middleware

import (
	"errors"
	"fmt"

	"github.com/amaanq/coc.go"
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"bot/commands/messages"
	"bot/commands/repos"
	"bot/store/postgres/models"
)

type AuthMiddleware struct {
	guilds repos.IGuildsRepo
	clans  repos.IClansRepo
	users  repos.IUsersRepo
}

func NewAuthMiddleware(guilds repos.IGuildsRepo, clans repos.IClansRepo, users repos.IUsersRepo) AuthMiddleware {
	return AuthMiddleware{
		guilds: guilds,
		clans:  clans,
		users:  users,
	}
}

// NewHandler returns a middleware handler that checks if the user has the specified role in the specified clan.
func (m AuthMiddleware) NewHandler(clanTag string, role coc.Role) InteractionMiddleware {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) error {
		user := models.UserFromGuildMember(i.Member)
		if err := m.users.CreateOrUpdateUser(user); err != nil {
			return err
		}
		if user.IsAdmin {
			return nil
		}

		guild, err := m.guilds.GuildByClanTag(i.GuildID, clanTag)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				m.sendClanNotInGuildError(s, i, clanTag)
				return err
			}
			messages.SendUnknownError(s, i)
			return err
		}

		switch role {
		case coc.Member:
			if guild.IsMember(i.Member.Roles) {
				return nil
			}
		case coc.Elder:
			if guild.IsElder(i.Member.Roles) {
				return nil
			}
		case coc.CoLeader:
			if guild.IsCoLeader(i.Member.Roles) {
				return nil
			}
		case coc.Leader:
			if guild.IsLeader(i.Member.Roles) {
				return nil
			}
		}

		m.sendNoPermissionError(s, i, clanTag, role)
		return errors.New("member is not a leader")
	}
}

func (m AuthMiddleware) sendClanNotInGuildError(s *discordgo.Session, i *discordgo.InteractionCreate, clanTag string) {
	messages.SendEmbed(s, i, messages.NewEmbed(
		"Ungültiger Clan",
		fmt.Sprintf("Der Clan mit dem ClanTag '%s' ist nicht Teil dieses Discord Servers.", clanTag),
		messages.ColorRed,
	))
}

func (m AuthMiddleware) sendNoPermissionError(s *discordgo.Session, i *discordgo.InteractionCreate, clanTag string, role coc.Role) {
	clan, err := m.clans.ClanByTag(clanTag)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			messages.SendClanNotFound(s, i, clanTag)
			return
		}
		messages.SendUnknownError(s, i)
		return
	}

	messages.SendEmbed(s, i, messages.NewEmbed(
		"Keine Berechtigung",
		fmt.Sprintf("Um diesen Befehl ausführen zu können, musst du %s in %s sein.", role.String(), clan.Name),
		messages.ColorRed,
	))
}
