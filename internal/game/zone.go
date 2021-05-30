package game

import (
	"RainbowRunner/internal/byter"
	"net"
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
)

func handleZoneUnk6(conn net.Conn, clientID uint32) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ZoneChannel))
	body.WriteByte(0x01)
	//body.WriteByte(0x02) // Other acceptable values
	//body.WriteByte(0x05) // Other acceptable values
	body.WriteUInt32(0xFEEDBABA) // One of these is the world ID?
	WriteCompressedA(clientID, 0x01, 0x0f, body, conn)

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ZoneChannel))
	body.WriteByte(0x05)

	// Adds two separate values into the ZoneClient
	body.WriteUInt32(0xFEEDBABA) // One of these is the world ID?
	body.WriteUInt32(0xFEEDBABA) // One of these is the world ID?
	WriteCompressedA(clientID, 0x01, 0x0f, body, conn)

	// Creating Player Entity
	sendCreateNewPlayerEntity(conn, clientID, body)
}

func sendGoToZone(conn net.Conn, clientID uint32, body *byter.Byter, zone string) {
	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ZoneChannel))
	body.WriteByte(0x00)
	//body.WriteCString("TheHub")
	//body.WriteCString("Tutorial")
	body.WriteCString(zone)
	WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
}
