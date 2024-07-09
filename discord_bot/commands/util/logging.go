package util

import (
	"log/slog"
	"os"

	"bot/env"
)

func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: getLogLevel(),
	}
	if opts.Level == slog.LevelDebug {
		opts.AddSource = true
	}
	return slog.New(slog.NewTextHandler(os.Stderr, opts))
}

func getLogLevel() slog.Level {
	switch env.MODE.Value() {
	case "PROD":
		return slog.LevelWarn
	case "DEBUG":
		return slog.LevelDebug
	}
	return slog.LevelWarn
}
