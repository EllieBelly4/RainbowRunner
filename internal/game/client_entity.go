package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/client_entity"
	"RainbowRunner/internal/game/components"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/logging"
	"RainbowRunner/internal/managers"
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/byter"
	"encoding/hex"
	"fmt"
)

func handleClientEntityChannelMessages(conn *connections.RRConn, msgType byte, reader *byter.Byter) error {
	switch messages.ClientEntityMessage(msgType) {
	case messages.ClientRequestRespawn:
		handleClientEntityUnk4(conn, reader)
	case messages.ClientEntityThings:
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
				fmt.Printf("unhandled inventory message %x\n%s", inventoryMessageType, hex.Dump(reader.Data()))
				return UnhandledChannelMessageError
			}
		case 0x04:
			fmt.Printf("Player tried to put something on hotbar\n%s", hex.Dump(reader.Data()))
		case 0x05:
			return handleClientEntityMovement(conn, reader)
		case 0x0b:
			return handleClientEntityMovement(conn, reader)
		case 0x0a:
			handleSelectEquipment(conn, reader)
		default:
			fmt.Printf("unhandled client entity sub message %x", clientEntitySubMessage)
			return UnhandledChannelMessageError
		}
	default:
		return UnhandledChannelMessageError
	}
	return nil
}

func handleSelectEquipment(conn *connections.RRConn, reader *byter.Byter) {
	body := byter.NewLEByter(make([]byte, 0, 1024))

	body.WriteByte(byte(messages.ClientEntityChannel))
	body.WriteByte(0x35) // ComponentUpdate

	equipID := managers.Players.Players[conn.GetID()].CurrentCharacter.GetChildByGCNativeType("Avatar").GetChildByGCNativeType("Equipment").RREntityProperties().ID

	body.WriteUInt16(equipID) // Equipment ComponentID
	body.WriteByte(0x28)      // Add item

	addEquippedItem(body, "PlateMythicPAL.PlateMythicBoots1", EquipmentSlotFoot, true, "PlateMythicPAL.PlateMythicBoots1.Mod1")

	AddSynch(conn, body)
	AddEntityUpdateStreamEnd(body)

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	fmt.Printf("Player tried to select equipment in inventory\n%s", hex.Dump(reader.Data()))
}

func handleClientEntityMovement(conn *connections.RRConn, reader *byter.Byter) error {
	subMessage := reader.Byte()
	switch subMessage {
	case 0x65:
		handleClientMove(conn, reader)
	// Potentially requesting current position because starting a new path
	case 0x03:
		fmt.Printf("player sent pre-path")
		managers.Players.Players[conn.GetID()].CurrentCharacter.SendPosition()
	default:
		fmt.Printf("unhandled client entity sub message %x", subMessage)
		return UnhandledChannelMessageError
	}

	return nil
}

func handleClientMove(conn *connections.RRConn, reader *byter.Byter) {
	// This increments each time the server sends a MoveTo message
	// The client will then increment by 1 for every individual movement performed (clicking)
	updateNumber := reader.Byte()
	count := int(reader.Byte())
	pos := pkg.Vector2{}

	if logging.LoggingOpts.LogMoves {
		fmt.Printf("Received %d player moves unk val: %x\n", count, updateNumber)
	}

	for i := 0; i < count; i++ {
		unk := reader.Byte()       // Unk
		rotation := reader.Int32() // Seems to be rotation

		degrees := float32((float64(rotation) / 0x17000) * 360)

		pos.X = reader.Int32()
		pos.Y = reader.Int32()

		managers.Players.Players[conn.GetID()].CurrentCharacter.ClientUpdateNumber = updateNumber
		if logging.LoggingOpts.LogMoves {
			fmt.Printf(
				"Player move 0x%x rotation 0x%x(%.2fdeg) (%d, %d) Hex (%x, %x)\n",
				unk, rotation, degrees, pos.X, pos.Y, pos.X, pos.Y,
			)
		}

		conn.Client.LastPosition = managers.Players.Players[conn.GetID()].CurrentCharacter.Position

		managers.Players.Players[conn.GetID()].CurrentCharacter.Position.X = pos.X
		managers.Players.Players[conn.GetID()].CurrentCharacter.Position.Y = pos.Y
		managers.Players.Players[conn.GetID()].CurrentCharacter.Position.Z = 0
		managers.Players.Players[conn.GetID()].CurrentCharacter.Rotation = rotation

		//conn.Player.SendPosition(unk)

		//conn.Player.MoveQueue.Add(MovementUpdate{
		//	Position: pos,
		//	Rotation: rotation,
		//	Tick:     Tick,
		//})

		if unk&0x02 > 0 {
			if logging.LoggingOpts.LogMoves {
				fmt.Println("player started moving")
			}
			managers.Players.Players[conn.GetID()].CurrentCharacter.IsMoving = true
			//conn.Player.SendPosition(0x02)
		}

		if unk&0x01 > 0 {
			if logging.LoggingOpts.LogMoves {
				fmt.Println("player finished moving")
			}
			managers.Players.Players[conn.GetID()].CurrentCharacter.IsMoving = false
			managers.Players.Players[conn.GetID()].CurrentCharacter.SendPosition()
		}
	}

	if managers.Players.Players[conn.GetID()].CurrentCharacter.MoveUpdate >= 0x2D {
		//fmt.Printf(
		//	"sending move update %d, %d || %x, %x!\n",
		//	pos.X, pos.Y,
		//	pos.X, pos.Y,
		//)
		//conn.Player.Move(pos.X, pos.Y)
		//conn.Player.SendFollowClient()
		managers.Players.Players[conn.GetID()].CurrentCharacter.MoveUpdate = 0
	}

	if logging.LoggingOpts.LogMoves {
		fmt.Printf("%s\n", hex.Dump(reader.Data()))
	}
}

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

func sendCreateNewPlayerEntity(conn *connections.RRConn, body *byter.Byter) {
	body = byter.NewLEByter(make([]byte, 0, 2048))

	player := managers.Players.Players[conn.GetID()].CurrentCharacter

	clientEntityWriter := client_entity.NewClientEntityWriter(body)

	clientEntityWriter.Start()

	clientEntityWriter.Create(player)
	clientEntityWriter.Init(player)
	clientEntityWriter.Update(player)

	questManager := player.GetChildByGCType("QuestManager")
	if questManager == nil {
		questManager = objects.NewQuestManager()
		managers.Entities.RegisterAll(conn.Client, questManager)
	}
	clientEntityWriter.CreateComponent(questManager, player)

	dialogManager := player.GetChildByGCType("DialogManager")
	if dialogManager == nil {
		dialogManager = objects.NewDialogManager()
		managers.Entities.RegisterAll(conn.Client, dialogManager)
	}
	clientEntityWriter.CreateComponent(dialogManager, player)

	//addCreateComponent(body, 0x01, 0x0C, "AvatarMetrics")
	//addCreateComponent(body, 0x01, 0x0B, "QuestManager")
	//addCreateComponent(body, 0x01, 32, "DialogManager")

	avatar := player.GetChildByGCNativeType("Avatar")
	clientEntityWriter.Create(avatar)

	//// CREATE AVATAR /////////////////////////////////////////
	//body.WriteByte(0x01)     // Create
	//body.WriteUInt16(0x0002) // Entity ID
	//body.WriteByte(0xFF)
	//body.WriteCString("avatar.classes.FighterFemale")
	//body.WriteCString("avatar.classes.FighterMale")

	addCreateComponent(body, avatar.RREntityProperties().ID, managers.NewID(), "avatar.base.Equipment")

	itemCount := byte(0x01)
	body.WriteByte(itemCount) // Item Count

	//addEquippedItem(body, "1HAxe1PAL.1HAxe1-1", EquipmentSlotWeapon)
	addEquippedItem(
		body,
		"ScaleArmor1PAL.ScaleArmor1-1",
		//"LeatherArmor1PAL.LeatherArmor1-1",
		EquipmentSlotTorso,
		true,
		//"LeatherModPAL.Unique.Mod0",
		"ScaleModPAL.Rare.Mod0",
	)

	//addInitEquipment(body, 0x0A)

	// UNITCONTAINER ////////////////////////////////////
	addCreateComponent(body, avatar.RREntityProperties().ID, managers.NewID(), "UnitContainer")

	// Container::readInit()
	body.WriteUInt32(1)
	body.WriteUInt32(1)
	body.WriteByte(0x03) // Inventory Count?

	body.WriteByte(0xFF)
	body.WriteCString("avatar.base.Inventory")
	body.WriteByte(0x01)
	body.WriteByte(0x01)

	// GCObject::ReadChildData<Item>()
	inventoryItemCount := 0x01
	body.WriteByte(byte(inventoryItemCount)) // Item count?
	AddInventoryItem(body, "PlateMythicPAL.PlateMythicBoots1", 0, 0, "PlateMythicPAL.PlateMythicBoots1.Mod1")

	// Items with PAL seem to be for players
	//for i := 0; i < inventoryItemCount; i++ {
	//AddInventoryItem(body, "1HAxe2PAL.1HAxe2-1", 0, 0)
	//AddInventoryItem(body, "LeatherArmor1PAL.LeatherArmor1-1", 2, 0, true)
	//AddInventoryItem(body, "CrystalHelm1PAL.CrystalHelm1-1", 2, 0)
	//AddInventoryItem(body, "CrystalMythicPAL.CrystalMythicArmor2", 2, 0)
	//}

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

	// UNITCONTAINER UPDATE
	//addUnitContainerUpdate(body, 0x01)

	// MODIFIERS //////////////////////////////////
	// Modifiers are for modifying damage and defences
	addCreateComponent(body, avatar.RREntityProperties().ID, managers.NewID(), "Modifiers")

	// Modifiers::readInit
	body.WriteUInt32(0x00) //
	body.WriteUInt32(0x00) //

	// GCObject::readChildData<Modifier>
	body.WriteByte(0x00)

	// MANIPULATORS //////////////////////////////////
	addCreateComponent(body, avatar.RREntityProperties().ID, managers.NewID(), "Manipulators")

	// Manipulators::readInit
	body.WriteByte(0x00) // Some count

	// SKILLS //////////////////////////////////
	addCreateComponent(body, avatar.RREntityProperties().ID, managers.NewID(), "avatar.base.skills")

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

	unitBehaviour := avatar.GetChildByGCNativeType("UnitBehavior")

	if behaviorName == "avatar.base.UnitBehavior" {
		addCreateComponent(body, avatar.RREntityProperties().ID, unitBehaviour.RREntityProperties().ID, "avatar.base.UnitBehavior")

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
		addCreateComponent(body, avatar.RREntityProperties().ID, managers.NewID(), "base.MeleeUnit.Behavior")

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
	body.WriteUInt16(avatar.RREntityProperties().ID)

	//WorldEntity::readInit
	// Flags
	// 0x01 Static object?
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
	// 0x800 Breaks everything
	// 0x1000 Makes the character invisible
	// 0x2000 Makes movement very jumpy
	// 0x4000 Unk
	// 0x8000 Unk
	// 0x10000 Unk
	// One of these flags stops the below positions from working
	// With only 0x04 the character can be moved and is the least broken
	body.WriteUInt32(
		0x04, // With this one alone it was working
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
	// 0x04 Makes movement smoother, interpolated position?
	// 0x08 Unk
	//body.WriteByte(1 | 2 | 4 | 8)
	// When this is set to 0 the character is slightly less broken
	// With 1 | 2 | 4 | 8 it was causing the character to have no animations and
	// eventually collapse into itself
	//worldEntityInitFlag := 0x04 | 0x08
	worldEntityInitFlag := 0xFF
	body.WriteByte(byte(worldEntityInitFlag))

	if worldEntityInitFlag&0x01 > 0 {
		// 0x01
		body.WriteUInt16(0x00)
	}

	if worldEntityInitFlag&0x02 > 0 {
		// Ox02
		body.WriteByte(0xFF)
	}

	if worldEntityInitFlag&0x04 > 0 {
		// 0x04
		body.WriteUInt32(0xFFFFFFFF)
	}

	if worldEntityInitFlag&0x08 > 0 {
		// 0x08
		body.WriteUInt32(0xFFFFFFFF)
	}

	// Unit::readInit()
	// Next 4 values always used
	// Same flag as above? + has extras
	// 0x01 - has parent/player owner?
	// 0x02 - add HP
	// 0x04 -
	//body.WriteByte(0x07) // HasParent + Unk
	//unitReadinitFlag := 0x01 | 0x02 | 0x04 | 0x10 | 0x20 | 0x40 | 0x80
	unitReadinitFlag := 0x01 | 0x02 | 0x04
	body.WriteByte(byte(unitReadinitFlag))
	body.WriteByte(50) // Level
	body.WriteUInt16(0x01)
	body.WriteUInt16(0x02)

	if unitReadinitFlag&0x01 > 0 {
		// 0x01 case
		body.WriteUInt16(0x01) // Parent ID!!!!!
	}

	if unitReadinitFlag&0x02 > 0 {
		managers.Players.Players[conn.GetID()].CurrentCharacter.CurrentHP = 1150 * 256
		// 0x02 case
		// Multiply HP by 256
		body.WriteUInt32(managers.Players.Players[conn.GetID()].CurrentCharacter.CurrentHP) // Current HP
	}

	if unitReadinitFlag&0x04 > 0 {
		// 0x04 case
		// Multiply MP by 256
		body.WriteUInt32(505 * 256) // MP
	}

	if unitReadinitFlag&0x010 > 0 {
		// 0x10 case
		body.WriteByte(0x04) // Unk
	}

	if unitReadinitFlag&0x020 > 0 {
		// 0x20 case
		body.WriteUInt16(0x01) // Entity ID, Includes a call to IsKindOf<EncounterObject,Entity>(Entity *)
	}

	if unitReadinitFlag&0x040 > 0 {
		// 0x40 case
		body.WriteUInt16(0x02) // Unk
		body.WriteUInt16(0x03) // Unk
		body.WriteUInt16(0x04) // Unk
		body.WriteByte(0x02)
	}

	if unitReadinitFlag&0x080 > 0 {
		//0x80 case
		body.WriteByte(0x05)
	}

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
	body.WriteUInt32(managers.Players.Players[conn.GetID()].CurrentCharacter.CurrentHP)
}

func AddComponentUpdate(body *byter.Byter, comp components.Component) {
	body.WriteByte(byte(messages.ClientEntityChannel))
	//body.WriteByte(0x36) // UpdateComponent - only synch
	body.WriteByte(0x35) // ComponentUpdate - component specific handler + synch
	comp.AddUpdate(body)
}

func AddEntityUpdateStreamEnd(body *byter.Byter) error {
	return body.WriteByte(0x06)
}

func SendMoveTo(conn *connections.RRConn, unk uint8, compID uint16, posX, posY int32) {
	body := byter.NewLEByter(make([]byte, 0))

	body.WriteByte(byte(messages.ClientEntityChannel))
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

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	if logging.LoggingOpts.LogMoves {
		fmt.Printf("Send MoveTo %x (%d, %d) (%x, %x)\n", unk, posX, posY, posX, posY)
	}
}
