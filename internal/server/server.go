package server

import (
	broadcastUc "github.com/kkcaz/shu-dades-server/internal/broadcast"
	"github.com/kkcaz/shu-dades-server/internal/config"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	r "github.com/kkcaz/shu-dades-server/internal/router"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	slog.Info("Initialising server")

	// Create context that listens for interrupt signal from the OS
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to initialise config: %v", err)
	}

	router, broadcaster, encryptor, err := Inject(cfg)
	if err != nil {
		log.Fatalf("failed to inject dependencies: %v", err)
	}

	server, err := net.Listen(cfg.Service.SocketType, cfg.Service.Host+":"+cfg.Service.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer server.Close()
	slog.Info("Listening on " + cfg.Service.Host + ":" + cfg.Service.Port)

	go func() {
		for {
			connect, err := server.Accept()
			if err != nil {
				slog.Error("failed to accept connection: " + err.Error())
				continue
			}

			slog.Info("Accepted connection from " + connect.RemoteAddr().String())
			go handleConnection(connect, router, *encryptor, broadcaster)
			if err != nil {
				slog.Error("failed to handle connection: " + err.Error())
			}
		}
	}()

	<-interrupt
	slog.Info("Shutting down server")
}

func handleConnection(conn net.Conn, router *r.RouterUseCase, encryptor domain.EncryptionUseCase, broadcaster *broadcastUc.BroadcastUseCase) {
	buffer := make([]byte, 1024)

	for {
		mLen, err := conn.Read(buffer)
		if err != nil {
			slog.Info("connection closed", "remoteAddress", conn.RemoteAddr().String())
			break
		}

		decryptedMessage, err := encryptor.Decrypt(buffer[:mLen])

		response, err := router.Handle(decryptedMessage, conn.RemoteAddr().String())
		if err != nil {
			slog.Error("failed to handle message", "error", err)
			continue
		}

		encryptedResponse, err := encryptor.Encrypt(*response)
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

	broadcaster.RemoveConnection(conn.RemoteAddr().String())
	err := conn.Close()
	if err != nil {
		slog.Error("failed to close connection", "error", err)
	}
}
