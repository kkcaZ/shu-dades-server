package domain

import "net"

type RouterUseCase interface {
	Handle(conn net.Conn) error
}
