package cron

import (
	"github.com/go-co-op/gocron"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"log/slog"
	"time"
)

type cronManager struct {
	scheduler *gocron.Scheduler
	logger    slog.Logger
	product   domain.ProductUseCase
}

func NewCronManager(product domain.ProductUseCase, logger slog.Logger) domain.CronManager {
	return &cronManager{
		scheduler: gocron.NewScheduler(time.UTC),
		logger:    logger,
		product:   product,
	}
}

func (c *cronManager) Start() {
	c.logger.Info("starting cron manager")

	c.scheduler.Every(1).Hour().Do(func() {
		err := c.product.SendProductNotifications("hourly")
		if err != nil {
			c.logger.Error("failed whilst sending hourly product notifications: %s", err.Error())
		}
	})

	c.scheduler.Every(1).Day().At("09:00").Do(func() {
		err := c.product.SendProductNotifications("daily")
		if err != nil {
			c.logger.Error("failed whilst sending daily product notifications: %s", err.Error())
		}
	})

	c.scheduler.StartAsync()
}
