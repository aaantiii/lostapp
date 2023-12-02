package commands

import (
	"github.com/bwmarrin/discordgo"

	"bot/client"
	"bot/commands/messages"
	"bot/env"
	"bot/store/postgres"
)

func Setup(session *discordgo.Session) error {
	db, err := postgres.NewClient()
	if err != nil {
		return err
	}

	cocClient, err := client.NewCocClient()
	if err != nil {
		return err
	}

	interactions := createInteractions(db, cocClient)
	if _, err = session.ApplicationCommandBulkOverwrite(env.DISCORD_CLIENT_ID.Value(), env.DISCORD_GUILD_ID.Value(), interactions.ApplicationCommands()); err != nil {
		return err
	}
	session.AddHandler(interactionHandler(interactions))

	return nil
}

func sendCommandNotFound(s *discordgo.Session, i *discordgo.InteractionCreate) {
	messages.SendEmbed(s, i, messages.NewEmbed(
		"Fehler - Unbekannter Befehl",
		"Dieser Befehl wurde nicht gefunden.",
		messages.ColorRed,
	))
}

func sendDMNotSupported(s *discordgo.Session, i *discordgo.InteractionCreate) {
	messages.SendEmbed(s, i, messages.NewEmbed(
		"Fehler",
		"DMs werden vom Bot nicht unterstüzt. Bitte führe alle Befehle in einem Server aus.",
		messages.ColorRed,
	))
}
