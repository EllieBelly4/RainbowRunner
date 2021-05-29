package game

import (
	"RainbowRunner/internal/byter"
	"net"
)

type GroupChannelMessage byte

const (
	GroupConnected GroupChannelMessage = iota
)

func handleGroupConnected(conn net.Conn, clientID uint32) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(GroupChannel))
	body.WriteByte(48)

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ZoneChannel))
	body.WriteByte(0x00)
	//body.WriteCString("TheHub")
	//body.WriteCString("Tutorial")
	body.WriteCString("Town")
	WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
}
