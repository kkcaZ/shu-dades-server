package usecase

import (
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"log/slog"
	"net"
)

type routerUseCase struct {
	Logger slog.Logger
}

func NewRouterUseCase(logger slog.Logger) domain.RouterUseCase {
	return &routerUseCase{
		Logger: logger,
	}
}

func (r *routerUseCase) Handle(conn net.Conn) error {
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		r.Logger.Error("failed to read from connection: " + err.Error())
		return err
	}

	r.Logger.Info("received message: " + string(buffer[:mLen]))

	_, err = conn.Write([]byte("Hello from server"))
	if err != nil {
		return err
	}

	_ = conn.Close()
	return nil
}
