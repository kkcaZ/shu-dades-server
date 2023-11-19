package chat

import (
	"encoding/json"
	"fmt"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"log/slog"
	"os"
)

type chatData struct {
	Chats []models.Chat `json:"chat"`
}

type chatRepository struct {
	Logger slog.Logger
	chats  []models.Chat
}

func NewChatRepository(logger slog.Logger) domain.ChatRepository {
	chats, err := readChats()
	if err != nil {
		panic(err)
	}

	return &chatRepository{
		Logger: logger,
		chats:  chats,
	}
}

func readChats() ([]models.Chat, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	dat, err := os.ReadFile(fmt.Sprintf("%s/internal/data/chat/chats.json", currentDir))
	if err != nil {
		return nil, err
	}

	var productData chatData
	err = json.Unmarshal(dat, &productData)
	if err != nil {
		return nil, err
	}

	return productData.Chats, nil
}

func (c *chatRepository) GetAllChatThumbnails() ([]models.ChatThumbnail, error) {
	chatThumbnails := make([]models.ChatThumbnail, 0)
	for _, chat := range c.chats {
		var message *models.Message
		if len(chat.Messages) != 0 {
			message = &chat.Messages[len(chat.Messages)-1]
		}

		chatThumbnails = append(chatThumbnails, models.ChatThumbnail{
			ChatId:       chat.Id,
			LastMessage:  message,
			Participants: chat.Participants,
		})
	}
	return chatThumbnails, nil
}

func (c *chatRepository) GetChat(chatId string) (*models.Chat, error) {
	for _, chat := range c.chats {
		if chat.Id == chatId {
			c.Logger.Info("retrieved chat", "chat", chat)
			return &chat, nil
		}
	}

	return nil, nil
}

func (c *chatRepository) CreateChat(chat models.Chat) error {
	c.Logger.Info("creating chat", "chat", chat)
	c.chats = append(c.chats, chat)
	return nil
}

func (c *chatRepository) AddMessage(chatId string, message models.Message) error {
	c.Logger.Info("adding message to chat", "chatId", chatId, "message", message)
	var foundChat *models.Chat
	for i := range c.chats {
		if c.chats[i].Id == chatId {
			foundChat = &c.chats[i]
		}
	}

	if foundChat == nil {
		return fmt.Errorf("chat not found")
	}

	messages := append([]models.Message{}, foundChat.Messages...)
	messages = append(messages, message)
	foundChat.Messages = messages

	return nil
}
