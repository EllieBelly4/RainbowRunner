package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg"
	byter "RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/math"
)

type ZoneChannelMessage byte

const (
	ZoneUnk0 ZoneChannelMessage = iota
	ZoneUnk1
	ZoneUnk2
	ZoneUnk3
	ZoneUnk4
	ZoneUnk5
	ZoneJoin
	ZoneUnk7
	ZoneUnk8
)

func handleZoneChannelMessages(conn *connections.RRConn, msgSubType uint8, reader *byter.Byter) error {
	switch ZoneChannelMessage(msgSubType) {
	case ZoneJoin:
		handleZoneJoin(conn)
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

func handleZoneJoin(conn *connections.RRConn) {
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

	player := objects.Players.Players[conn.GetID()]

	createNPC(conn, player.Zone, pkg.Transform{
		Position: pkg.Vector3{0, 0, 15000},
		Rotation: 180 * math.DRDegToRot,
	})

	createNPC(conn, player.Zone, pkg.Transform{
		Position: pkg.Vector3{20 * 256, 20 * 256, 15000},
		Rotation: 270 * math.DRDegToRot,
	})

	createNPC(conn, player.Zone, pkg.Transform{
		Position: pkg.Vector3{0 * 256, 40 * 256, 15000},
		Rotation: 360 * math.DRDegToRot,
	})

	createNPC(conn, player.Zone, pkg.Transform{
		Position: pkg.Vector3{-20 * 256, 20 * 256, 15000},
		Rotation: 90 * math.DRDegToRot,
	})

	player.CurrentCharacter.OnZoneJoin()

	// Creating Player Entity
	//objects.SendCreateNewPlayerEntity(conn, body)

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
	rrPlayer := objects.Players.Players[conn.GetID()]

	avatar := objects.Players.Players[conn.GetID()].CurrentCharacter.GetChildByGCNativeType("Avatar").(*objects.Avatar)
	if rrPlayer.Zone.Name == "town" {
		avatar.Warp(106342, -46263, 12778)
		avatar.SendPosition()
	} else if rrPlayer.Zone.Name == "dungeon16_level00" {
		avatar.Warp(0, 0, 15000)
		avatar.SendPosition()
	}

	avatar.SendFollowClient()
	avatar.IsSpawned = true

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

func createNPC(conn *connections.RRConn, zone *objects.Zone, transform pkg.Transform) {
	npc := objects.NewNPC("world.town.npc.HelperNoobosaur01")

	npc.WorldPosition = transform.Position
	npc.Rotation = transform.Rotation

	unitBehavior := objects.NewMonsterBehavior2("npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior")
	npc.AddChild(unitBehavior)

	skills := objects.NewSkills("skills")
	npc.AddChild(skills)

	manipulators := objects.NewManipulators("manipulators")
	npc.AddChild(manipulators)

	modifiers := objects.NewModifiers("modifiers")
	npc.AddChild(modifiers)

	zone.Spawn(npc)

	clientEntityWriter := objects.NewClientEntityWriterWithByter()

	clientEntityWriter.BeginStream()
	clientEntityWriter.Create(npc)
	clientEntityWriter.CreateComponent(skills, npc)
	clientEntityWriter.CreateComponent(manipulators, npc)
	clientEntityWriter.CreateComponent(modifiers, npc)
	// Adding unit behavior makes the NPC move in a random direction, missing something here
	//clientEntityWriter.CreateComponent(unitBehavior, npc)
	clientEntityWriter.Init(npc)
	clientEntityWriter.EndStream()

	connections.WriteCompressedASimple(conn, clientEntityWriter.Body)

	//unitBehavior.Warp(106342, -46263, 12778)
	//unitBehavior.Warp(0, 0, 0)
	//unitBehavior.SendPosition()
}

func sendGoToZone(conn *connections.RRConn, body *byter.Byter, zone string) {
	objects.Zones.PlayerJoin(zone, objects.Players.Players[conn.GetID()])

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(0x00)
	//body.WriteCString("TheHub")
	//body.WriteCString("Tutorial")
	body.WriteCString(zone)
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}
