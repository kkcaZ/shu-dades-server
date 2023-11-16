package server

import (
	"github.com/kkcaz/shu-dades-server/internal/auth"
	"github.com/kkcaz/shu-dades-server/internal/broadcast"
	"github.com/kkcaz/shu-dades-server/internal/config"
	"github.com/kkcaz/shu-dades-server/internal/product"
	routerUc "github.com/kkcaz/shu-dades-server/internal/router"
	"github.com/pkg/errors"
	"log/slog"
	"os"
)

func Inject(cfg *config.Config) (*routerUc.RouterUseCase, *broadcast.BroadcastUseCase, error) {
	logger, err := initLogger(cfg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed whilst initialising logger")
	}

	logger.Info("Logger initialised")

	authUseCase := auth.NewAuthUseCase()
	productRepository := product.NewProductRepository(*logger)
	productUseCase := product.NewProductUseCase(productRepository)

	router := routerUc.NewRouterUseCase(*logger)
	broadcastUseCase := broadcast.NewBroadcastUseCase(*logger)
	product.NewProductHandler(router, productUseCase)
	auth.NewAuthHandler(router, authUseCase)
	broadcast.NewBroadcastHandler(router, broadcastUseCase)

	return router, broadcastUseCase, nil
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
