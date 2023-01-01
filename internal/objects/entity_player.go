package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/serverconfig"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"crypto/md5"
	"encoding/binary"
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

func (p *Player) Type() drobjecttypes.DRObjectType {
	return drobjecttypes.DRObjectOther
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

func (p *Player) SendCreateNewPlayerEntity() {
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

	manipulators := avatar.GetChildByGCNativeType("Manipulators")
	clientEntityWriter.CreateComponentAndInit(manipulators, avatar)

	equipment := avatar.GetChildByGCType("avatar.base.Equipment")
	addCreateComponent(body, uint16(avatar.(IRREntityPropertiesHaver).GetRREntityProperties().ID), uint16(equipment.(IRREntityPropertiesHaver).GetRREntityProperties().ID), "avatar.base.Equipment")

	body.WriteByte(byte(len(equippedItems)))

	for _, equippedItem := range equippedItems {
		equippedItem.GetEquipment().WriteInit(body)
	}

	questManager := p.GetChildByGCType("QuestManager")
	clientEntityWriter.CreateComponentAndInit(questManager, p)

	dialogManager := p.GetChildByGCType("DialogManager")
	clientEntityWriter.CreateComponentAndInit(dialogManager, p)

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

	// UnitContainer::readInit()
	body.WriteByte(0x00) // If >0 it tries to read more, something to do with item

	// UNITCONTAINER UPDATE
	//addUnitContainerUpdate(body, 0x01)

	// MODIFIERS //////////////////////////////////
	// Modifiers are for modifying damage and defences
	addCreateComponent(body, uint16(avatar.(IRREntityPropertiesHaver).GetRREntityProperties().ID), NewID(), "Modifiers")

	// Modifiers::readInit
	body.WriteUInt32(0x00) //
	body.WriteUInt32(0x00) //

	// GCObject::readChildData<Modifier>
	body.WriteByte(0x00)

	// SKILLS //////////////////////////////////
	skills := avatar.GetChildByGCNativeType("Skills")
	clientEntityWriter.CreateComponentAndInit(skills, avatar)

	// UnitBehaviour//////////////////////////////////
	unitBehaviour := avatar.GetChildByGCNativeType("UnitBehavior")

	addCreateComponent(body, uint16(avatar.(IRREntityPropertiesHaver).GetRREntityProperties().ID), uint16(unitBehaviour.(IRREntityPropertiesHaver).GetRREntityProperties().ID), "avatar.base.UnitBehavior")

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

	clientEntityWriter.Init(avatar)

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

	body.WriteByte(70) // Now connected
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	//log.Info(fmt.Sprintf("Sent: \n%s", hex.Dump(body.Data())))
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func AddRandomEquipment(equipment database.EquipmentMap, t ItemType) drobjecttypes.DRObject {
	i := 0

	target := int(r.Int63()) % len(equipment)

	for key, class := range equipment {
		if i == target {
			slot, err := class.Slot()

			if err != nil {
				log.Error(err)
				break
			}

			if t == ItemMeleeWeapon {
				return NewMeleeWeapon(
					key, "ScaleModPAL.Rare.Mod1",
				)
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
	rrPlayer := Players.GetPlayer(p.OwnerID())

	p.Zone = tZone
	tZone.AddPlayer(rrPlayer)

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageConnected))
	//body.WriteCString("TheHub")
	//body.WriteCString("Tutorial")
	body.WriteCString(tZone.Name)

	md5Seed := md5.Sum([]byte(serverconfig.Config.ZoneOptions.Seed))

	zoneSeed := binary.LittleEndian.Uint32(md5Seed[:])

	if serverconfig.Config.ZoneOptions.UseRandomSeed {
		zoneSeed = r.Uint32()
	}

	body.WriteUInt32(zoneSeed)
	body.WriteByte(0x01)
	body.WriteByte(0xFF)
	body.WriteCString("world.town.quest.Q01_a1")
	body.WriteUInt32(0x01)
	connections.WriteCompressedA(rrPlayer.Conn, 0x01, 0x0f, body)

	if serverconfig.Config.Logging.LogChangeZone {
		log.Info(fmt.Sprintf("Sent Change Zone: \n%s", hex.Dump(body.Data())))
	}
}

func (p *Player) OnZoneJoin() {
	rrplayer := Players.GetPlayer(p.OwnerID())

	p.Spawned = true
	entities := p.Zone.Entities()

	for _, entity := range entities {
		if int(entity.OwnerID()) == rrplayer.Conn.GetID() {
			continue
		}

		if _, ok := entity.(IPlayer); ok {
			continue
		}

		CEWriter := NewClientEntityWriterWithByter()
		CEWriter.Create(entity)

		entity.WalkChildren(func(object drobjecttypes.DRObject) {
			switch object.(type) {
			case IComponent:
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

	rrplayer := Players.GetPlayer(uint16(p.OwnerID()))
	rrplayer.MessageQueue.Clear(message.QueueTypeClientEntity)

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageDisconnected))
	body.WriteCString("zoneleaveuhh")
	connections.WriteCompressedA(rrplayer.Conn, 0x01, 0x0f, body)

	if serverconfig.Config.Logging.LogChangeZone {
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
