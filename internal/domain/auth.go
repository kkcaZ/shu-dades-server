package domain

import (
	"github.com/kkcaz/shu-dades-server/pkg/models"
	"net"
)

type AuthHandler interface {
	Handle(conn net.Conn) error
}

type AuthUseCase interface {
	Authenticate(username string, password string) (*models.UserClaim, error)
	TokenIsValid(token string) bool
}
