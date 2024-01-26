package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"

	"bot/commands"
	"bot/commands/util"
	"bot/env"
)

func init() {
	log.SetPrefix("[BOT] ")
	log.SetFlags(log.Ldate | log.Ltime)

	if err := env.Load(); err != nil {
		log.Fatalf("Failed to init environment variables: %v", err)
	}
}

func main() {
	s, err := discordgo.New("Bot " + env.DISCORD_CLIENT_SECRET.Value())
	if err != nil {
		log.Fatalf("Failed to create discord session: %v", err)
	}

	s.Identify.Intents = discordgo.IntentsAll
	if err = s.Open(); err != nil {
		log.Fatalf("Failed to open discord session: %v", err)
	}
	log.Printf("Bot is now logged in as %s and running. Press CTRL-C to exit.", s.State.User.Username)

	go autoUpdateStatus(s)
	util.Session = s

	cmds, err := commands.Setup(s)
	if err != nil {
		log.Fatalf("Failed to add commands: %v", err)
	}
	log.Printf("Successfully added %d commands.", len(cmds))

	shutdownSig := make(chan os.Signal, 1)
	signal.Notify(shutdownSig, os.Interrupt)
	<-shutdownSig

	log.Println("Gracefully shutting down...")
	if err = removeCommands(s, cmds); err != nil {
		log.Printf("Failed to remove commands: %v", err)
	}
	if err = s.Close(); err != nil {
		log.Fatalf("Failed to close discord session: %v", err)
	}
}

// removeCommands removes all commands from the discord server.
func removeCommands(s *discordgo.Session, cmds []*discordgo.ApplicationCommand) error {
	if env.MODE.Value() == "DEBUG" {
		return nil
	}

	errChan := make(chan error)
	log.Printf("Removing %d commands, this takes about a minute...", len(cmds))
	for _, cmd := range cmds {
		go func(cmd *discordgo.ApplicationCommand) {
			errChan <- s.ApplicationCommandDelete(env.DISCORD_CLIENT_ID.Value(), env.DISCORD_GUILD_ID.Value(), cmd.ID)
		}(cmd)
	}

	for range cmds {
		if err := <-errChan; err != nil {
			return err
		}
	}

	log.Printf("Successfully removed %d commands.", len(cmds))
	return nil
}

// autoUpdateStatus updates the game status every hour.
func autoUpdateStatus(s *discordgo.Session) {
	for {
		if err := s.UpdateGameStatus(0, "mit deinen Kickpunkten"); err != nil {
			log.Printf("Failed to update game status: %v", err)
		}
		time.Sleep(time.Hour)
	}
}
