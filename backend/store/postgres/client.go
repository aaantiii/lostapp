package postgres

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"backend/env"
	"backend/store/postgres/models"
)

const (
	maxRetries   = 10
	retryTimeout = time.Second * 15
)

func NewClient() (*gorm.DB, error) {
	db, err := newGormClient()
	if err != nil {
		return nil, err
	}

	if err = migrate(db); err != nil {
		return nil, err
	}

	log.Printf("Auto-migrated postgres models.")
	return db, nil
}

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Clan{},
		&models.ClanMember{},
		&models.ComparableStats{},
		&models.Guild{},
		&models.Kickpoint{},
		&models.LostClan{},
		&models.LostClanSettings{},
		&models.Member{},
		&models.PlayerClan{},
		&models.Player{},
		&models.PlayerStats{},
		&models.User{},
	); err != nil {
		return err
	}

	if err := models.SeedLostClans(db); err != nil {
		return err
	}
	if err := models.SeedComparableStats(db); err != nil {
		return err
	}

	return nil
}

func newGormClient() (client *gorm.DB, err error) {
	dsn := env.POSTGRES_URL.Value()
	loggerMode := logger.Silent
	if env.MODE.Value() != "PROD" {
		loggerMode = logger.Warn
	}

	for i := 0; i < maxRetries; i++ {
		if client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(loggerMode),
		}); err != nil {
			log.Printf("Failed to connect to database: %v\nRetrying in %s...", err, retryTimeout.String())
			time.Sleep(retryTimeout)
			continue
		}

		log.Println("Connected to postgres database.")
		return
	}

	return nil, errors.New("max retries reached")
}
