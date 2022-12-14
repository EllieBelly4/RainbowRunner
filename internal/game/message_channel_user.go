package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	byter "RainbowRunner/pkg/byter"
)

func handleUserChannelMessages(conn *connections.RRConn, msgSubType byte, reader *byter.Byter) error {
	switch msgSubType {
	case 0x00: // Request rosters
		handleUserUnk0(conn)
	case 0x01: // Rosters response
		handleUserUnk1(conn)
	default:
		return UnhandledChannelMessageError
	}

	return nil
}

func handleUserUnk0(conn *connections.RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.UserChannel))
	body.WriteByte(0x00)
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func handleUserUnk1(conn *connections.RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.UserChannel))
	body.WriteByte(0x01)

	body.WriteByte(0x01) // Unk
	body.WriteByte(0x01) // Unk
	body.WriteInt32(1)   // Some count, Must be non-negative, non-0 value

	body.WriteCString("Hello")

	//body.WriteInt32(0x10000001) // Unk
	body.WriteInt32(1) // Friend count

	body.WriteCString("Goodbye") // Friend Name
	body.WriteByte(0x01)         // Unk
	body.WriteUInt32(0x01)       // Unk

	body.WriteCString("ILikeTrains")
	body.WriteCString("AndBrains")

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}
