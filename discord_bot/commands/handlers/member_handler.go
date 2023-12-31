package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"bot/client"
	"bot/commands/messages"
	"bot/commands/middleware"
	"bot/commands/repos"
	"bot/commands/util"
	"bot/commands/validation"
	"bot/store/postgres/models"
	"bot/types"
)

type IMemberHandler interface {
	ListMembers(s *discordgo.Session, i *discordgo.InteractionCreate)
	ClanMemberStatus(s *discordgo.Session, i *discordgo.InteractionCreate)
	AddMember(s *discordgo.Session, i *discordgo.InteractionCreate)
	RemoveMember(s *discordgo.Session, i *discordgo.InteractionCreate)
	EditMember(s *discordgo.Session, i *discordgo.InteractionCreate)
	HandleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type MemberHandler struct {
	members        repos.IMembersRepo
	clans          repos.IClansRepo
	players        repos.IPlayersRepo
	guilds         repos.IGuildsRepo
	authMiddleware middleware.AuthMiddleware
	cocClient      *client.CocClient
}

func NewMemberHandler(members repos.IMembersRepo, clans repos.IClansRepo, players repos.IPlayersRepo, guilds repos.IGuildsRepo, authMiddleware middleware.AuthMiddleware, cocClient *client.CocClient) IMemberHandler {
	return &MemberHandler{
		members:        members,
		clans:          clans,
		players:        players,
		guilds:         guilds,
		authMiddleware: authMiddleware,
		cocClient:      cocClient,
	}
}

func (h *MemberHandler) ListMembers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	if clanTag == "" {
		messages.SendInvalidInputError(s, i, "Bitte gib einen Clan an.")
		return
	}

	clan, err := h.clans.ClanByTagPreload(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	messages.SendClanMembers(s, i, clan)
}

func (h *MemberHandler) ClanMemberStatus(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	if clanTag == "" {
		messages.SendInvalidInputError(s, i, "Bitte gib einen Clan an.")
		return
	}

	members, err := h.members.MembersByClanTag(clanTag)
	if err != nil {
		messages.SendUnknownError(s, i)
		return
	}

	clan, err := h.cocClient.GetClan(clanTag)
	if err != nil {
		messages.SendCocApiError(s, i)
		return
	}

	messages.SendClansMembersStatus(s, i, members, clan)
}

func (h *MemberHandler) AddMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	playerTag := util.StringOptionByName(PlayerTagOptionName, opts)
	role := models.ClanRole(util.StringOptionByName(RoleOptionName, opts))

	if clanTag == "" || playerTag == "" || role == "" {
		messages.SendInvalidInputError(s, i, "Bitte gib einen Clan, ein Mitglied und eine Rolle an.")
		return
	}

	if !validation.ValidateClanRole(role) {
		messages.SendInvalidInputError(s, i, fmt.Sprintf("Die Rolle %s ist ungültig.", string(role)))
		return
	}

	requiredAuthRole := types.AuthRoleAdmin
	if role == models.RoleMember || role == models.RoleElder {
		requiredAuthRole = types.AuthRoleCoLeader
	} else if role == models.RoleCoLeader {
		requiredAuthRole = types.AuthRoleLeader
	}

	if err := h.authMiddleware.NewClanHandler(clanTag, requiredAuthRole)(s, i); err != nil {
		return
	}

	player, err := h.players.PlayerByTag(playerTag)
	if err != nil || player.DiscordID == "" {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Mitglied nicht verifiziert",
			"Das Mitglied muss sich zuerst verifizieren, bevor es zu einem Clan hinzugefügt werden kann.",
			messages.ColorRed,
		))
		return
	}

	if err = h.members.CreateMember(&models.ClanMember{
		PlayerTag:        playerTag,
		ClanTag:          clanTag,
		ClanRole:         role,
		AddedByDiscordID: i.Member.User.ID,
	}); err != nil {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Es ist ein Fehler aufgetreten",
			"Beim Speichern des Mitglieds ist ein Fehler aufgetreten. Dies kann daran liegen, dass das Mitglied bereits existiert oder ungültige Daten angegeben wurden.",
			messages.ColorRed,
		))
		return
	}

	guild, roleErr := h.guilds.GuildByClanTag(i.GuildID, clanTag)
	if roleErr == nil {
		roleErr = s.GuildMemberRoleAdd(i.GuildID, player.DiscordID, guild.MemberRoleID)
	}

	desc := fmt.Sprintf("Das Mitglied wurde erfolgreich als %s zum Clan hinzugefügt.", role.Format())
	if roleErr != nil {
		desc += "\n\n**ACHTUNG**: Dem Mitglied konnte die Mitglieder-Rolle nicht zugewiesen werden. Bitte weise ihm die Rolle manuell zu."
	}

	messages.SendEmbed(s, i, messages.NewEmbed(
		"Mitglied hinzugefügt",
		desc,
		messages.ColorGreen,
	))
}

func (h *MemberHandler) RemoveMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	memberTag := util.StringOptionByName(MemberTagOptionName, opts)

	if clanTag == "" || memberTag == "" {
		messages.SendInvalidInputError(s, i, "Bitte gib einen Clan und ein Mitglied an.")
		return
	}

	member, err := h.members.MemberByID(memberTag, clanTag)
	if err != nil {
		messages.SendMemberNotFound(s, i, memberTag, clanTag)
		return
	}

	requiredAuthRole := types.AuthRoleAdmin
	if member.ClanRole == models.RoleMember || member.ClanRole == models.RoleElder {
		requiredAuthRole = types.AuthRoleCoLeader
	} else if member.ClanRole == models.RoleCoLeader {
		requiredAuthRole = types.AuthRoleLeader
	}

	if err = h.authMiddleware.NewClanHandler(clanTag, requiredAuthRole)(s, i); err != nil {
		return
	}

	if err = h.members.DeleteMember(member.PlayerTag, member.ClanTag); err != nil {
		messages.SendUnknownError(s, i)
		return
	}

	guild, roleErr := h.guilds.GuildByClanTag(i.GuildID, clanTag)
	if roleErr == nil {
		roleErr = s.GuildMemberRoleRemove(i.GuildID, member.Player.DiscordID, guild.MemberRoleID)
	}

	desc := fmt.Sprintf("Das Mitglied %s wurde aus %s entfernt.", member.Player.Name, member.Clan.Name)
	if roleErr != nil {
		desc += "\n\n**ACHTUNG**: Dem Mitglied konnte die Mitglieder-Rolle nicht entfernt werden. Bitte entferne ihm die Rolle manuell."
	}
	if member.ClanRole == models.RoleElder {
		desc += "\n\n**ACHTUNG**: Das Mitglied war Ältester. Bitte entferne ihm die Ältesten-Rolle manuell."
	}

	messages.SendEmbed(s, i, messages.NewEmbed(
		"Mitglied entfernt",
		desc,
		messages.ColorGreen,
	))
}

func (h *MemberHandler) EditMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	clanTag := util.StringOptionByName(ClanTagOptionName, opts)
	memberTag := util.StringOptionByName(MemberTagOptionName, opts)
	role := models.ClanRole(util.StringOptionByName(RoleOptionName, opts))

	if clanTag == "" || memberTag == "" || role == "" {
		messages.SendInvalidInputError(s, i, "Bitte gib einen Clan, ein Mitglied und eine Rolle an.")
		return
	}

	if !validation.ValidateClanRole(role) {
		messages.SendInvalidInputError(s, i, fmt.Sprintf("Die Rolle %s ist ungültig.", string(role)))
		return
	}

	requiredAuthRole := types.AuthRoleAdmin
	if role == models.RoleMember || role == models.RoleElder {
		requiredAuthRole = types.AuthRoleCoLeader
	} else if role == models.RoleCoLeader {
		requiredAuthRole = types.AuthRoleLeader
	}

	if err := h.authMiddleware.NewClanHandler(clanTag, requiredAuthRole)(s, i); err != nil {
		return
	}

	member, err := h.members.MemberByID(memberTag, clanTag)
	if err != nil {
		messages.SendMemberNotFound(s, i, memberTag, clanTag)
		return
	}

	if member.ClanRole == role {
		messages.SendEmbed(s, i, messages.NewEmbed(
			"Mitglied nicht geändert",
			fmt.Sprintf("Das Mitglied %s hat bereits die Rolle %s.", member.Player.Name, role.Format()),
			messages.ColorRed,
		))
		return
	}

	if err = h.members.UpdateMemberRole(member.PlayerTag, member.ClanTag, role); err != nil {
		messages.SendUnknownError(s, i)
		return
	}

	messages.SendEmbed(s, i, messages.NewEmbed(
		"Mitglied geändert",
		fmt.Sprintf("Das Mitglied %s hat nun die Rolle %s.", member.Player.Name, role.Format()),
		messages.ColorGreen,
	))
}

func (h *MemberHandler) HandleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options

	for _, opt := range opts {
		if !opt.Focused {
			continue
		}

		switch opt.Name {
		case ClanTagOptionName:
			autocompleteClans(h.clans, opt.StringValue())(s, i)
		case MemberTagOptionName:
			autocompleteMembers(h.players, opt.StringValue(), util.StringOptionByName(ClanTagOptionName, opts))(s, i)
		case PlayerTagOptionName:
			autocompletePlayers(h.players, opt.StringValue())(s, i)
		}
	}
}
