package main

import (
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/bwmarrin/discordgo"

	"bot/commands"
	"bot/commands/util"
	"bot/env"
)

func init() {
	if err := env.Load(); err != nil {
		slog.Error("Failed to init environment variables.", slog.Any("err", err))
		os.Exit(1)
	}
	slog.SetDefault(util.NewLogger())
}

func main() {
	s, err := discordgo.New("Bot " + env.DISCORD_CLIENT_SECRET.Value())
	if err != nil {
		slog.Error("Failed to create discord session.", slog.Any("err", err))
		os.Exit(1)
	}

	s.Identify.Intents = discordgo.IntentsAll
	if err = s.Open(); err != nil {
		slog.Error("Failed to open discord session.", slog.Any("err", err))
		os.Exit(1)
	}
	slog.Info("Bot is now logged in and running. Press CTRL-C to exit.", slog.String("username", s.State.User.Username), slog.String("id", s.State.User.ID))

	go autoUpdateStatus(s)
	util.Session = s
	cmds, err := commands.Setup(s)
	if err != nil {
		slog.Error("Failed to add commands. Shutting down...", slog.Any("err", err))
		s.Close()
		os.Exit(1)
	}
	slog.Info("Commands were added successfully.", slog.Int("amount", len(cmds)))

	shutdownSig := make(chan os.Signal, 1)
	signal.Notify(shutdownSig, os.Interrupt)
	<-shutdownSig

	slog.Info("Gracefully shutting down...")
	if err = removeCommands(s, cmds); err != nil {
		slog.Warn("Failed to remove commands.", slog.Any("err", err))
	}
	if err = s.Close(); err != nil {
		slog.Error("Failed to close discord session.", slog.Any("err", err))
		os.Exit(1)
	}
}

// removeCommands removes all commands from the discord server.
func removeCommands(s *discordgo.Session, cmds []*discordgo.ApplicationCommand) error {
	if env.MODE.Value() == "DEBUG" {
		return nil
	}

	slog.Info("Removing all commands, this can take a few minutes...", slog.Int("amount", len(cmds)))
	errChan := make(chan error)
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

	slog.Info("Successfully removed all commands.", slog.Int("amount", len(cmds)))
	return nil
}

// autoUpdateStatus updates the bot's game status every hour.
func autoUpdateStatus(s *discordgo.Session) {
	for {
		if err := s.UpdateGameStatus(0, "mit deinen Kickpunkten"); err != nil {
			slog.Warn("Failed to update game status.", slog.Any("err", err))
		}
		time.Sleep(time.Hour)
	}
}
