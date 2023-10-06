package server

import (
	"context"
	"github.com/kkcaz/shu-dades-server/internal/config"
	"log"
	"log/slog"
	"os/signal"
	"syscall"
)

func Run() {
	slog.Info("Initialising server")

	// Create context that listens for interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to initialise config: %v", err)
	}

	err = Inject(cfg)
	if err != nil {
		log.Fatalf("failed to inject dependencies: %v", err)
	}

	<-ctx.Done()
}
