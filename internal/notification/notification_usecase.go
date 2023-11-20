package notification

import (
	"github.com/google/uuid"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"log/slog"
)

type notificationUseCase struct {
	Repository domain.NotificationRepository
	Auth       domain.AuthUseCase
	Broadcast  domain.BroadcastUseCase
	Logger     slog.Logger
}

func NewNotificationUseCase(repository domain.NotificationRepository, auth domain.AuthUseCase, broadcast domain.BroadcastUseCase, logger slog.Logger) domain.NotificationUseCase {
	return &notificationUseCase{
		Repository: repository,
		Auth:       auth,
		Broadcast:  broadcast,
		Logger:     logger,
	}
}

func (n *notificationUseCase) Get(userId string) ([]models.Notification, error) {
	notifications, err := n.Repository.Get(userId)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (n *notificationUseCase) Add(userId string, message string) error {
	notification := models.Notification{
		Id:      uuid.New().String(),
		UserId:  userId,
		Message: message,
	}

	err := n.Repository.Add(notification)
	if err != nil {
		return err
	}

	return nil
}

func (n *notificationUseCase) AddAll(message string) error {
	users := n.Auth.GetAllUserIds()
	err := n.AddForUsers(message, users)
	if err != nil {
		return err
	}

	return nil
}

func (n *notificationUseCase) AddForUsers(message string, users []string) error {
	for _, user := range users {
		err := n.Add(user, message)
		if err != nil {
			n.Logger.Warn("failed to add notification to user", "userId", user, "error", err)
		}
	}

	err := n.Broadcast.PublishToUsers(message, "notification", users)
	if err != nil {
		n.Logger.Error("failed to broadcast notification", "error", err)
		return err
	}

	return nil
}

func (n *notificationUseCase) Delete(userId string, notificationId string) error {
	err := n.Repository.Delete(userId, notificationId)
	if err != nil {
		return err
	}

	return nil
}
