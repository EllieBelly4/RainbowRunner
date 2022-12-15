package objects

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

//go:generate go run ../../scripts/generatelua -type=Player -extends=GCObject
type Player struct {
	*GCObject
	Name      string
	CurrentHP uint32 // This is probably a DRFloat
	Spawned   bool
	Zone      *Zone
}

func (p *Player) Type() DRObjectType {
	return DRObjectOther
}

func (p *Player) WriteInit(b *byter.Byter) {
	rrPlayer := Players.Players[int(p.OwnerID())]

	// Init PLAYER /////////////////////////////////////////
	b.WriteCString("Ellie")
	b.WriteUInt32(0x01)
	b.WriteUInt32(0x01)
	b.WriteByte(0x01)

	b.WriteUInt32(rrPlayer.Zone().ID) // World ID
	b.WriteUInt32(1001)               // PvP wins
	b.WriteUInt32(1000)               // PvP rating?, 0 = ???

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

func (p *Player) SendCreateNewPlayerEntity(rrplayer *RRPlayer) {
	zone := rrplayer.Zone()
	//clientEntityWriter := rrplayer.ClientEntityWriter
	//equippedItems := getRandomEquipment()
	avatar := p.GetChildByGCNativeType("Avatar")
	inventoryEquipment := avatar.GetChildByGCNativeType("Equipment")

	equippedItems := inventoryEquipment.(*EquipmentInventory).GetEquipment()
	fmt.Printf("%+v\n", inventoryEquipment)
	body := byter.NewLEByter(make([]byte, 0, 2048))

	conn := p.RREntityProperties().Conn
	clientEntityWriter := NewClientEntityWriter(body)
	clientEntityWriter.BeginStream()

	clientEntityWriter.Create(avatar)

	clientEntityWriter.Create(p)
	clientEntityWriter.Init(p)
	clientEntityWriter.Update(p)

	//creatures.humanoid.base.MeleeBase.Manipulators
	//avatar.base.Equipment
	//2HMace5PAL.2HMace5-7
	//ScaleModPAL.Rare.Mod1
	// MANIPULATORS //////////////////////////////////
	//addCreateComponent(body, avatar.RREntityProperties().ID, NewID(), "Manipulators")
	manipulators := avatar.GetChildByGCNativeType("Manipulators")
	clientEntityWriter.CreateComponentAndInit(manipulators, avatar)

	equipment := avatar.GetChildByGCType("avatar.base.Equipment")
	addCreateComponent(body, uint16(avatar.RREntityProperties().ID), uint16(equipment.RREntityProperties().ID), "avatar.base.Equipment")

	body.WriteByte(byte(len(equippedItems)))

	for _, equippedItem := range equippedItems {
		equippedItem.WriteInit(body)
	}

	// Invalid component type from server
	//fighterEquipment := NewGCObject("avatar.classes.FighterStartingEquipment")
	//Entities.RegisterAll(conn, fighterEquipment)
	//clientEntityWriter.CreateComponent(fighterEquipment, avatar)

	questManager := p.GetChildByGCType("QuestManager")
	ownerID := types.UInt16(uint16(rrplayer.Conn.GetID()))

	if questManager == nil {
		questManager = NewQuestManager()
	}

	zone.AddEntity(ownerID, questManager)
	clientEntityWriter.CreateComponentAndInit(questManager, p)

	dialogManager := p.GetChildByGCType("DialogManager")
	if dialogManager == nil {
		dialogManager = NewDialogManager()
	}

	zone.AddEntity(ownerID, dialogManager)
	clientEntityWriter.CreateComponentAndInit(dialogManager, p)

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
	unitContainer := avatar.(*Avatar).GetUnitContainer()
	clientEntityWriter.CreateComponentAndInit(unitContainer, avatar)
	//addCreateComponent(body, avatar.RREntityProperties().ID, NewID(), "UnitContainer")

	// Container::readInit()
	body.WriteUInt32(1)
	body.WriteUInt32(1)
	body.WriteByte(0x03) // Inventory Count?

	baseInventory := unitContainer.GetChildByGCType("avatar.base.Inventory")
	baseInventory.WriteInit(clientEntityWriter.Body)

	bankInventory := unitContainer.GetChildByGCType("avatar.base.Bank")
	bankInventory.WriteInit(clientEntityWriter.Body)

	tradeInventory := unitContainer.GetChildByGCType("avatar.base.TradeInventory")
	tradeInventory.WriteInit(clientEntityWriter.Body)

	// Items with PAL seem to be for players
	//for i := 0; i < inventoryItemCount; i++ {
	//AddInventoryItem(body, "1HAxe2PAL.1HAxe2-1", 0, 0)
	//AddInventoryItem(body, "LeatherArmor1PAL.LeatherArmor1-1", 2, 0, true)
	//AddInventoryItem(body, "CrystalHelm1PAL.CrystalHelm1-1", 2, 0)
	//AddInventoryItem(body, "CrystalMythicPAL.CrystalMythicArmor2", 2, 0)
	//}

	// UnitContainer::readInit()
	body.WriteByte(0x00) // If >0 it tries to read more, something to do with item

	// UNITCONTAINER UPDATE
	//addUnitContainerUpdate(body, 0x01)

	// MODIFIERS //////////////////////////////////
	// Modifiers are for modifying damage and defences
	addCreateComponent(body, uint16(avatar.RREntityProperties().ID), NewID(), "Modifiers")

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
	addCreateComponent(body, uint16(avatar.RREntityProperties().ID), NewID(), "avatar.base.skills")

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
		addCreateComponent(body, uint16(avatar.RREntityProperties().ID), uint16(unitBehaviour.RREntityProperties().ID), "avatar.base.UnitBehavior")

		behav := behavior.NewBehavior()
		behav.Init(body, nil, nil, 0xFF)

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
	body.WriteUInt16(uint16(avatar.RREntityProperties().ID))

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
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	//log.Info(fmt.Sprintf("Sent: \n%s", hex.Dump(body.Data())))
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func AddRandomEquipment(equipment database.EquipmentMap, t ItemType) *Equipment {
	i := 0

	target := int(r.Int63()) % len(equipment)

	for key, class := range equipment {
		if i == target {
			slot, err := class.Slot()

			if err != nil {
				log.Error(err)
				break
			}

			return NewEquipment(
				key, "ScaleModPAL.Rare.Mod1",
				t, slot,
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

func (p *Player) ChangeZone(zoneName string) {
	tZone := Zones.GetOrCreateZone(zoneName)

	if tZone == nil {
		log.Errorf("could not find zone %s", zoneName)
		return
	}

	if p.Zone != nil {
		p.LeaveZone()
	}

	p.JoinZone(tZone)
}

func (p *Player) JoinZone(tZone *Zone) {
	rrPlayer := Players.GetPlayer(uint16(p.ID()))

	p.Zone = tZone
	tZone.AddPlayer(rrPlayer)

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageConnected))
	//body.WriteCString("TheHub")
	//body.WriteCString("Tutorial")
	body.WriteCString(tZone.Name)
	body.WriteUInt32(0xBEEFBEEF)
	body.WriteByte(0x01)
	body.WriteByte(0xFF)
	body.WriteCString("world.town.quest.Q01_a1")
	body.WriteUInt32(0x01)
	connections.WriteCompressedA(rrPlayer.Conn, 0x01, 0x0f, body)

	if config.Config.Logging.LogChangeZone {
		log.Info(fmt.Sprintf("Sent Change Zone: \n%s", hex.Dump(body.Data())))
	}
}

func (p *Player) OnZoneJoin() {
	rrplayer := Players.GetPlayer(uint16(p.ID()))

	p.Spawned = true
	entities := p.Zone.Entities()

	for _, entity := range entities {
		if int(entity.OwnerID()) == rrplayer.Conn.GetID() {
			continue
		}

		CEWriter := NewClientEntityWriterWithByter()
		CEWriter.Create(entity)

		entity.WalkChildren(func(object DRObject) {
			switch object.Type() {
			case DRObjectComponent:
				//if mb2, ok := object.(*MonsterBehavior2); ok {
				//	CEWriter.CreateComponentAndInit(object, entity)
				//}
				CEWriter.CreateComponentAndInit(object, entity)
			}
		})

		CEWriter.Init(entity)

		if unitBehavior, ok := entity.GetChildByGCNativeType("UnitBehavior").(IUnitBehavior); unitBehavior != nil && ok {
			unitBehavior.GetUnitBehavior().WriteWarp(CEWriter)
		}

		rrplayer.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeCreateNPC)
	}
}

func (p *Player) OnZoneLeave() {

}

func (p *Player) LeaveZone() {
	p.Spawned = false
	p.Zone.RemovePlayer(int(p.ID()))

	rrplayer := Players.GetPlayer(uint16(p.ID()))
	rrplayer.MessageQueue.Clear(message.QueueTypeClientEntity)

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageDisconnected))
	body.WriteCString("zoneleaveuhh")
	connections.WriteCompressedA(rrplayer.Conn, 0x01, 0x0f, body)

	if config.Config.Logging.LogChangeZone {
		log.Info(fmt.Sprintf("Sent Leave Zone: \n%s", hex.Dump(body.Data())))
	}

	p.OnZoneLeave()
}

func NewPlayer(name string) (p *Player) {
	p = &Player{
		Name: name,
	}

	p.GCObject = NewGCObject("Player")
	p.GCLabel = "ElliePlayer"
	p.GCType = "player"

	return
}
