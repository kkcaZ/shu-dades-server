package domain

import "net"

type AuthHandler interface {
	Handle(conn net.Conn) error
}

type AuthUseCase interface {
	Authenticate(username string, password string) (*string, error)
	TokenIsValid(token string) bool
}
