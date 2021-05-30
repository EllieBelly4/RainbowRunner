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

	sendGoToZone(conn, clientID, body, "dungeon00_level01")
}
