package game

import (
	"RainbowRunner/internal/byter"
)

type GroupChannelMessage byte

const (
	GroupConnected GroupChannelMessage = iota
)

func handleGroupConnected(conn *RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(GroupChannel))
	body.WriteByte(48)

	sendGoToZone(conn, body, "dungeon00_level01")
}
