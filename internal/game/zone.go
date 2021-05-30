package game

import (
	"RainbowRunner/internal/byter"
)

type ZoneChannelMessage byte

const (
	ZoneUnk0 ZoneChannelMessage = iota
	ZoneUnk1
	ZoneUnk2
	ZoneUnk3
	ZoneUnk4
	ZoneUnk5
	ZoneUnk6
	ZoneUnk7
	ZoneUnk8
)

func handleZoneChannelMessages(conn *RRConn, msgSubType uint8, reader *byter.Byter) error {
	switch ZoneChannelMessage(msgSubType) {
	case ZoneUnk6:
		handleZoneUnk6(conn)
	case ZoneUnk8:
		body := byter.NewLEByter(make([]byte, 0, 1024))
		body.WriteByte(byte(ZoneChannel))
		body.WriteByte(0x08)
		WriteCompressedA(conn, 0x01, 0x0f, body)
	default:
		return UnhandledChannelMessageError
	}
	return nil
}

func handleZoneUnk6(conn *RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ZoneChannel))
	body.WriteByte(0x01)
	//body.WriteByte(0x02) // Other acceptable values
	//body.WriteByte(0x05) // Other acceptable values
	body.WriteUInt32(0xFEEDBABA) // One of these is the world ID?
	WriteCompressedA(conn, 0x01, 0x0f, body)

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ZoneChannel))
	body.WriteByte(0x05)

	// Adds two separate values into the ZoneClient
	body.WriteUInt32(0xFEEDBABA) // One of these is the world ID?
	body.WriteUInt32(0xFEEDBABA) // One of these is the world ID?
	WriteCompressedA(conn, 0x01, 0x0f, body)

	// Creating Player Entity
	sendCreateNewPlayerEntity(conn, body)
}

func sendGoToZone(conn *RRConn, body *byter.Byter, zone string) {
	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ZoneChannel))
	body.WriteByte(0x00)
	//body.WriteCString("TheHub")
	//body.WriteCString("Tutorial")
	body.WriteCString(zone)
	WriteCompressedA(conn, 0x01, 0x0f, body)
}
