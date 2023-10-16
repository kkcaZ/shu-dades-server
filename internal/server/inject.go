package server

import (
	"github.com/kkcaz/shu-dades-server/internal/config"
	"github.com/pkg/errors"
	"log/slog"
	"os"
)

func Inject(cfg *config.Config) error {
	logger, err := initLogger(cfg)
	if err != nil {
		return errors.Wrap(err, "failed whilst initialising logger")
	}

	logger.Info("Logger initialised")
	return nil
}

func initLogger(cfg *config.Config) (*slog.Logger, error) {
	var logLevel slog.Level
	err := logLevel.UnmarshalText([]byte(cfg.Service.LogLevel))
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	return logger, nil
}
