package connections

import (
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/pkg/byter"
	"net"
)

type RRConn struct {
	NetConn       net.Conn
	Client        *RRConnClient
	IsConnected   bool
	MessageBuffer *byter.Byter
	LoginName     string
}

func (R *RRConn) Send(b *byter.Byter) error {
	_, err := R.NetConn.Write(b.Data())

	return err
}

func (R *RRConn) GetID() int {
	return int(R.Client.ID)
}

func (R *RRConn) SendMessage(message messages.DRMessage) {
	R.MessageBuffer.Clear()

	message.Write(R.MessageBuffer)
	WriteCompressedASimple(R, R.MessageBuffer)
}

func NewRRConn(conn net.Conn) *RRConn {
	return &RRConn{
		NetConn:       conn,
		IsConnected:   true,
		MessageBuffer: byter.NewLEByter(make([]byte, 1024*1000)),
	}
}
