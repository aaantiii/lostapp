package main

import (
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/aaantiii/lostapp/services/clashsync"
	"github.com/aaantiii/lostapp/services/clashsync/env"
)

func init() {
	if err := env.Load(); err != nil {
		panic(err)
	}
	slog.SetDefault(clashsync.NewLogger())
}

func main() {
	db, err := clashsync.NewGormClient()
	if err != nil {
		slog.Error("Error while connecting to database.", slog.Any("err", err))
		os.Exit(1)
	}

	clashClient, err := clashsync.NewCocClient()
	if err != nil {
		slog.Error("Error while connecting to COC-API.", slog.Any("err", err))
		os.Exit(1)
	}

	playersNamesScheduler, err := clashsync.NewUpdatePlayersScheduler(db, clashClient)
	if err != nil {
		slog.Error("Error while creating scheduler.", slog.Any("err", err))
		os.Exit(1)
	}
	playersNamesScheduler.RunEvery(time.Hour * 24)

	shutdownSig := make(chan os.Signal, 1)
	signal.Notify(shutdownSig, os.Interrupt, os.Kill)
	<-shutdownSig
}
