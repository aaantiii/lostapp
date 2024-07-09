package clashsync

import (
	"errors"
	"log/slog"
	"time"

	"github.com/aaantiii/goclash"
	"gorm.io/gorm"

	"github.com/aaantiii/lostapp/services/clashsync/models"
)

func NewUpdatePlayersScheduler(db *gorm.DB, client *goclash.Client) (*Scheduler, error) {
	factory := newUpdatePlayersTaskFactory(db, client)
	if factory == nil {
		return nil, errors.New("failed to create task factory")
	}

	return NewScheduler(factory), nil
}

func newUpdatePlayersTaskFactory(db *gorm.DB, client *goclash.Client) func() Task {
	var dbPlayers models.Players
	if err := db.Find(&dbPlayers).Error; err != nil {
		slog.Error("Failed to get players from database.", slog.Any("err", err))
		return nil
	}

	return func() Task {
		return NewBatchedTask[*models.Player](func([]*models.Player) error {
			players := client.GetPlayers(dbPlayers.Tags()...)
			changes := make(models.Players, 0)
			for i, player := range players {
				if player == nil {
					slog.Warn("Player not found or failed to fetch.", slog.String("tag", dbPlayers[i].CocTag))
					continue
				}
				if player.Name != dbPlayers[i].Name {
					dbPlayers[i].Name = player.Name
					changes = append(changes, dbPlayers[i])
				}
			}

			if len(changes) == 0 {
				slog.Info("No player names to update.")
				return nil
			}
			if err := db.Save(changes).Error; err != nil {
				return err
			}

			slog.Info("Saved updated player names to database.", slog.Int("amount", len(changes)))
			return nil
		}, time.Minute, 100, dbPlayers)
	}
}
