package game

import (
	"net"
)

type RRConn struct {
	NetConn     net.Conn
	Client      *RRConnClient
	IsConnected bool
}
