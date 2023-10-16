package server

import (
	"github.com/kkcaz/shu-dades-server/internal/config"
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

	router, err := Inject(cfg)
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
			err = router.Handle(connect)
			if err != nil {
				slog.Error("failed to handle connection: " + err.Error())
			}
		}
	}()

	<-interrupt
	slog.Info("Shutting down server")
}
