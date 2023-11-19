package models

import "time"

type Chat struct {
	Id           string        `json:"id"`
	Participants []Participant `json:"participants"`
	Messages     []Message     `json:"messages"`
}

type ChatResponse struct {
	StatusCode int   `json:"statusCode"`
	Chat       *Chat `json:"chat"`
}

type Participant struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
}

type Message struct {
	UserId   string    `json:"userId"`
	Username string    `json:"username"`
	Content  string    `json:"content"`
	SentAt   time.Time `json:"sentAt"`
}

type MessageEvent struct {
	Message
	ChatId string `json:"chatId"`
}

type ChatThumbnailsResponse struct {
	StatusCode int             `json:"statusCode"`
	Chats      []ChatThumbnail `json:"chats"`
}

type ChatThumbnail struct {
	ChatId       string        `json:"chatId"`
	Participants []Participant `json:"participants"`
	LastMessage  *Message      `json:"lastMessage"`
}

type CreateChatRequest struct {
	UserIds []string `json:"userIds"`
}

type SendMessageRequest struct {
	ChatId  string `json:"chatId"`
	Message string `json:"message"`
}
