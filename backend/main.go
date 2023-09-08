package main

import (
	"fmt"
	"log"

	"backend/api"
)

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

	log.Println("Server shutted down successfully.")
}
