package notification

import (
	"encoding/json"
	"fmt"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"log/slog"
	"os"
)

type notificationData struct {
	Notifications []models.Notification `json:"notifications"`
}

type notificationRepository struct {
	Logger        slog.Logger
	notifications []models.Notification
}

func NewNotificationRepository(logger slog.Logger) domain.NotificationRepository {
	notifications, err := readNotifications()
	if err != nil {
		panic(err)
	}

	return &notificationRepository{
		notifications: notifications,
		Logger:        logger,
	}
}

func readNotifications() ([]models.Notification, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dat, err := os.ReadFile(fmt.Sprintf("%s/internal/data/product/products.json", currentDir))
	if err != nil {
		return nil, err
	}

	var productData notificationData
	err = json.Unmarshal(dat, &productData)
	if err != nil {
		return nil, err
	}

	return productData.Notifications, nil
}

func (n *notificationRepository) Get(userId string) ([]models.Notification, error) {
	notifs := make([]models.Notification, 0)
	for _, notif := range n.notifications {
		if notif.UserId == userId {
			notifs = append(notifs, notif)
		}
	}
	return notifs, nil
}

func (n *notificationRepository) Add(notification models.Notification) error {
	n.Logger.Info("adding notification", "notification", notification)
	n.notifications = append(n.notifications, notification)
	return nil
}

func (n *notificationRepository) Delete(userId string, notificationId string) error {
	n.Logger.Info("deleting notification", "userId", userId, "notificationId", notificationId)
	for i, notif := range n.notifications {
		if notif.UserId == userId && notif.Id == notificationId {
			n.notifications = append(n.notifications[:i], n.notifications[i+1:]...)
		}
	}
	return nil
}
