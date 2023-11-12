package client

import (
	"github.com/bwmarrin/discordgo"

	"bot/commands"
	"bot/env"
)

func NewDiscordSession() (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + env.DISCORD_CLIENT_SECRET.Value())
	if err != nil {
		return nil, err
	}

	if err = commands.Setup(session); err != nil {
		return nil, err
	}

	if err = session.Open(); err != nil {
		return nil, err
	}

	if err = session.UpdateGameStatus(0, "mit deinen Kickpunkten"); err != nil {
		return nil, err
	}

	return session, nil
}

func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {

	}
}
