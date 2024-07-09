package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/aaantiii/lostapp/backend/api"
	"github.com/aaantiii/lostapp/backend/api/utils"
	"github.com/aaantiii/lostapp/backend/env"
)

func main() {
	if err := env.Load(); err != nil {
		fmt.Printf("Failed to initialize env: %v", err)
		os.Exit(1)
	}
	slog.SetDefault(utils.NewLogger())

	router, err := api.NewRouter()
	if err != nil {
		slog.Error("Error while initializing router", slog.Any("err", err))
		os.Exit(1)
	}

	fmt.Println(`
██       ██████  ███████ ████████        ██████ ██       █████  ███    ██ ███████    ██████  ███████ 
██      ██    ██ ██         ██          ██      ██      ██   ██ ████   ██ ██         ██   ██ ██      
██      ██    ██ ███████    ██    █████ ██      ██      ███████ ██ ██  ██ ███████    ██   ██ █████   
██      ██    ██      ██    ██          ██      ██      ██   ██ ██  ██ ██      ██    ██   ██ ██      
███████  ██████  ███████    ██           ██████ ███████ ██   ██ ██   ████ ███████ ██ ██████  ███████
                                      SERVER STARTING...
	`)

	if err = api.ListenAndServe(router); err != nil {
		slog.Error("Error while starting server", slog.Any("err", err))
	}
	slog.Info("Server shutted down.")
}
