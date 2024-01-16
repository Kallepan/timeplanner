package app

import (
	"log/slog"
	"planner-backend/config"
	"time"
)

func InitalizeSynchronization(injector *config.Injector) {
	interval := 24 * time.Hour
	slog.Info("Initializing synchronization")

	if err := injector.SynchronizeRepo.Synchronize(60); err != nil {
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
