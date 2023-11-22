package server

import (
	"github.com/kkcaz/shu-dades-server/internal/auth"
	"github.com/kkcaz/shu-dades-server/internal/broadcast"
	"github.com/kkcaz/shu-dades-server/internal/chat"
	"github.com/kkcaz/shu-dades-server/internal/config"
	"github.com/kkcaz/shu-dades-server/internal/cron"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	encryption2 "github.com/kkcaz/shu-dades-server/internal/encryption"
	"github.com/kkcaz/shu-dades-server/internal/front_controller"
	"github.com/kkcaz/shu-dades-server/internal/notification"
	"github.com/kkcaz/shu-dades-server/internal/product"
	routerUc "github.com/kkcaz/shu-dades-server/internal/router"
	"github.com/pkg/errors"
	"log/slog"
	"os"
)

func Inject(cfg *config.Config) (domain.FrontController, error) {
	logger, err := initLogger(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed whilst initialising logger")
	}

	logger.Info("Logger initialised")

	encryption := encryption2.NewEncryptionUseCase(*logger)

	authUseCase := auth.NewAuthUseCase()
	broadcastUseCase := broadcast.NewBroadcastUseCase(*logger, encryption)

	notificationRepository := notification.NewNotificationRepository(*logger)
	notificationUseCase := notification.NewNotificationUseCase(notificationRepository, authUseCase, broadcastUseCase, *logger)

	productRepository := product.NewProductRepository(*logger)
	productUseCase := product.NewProductUseCase(productRepository, notificationUseCase, *logger)

	chatRepository := chat.NewChatRepository(*logger)
	chatUseCase := chat.NewChatUseCase(chatRepository, authUseCase, broadcastUseCase, *logger)

	router := routerUc.NewRouterUseCase(*logger)
	product.NewProductHandler(router, productUseCase, authUseCase)
	auth.NewAuthHandler(router, authUseCase)
	broadcast.NewBroadcastHandler(router, broadcastUseCase, authUseCase)
	notification.NewNotificationHandler(router, notificationUseCase, authUseCase)
	chat.NewChatHandler(router, chatUseCase, authUseCase)

	frontController := front_controller.NewFrontController(*router, encryption, broadcastUseCase)

	cronManager := cron.NewCronManager(productUseCase, *logger)
	cronManager.Start()

	return frontController, nil
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
