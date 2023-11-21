package domain

import "github.com/kkcaz/shu-dades-server/pkg/models"

type ChatUseCase interface {
	GetChatThumbnails(userId string) ([]models.ChatThumbnail, error)
	GetChat(chatId string) (*models.Chat, error)
	GetChatParticipantIds(chatId string) ([]string, error)
	CreateChat(participants []string) (*models.Chat, error)
	SendMessage(chatId string, message string, userId string) error
}

type ChatRepository interface {
	GetAllChatThumbnails() ([]models.ChatThumbnail, error)
	GetChat(chatId string) (*models.Chat, error)
	CreateChat(chat models.Chat) error
	AddMessage(chatId string, message models.Message) error
}
