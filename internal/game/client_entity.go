package game

import (
	"RainbowRunner/internal/byter"
)

type ClientEntityMessage byte

const (
	ClientEntityUnk0 ClientEntityMessage = iota
	ClientEntityUnk1
	ClientEntityUnk2
	ClientEntityUnk3
	ClientEntityUnk4
	ClientEntityUnk5
	ClientEntityUnk6
	ClientEntityUnk7
	ClientEntityUnk8
	ClientEntityUnk9
	ClientEntityUnk34 = 0x34
)

//07 34 05 00 64 01

func handleClientEntityUnk4(conn *RRConn, reader *byter.Byter) {
	reader.UInt16()
	reader.Byte()

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(ClientEntityChannel))
	body.WriteByte(0x04)
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
	body.WriteUInt32(0x00000000)
	body.WriteUInt32(0x00000000)
	body.WriteByte(0x00)

	body.WriteUInt32(0xFEEDBABA) // World ID
	body.WriteUInt32(1001)       // PvP wins
	body.WriteUInt32(1000)       // PvP rating?, 0 = ???

	// Here goes PvP Team
	// Null string
	body.WriteByte(0x0)

	body.WriteCString("Hello")
	body.WriteUInt32(0x0)

	// UPDATE PLAYER /////////////////////////////////////////
	body.WriteByte(0x03)   // MsgType Update
	body.WriteUInt16(0x01) // Entity ID

	// This maps to a specific event type for Player::processUpdate()
	// 0x01 - do nothing
	// 0x03 - Unk
	body.WriteByte(0x03)

	// 0x03 case
	body.WriteUInt16(0x00)

	// EntitySynchInfo
	// Flags
	body.WriteByte(0x0)
	//body.WriteUInt32(0x1)

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
	body.WriteCString("avatar.classes.FighterMale")

	addCreateComponent(body, 0x02, 0x0A, "avatar.base.Equipment")
	body.WriteByte(0x01)
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
	body.WriteUInt16(0x01)

	// GCObject::readChildData<ItemModifier>
	body.WriteByte(0x00) // Count

	// UNITCONTAINER ////////////////////////////////////
	addCreateComponent(body, 0x02, 0x01, "UnitContainer")
	body.WriteByte(0x01)

	// Container::readInit()
	body.WriteUInt32(0x01)
	body.WriteUInt32(0x01)
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

	// Behavior::readInit()
	body.WriteByte(0x0)
	body.WriteByte(0x0)
	body.WriteByte(0x0)
	body.WriteByte(0x0)
	body.WriteByte(0x1)

	// UnitMover::readInit()
	body.WriteByte(0x0)

	// UnitBehavior::readInit()
	body.WriteUInt32(0x0)
	body.WriteUInt32(0x0)
	body.WriteUInt32(0x0)

	// AVATAR ////////////////////////////////////////

	// Init
	body.WriteByte(0x02)
	body.WriteUInt16(0x0002)

	//WorldEntity::readInit
	// Flags
	// 0x800 Alive?
	body.WriteUInt32(0x800)
	body.WriteInt32(100000) // Pos X
	body.WriteInt32(-50000) // Pos Y
	body.WriteInt32(15000)  // Pos Z
	body.WriteInt32(0x0)    // Unk

	// Flags
	// Each flag adds one more section of data to read sequentially
	// 0x01 Has Parent?
	// 0x02 Unk
	// 0x04 Unk
	// 0x08 Unk
	body.WriteByte(0x01)

	// 0x01
	body.WriteUInt16(0x00)

	// Ox02
	//body.WriteByte(0xFF)

	// 0x04
	//body.WriteUInt32(0xFFFFFFFF)

	// 0x08
	//body.WriteUInt32(0xFFFFFFFF)

	// Unit::readInit()
	// Same flag as above? + has extras
	body.WriteByte(0x07) // HasParent + Unk

	body.WriteByte(50)     // Level
	body.WriteUInt16(0x01) // Unk
	body.WriteUInt16(0x01) // Unk

	body.WriteUInt16(0x01) // Parent ID!!!!!
	body.WriteUInt32(0x01) // Unk
	body.WriteUInt32(0x01) // Unk

	// Hero::readInit()
	// The actual EXP value you want to add needs to be multiplied by 20
	// Probably a homemade float with 2 points of precision
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
	body.WriteByte(0x01)
	body.WriteByte(0x01)
	body.WriteByte(0x01)

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
