package server

import (
	"github.com/kkcaz/shu-dades-server/internal/config"
	routerUc "github.com/kkcaz/shu-dades-server/internal/router"
	"github.com/pkg/errors"
	"log/slog"
	"os"
)

func Inject(cfg *config.Config) (*routerUc.RouterUseCase, error) {
	logger, err := initLogger(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed whilst initialising logger")
	}

	logger.Info("Logger initialised")

	router := routerUc.NewRouterUseCase(*logger)
	return router, nil
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
