/* Here there are functions that are used to synchronize the database periodically */
package app

import (
	"log/slog"
	"planner-backend/config"
	"time"
)

func InitalizeSynchronization(injector *config.Injector) {
	// ever week
	interval := 7 * 24 * time.Hour
	slog.Info("Initializing synchronization")

	if err := injector.SynchronizeRepo.Synchronize(5); err != nil {
		slog.Error("Error synchronizing", "error", err)
	}

	// Create a ticker that ticks every 24 hours
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			slog.Info("Synchronizing")
			if err := injector.SynchronizeRepo.Synchronize(60); err != nil {
				slog.Error("Error synchronizing", "error", err)
			}
		}
	}()
}
