package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	byter "RainbowRunner/pkg/byter"
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
		body.WriteByte(byte(messages.ZoneChannel))
		body.WriteByte(0x08)
		connections.WriteCompressedA(conn, 0x01, 0x0f, body)
	default:
		return UnhandledChannelMessageError
	}
	return nil
}

func handleZoneUnk6(conn *RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(0x01)
	//body.WriteByte(0x02) // Other acceptable values
	//body.WriteByte(0x05) // Other acceptable values
	body.WriteUInt32(0xFEEDBABA) // World ID

	// MiniMapExplored::ReadExploredBits
	// dungeon00_level01 - 0x12
	exploredBitCount := uint16(0x12)
	body.WriteUInt16(exploredBitCount)

	for i := 0; i < int(exploredBitCount); i++ {
		body.WriteUInt32(0x01)
	}

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(0x05)

	// Adds two separate values into the ZoneClient
	body.WriteUInt32(0x01)
	body.WriteUInt32(0x01)
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	// Creating Player Entity
	sendCreateNewPlayerEntity(conn, body)

	//	WriteCompressedA(conn, 0x01, 0x0f, NewLEByterFromCommandString(`
	//07
	//# ClientEntityManager::processInterval
	//# This seems to log messages about the pathManager budget
	//# I don't think this is meant to sync each tick
	//0D # ID
	//
	//# ClientEntityManager::processInterval
	//01 00 00 00
	//01 00 00 00
	//01 00 00 00
	//
	//# PathManager::ReadBudget
	//# These values are syncing the path budget
	//FF FF FF FF
	//FF FF # Per Update (idk)
	//FF FF # Per Path (idk)
	//
	//06 #end`))

	if conn.Client.Zone == "town" {
		conn.Client.CurrentCharacter.Warp(106342, -46263, 12778)
		conn.Client.CurrentCharacter.SendPosition()
	} else if conn.Client.Zone == "dungeon16_level00" {
		conn.Client.CurrentCharacter.Warp(0, 0, 15000)
		conn.Client.CurrentCharacter.SendPosition()
	}

	conn.Client.CurrentCharacter.SendFollowClient()
	conn.Client.CurrentCharacter.IsSpawned = true

	////	Some client control message
	//body = byter.NewLEByter(make([]byte, 0, 1024))
	//
	//unitBehavior := behavior.NewUnitBehavior(0x05)
	//unitBehavior.Action = &behavior.MoveTo{
	//	PosX: 10,
	//	PosY: 10,
	//}
	//
	//AddComponentUpdate(body, unitBehavior)
	//AddSynch(conn, body)
	//// END STREAM /////////////////////////////////
	//AddEntityUpdateStreamEnd(body)

	//cmd := NewLEByterFromCommandString(`
	//07 # ClientEntityChannel
	//# UnitBehavior - Activate::readData
	//35 # ComponentUpdate
	//05 00 # Component ID
	//# Command
	//04 # CreateAction1
	//06 # ActionID
	//
	//# Activate::readData
	//01
	//02 00
	//`)
	//
	//AddSynch(conn, cmd)
	//cmd.WriteByte(0x06) // End
	//WriteCompressedA(conn, 0x01, 0x0f, cmd)

	//WriteCompressedA(conn, 0x01, 0x0f, body)
}

func sendGoToZone(conn *RRConn, body *byter.Byter, zone string) {
	conn.Client.Zone = zone

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(0x00)
	//body.WriteCString("TheHub")
	//body.WriteCString("Tutorial")
	body.WriteCString(zone)
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}
