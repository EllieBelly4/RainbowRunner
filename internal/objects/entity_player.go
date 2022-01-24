package objects

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/helpers"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	*GCObject
	Name      string
	CurrentHP uint32
}

func (p *Player) Type() DRObjectType {
	return DRObjectOther
}

func (p *Player) WriteInit(b *byter.Byter) {
	// Init PLAYER /////////////////////////////////////////
	b.WriteCString("Ellie")
	b.WriteUInt32(0x01)
	b.WriteUInt32(0x01)
	b.WriteByte(0x01)

	b.WriteUInt32(0xFEEDBABA) // World ID
	b.WriteUInt32(1001)       // PvP wins
	b.WriteUInt32(1000)       // PvP rating?, 0 = ???

	// Here goes PvP Team
	// Null string
	b.WriteByte(0x00)

	// If player is in a PvP team then Avatar respawn will look for the team waypoints
	//b.WriteByte(0xFF)
	//b.WriteCString("pvp.DefaultTeamList.BlueTeam")

	b.WriteCString("Hello")
	b.WriteUInt32(0x01)

}

func (p *Player) WriteUpdate(b *byter.Byter) {
	// This maps to a specific event type for Player::processUpdate()
	// 0x01 - do nothing
	// 0x03 - Unk
	b.WriteByte(0x03)

	// 0x03 case
	b.WriteUInt16(0x02)
}

func (p *Player) WriteFullGCObject(byter *byter.Byter) {
	p.Properties = []GCObjectProperty{
		StringProp("Name", p.Name),
	}

	p.GCObject.WriteFullGCObject(byter)

	byter.WriteCString("Unk")  // Specific to player::readObject
	byter.WriteCString("Unk2") // Specific to player::readObject
	byter.WriteUInt32(0x01)    // Specific to player::readObject
	byter.WriteUInt32(0x01)    // Specific to player::readObject
}

func (p *Player) WriteSynch(b *byter.Byter) {
	b.WriteByte(0x02)
	b.WriteUInt32(p.CurrentHP)
}

func (p *Player) OnZoneJoin(rrPlayer *RRPlayer) {
	SendCreateNewPlayerEntity(rrPlayer, p)
}

func SendCreateNewPlayerEntity(rrplayer *RRPlayer, p *Player) {
	//clientEntityWriter := rrplayer.ClientEntityWriter
	equippedItems := getRandomEquipment()

	body := byter.NewLEByter(make([]byte, 0, 2048))

	conn := p.RREntityProperties().Conn
	player := Players.Players[conn.GetID()].CurrentCharacter
	clientEntityWriter := NewClientEntityWriter(body)
	clientEntityWriter.BeginStream()

	avatar := player.GetChildByGCNativeType("Avatar")
	clientEntityWriter.Create(avatar)

	clientEntityWriter.Create(player)
	clientEntityWriter.Init(player)
	clientEntityWriter.Update(player)

	// MANIPULATORS //////////////////////////////////
	addCreateComponent(body, avatar.RREntityProperties().ID, NewID(), "creatures.humanoid.base.MeleeBase.Manipulators")

	// Manipulators::readInit
	//manipCount := byte(0x01)
	body.WriteByte(byte(len(equippedItems))) // Some count

	for _, equippedItem := range equippedItems {
		equippedItem.WriteInit(body)
		equippedItem.WriteManipulatorInit(body)
	}

	equipment := avatar.GetChildByGCType("avatar.base.Equipment")
	addCreateComponent(body, avatar.RREntityProperties().ID, equipment.RREntityProperties().ID, "avatar.base.Equipment")

	body.WriteByte(byte(len(equippedItems)))

	for _, equippedItem := range equippedItems {
		equippedItem.WriteInit(body)
	}

	// Invalid component type from server
	//fighterEquipment := NewGCObject("avatar.classes.FighterStartingEquipment")
	//Entities.RegisterAll(conn, fighterEquipment)
	//clientEntityWriter.CreateComponent(fighterEquipment, avatar)

	questManager := player.GetChildByGCType("QuestManager")
	if questManager == nil {
		questManager = NewQuestManager()
		Entities.RegisterAll(conn, questManager)
	}
	clientEntityWriter.CreateComponent(questManager, player)

	dialogManager := player.GetChildByGCType("DialogManager")
	if dialogManager == nil {
		dialogManager = NewDialogManager()
		Entities.RegisterAll(conn, dialogManager)
	}

	clientEntityWriter.CreateComponent(dialogManager, player)

	//addCreateComponent(body, 0x01, 0x0C, "AvatarMetrics")
	//addCreateComponent(body, 0x01, 0x0B, "QuestManager")
	//addCreateComponent(body, 0x01, 32, "DialogManager")

	//// CREATE AVATAR /////////////////////////////////////////
	//body.WriteByte(0x01)     // Create
	//body.WriteUInt16(0x0002) // Entity ID
	//body.WriteByte(0xFF)
	//body.WriteCString("avatar.classes.FighterFemale")
	//body.WriteCString("avatar.classes.FighterMale")

	// UNITCONTAINER ////////////////////////////////////
	addCreateComponent(body, avatar.RREntityProperties().ID, NewID(), "UnitContainer")

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
	addCreateComponent(body, avatar.RREntityProperties().ID, NewID(), "Modifiers")

	// Modifiers::readInit
	body.WriteUInt32(0x00) //
	body.WriteUInt32(0x00) //

	// GCObject::readChildData<Modifier>
	body.WriteByte(0x00)

	//for i := 0; i < int(manipCount); i++ {
	//	equipBoots := NewEquipment(
	//		"PlateMythicPAL.PlateMythicBoots1",
	//		"PlateMythicPAL.PlateMythicBoots1.Mod1",
	//		EquipmentItemArmour, EquipmentSlotFoot, 5,
	//	)
	//
	//	equipBoots.WriteInit(body)
	//}

	// SKILLS //////////////////////////////////
	addCreateComponent(body, avatar.RREntityProperties().ID, NewID(), "avatar.base.skills")

	// Skills::readInit()
	body.WriteUInt32(0xFFFFFFFF)

	// GCObject::readChildData<Skill>
	body.WriteByte(0x04) // Count

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
		addCreateComponent(body, avatar.RREntityProperties().ID, NewID(), "base.MeleeUnit.Behavior")

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
		Players.Players[conn.GetID()].CurrentCharacter.CurrentHP = 1150 * 256
		// 0x02 case
		// Multiply HP by 256
		body.WriteUInt32(Players.Players[conn.GetID()].CurrentCharacter.CurrentHP) // Current HP
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
	helpers.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func getRandomEquipment() []*Equipment {
	equippedItems := []*Equipment{
		//NewEquipment(
		//	"PlateMythicPAL.PlateMythicBoots1",
		//	"PlateMythicPAL.PlateMythicBoots1.Mod1",
		//	EquipmentItemArmour, EquipmentSlotFoot, 5,
		//),
		//NewEquipment(
		//	"PlateMythicPAL.PlateMythicArmor1",
		//	"PlateMythicPAL.PlateMythicArmor1.Mod1",
		//	EquipmentItemArmour, EquipmentSlotTorso, 5,
		//),
		//NewEquipment(
		//	"PlateMythicPAL.PlateMythicGloves1",
		//	"PlateMythicPAL.PlateMythicGloves1.Mod1",
		//	EquipmentItemArmour, EquipmentSlotHand, 5,
		//),
		//NewEquipment(
		//	"PlateMythicPAL.PlateMythicHelm1",
		//	"PlateMythicPAL.PlateMythicHelm1.Mod1",
		//	EquipmentItemArmour, EquipmentSlotHead, 5,
		//),
		//NewEquipment(
		//	"PlateMythicPAL.PlateMythicShield1",
		//	"PlateMythicPAL.PlateMythicShield1.Mod1",
		//	EquipmentItemArmour, EquipmentSlotOffhand, 5,
		//),
		//NewEquipment(
		//	"1HSwordMythicPAL.1HSwordMythic6",
		//	"1HSwordMythicPAL.1HSwordMythic6.Mod1",
		//	EquipmentItemMeleeWeapon, EquipmentSlotWeapon, 6,
		//),

		//NewEquipment(
		//	"1HSwordMythicPAL.1HSwordMythic6",
		//	"1HSwordMythicPAL.1HSwordMythic6.Mod1",
		//	EquipmentItemMeleeWeapon, EquipmentSlotWeapon,
		//),

		//NewEquipment(
		//	"1HSwordMythicPAL.1HSwordMythic1",
		//	"1HSwordMythicPAL.1HSwordMythic1.Mod1",
		//	EquipmentItemMeleeWeapon, EquipmentSlotWeapon, 5,
		//),
		//NewEquipment(
		//	"ScaleArmor1PAL.ScaleArmor1-1",
		//	"ScaleModPAL.Binder.Mod1",
		//	EquipmentItemArmour, EquipmentSlotTorso, 1,
		//),

		//NewEquipment(
		//	"PlateArmor3PAL.PlateArmor3-7",
		//	"ScaleModPAL.Rare.Mod1",
		//	EquipmentItemArmour, EquipmentSlotTorso,
		//),
		//
		//NewEquipment(
		//	"PlateBoots3PAL.PlateBoots3-7",
		//	"ScaleModPAL.Rare.Mod1",
		//	EquipmentItemArmour, EquipmentSlotFoot,
		//),
		//
		//NewEquipment(
		//	"PlateHelm3PAL.PlateHelm3-7",
		//	"ScaleModPAL.Rare.Mod1",
		//	EquipmentItemArmour, EquipmentSlotHead,
		//),
		//
		//NewEquipment(
		//	"PlateGloves3PAL.PlateGloves3-7",
		//	"ScaleModPAL.Rare.Mod1",
		//	EquipmentItemArmour, EquipmentSlotHand,
		//),
		//
		//NewEquipment(
		//	"CrystalMythicPAL.CrystalMythicShield1",
		//	"ScaleModPAL.Rare.Mod1",
		//	EquipmentItemArmour, EquipmentSlotOffhand,
		//),

		//NewEquipment(
		//	"ScaleArmor2PAL.ScaleArmor2-4",
		//	"ScaleModPAL.Binder.Mod1",
		//	EquipmentItemArmour, EquipmentSlotTorso, 1,
		//),
	}

	//return equippedItems

	//PlateHelm1PAL.PlateHelm1-6

	// One of these failed
	//PlateHelm1PAL.PlateHelm1-7
	//ChainArmor3PAL.ChainArmor3-4
	//ScaleGloves2PAL.ScaleGloves2-6
	//LeatherBoots1PAL.LeatherBoots1-10
	//2HPickMythicPAL.2HPickMythic1

	equippedItems = append(equippedItems, addRandomEquipment(database.Helmets, EquipmentItemArmour))
	equippedItems = append(equippedItems, addRandomEquipment(database.Armours, EquipmentItemArmour))
	equippedItems = append(equippedItems, addRandomEquipment(database.Gloves, EquipmentItemArmour))
	equippedItems = append(equippedItems, addRandomEquipment(database.Boots, EquipmentItemArmour))

	equippedItems = append(equippedItems, addRandomEquipment(database.MeleeWeapons, EquipmentItemMeleeWeapon))
	//equippedItems = append(equippedItems, addRandomEquipment(database.RangedWeapons, EquipmentItemRangedWeapon))

	//randomArmour := database.Armours[int(r.Int63())%len(database.Armours)]
	//equippedItems = append(equippedItems, NewEquipment(
	//	randomArmour.Name, "ScaleModPAL.Rare.Mod1",
	//	EquipmentItemArmour, types.EquipmentSlotTorso,
	//))
	//
	//// failed
	//// ChainBoots3PAL.ChainBoots3-1 // Rare Only
	//randomBoots := ArmourMap["boots"][int(r.Int63())%len(ArmourMap["boots"])]
	//equippedItems = append(equippedItems, NewEquipment(
	//	randomBoots, "ScaleModPAL.Rare.Mod1",
	//	EquipmentItemArmour, types.EquipmentSlotFoot,
	//))
	//
	//randomHelm := ArmourMap["helm"][int(r.Int63())%len(ArmourMap["helm"])]
	//equippedItems = append(equippedItems, NewEquipment(
	//	randomHelm, "ScaleModPAL.Rare.Mod1",
	//	EquipmentItemArmour, types.EquipmentSlotHead,
	//))
	//
	//randomGloves := ArmourMap["gloves"][int(r.Int63())%len(ArmourMap["gloves"])]
	//equippedItems = append(equippedItems, NewEquipment(
	//	randomGloves, "ScaleModPAL.Rare.Mod1",
	//	EquipmentItemArmour, types.EquipmentSlotHand,
	//))

	output := ""

	for _, item := range equippedItems {
		output += fmt.Sprintf("%s\n", item.GCType)
	}

	fmt.Printf("Random equipment for today is:\n%s\n", output)

	return equippedItems
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func addRandomEquipment(equipment database.EquipmentMap, t ItemType) *Equipment {
	i := 0

	target := int(r.Int63()) % len(equipment)

	for key, class := range equipment {
		if i == target {
			return NewEquipment(
				key, "ScaleModPAL.Rare.Mod1",
				t, class.Slot(),
			)
		}
		i++
	}

	return nil
}

func addCreateComponent(body *byter.Byter, parentID uint16, componentID uint16, typeString string) {
	body.WriteByte(0x32)          // Create Component
	body.WriteUInt16(parentID)    // Parent Entity ID
	body.WriteUInt16(componentID) // Component ID
	body.WriteByte(0xFF)          // Unk
	body.WriteCString(typeString) // Component Type
	body.WriteByte(0x01)          // Unk
}

func NewPlayer(name string) (p *Player) {
	p = &Player{
		Name: name,
	}

	p.GCObject = NewGCObject("Player")
	p.GCName = "ElliePlayer"
	p.GCType = "player"

	return
}

func AddEquippedItem(
	body *byter.Byter,
	item string,
	slot types.EquipmentSlot,
	armour bool,
	mod string,
) {
	body.WriteByte(0xFF) // GetType
	body.WriteCString(item)

	// Item::readData
	body.WriteUInt32(uint32(slot))
	body.WriteByte(0xF0)
	body.WriteByte(0xF0)
	body.WriteByte(0x01)   // Item count
	body.WriteByte(50 + 5) // Required level + 5

	// Flag?
	// 0x01 - Soulbound in 9 minutes, no idea where the time comes from
	// 0x02 - Not Sellable
	// 0x04 - +0x01 = Soulbound timer
	// 0x08 - Requires Membership
	itemFlag := 0x01 | 0x04 | 0x08

	body.WriteByte(byte(itemFlag))

	if itemFlag&0x04 > 0 {
		// Soulbind time
		// Minutes * 0x800 max 9
		body.WriteUInt16(0x800 * 7)
	}

	if item == "LeatherArmor1PAL.LeatherArmor1-1" || item == "ScaleArmor1PAL.ScaleArmor1-1" {
		// Required modifiers?
		// ItemModifier?
		itemModifierFlag1 := 0x01 | 0x02

		body.WriteByte(byte(itemModifierFlag1))

		if itemModifierFlag1&0x01 > 0 {
			body.WriteByte(0xFF)
		}

		if itemModifierFlag1&0x02 > 0 {
			body.WriteUInt32(0xFFFFFFFF)
		}

		//if mod != "" {
		// GCObject::readChildData<ItemModifier>
		body.WriteByte(0x01) // Count

		body.WriteByte(0xFF)
		body.WriteCString(mod)

		// ItemModifier?
		itemModifierFlag := 0x01 | 0x02

		body.WriteByte(byte(itemModifierFlag))

		if itemModifierFlag&0x01 > 0 {
			body.WriteByte(0x15)
		}

		if itemModifierFlag&0x02 > 0 {
			body.WriteUInt32(0x11111111)
		}
		//} else {
		//	body.WriteByte(0x00) // Count
		//}
	} else if item == "PlateMythicPAL.PlateMythicArmor1" || item == "PlateMythicPAL.PlateMythicBoots1" {
		// Required modifiers?
		// ItemModifier?
		itemModifierFlag1 := 0x00

		// Each item has different numbers of required modifiers
		for i := 0; i < 5; i++ {
			body.WriteByte(byte(itemModifierFlag1))

			if itemModifierFlag1&0x01 > 0 {
				body.WriteByte(0xFF)
			}

			if itemModifierFlag1&0x02 > 0 {
				body.WriteUInt32(0xFFFFFFFF)
			}
		}

		//if mod != "" {
		// GCObject::readChildData<ItemModifier>
		body.WriteByte(0x01) // Count

		body.WriteByte(0xFF)
		body.WriteCString(mod)

		// ItemModifier?
		itemModifierFlag := 0x01 | 0x02

		body.WriteByte(byte(itemModifierFlag))

		if itemModifierFlag&0x01 > 0 {
			body.WriteByte(0x15)
		}

		if itemModifierFlag&0x02 > 0 {
			body.WriteUInt32(0x11111111)
		}
	} else {
		panic("unhandled equipment")
	}
}
