package clashsync

import (
	"log/slog"
	"os"

	"github.com/aaantiii/lostapp/services/clashsync/env"
)

func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: getLogLevel(),
	}
	if opts.Level == slog.LevelDebug {
		opts.AddSource = true
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
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
