package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/database"
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

func (p *Player) GetRRPlayer() *RRPlayer {
	return Players.GetPlayer(p.OwnerID())
}

func (p *Player) Type() drobjecttypes.DRObjectType {
	return drobjecttypes.DRObjectOther
}

func (p *Player) WriteInit(b *byter.Byter) {
	// Init PLAYER /////////////////////////////////////////
	b.WriteCString(p.Name)
	b.WriteUInt32(0x00)
	b.WriteUInt32(0x00)
	b.WriteByte(0xFF)

	b.WriteUInt32(p.Zone.ID) // World ID
	b.WriteUInt32(1001)      // PvP wins
	b.WriteUInt32(1000)      // PvP rating?, 0 = ???

	// Here goes PvP Team
	// Null string
	b.WriteByte(0x00)

	// If player is in a PvP team then Avatar respawn will look for the team waypoints
	//b.WriteByte(0xFF)
	//b.WriteCString("pvp.DefaultTeamList.BlueTeam")

	b.WriteCString("RainbowRunners") // Posse Name
	b.WriteUInt32(0x00)
}

func (p *Player) WriteUpdate(b *byter.Byter) {
	// This maps to a specific event type for Player::processUpdate()
	b.WriteByte(0x01)

	// 0x03 case
	//b.WriteUInt16(0x02)
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

func (p *Player) WriteCreateNewPlayerEntity(clientEntityWriter *ClientEntityWriter, owned bool) {
	avatar := p.GetChildByGCNativeType("Avatar")
	avatarObj := avatar.(IAvatar).GetAvatar()

	clientEntityWriter.Create(avatar)

	clientEntityWriter.Create(p)
	clientEntityWriter.Init(p)

	manipulators := avatar.GetChildByGCNativeType("Manipulators")
	clientEntityWriter.CreateComponentAndInit(manipulators, avatar)

	equipment := avatar.GetChildByGCType("avatar.base.Equipment")
	clientEntityWriter.CreateComponentAndInit(equipment, avatar)

	questManager := p.GetChildByGCType("QuestManager")
	clientEntityWriter.CreateComponentAndInit(questManager, p)

	dialogManager := p.GetChildByGCType("DialogManager")
	clientEntityWriter.CreateComponentAndInit(dialogManager, p)

	unitContainer := avatar.(*Avatar).GetUnitContainer()
	clientEntityWriter.CreateComponentAndInit(unitContainer, avatar)

	modifiers := avatar.GetChildByGCNativeType("Modifiers")
	clientEntityWriter.CreateComponentAndInit(modifiers, avatar)

	skills := avatar.GetChildByGCNativeType("Skills")
	clientEntityWriter.CreateComponentAndInit(skills, avatar)

	ub := avatarObj.GetUnitBehaviour()
	umf := ub.UnitMoverFlags

	if !owned {
		ub.UnitMoverFlags = 0x01 | 0x04 | 0x80
	}

	unitBehaviour := avatar.GetChildByGCNativeType("UnitBehavior")
	clientEntityWriter.CreateComponentAndInit(unitBehaviour, avatar)

	ub.UnitMoverFlags = umf

	clientEntityWriter.Init(avatar)
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

		CEWriter := NewClientEntityWriterWithByter()

		WriteCreateExistingEntity(entity, CEWriter)

		rrplayer.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeCreateEntity)
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

func (p *Player) GetAvatar() *Avatar {
	return p.GetChildByGCNativeType("Avatar").(IAvatar).GetAvatar()
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
