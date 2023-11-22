package front_controller

import (
	broadcastUc "github.com/kkcaz/shu-dades-server/internal/broadcast"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	routerUc "github.com/kkcaz/shu-dades-server/internal/router"
	"log/slog"
	"net"
)

type frontController struct {
	Router      routerUc.RouterUseCase
	Encryptor   domain.EncryptionUseCase
	Broadcaster *broadcastUc.BroadcastUseCase
}

func NewFrontController(router routerUc.RouterUseCase, encryptor domain.EncryptionUseCase, broadcaster *broadcastUc.BroadcastUseCase) domain.FrontController {
	return &frontController{
		Router:      router,
		Encryptor:   encryptor,
		Broadcaster: broadcaster,
	}
}

func (f *frontController) HandleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)

	for {
		mLen, err := conn.Read(buffer)
		if err != nil {
			slog.Info("connection closed", "remoteAddress", conn.RemoteAddr().String())
			break
		}

		decryptedMessage, err := f.Encryptor.Decrypt(buffer[:mLen])
		slog.Info("received message: " + string(decryptedMessage))

		response, err := f.Router.Handle(decryptedMessage, conn.RemoteAddr().String())
		if err != nil {
			slog.Error("failed to handle message", "error", err)
			continue
		}

		slog.Info("sending message", "message", *response, "remoteAddress", conn.RemoteAddr().String())

		encryptedResponse, err := f.Encryptor.Encrypt(*response)
		slog.Info(encryptedResponse)
		if err != nil {
			slog.Error("failed to encrypt response", "error", err)
			continue
		}

		_, err = conn.Write([]byte(encryptedResponse))
		if err != nil {
			slog.Error("failed to write to connection", "error", err)
			continue
		}
	}

	f.Broadcaster.RemoveConnection(conn.RemoteAddr().String())
	err := conn.Close()
	if err != nil {
		slog.Error("failed to close connection", "error", err)
	}
}
