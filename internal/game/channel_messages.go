package game

import (
	"RainbowRunner/internal/byter"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

func handleChannelMessage(conn net.Conn, reader *byter.Byter, clientID uint32) {
	msgChan := reader.UInt8()    // Channel
	msgSubType := reader.UInt8() // Message Type

	switch Channel(msgChan) {
	case Unk2:
		body := byter.NewLEByter(make([]byte, 0, 1024))
		body.WriteByte(byte(Unk2)) // Character channel
		body.WriteByte(0x00)
		WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
	case CharacterChannel:
		switch CharacterMessage(msgSubType) {
		case CharacterConnected:
			body := byter.NewLEByter(make([]byte, 0, 1024))
			body.WriteByte(byte(CharacterChannel))   // Character channel
			body.WriteByte(byte(CharacterConnected)) // Connected
			WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
		case CharacterPlay:
			body := byter.NewLEByter(make([]byte, 0, 1024))
			body.WriteByte(byte(CharacterChannel))
			body.WriteByte(byte(CharacterPlay))
			WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
		case CharacterGetList:
			sendCharacterList(conn, clientID)
		case CharacterCreate:
			name := reader.String()
			class := reader.String()
			reader.UInt8() // Unk
			reader.UInt8() // Face
			reader.UInt8() // Hair
			reader.UInt8() // Hair Colour

			log.Infof("New character created %s (%s)", name, class)

			body := byter.NewLEByter(make([]byte, 0, 1024))
			body.WriteByte(byte(CharacterChannel)) // Character channel
			body.WriteByte(byte(CharacterCreate))
			body.WriteUInt32(0x01)

			body.WriteCString("Ellie")

			sendPlayer(body)

			WriteCompressedA(clientID, 0x01, 0x0f, body, conn)

			//sendCharacterList(conn, clientID)
		default:
			log.Panicf("Unhandled chan %x msgSubType %x", msgChan, msgSubType)
		}
	case GroupChannel:
		switch GroupChannelMessage(msgSubType) {
		case GroupConnected:
			body := byter.NewLEByter(make([]byte, 0, 1024))
			body.WriteByte(byte(GroupChannel))
			body.WriteByte(48)

			body = byter.NewLEByter(make([]byte, 0, 1024))
			body.WriteByte(byte(ZoneChannel))
			body.WriteByte(0x00)
			//body.WriteCString("TheHub")
			body.WriteCString("Tutorial")
			WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
		}
	case ZoneChannel:
		switch ZoneChannelMessage(msgSubType) {
		case ZoneUnk6:
			// This cannot continue because the game cannot find any players I think
			// Search for how to add players
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

			body = byter.NewLEByter(make([]byte, 0, 1024))
			body.WriteByte(byte(ClientEntityChannel))
			//body.WriteByte(0x01) // Create
			body.WriteByte(0x08) // CreateInit
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
			body.WriteCString("Ellie")
			body.WriteUInt32(0xF00DF00D)
			body.WriteUInt32(0xBABAFAAB)
			body.WriteByte(0x01)

			body.WriteUInt32(0xFEEDBABA) // One of these is the world ID?
			body.WriteUInt32(0xFEEDBABA) // One of these is the world ID?
			body.WriteUInt32(0xFEEDBABA) // One of these is the world ID?

			// Here goes Team
			// Null string
			body.WriteByte(0x0)

			body.WriteCString("Hello")
			body.WriteUInt32(0x01)

			// UPDATE PLAYER /////////////////////////////////////////
			body.WriteByte(0x03)   // MsgType Update
			body.WriteUInt16(0x01) // Entity ID
			body.WriteByte(0x01)

			// EntitySynchInfo
			body.WriteByte(0x0)
			//body.WriteUInt32(0x0) // If above is even

			//body.WriteByte(0x32) // Create Component
			//body.WriteUInt16(0x01)   // Entity ID
			//body.WriteUInt16(0xDAAD) // Unk
			//body.WriteByte(0xFF)     // Unk
			//body.WriteCString() // Component Type

			//body.WriteByte(0x20) // Create Subentity
			//body.WriteByte(0xFF)
			//body.WriteCString("Avatar)

			// QUEST MANAGER ////////////////////////////////////////////////////////
			addCreateComponent(body, 0x01, 0x0B, "QuestManager")
			body.WriteByte(0x01)

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
			body.WriteByte(0x01)

			// CREATE AVATAR /////////////////////////////////////////
			body.WriteByte(0x01)     // Create
			body.WriteUInt16(0x0002) // Entity ID
			// Type of lookup?
			// 0x04 by ID?
			// 0xFF by string?
			body.WriteByte(0xFF)
			body.WriteCString("avatar.classes.fighterfemale")

			addCreateComponent(body, 0x02, 0x0A, "avatar.base.Equipment")
			body.WriteByte(0x01)
			body.WriteByte(0x00) // Item Count

			// UNITCONTAINER ////////////////////////////////////
			addCreateComponent(body, 0x02, 0x01, "UnitContainer")
			body.WriteByte(0x01)

			// Container::readInit()
			body.WriteUInt32(0x01)
			body.WriteUInt32(0x01)
			body.WriteByte(0x02)
			body.WriteByte(0xFF)
			body.WriteCString("avatar.base.Inventory")
			body.WriteByte(0x01)
			body.WriteByte(0x01)

			// GCObject::ReadChildData<Item>()
			body.WriteByte(0x00) // Item count?

			body.WriteByte(0xFF)
			body.WriteCString("avatar.base.TradeInventory") // Copied from above, may not work
			body.WriteByte(0x01)                            // Copied from above, may not work
			body.WriteByte(0x01)                            // Copied from above, may not work
			// GCObject::ReadChildData<Item>() // Copied from above, may not work
			body.WriteByte(0x00) // Item count? // Copied from above, may not work

			// UnitContainer::readInit()
			body.WriteByte(0x00) // If >0 it tries to read more, something to do with item

			// MODIFIERS //////////////////////////////////
			addCreateComponent(body, 0x02, 0x0002, "Modifiers")
			body.WriteByte(0x01) // Unk

			// Modifiers::readInit
			body.WriteUInt32(0x01) //
			body.WriteUInt32(0x01) //

			// GCObject::readChildData<Modifier>
			body.WriteByte(0x00)

			// MANIPULATORS //////////////////////////////////
			addCreateComponent(body, 0x02, 0x003, "Manipulators")
			body.WriteByte(0x0A) // Unk

			// Manipulators::readInit
			body.WriteByte(0x00) // Some count
			// body.WriteCString() // Some ManipulatorClass

			// SKILLS //////////////////////////////////
			addCreateComponent(body, 0x02, 0x004, "avatar.base.skills")
			body.WriteByte(0x0A) // Unk

			// Skills:readInit()
			body.WriteUInt32(0x00)

			// GCObject::readChildData<Skill>
			body.WriteByte(0x00)

			// GCObject::readChildData<SkillProfession>
			body.WriteByte(0x00)

			// UnitBehaviour//////////////////////////////////
			addCreateComponent(body, 0x02, 0x005, "avatar.base.UnitBehavior")

			// Behaviour::readInit()
			body.WriteByte(0x0)
			body.WriteByte(0x0)
			body.WriteByte(0x0)
			body.WriteByte(0x0)
			body.WriteByte(0x1)

			// UnitMover::readInit()
			body.WriteByte(0x0)

			// UnitBehaviour::readInit()
			body.WriteUInt32(0x0)
			body.WriteUInt32(0x0)
			body.WriteUInt32(0x0)

			// AVATAR ////////////////////////////////////////

			// Init
			body.WriteByte(0x02)
			body.WriteUInt16(0x0002)

			//WorldEntity::readInit
			body.WriteUInt32(0xFEEDBABA)
			body.WriteUInt32(0xFEEDBABA)
			body.WriteUInt32(0xFEEDBABA)
			body.WriteUInt32(0xFEEDBABA)
			body.WriteUInt32(0xFEEDBABA)

			// Flags
			// Each flag adds one more section of data to read sequentially
			// 0x01 Has Parent? (hopefully)
			// 0x02 Unk
			// 0x04 Unk
			// 0x08 Unk
			body.WriteByte(0x01)

			// 0x01
			body.WriteUInt16(0x0001)

			// Ox02
			//body.WriteByte(0x01)

			// 0x04
			//body.WriteUInt32(0x01)

			// 0x08
			//body.WriteUInt32(0x01)

			// Unit::readInit()
			// Same flag as above? + has extras
			body.WriteByte(0x07) // HasParent + Unk

			body.WriteByte(0x01)
			body.WriteUInt16(0x01)
			body.WriteUInt16(0x01)

			// Parent ID!!!!!
			body.WriteUInt16(0x01)
			body.WriteUInt32(0x01)
			body.WriteUInt32(0x01)

			// Hero::readInit()
			body.WriteUInt32(0x01)
			body.WriteUInt16(0x01)
			body.WriteUInt16(0x01)
			body.WriteUInt16(0x01)
			body.WriteUInt16(0x01)
			body.WriteUInt16(0x01)
			body.WriteUInt16(0x01)
			body.WriteUInt32(0x01)
			body.WriteUInt32(0x01)

			// Avatar::readInit()
			body.WriteByte(0x01)
			body.WriteByte(0x01)
			body.WriteByte(0x01)

			body.WriteByte(70) // Now connected
			WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
		}
	case UserChannel:
		switch msgSubType {
		case 0x00: // Request rosters
			body := byter.NewLEByter(make([]byte, 0, 1024))
			body.WriteByte(byte(UserChannel))
			body.WriteByte(0x00)
			WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
		case 0x01: // Rosters response
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
		default:
			log.Panicf("Unhandled chan %x msgSubType %x", msgChan, msgSubType)
		}
	default:
		panic(fmt.Sprintf("Unhandled channel message %x", msgChan))
	}
}

func addCreateComponent(body *byter.Byter, parentID uint16, componentID uint16, typeString string) {
	body.WriteByte(0x32)          // Create Component
	body.WriteUInt16(parentID)    // Parent Entity ID
	body.WriteUInt16(componentID) // Component ID
	body.WriteByte(0xFF)          // Unk
	body.WriteCString(typeString) // Component Type
}
