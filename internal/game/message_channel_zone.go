package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/serverconfig"
	byter "RainbowRunner/pkg/byter"
	log "github.com/sirupsen/logrus"
	"strings"
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
	player := objects.Players.GetPlayer(uint16(conn.GetID()))

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageReady))
	//body.WriteByte(0x02) // Other acceptable values
	//body.WriteByte(0x05) // Other acceptable values
	body.WriteUInt32(player.Zone().ID) // World ID

	// MiniMapExplored::ReadExploredBits
	// dungeon00_level01 - 0x12
	exploredBitCount := uint16(0x12)
	body.WriteUInt16(exploredBitCount)

	for i := 0; i < int(exploredBitCount); i++ {
		body.WriteUInt32(0xFFFFFFFF)
	}

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageInstanceCount))

	// Adds two separate values into the ZoneClient
	body.WriteUInt32(0x01)
	body.WriteUInt32(0x01)
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	//entitiesToSpawn := [][]string{
	//	{"npc.Avatar.Female.base.NPC_Amazon1_Base", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Amazon_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Fighter_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Gladiator_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Mage_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Mage_002", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Ninja_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Officer_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Ranger_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Ranger_002", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Ranger_003", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"npc.Avatar.Female.Basic.Scout_001", "npc.Avatar.Female.base.NPC_Amazon1_Base.Behavior"},
	//	//{"world.town.npc.TownCommander", "world.town.npc.TownCommander.Behavior"},
	//	//{"world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior"},
	//	//{"world.dungeon16.mob.boss_manager01", "world.dungeon16.mob.boss_manager01.Behavior"},
	//	//{"world.dungeon15.mob.boss", "world.dungeon15.mob.boss.Behavior"},
	//}

	SendInterval(conn)

	player.CurrentCharacter.SendCreateNewPlayerEntity(player)
	player.CurrentCharacter.OnZoneJoin()

	if serverconfig.Config.Welcome.SendWelcomeMessage {
		SendWelcomeMessage(player)
	}

	if player.Zone().Name == "town" {
		//for i, entityStrings := range entitiesToSpawn {
		//	objects.CreateNPC(player, player.Zone, datatypes.Transform{
		//		Position: datatypes.Vector3Float32{float32(106342+2048*int32(i)) / 256, -140, 49},
		//		Rotation: 180,
		//	}, entityStrings[0], entityStrings[1])
		//}

	}
	//else if player.Zone.Name == "dungeon16_level00" {
	//	objects.CreateNPC(player, player.Zone, datatypes.Transform{
	//		Position: datatypes.Vector3Float32{0, 0, 15000},
	//		Rotation: 180 * math.DRDegToRot,
	//	}, "world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior")
	//
	//	objects.CreateNPC(player, player.Zone, datatypes.Transform{
	//		Position: datatypes.Vector3Float32{20 * 256, 20 * 256, 15000},
	//		Rotation: 270 * math.DRDegToRot,
	//	}, "world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior")
	//
	//	objects.CreateNPC(player, player.Zone, datatypes.Transform{
	//		Position: datatypes.Vector3Float32{0 * 256, 40 * 256, 15000},
	//		Rotation: 360 * math.DRDegToRot,
	//	}, "world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior")
	//
	//	objects.CreateNPC(player, player.Zone, datatypes.Transform{
	//		Position: datatypes.Vector3Float32{-20 * 256, 20 * 256, 15000},
	//		Rotation: 90 * math.DRDegToRot,
	//	}, "world.town.npc.HelperNoobosaur01", "npc.misc.HelperNoobosaur.base.HelperNoobosaur_Base.Behavior")
	//}

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

	avatar := player.CurrentCharacter.GetChildByGCNativeType("Avatar").(*objects.Avatar)
	lcZoneName := strings.ToLower(player.Zone().Name)
	if lcZoneName == "town" {
		avatar.Warp(106342/256, -46263/256, 12778/256)
		//avatar.SendPosition()
	} else if lcZoneName == "dungeon03_level00" {
		avatar.Warp(100, 150, 7700/256)
		//avatar.SendPosition()
	} else if lcZoneName == "dungeon16_level00" {
		avatar.Warp(0, 0, 15000/256)
	} else if lcZoneName == "dungeon00_level01" {
		avatar.Warp(-200, 0, 10)
	} else if lcZoneName == "dungeon02_level00" {
		avatar.Warp(-150, 500, 2700/256)
	} else if lcZoneName == "dungeon04_level00" {
		avatar.Warp(100, 500, 2700/256)
	} else if lcZoneName == "dungeon05_level00" {
		avatar.Warp(0, -50, 10000/256)
	} else if lcZoneName == "dungeon06_level00" {
		avatar.Warp(600, 0, 6500/256)
	} else if lcZoneName == "dungeon09_level00" {
		avatar.Warp(75, -50, 12500/256)
	} else if lcZoneName == "dungeon11_level00" {
		avatar.Warp(75, 150, 2000/256)
	} else if lcZoneName == "dungeon15_level00" {
		avatar.Warp(75, 150, 2000/256)
	} else if lcZoneName == "dungeon15_level01" {
		avatar.Warp(-200, 150, 20)
	} else if lcZoneName == "tutorial" {
		avatar.Warp(750, 450, 10000/256)
	} else if lcZoneName == "testvendorlevelspecarmor" {
		avatar.Warp(0, 100, 5000/256)
	} else if lcZoneName == "thehubportals_dungeon01" {
		avatar.Warp(600, -100, 5000/256)
	} else if lcZoneName == "dungeon08_level00" {
		avatar.Warp(0, 0, 5000/256)
	} else if lcZoneName == "pvp_start" {
		avatar.Warp(-200, -200, 5000/256)
	} else if lcZoneName == "dungeon02_level08_boss" {
		avatar.Warp(-200, -200, 5000/256)
	} else if lcZoneName == "d06_l01_q05" {
		avatar.Warp(150, -200, 15000/256)
	} else if lcZoneName == "d06_l07_q05" {
		avatar.Warp(150, -200, 15000/256)
	} else if lcZoneName == "epic01_central" {
		avatar.Warp(-250, 250, 75000/256)
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

func sendGoToZone(conn *connections.RRConn, zoneName string) {
	rrPlayer := objects.Players.Players[conn.GetID()]

	tZone := objects.Zones.GetOrCreateZone(zoneName)

	if tZone == nil {
		log.Errorf("could not find zone %s", zoneName)
		return
	}

	rrPlayer.CurrentCharacter.JoinZone(tZone)
}

func SendWelcomeMessage(player *objects.RRPlayer) {
	msg := messages.ChatMessage{
		Channel: messages.MessageChannelSourceGlobalAnnouncement,
		Message: serverconfig.Config.Welcome.Message,
	}

	player.Conn.SendMessage(msg)
}
