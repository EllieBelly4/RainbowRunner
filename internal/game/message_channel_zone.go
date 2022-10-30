package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/helpers"
	"RainbowRunner/internal/objects"
	byter "RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
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
		helpers.WriteCompressedA(conn, 0x01, 0x0f, body)
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

	helpers.WriteCompressedA(conn, 0x01, 0x0f, body)

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(0x05)

	// Adds two separate values into the ZoneClient
	body.WriteUInt32(0x01)
	body.WriteUInt32(0x01)
	helpers.WriteCompressedA(conn, 0x01, 0x0f, body)

	player := objects.Players.Players[conn.GetID()]

	entitiesToSpawn := [][]string{
		{"npc.Avatar.Female.base.NPC_Amazon1_Base", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Amazon_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Fighter_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Gladiator_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Mage_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Mage_002", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Ninja_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Officer_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Ranger_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Ranger_002", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Ranger_003", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"npc.Avatar.Female.Basic.Scout_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
		//{"world.town.npc.TownCommander", "world.town.npc.TownCommander.Behavior"},
		//{"world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior"},
		//{"world.dungeon16.mob.boss_manager01", "world.dungeon16.mob.boss_manager01.Behavior"},
		//{"world.dungeon15.mob.boss", "world.dungeon15.mob.boss.Behavior"},
	}

	SendInterval(conn)

	player.CurrentCharacter.OnZoneJoin(player)

	if player.Zone.Name == "town" {
		for i, entityStrings := range entitiesToSpawn {
			objects.CreateNPC(player, player.Zone, datatypes.Transform{
				Position: datatypes.Vector3{106342 + 2048*int32(i), -36000, 12778},
				Rotation: 180 * math.DRDegToRot,
			}, entityStrings[0], entityStrings[1])
		}

	} else if player.Zone.Name == "dungeon16_level00" {
		objects.CreateNPC(player, player.Zone, datatypes.Transform{
			Position: datatypes.Vector3{0, 0, 15000},
			Rotation: 180 * math.DRDegToRot,
		}, "world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior")

		objects.CreateNPC(player, player.Zone, datatypes.Transform{
			Position: datatypes.Vector3{20 * 256, 20 * 256, 15000},
			Rotation: 270 * math.DRDegToRot,
		}, "world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior")

		objects.CreateNPC(player, player.Zone, datatypes.Transform{
			Position: datatypes.Vector3{0 * 256, 40 * 256, 15000},
			Rotation: 360 * math.DRDegToRot,
		}, "world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior")

		objects.CreateNPC(player, player.Zone, datatypes.Transform{
			Position: datatypes.Vector3{-20 * 256, 20 * 256, 15000},
			Rotation: 90 * math.DRDegToRot,
		}, "world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior")
	}

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
		//avatar.SendPosition()
	} else if rrPlayer.Zone.Name == "dungeon16_level00" {
		avatar.Warp(0, 0, 15000)
		//avatar.SendPosition()
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

func sendGoToZone(conn *connections.RRConn, body *byter.Byter, zone string) {
	objects.Zones.PlayerJoin(zone, objects.Players.Players[conn.GetID()])

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(0x00)
	//body.WriteCString("TheHub")
	//body.WriteCString("Tutorial")
	body.WriteCString(zone)
	body.WriteUInt32(0x01)
	body.WriteByte(0x01)
	body.WriteByte(0xFF)
	body.WriteCString("world.town.quest.Q01_a1")
	body.WriteUInt32(0x01)
	helpers.WriteCompressedA(conn, 0x01, 0x0f, body)
}