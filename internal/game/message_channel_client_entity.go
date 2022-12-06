package game

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/global"
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg/byter"
	"fmt"
)

func handleClientEntityChannelMessages(conn *connections.RRConn, msgType byte, reader *byter.Byter) error {
	switch messages.ClientEntityMessage(msgType) {
	case messages.ClientRequestRespawn:
		handleClientEntityUnk4(conn, reader)
	case messages.ClientEntityComponentUpdate:
		componentID := reader.UInt16()

		entity := objects.Entities.FindByID(componentID)

		if entity != nil {
			err := entity.ReadUpdate(reader)

			if err != nil {
				fmt.Printf("failed to ReadUpdate for component:\n%s", err.Error())
				return UnhandledChannelMessageError
			}
			//if entity.GetGCObject().EntityHandler != nil {
			//	err := entity.GetGCObject().EntityHandler.ReadUpdate(reader)
			//
			//	if err != nil {
			//		panic(err)
			//	}
			//}else{
			//	fmt.Printf("Component %x does not have an entityHandler!\n", entity.RREntityProperties().ID)
			//}
		} else {
			fmt.Printf("Component %x does not exist!\n", componentID)
			return UnhandledChannelMessageError
		}

		//switch componentID {
		//// Inventory Message?
		//case 0x01:
		//	inventoryMessageType := reader.UInt8()
		//	switch inventoryMessageType {
		//	case 0x21:
		//		fmt.Printf("Player opened inventory\n%s", hex.Dump(reader.Data()))
		//	case 0x22:
		//		fmt.Printf("Player closed inventory\n%s", hex.Dump(reader.Data()))
		//	default:
		//		fmt.Printf("unhandled inventory message %x\n%s", inventoryMessageType, hex.Dump(reader.Data()))
		//		return UnhandledChannelMessageError
		//	}
		//case 0x04:
		//	fmt.Printf("Player tried to put something on hotbar\n%s", hex.Dump(reader.Data()))
		//case 0x05:
		//	return handleClientEntityMovement(conn, reader)
		//case 0x09:
		//	return handleClientEntityMovement(conn, reader)
		//case 0x12:
		//	return handleClientEntityMovement(conn, reader)
		//case 0x0b:
		//	return handleClientEntityMovement(conn, reader)
		//case 0x0a:
		//	handleSelectEquipment(conn, reader)
		//default:
		//	fmt.Printf("unhandled client entity sub message %x", componentID)
		//	return UnhandledChannelMessageError
		//}
	default:
		return UnhandledChannelMessageError
	}
	return nil
}

//func handleSelectEquipment(conn *connections.RRConn, reader *byter.Byter) {
//	body := byter.NewLEByter(make([]byte, 0, 1024))
//
//	body.WriteByte(byte(messages.ClientEntityChannel))
//	body.WriteByte(0x35) // ComponentUpdate
//
//	equipID := objects.Players.Players[conn.GetID()].CurrentCharacter.GetChildByGCNativeType("Avatar").GetChildByGCNativeType("Equipment").RREntityProperties().ID
//
//	body.WriteUInt16(equipID) // Equipment ComponentID
//	body.WriteByte(0x28)      // Add item
//
//	objects.AddEquippedItem(body, "PlateMythicPAL.PlateMythicBoots1", types.EquipmentSlotFoot, true, "PlateMythicPAL.PlateMythicBoots1.Mod1")
//
//	AddSynch(conn, body)
//	AddEntityUpdateStreamEnd(body)
//
//	helpers.WriteCompressedA(conn, 0x01, 0x0f, body)
//
//	fmt.Printf("Player tried to select equipment in inventory\n%s", hex.Dump(reader.Data()))
//}

func handleClientEntityUnk4(conn *connections.RRConn, reader *byter.Byter) {
	id := reader.UInt16()
	event := reader.Byte() // Guessing here

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ClientEntityChannel))
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

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func addUnitContainerUpdate(body *byter.Byter, ID uint16) {
	body.WriteByte(0x35)
	body.WriteUInt16(ID)
	body.WriteByte(0x1E)

	//Container::processAddItem
	body.WriteByte(0x01)

	body.WriteByte(0xFF)
	body.WriteCString("PlateMythicPAL.PlateMythicBoots1")
}

func addInitEquipment(body *byter.Byter, componentID uint16) {
	body.WriteByte(0x33)          // InitComponent
	body.WriteUInt16(componentID) // Parent Entity ID
}

func AddSynch(conn *connections.RRConn, body *byter.Byter) {
	// EntitySynchInfo::readFromStream
	body.WriteByte(0x02)
	body.WriteUInt32(objects.Players.Players[conn.GetID()].CurrentCharacter.CurrentHP)
}

func AddEntityUpdateStreamEnd(body *byter.Byter) error {
	return body.WriteByte(0x06)
}

func SendMoveTo(conn *connections.RRConn, unk uint8, compID uint16, posX, posY int32) {
	body := byter.NewLEByter(make([]byte, 0))

	body.WriteByte(byte(messages.ClientEntityChannel))
	body.WriteByte(0x35)     // MoveTo
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

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	if config.Config.Logging.LogMoves {
		fmt.Printf("Send MoveTo %x (%d, %d) (%x, %x)\n", unk, posX, posY, posX, posY)
	}
}

// Interval does not seem to be specific to an entity because it sets some global values
func SendInterval(conn *connections.RRConn) {
	body := byter.NewLEByter(make([]byte, 0))

	body.WriteByte(byte(messages.ClientEntityChannel))
	body.WriteByte(0x0D) // IntervalMessage

	// ClientEntityManager::processInterval
	// Current Server Tick
	body.WriteInt32(int32(global.Tick)) // Unk - Stored in ClientEntityManager::vftable + 0xa94

	body.WriteInt32(global.TickInterval) // TickInterval - Stored in ClientEntityManager::vftable + 0xa80

	// Seems to be a message queue limit? If this is too low it seems to break smooth movement
	// 10 = 8 moves per message
	// 11 = 8 moves per message
	// 12 = 9 moves per message
	// 13 = 9 moves per message
	// 15 = 10 moves per message
	// 2 = 3 moves per message
	// Moves per message = max(thisNumber / 3, 3)
	body.WriteInt32(0) // Unk - Stored in ClientEntityManager::vftable + 0xa84 - Movement prediction buffer/ticks ahead of server?/max ticks behind server

	// PathManager::readBudget
	body.WriteInt32(0)    // Unk
	body.WriteUInt16(100) // Budget Per Update
	body.WriteUInt16(20)  // Budget Per Path

	AddEntityUpdateStreamEnd(body)

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}
