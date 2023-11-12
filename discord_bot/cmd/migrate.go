package main

import (
	"log"

	"bot/env"
	"bot/store/postgres"
	"bot/store/postgres/models"
)

func main() {
	if err := env.Init(); err != nil {
		panic(err)
	}

	db, err := postgres.NewClient()
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(
		&models.Penalty{},
		&models.LostClan{},
		&models.Member{},
		&models.Player{}); err != nil {
		panic(err)
	}

	log.Println("Migrated successfully")
}
