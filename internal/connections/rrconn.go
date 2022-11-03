package connections

import (
	"RainbowRunner/pkg/byter"
	"net"
)

type RRConn struct {
	NetConn     net.Conn
	Client      *RRConnClient
	IsConnected bool
}

func (R *RRConn) Send(b *byter.Byter) error {
	_, err := R.NetConn.Write(b.Data())

	return err
}

func (R *RRConn) GetID() int {
	return int(R.Client.ID)
}
