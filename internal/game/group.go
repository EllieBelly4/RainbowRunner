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

	//sendGoToZone(conn, body, "dungeon00_level01")
	//sendGoToZone(conn, body, "TestTilesets")
	//sendGoToZone(conn, body, "thehub")
	//sendGoToZone(conn, body, "dungeon01_level01")
	//sendGoToZone(conn, body, "dungeon15_level01")
	sendGoToZone(conn, body, "dungeon16_level00")
	//sendGoToZone(conn, body, "town")
	//sendGoToZone(conn, body, "dungeon02_level01")
}
