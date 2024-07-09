package utils

import (
	"log/slog"
	"os"

	"github.com/aaantiii/lostapp/backend/env"
)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: getLogLevel(),
	}))
}

func getLogLevel() slog.Level {
	switch env.MODE.Value() {
	case "PROD":
		return slog.LevelInfo
	case "DEBUG":
		return slog.LevelDebug
	default:
		return slog.LevelDebug
	}
}
