package main

import (
	"fmt"
	"log"

	"github.com/aaantiii/lostapp/backend/api"
	"github.com/aaantiii/lostapp/backend/env"
)

func init() {
	log.SetFlags(log.Ldate + log.Ltime)
	log.SetPrefix("[SERVER] ")

	if err := env.Load(); err != nil {
		log.Fatalf("Failed to initialize env: %v", err)
	}
}

func main() {
	router, err := api.NewRouter()
	if err != nil {
		log.Fatalf("Error while initializing router: %v", err)
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
		log.Fatalf("Error while serving: %v", err)
	}

	log.Println("Server shutdown successfully.")
}
