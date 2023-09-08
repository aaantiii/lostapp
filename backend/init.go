package main

import (
	"log"
	"time"

	"backend/api"
	"backend/env"
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
