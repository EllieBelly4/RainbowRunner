package game

import (
	"RainbowRunner/internal/byter"
	"net"
)

func handleUserUnk0(conn net.Conn, clientID uint32) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(UserChannel))
	body.WriteByte(0x00)
	WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
}

func handleUserUnk1(conn net.Conn, clientID uint32) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(UserChannel))
	body.WriteByte(0x01)

	body.WriteByte(0x01) // Unk
	body.WriteByte(0x01) // Unk
	body.WriteInt32(1)   // Some count, Must be non-negative, non-0 value

	body.WriteCString("Hello")

	//body.WriteInt32(0x10000001) // Unk
	body.WriteInt32(1) // Some count, Must be non-negative, non-0 value

	body.WriteCString("Goodbye")
	body.WriteByte(0x01)   // Unk
	body.WriteUInt32(0x01) // Unk

	body.WriteCString("ILikeTrains")
	body.WriteCString("AndBrains")

	WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
}
