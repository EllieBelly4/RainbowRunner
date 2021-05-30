package game

import (
	"RainbowRunner/internal/byter"
)

type GroupChannelMessage byte

const (
	GroupConnected GroupChannelMessage = iota
)

func handleGroupChannelMessages(conn *RRConn, msgType byte, reader *byter.Byter) error {
	switch GroupChannelMessage(msgType) {
	case GroupConnected:
		handleGroupConnected(conn)
	default:
		return UnhandledChannelMessageError
	}

	return nil
}

func handleGroupConnected(conn *RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(GroupChannel))
	body.WriteByte(48)

	sendGoToZone(conn, body, "dungeon00_level01")
}
