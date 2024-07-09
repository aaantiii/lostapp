package clashsync

import (
	"errors"
	"log/slog"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/aaantiii/lostapp/services/clashsync/env"
)

const (
	maxRetries   = 3
	retryTimeout = time.Second * 15
)

func NewGormClient() (*gorm.DB, error) {
	db, err := newGormClient()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newGormClient() (client *gorm.DB, err error) {
	dsn := env.POSTGRES_URL.Value()
	loggerMode := logger.Silent
	if env.MODE.Value() != "PROD" {
		loggerMode = logger.Info
	}

	for i := 0; i < maxRetries; i++ {
		if client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(loggerMode),
		}); err != nil {
			slog.Error("Failed to connect to database.", slog.Any("err", err), slog.String("retryIn", retryTimeout.String()))
			time.Sleep(retryTimeout)
			continue
		}

		slog.Info("Connected to postgres database.")
		return client, nil
	}

	return nil, errors.New("max retries reached")
}
