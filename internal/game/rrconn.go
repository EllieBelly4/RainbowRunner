package game

import (
	"net"
)

type RRConn struct {
	NetConn     net.Conn
	ClientID    uint32
	Player      *RRConnClient
	IsConnected bool
}
