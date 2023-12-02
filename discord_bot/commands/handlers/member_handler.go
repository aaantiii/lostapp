package handlers

import (
	"github.com/bwmarrin/discordgo"

	"bot/commands/messages"
	"bot/commands/repos"
)

type IMemberHandler interface {
	ListMembers(s *discordgo.Session, i *discordgo.InteractionCreate)
	AddMember(s *discordgo.Session, i *discordgo.InteractionCreate)
	RemoveMember(s *discordgo.Session, i *discordgo.InteractionCreate)
	EditMember(s *discordgo.Session, i *discordgo.InteractionCreate)
	HandleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type MemberHandler struct {
	members repos.IMembersRepo
	clans   repos.IClansRepo
	players repos.IPlayersRepo
}

func NewMemberHandler(members repos.IMembersRepo, clans repos.IClansRepo, players repos.IPlayersRepo) IMemberHandler {
	return &MemberHandler{
		members: members,
		clans:   clans,
		players: players,
	}
}

func (h *MemberHandler) ListMembers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	opts := i.ApplicationCommandData().Options
	if len(opts) != 1 {
		messages.SendInvalidInputError(s, i, "Du musst einen Clan angeben.")
		return
	}

	clanTag := opts[0].StringValue()
	clan, err := h.clans.ClanByTagPreload(clanTag)
	if err != nil {
		messages.SendClanNotFound(s, i, clanTag)
		return
	}

	messages.SendClanMembers(s, i, clan)
}

func (h *MemberHandler) AddMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//TODO implement me
	panic("implement me")
}

func (h *MemberHandler) RemoveMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//TODO implement me
	panic("implement me")
}

func (h *MemberHandler) EditMember(s *discordgo.Session, i *discordgo.InteractionCreate) {
	//TODO implement me
	panic("implement me")
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
		case PlayerTagOptionName:
			autocompleteMembers(h.players, opt.StringValue(), opts[0].StringValue())(s, i)
		}
	}
}
