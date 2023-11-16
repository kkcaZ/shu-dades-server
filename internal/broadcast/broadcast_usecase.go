package broadcast

import (
	"github.com/kkcaz/shu-dades-server/pkg/models"
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

func (b *BroadcastUseCase) Publish(message string, sender string) error {
	for _, conn := range b.Connections {
		b.Logger.Info("sending message", "message", message, "remoteAddress", conn)
		connClient, err := net.Dial("tcp", conn.SubscribeAddress)
		if err != nil {
			b.Logger.Error("failed to dial connection", "error", err)
			continue
		}

		_, err = connClient.Write([]byte(message))
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
