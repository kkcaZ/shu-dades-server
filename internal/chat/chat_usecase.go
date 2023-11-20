package chat

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"github.com/pkg/errors"
	"log/slog"
	"time"
)

type chatUseCase struct {
	Repository domain.ChatRepository
	Auth       domain.AuthUseCase
	Broadcast  domain.BroadcastUsecase
	Logger     slog.Logger
}

func NewChatUseCase(chatRepository domain.ChatRepository, auth domain.AuthUseCase, broadcast domain.BroadcastUsecase, logger slog.Logger) domain.ChatUseCase {
	return &chatUseCase{
		Repository: chatRepository,
		Auth:       auth,
		Broadcast:  broadcast,
		Logger:     logger,
	}
}

func (c *chatUseCase) GetChatThumbnails(userId string) ([]models.ChatThumbnail, error) {
	thumbnails, err := c.Repository.GetAllChatThumbnails()
	if err != nil {
		return nil, err
	}

	var userThumbnails = make([]models.ChatThumbnail, 0)
	for _, thumbnail := range thumbnails {
		for _, participant := range thumbnail.Participants {
			if participant.UserId == userId {
				userThumbnails = append(userThumbnails, thumbnail)
				break
			}
		}
	}

	return userThumbnails, nil
}

func (c *chatUseCase) GetChat(chatId string) (*models.Chat, error) {
	chat, err := c.Repository.GetChat(chatId)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (c *chatUseCase) GetChatParticipantIds(chatId string) ([]string, error) {
	chat, err := c.GetChat(chatId)
	if err != nil {
		return nil, err
	}

	var participantIds = make([]string, 0)
	for _, participant := range chat.Participants {
		participantIds = append(participantIds, participant.UserId)
	}

	return participantIds, nil
}

func (c *chatUseCase) CreateChat(userIds []string) (*models.Chat, error) {
	participants := make([]models.Participant, 0)
	for _, userId := range userIds {
		user, err := c.Auth.GetUserById(userId)
		if err != nil {
			return nil, err
		}

		participant := models.Participant{
			UserId:   user.Id,
			Username: user.Username,
		}
		participants = append(participants, participant)
	}

	chat := models.Chat{
		Id:           uuid.New().String(),
		Participants: participants,
		Messages:     make([]models.Message, 0),
	}

	err := c.Repository.CreateChat(chat)
	if err != nil {
		return nil, err
	}

	err = c.Broadcast.PublishToUsers(chat.Id, "newChat", userIds)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to publish message to users")
	}

	return &chat, nil
}

func (c *chatUseCase) SendMessage(chatId string, content string, userId string) error {
	user, err := c.Auth.GetUserById(userId)
	if err != nil {
		return err
	}

	message := models.Message{
		UserId:   user.Id,
		Username: user.Username,
		Content:  content,
		SentAt:   time.Now().UTC(),
	}

	err = c.Repository.AddMessage(chatId, message)
	if err != nil {
		return err
	}

	err = c.NotifyUsersAboutNewMessage(chatId, message)
	if err != nil {
		return errors.Wrap(err, "failed to notify users")
	}

	return nil
}

func (c *chatUseCase) NotifyUsersAboutNewMessage(chatId string, message models.Message) error {
	participantIds, err := c.GetChatParticipantIds(chatId)
	if err != nil {
		return errors.Wrapf(err, "failed to get chat participant ids - unable to send socket update")
	}

	messageEvent := models.MessageEvent{
		Message: message,
		ChatId:  chatId,
	}

	msgBytes, err := json.Marshal(messageEvent)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal message")
	}

	err = c.Broadcast.PublishToUsers(string(msgBytes), "message", participantIds)
	if err != nil {
		return errors.Wrapf(err, "failed to publish message to users")
	}
	return nil
}
