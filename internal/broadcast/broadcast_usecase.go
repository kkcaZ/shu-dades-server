package broadcast

import (
	"encoding/json"
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"github.com/pkg/errors"
	"log/slog"
	"net"
)

type BroadcastUseCase struct {
	Logger      slog.Logger
	Connections []models.BroadcastConnection
}

func NewBroadcastUseCase(logger slog.Logger) *BroadcastUseCase {
	return &BroadcastUseCase{
		Logger:      logger,
		Connections: make([]models.BroadcastConnection, 0),
	}
}

func (b *BroadcastUseCase) PublishToUsers(message string, eventType string, users []string) error {
	formattedMessage := models.BroadcastRequest{
		Message: message,
		Type:    eventType,
	}
	messageBytes, err := json.Marshal(formattedMessage)
	if err != nil {
		return err
	}

	for _, user := range users {
		b.Logger.Info("sending message to user", "message", message, "user", user)
		for _, conn := range b.Connections {
			if conn.UserId == user {
				err := b.publishMessage(messageBytes, conn.SubscribeAddress)
				if err != nil {
					b.Logger.Error("failed to publish message", "error", err)
					continue
				}
			}
		}
	}

	return nil
}

func (b *BroadcastUseCase) publishMessage(msg []byte, address string) error {
	b.Logger.Info("sending message", "message", string(msg), "remoteAddress", address)
	connClient, err := net.Dial("tcp", address)
	if err != nil {
		return errors.Wrapf(err, "failed to dial connection: ")
	}

	_, err = connClient.Write(msg)
	if err != nil {
		return errors.Wrapf(err, "failed to write message: ")
	}

	return nil
}

func (b *BroadcastUseCase) AddConnection(subscribeAddress string, publishAddress string) {
	b.Logger.Info("adding connection to broadcast use case", "subscribeAddress", subscribeAddress, "publishAddress", publishAddress)
	b.Connections = append(b.Connections, models.BroadcastConnection{
		SubscribeAddress: subscribeAddress,
		PublishAddress:   publishAddress,
	})
}

func (b *BroadcastUseCase) RemoveConnection(addr string) {
	b.Logger.Info("removing connection from broadcast use case", "address", addr)
	for i, conn := range b.Connections {
		if conn.PublishAddress == addr || conn.SubscribeAddress == addr {
			b.Connections = append(b.Connections[:i], b.Connections[i+1:]...)
		}
	}
}

func (b *BroadcastUseCase) RegisterUser(addr string, userId string) {
	b.Logger.Info("registering user to broadcast use case", "address", addr, "userId", userId)
	for i, conn := range b.Connections {
		if conn.PublishAddress == addr {
			b.Connections[i].UserId = userId
		}
	}
}

func (b *BroadcastUseCase) RemoveUser(addr string) {
	b.Logger.Info("removing user from broadcast use case", "address", addr)
	for i, conn := range b.Connections {
		if conn.PublishAddress == addr {
			b.Connections[i].UserId = ""
		}
	}
}
