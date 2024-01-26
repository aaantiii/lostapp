package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"bot/client"
	"bot/commands/messages"
	"bot/env"
	"bot/store/postgres"
)

func Setup(s *discordgo.Session) ([]*discordgo.ApplicationCommand, error) {
	db, err := postgres.NewClient()
	if err != nil {
		return nil, err
	}

	cocClient, err := client.NewCocClient()
	if err != nil {
		log.Print("Bier")
		return nil, err
	}

	interactions := createInteractions(db, cocClient)
	log.Println("Overwriting application commands...")
	cmds, err := s.ApplicationCommandBulkOverwrite(env.DISCORD_CLIENT_ID.Value(), env.DISCORD_GUILD_ID.Value(), interactions.ApplicationCommands())
	if err != nil {
		return nil, err
	}

	s.AddHandler(interactionHandler(interactions))
	return cmds, nil
}

func sendCommandNotFound(i *discordgo.InteractionCreate) {
	messages.SendEmbedResponse(i, messages.NewEmbed(
		"Fehler - Unbekannter Befehl",
		"Dieser Befehl wurde nicht gefunden.",
		messages.ColorRed,
	))
}

func sendDMNotSupported(i *discordgo.InteractionCreate) {
	messages.SendEmbedResponse(i, messages.NewEmbed(
		"Fehler",
		"DMs werden vom Bot nicht unterstüzt. Bitte führe alle Befehle in einem Server aus.",
		messages.ColorRed,
	))
}
