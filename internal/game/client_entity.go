package game

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/game/components"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/logging"
	"encoding/hex"
	"fmt"
)

type ClientEntityMessage byte

const (
	ClientEntityUnk0 ClientEntityMessage = iota
	ClientEntityUnk1
	ClientEntityUnk2
	ClientEntityUnk3
	ClientRequestRespawn
	ClientEntityUnk5
	ClientEntityUnk6
	ClientEntityUnk7
	ClientEntityUnk8
	ClientEntityUnk9
	ClientEntityThings   = 0x34
	ClientEntityMovement = 0x35
)

func handleClientEntityChannelMessages(conn *RRConn, msgType byte, reader *byter.Byter) error {
	switch ClientEntityMessage(msgType) {
	case ClientRequestRespawn:
		handleClientEntityUnk4(conn, reader)
	case ClientEntityThings:
		clientEntitySubMessage := reader.UInt16()
		switch clientEntitySubMessage {
		// Inventory Message?
		case 0x01:
			inventoryMessageType := reader.UInt8()
			switch inventoryMessageType {
			case 0x21:
				fmt.Printf("Player opened inventory\n%s", hex.Dump(reader.Data()))
			case 0x22:
				fmt.Printf("Player closed inventory\n%s", hex.Dump(reader.Data()))
			default:
				fmt.Printf("unhandled inventory message %x", inventoryMessageType)
				return UnhandledChannelMessageError
			}
		case 0x04:
			fmt.Printf("Player tried to put something on hotbar\n%s", hex.Dump(reader.Data()))
		case 0x05:
			return handleClientEntityMovement(conn, reader)
		default:
			fmt.Printf("unhandled client entity sub message %x", clientEntitySubMessage)
			return UnhandledChannelMessageError
		}
	default:
		return UnhandledChannelMessageError
	}
	return nil
}

func handleClientEntityMovement(conn *RRConn, reader *byter.Byter) error {
	subMessage := reader.Byte()
	switch subMessage {
	case 0x65:
		// This increments each time the server sends a MoveTo message
		// The client will then increment by 1 for every individual movement performed (clicking)
		updateNumber := reader.Byte()
		count := int(reader.Byte())
		pos := Vector3{}

		if logging.LoggingOpts.LogMoves {
			fmt.Printf("Received %d player moves unk val: %x\n", count, updateNumber)
		}

		for i := 0; i < count; i++ {
			unk := reader.Byte()       // Unk
			rotation := reader.Int32() // Seems to be rotation

			degrees := float32((rotation / 0x17000) / 360)

			pos.X = reader.Int32()
			pos.Y = reader.Int32()

			conn.Player.Rotation = rotation

			conn.Player.LastMovementRequest = pos
			conn.Player.ClientUpdateNumber = updateNumber
			if logging.LoggingOpts.LogMoves {
				fmt.Printf(
					"Player move 0x%x rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
					unk, rotation, degrees, pos.X, pos.Y, pos.X, pos.Y,
				)
			}

			conn.Player.Position = pos
			conn.Player.SendPosition()
			conn.Player.MoveUpdate++

			if unk&0x02 > 0 {
				if logging.LoggingOpts.LogMoves {
					fmt.Println("player started moving")
				}
				conn.Player.IsMoving = true
			}

			if unk&0x01 > 0 {
				if logging.LoggingOpts.LogMoves {
					fmt.Println("player finished moving")
				}
				conn.Player.ServerUpdateNumber++
				conn.Player.IsMoving = false
			}
		}

		if conn.Player.MoveUpdate >= 0x2D {
			//fmt.Printf(
			//	"sending move update %d, %d || %x, %x!\n",
			//	pos.X, pos.Y,
			//	pos.X, pos.Y,
			//)
			//conn.Player.Move(pos.X, pos.Y)
			//conn.Player.SendFollowClient()
			conn.Player.MoveUpdate = 0
		}

		if logging.LoggingOpts.LogMoves {
			fmt.Printf("%s\n", hex.Dump(reader.Data()))
		}
	// Potentially requesting current position because starting a new path
	case 0x03:
		conn.Player.SendPosition()
	default:
		fmt.Printf("unhandled client entity sub message %x", subMessage)
		return UnhandledChannelMessageError
	}

	return nil
}

func handleClientEntityUnk4(conn *RRConn, reader *byter.Byter) {
	id := reader.UInt16()
	event := reader.Byte() // Guessing here

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ClientEntityChannel))
	// AVATAR UPDATE /////////////////////////////////////
	// This update is required to make the character alive
	body.WriteByte(0x03) // Update
	body.WriteUInt16(id) // ID

	// Avatar::processUpdate
	// 0x15 is special Avatar::processUpdate case(spawn entity?) anything else goes to Hero::processUpdate
	// Hero::processUpdate
	// 0x08 is Unit::processUseItemUpdate
	// 0x00 Hero::processUpdateAddExperience
	// 0x01 Hero::processUpdateRemoveExperience
	// 0x02 Hero::processUpdateSpendAttribPoint
	// 0x03 Hero::processUpdateReturnAttribPoint
	// 0x04 Hero::processUpdateRespectAttrbutes
	body.WriteByte(event)

	// EntitySynchInfo::ReadFromStream
	AddSynch(conn, body)

	AddEntityUpdateStreamEnd(body)

	WriteCompressedA(conn, 0x01, 0x0f, body)
}

func sendCreateNewPlayerEntity(conn *RRConn, body *byter.Byter) {
	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ClientEntityChannel))
	//body.WriteByte(0x01) // Create
	body.WriteByte(0x01) // CreateInit
	//body.WriteByte(0x02) // Init
	//body.WriteByte(0x03) // Update
	//body.WriteByte(21) // ClearEntityManager
	body.WriteUInt16(0x0001) // Entity ID
	// Type of lookup?
	// 0x04 by ID?
	// 0xFF by string?
	body.WriteByte(0xFF)
	// Examples
	// 0x0002
	// 0x1ADE
	// 0x1b08
	// 0x1af3
	body.WriteCString("Player") // Unk, might be used to lookup GCObject Type in registry

	// Init PLAYER /////////////////////////////////////////
	body.WriteByte(0x02)   // Init
	body.WriteUInt16(0x01) // ID
	body.WriteCString("Ellie")
	body.WriteUInt32(0x01)
	body.WriteUInt32(0x01)
	body.WriteByte(0x01)

	body.WriteUInt32(0xFEEDBABA) // World ID
	body.WriteUInt32(1001)       // PvP wins
	body.WriteUInt32(1000)       // PvP rating?, 0 = ???

	// Here goes PvP Team
	// Null string
	body.WriteByte(0x00)

	// If player is in a PvP team then Avatar respawn will look for the team waypoints
	//body.WriteByte(0xFF)
	//body.WriteCString("pvp.DefaultTeamList.BlueTeam")

	body.WriteCString("Hello")
	body.WriteUInt32(0x01)

	// UPDATE PLAYER /////////////////////////////////////////
	body.WriteByte(0x03)   // MsgType Update
	body.WriteUInt16(0x01) // Entity ID

	// This maps to a specific event type for Player::processUpdate()
	// 0x01 - do nothing
	// 0x03 - Unk
	body.WriteByte(0x03)

	// 0x03 case
	body.WriteUInt16(0x02)

	// EntitySynchInfo
	// Flags
	body.WriteByte(0x0)
	//body.WriteUInt32(0x1)

	// QUEST MANAGER ////////////////////////////////////////////////////////
	addCreateComponent(body, 0x01, 0x0B, "QuestManager")

	// QuestManager::readInit()
	body.WriteUInt32(0x01)
	body.WriteByte(0x01)
	body.WriteCString("Hello")
	body.WriteCString("HelloAgain")
	body.WriteUInt32(0x01)
	body.WriteByte(0x01)
	body.WriteCString("HelloAgainAgain")
	body.WriteCString("HelloAgainAgainAgain")
	body.WriteUInt32(0x01)
	body.WriteCString("Hi")
	body.WriteCString("HiAgain")
	body.WriteCString("HiAgainAgain")

	// QuestManager::ReadAvailableQuests()
	body.WriteByte(0x00) // Probably quest count

	// QuestManager::readInit()
	body.WriteUInt16(0x00) // Objectives count?
	body.WriteUInt16(0x00) // Some count

	// DIALOGUE MANAGER ///////////////////////////////////////
	addCreateComponent(body, 0x01, 0x08, "DialogManager")

	// CREATE AVATAR /////////////////////////////////////////
	body.WriteByte(0x01)     // Create
	body.WriteUInt16(0x0002) // Entity ID
	// Type of lookup?
	// 0x04 by ID?
	// 0xFF by string?
	body.WriteByte(0xFF)
	body.WriteCString("avatar.classes.FighterFemale")

	addCreateComponent(body, 0x02, 0x0A, "avatar.base.Equipment")
	body.WriteByte(0x01) // Item Count

	body.WriteByte(0xFF)
	body.WriteCString("1HAxe1PAL.1HAxe1-1")

	// Item::readData
	// Slot
	// 0x00 None
	// 0x01 Amulet
	// 0x02 Hand
	// 0x03 LRing
	// 0x04 RRing
	// 0x05 Head
	// 0x06 Torso
	// 0x07 Foot
	// 0x08 Shoulder
	// 0x09 None?
	// 0x0a Weapon
	// 0x0b Offhand/Shield
	body.WriteUInt32(0x0a)
	body.WriteByte(0x01)
	body.WriteByte(0x01)
	body.WriteByte(0x01)
	body.WriteByte(0x01)
	// Flag?
	// 0x04 read 2 more bytes
	body.WriteByte(0x04)

	// 0x04 case
	body.WriteUInt16(0x02)

	// GCObject::readChildData<ItemModifier>
	body.WriteByte(0x00) // Count

	// UNITCONTAINER ////////////////////////////////////
	addCreateComponent(body, 0x02, 0x01, "UnitContainer")

	// Container::readInit()
	body.WriteUInt32(0)
	body.WriteUInt32(0)
	body.WriteByte(0x03) // Inventory Count?
	body.WriteByte(0xFF)
	body.WriteCString("avatar.base.Inventory")
	body.WriteByte(0x01)
	body.WriteByte(0x01)

	// GCObject::ReadChildData<Item>()
	body.WriteByte(0x00) // Item count?

	body.WriteByte(0xFF)
	body.WriteCString("avatar.base.TradeInventory")
	body.WriteByte(0x01)
	body.WriteByte(0x01)
	// GCObject::ReadChildData<Item>()
	body.WriteByte(0x00) // Item count?

	body.WriteByte(0xFF)
	body.WriteCString("avatar.base.Bank")
	body.WriteByte(0x01)
	body.WriteByte(0x01)
	// GCObject::ReadChildData<Item>()
	body.WriteByte(0x00) // Item count?

	// UnitContainer::readInit()
	body.WriteByte(0x00) // If >0 it tries to read more, something to do with item

	// MODIFIERS //////////////////////////////////
	// Modifiers are for modifying damage and defences
	addCreateComponent(body, 0x02, 0x0002, "Modifiers")

	// Modifiers::readInit
	body.WriteUInt32(0x00) //
	body.WriteUInt32(0x00) //

	// GCObject::readChildData<Modifier>
	body.WriteByte(0x00)

	// MANIPULATORS //////////////////////////////////
	addCreateComponent(body, 0x02, 0x003, "Manipulators")

	// Manipulators::readInit
	body.WriteByte(0x00) // Some count

	// SKILLS //////////////////////////////////
	addCreateComponent(body, 0x02, 0x004, "avatar.base.skills")

	// Skills::readInit()
	body.WriteUInt32(0xFFFFFFFF)

	// GCObject::readChildData<Skill>
	body.WriteByte(0x04)
	body.WriteByte(0xFF)
	body.WriteCString("skills.generic.Butcher")
	body.WriteUInt32(0x02)
	body.WriteByte(0x03) // Level

	body.WriteByte(0xFF)
	body.WriteCString("skills.generic.Stomp")
	body.WriteUInt32(0x04)
	body.WriteByte(0x05) // Level

	body.WriteByte(0xFF)
	body.WriteCString("skills.generic.FighterClassPassive")
	body.WriteUInt32(0x06)
	body.WriteByte(0x07) // Level

	body.WriteByte(0xFF)
	body.WriteCString("skills.generic.MeleeAttackSpeedModPassive")
	body.WriteUInt32(0x08)
	body.WriteByte(0x09) // Level

	// GCObject::readChildData<SkillProfession>
	body.WriteByte(0x01)
	body.WriteByte(0xFF)
	body.WriteCString("skills.professions.Warrior")

	// UnitBehaviour//////////////////////////////////
	behaviorName := "avatar.base.UnitBehavior"
	if behaviorName == "avatar.base.UnitBehavior" {
		addCreateComponent(body, 0x02, 0x005, "avatar.base.UnitBehavior")

		behav := behavior.NewBehavior()
		behav.Init(body, nil, nil)

		// UnitMover::readInit()
		// Flags
		// 0x04
		// 0x01
		unitMover := byte(0x00)
		body.WriteByte(unitMover)

		if unitMover&0x04 > 0 {
			body.WriteByte(0xFF)
		}

		if unitMover&0x01 > 0 {
			// 0x01 case
			body.WriteUInt32(0x01)
			body.WriteUInt32(0x01)
		}

		body.WriteUInt32(0x00)
		body.WriteUInt32(0x00)

		if unitMover&0x80 > 0 {
			body.WriteUInt32(0x00)
		}

		// Set to 2 for waypoints
		unitMover2 := byte(0) // Could potentially be waypoints?

		body.WriteByte(unitMover2)

		if unitMover2 == 2 {
			waypointCount := uint16(0x0002)
			body.WriteUInt16(waypointCount)

			for i := 0; i < int(waypointCount); i++ {
				// Vector2
				body.WriteUInt32(uint32(1000 * i))   // X?
				body.WriteUInt32(uint32(100000 * i)) // Y?
			}
		}

		// UnitBehavior::readInit()
		body.WriteByte(0xFF)
		body.WriteByte(0xFF)
		body.WriteByte(0xFF)
	} else {
		// This is a monster behavior
		addCreateComponent(body, 0x02, 0x005, "base.MeleeUnit.Behavior")

		behav := behavior.NewBehavior()
		behav.Init(body, nil, nil)

		// UnitMover::readInit()
		// Flags
		// 0x04
		// 0x01
		unitMover := byte(0x00)
		body.WriteByte(unitMover)

		if unitMover&0x04 > 0 {
			body.WriteByte(0xFF)
		}

		if unitMover&0x01 > 0 {
			// 0x01 case
			body.WriteUInt32(0x01)
			body.WriteUInt32(0x01)
		}

		body.WriteUInt32(0x00)
		body.WriteUInt32(0x00)

		if unitMover&0x80 > 0 {
			body.WriteUInt32(0x00)
		}

		// Set to 2 for waypoints
		unitMover2 := byte(0) // Could potentially be waypoints?

		body.WriteByte(unitMover2)

		if unitMover2 == 2 {
			waypointCount := uint16(0x0002)
			body.WriteUInt16(waypointCount)

			for i := 0; i < int(waypointCount); i++ {
				// Vector2
				body.WriteUInt32(uint32(1000 * i))   // X?
				body.WriteUInt32(uint32(100000 * i)) // Y?
			}
		}

		// UnitBehavior::readInit()
		body.WriteByte(0xFF)
		body.WriteByte(0xFF)
		body.WriteByte(0xFF)
	}

	// AVATAR ////////////////////////////////////////

	// Init
	body.WriteByte(0x02)
	body.WriteUInt16(0x0002)

	//WorldEntity::readInit
	// Flags
	// 0x800 Alive? Can cause positions to not work
	// 0x01 Unk
	// 0x02 Unk
	// 0x04 Makes character appear
	// 0x08 Unk
	// 0x10 Unk
	// 0x20 Unk
	// 0x40 Unk
	// 0x80 Unk
	// 0x100 Unk
	// 0x200 Unk
	// 0x400 Unk
	// 0x1000 Makes the character invisible
	// 0x2000 Unk
	// 0x4000 Unk
	// 0x8000 Unk
	// 0x10000 Unk
	// One of these flags stops the below positions from working
	// With only 0x04 the character can be moved and is the least broken
	body.WriteUInt32(
		0x04,
	)
	// These positions stopped working at some point
	body.WriteInt32(0)    // Pos X
	body.WriteInt32(0)    // Pos Y
	body.WriteInt32(0)    // Pos Z
	body.WriteInt32(0x01) // Unk

	// Flags
	// Each flag adds one more section of data to read sequentially
	// 0x01 Has Parent?
	// 0x02 Unk
	// 0x04 Unk
	// 0x08 Unk
	//body.WriteByte(1 | 2 | 4 | 8)
	// When this is set to 0 the character is slightly less broken
	// With 1 | 2 | 4 | 8 it was causing the character to have no animations and
	// eventually collapse into itself
	body.WriteByte(0)

	// 0x01
	//body.WriteUInt16(0x00)

	// Ox02
	//body.WriteByte(0xFF)

	// 0x04
	//body.WriteUInt32(0xFFFFFFFF)

	// 0x08
	//body.WriteUInt32(0xFFFFFFFF)

	// Unit::readInit()
	// Next 4 values always used
	// Same flag as above? + has extras
	// 0x01 - has parent/player owner?
	// 0x02 - add HP
	// 0x04 -
	//body.WriteByte(0x07) // HasParent + Unk
	body.WriteByte(0x01 | 0x02 | 0x04 | 0x10 | 0x20)
	body.WriteByte(50) // Level
	body.WriteUInt16(0x01)
	body.WriteUInt16(0x02)

	// 0x01 case
	body.WriteUInt16(0x01) // Parent ID!!!!!

	conn.Player.CurrentHP = 1150 * 256
	// 0x02 case
	// Multiply HP by 256
	body.WriteUInt32(conn.Player.CurrentHP) // Current HP
	// 0x04 case
	// Multiply MP by 256
	body.WriteUInt32(505 * 256) // MP

	// 0x10 case
	body.WriteByte(0x04) // Unk

	// 0x20 case
	body.WriteUInt16(0x01) // Entity ID, Includes a call to IsKindOf<EncounterObject,Entity>(Entity *)

	// 0x40 case
	//body.WriteUInt16(0x02) // Unk
	//body.WriteUInt16(0x03) // Unk
	//body.WriteUInt16(0x04) // Unk
	//body.WriteByte(0x02)

	// 0x80 case
	//body.WriteByte(0x05)

	// Hero::readInit()
	// The actual EXP value you want to add needs to be multiplied by 20
	body.WriteUInt32(6000 * 20) // Current EXP this level

	// Stats
	// These stats are added to the base stats (seems to be 10)
	body.WriteUInt16(0x02) // Strength
	body.WriteUInt16(0x03) // Agility
	body.WriteUInt16(0x04) // Endurance
	body.WriteUInt16(0x05) // Intellect
	body.WriteUInt16(0x00) // Points remaining
	body.WriteUInt16(0x07) // Respec something or other
	body.WriteUInt32(0x01) // Unk
	body.WriteUInt32(0x01) // Unk

	// Avatar::readInit()
	body.WriteByte(10)  // Face variant
	body.WriteByte(10)  // Hair style
	body.WriteByte(100) // Hair colour

	// AVATAR UPDATE /////////////////////////////////////
	// This update is required to make the character alive
	//body.WriteByte(0x03)     // Update
	//body.WriteUInt16(0x0002) // ID

	// Avatar::processUpdate
	// 0x15 is special Avatar::processUpdate case(spawn entity?) anything else goes to Hero::processUpdate
	// Hero::processUpdate
	// 0x08 is Unit::processUseItemUpdate
	// 0x00 Hero::processUpdateAddExperience
	// 0x01 Hero::processUpdateRemoveExperience
	// 0x02 Hero::processUpdateSpendAttribPoint
	// 0x03 Hero::processUpdateReturnAttribPoint
	// 0x04 Hero::processUpdateRespectAttrbutes
	//body.WriteByte(0x15)
	//
	//// EntitySynchInfo::ReadFromStream
	//body.WriteByte(0x2)
	//body.WriteUInt32(147200) // HP

	body.WriteByte(70) // Now connected
	WriteCompressedA(conn, 0x01, 0x0f, body)
}

func AddSynch(conn *RRConn, body *byter.Byter) {
	// EntitySynchInfo::readFromStream
	body.WriteByte(0x02)
	body.WriteUInt32(conn.Player.CurrentHP)
}

func AddComponentUpdate(body *byter.Byter, comp components.Component) {
	body.WriteByte(byte(ClientEntityChannel))
	//body.WriteByte(0x36) // UpdateComponent - only synch
	body.WriteByte(0x35) // ComponentUpdate - component specific handler + synch
	comp.AddUpdate(body)
}

func AddEntityUpdateStreamEnd(body *byter.Byter) error {
	return body.WriteByte(0x06)
}

func SendWarpTo(conn *RRConn, compID uint16, posX, posY, posZ int32) {
	body := byter.NewLEByter(make([]byte, 0))

	body.WriteByte(byte(ClientEntityChannel))
	body.WriteByte(0x35)
	body.WriteUInt16(compID) // UnitBehavior
	body.WriteByte(0x04)     // CreateAction1
	body.WriteByte(17)
	body.WriteByte(0x00)
	body.WriteInt32(posX)
	body.WriteInt32(posY)
	body.WriteInt32(posZ)

	AddSynch(conn, body)
	AddEntityUpdateStreamEnd(body)

	WriteCompressedA(conn, 0x01, 0x0f, body)
}

func SendMoveTo(conn *RRConn, unk uint8, compID uint16, posX, posY int32) {
	body := byter.NewLEByter(make([]byte, 0))

	body.WriteByte(byte(ClientEntityChannel))
	body.WriteByte(0x35)
	body.WriteUInt16(compID) // UnitBehavior
	body.WriteByte(0x04)     // CreateAction1
	body.WriteByte(0x01)     // MoveTo
	body.WriteByte(unk)
	body.WriteInt32(posX)
	body.WriteInt32(posY)

	body.WriteByte(0x02)
	body.WriteUInt32(0x00)

	//AddSynch(conn, body)
	AddEntityUpdateStreamEnd(body)

	WriteCompressedA(conn, 0x01, 0x0f, body)

	if logging.LoggingOpts.LogMoves {
		fmt.Printf("Send MoveTo %x (%d, %d) (%x, %x)\n", unk, posX, posY, posX, posY)
	}
}
