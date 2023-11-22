package domain

import "net"

type FrontController interface {
	HandleConnection(conn net.Conn)
}
