package domain

import "github.com/kkcaz/shu-dades-server/pkg/models"

type NotificationRepository interface {
	Add(notification models.Notification) error
	Delete(userId string, notificationId string) error
	Get(userId string) ([]models.Notification, error)
}

type NotificationUseCase interface {
	Get(userId string) ([]models.Notification, error)
	Add(userId string, message string) error
	AddAll(message string) error
	Delete(userId string, notificationId string) error
}
