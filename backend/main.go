package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aaantiii/lostapp/backend/api"
	"github.com/aaantiii/lostapp/backend/env"
)

func init() {
	var err error
	start := time.Now()

	log.SetFlags(log.Ldate + log.Ltime)
	log.SetPrefix("[SERVER] ")

	if err = env.Init(); err != nil {
		log.Fatalf("Failed to initialize env: %v", err)
	}

	if err = api.Init(); err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	log.Printf("Initialized in %s.", time.Since(start).Round(time.Millisecond).String())
}

func main() {
	fmt.Println(`
██       ██████  ███████ ████████        ██████ ██       █████  ███    ██ ███████    ██████  ███████ 
██      ██    ██ ██         ██          ██      ██      ██   ██ ████   ██ ██         ██   ██ ██      
██      ██    ██ ███████    ██    █████ ██      ██      ███████ ██ ██  ██ ███████    ██   ██ █████   
██      ██    ██      ██    ██          ██      ██      ██   ██ ██  ██ ██      ██    ██   ██ ██      
███████  ██████  ███████    ██           ██████ ███████ ██   ██ ██   ████ ███████ ██ ██████  ███████
                                      SERVER STARTING...
	`)

	if err := api.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Server shutdown successfully.")
}
