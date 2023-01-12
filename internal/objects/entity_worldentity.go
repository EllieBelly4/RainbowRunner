package objects

import (
	actions2 "RainbowRunner/internal/actions"
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/script"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"RainbowRunner/pkg/datatypes/drfloat"
	log "github.com/sirupsen/logrus"
)

type WorldEntityFlags uint32

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
// 0x07 is the least required to get NPCs working
const (
	WorldEntityFlagStatic WorldEntityFlags = 1 << iota
	WorldEntityFlagCanBeActivated
	WorldEntityFlagVisible
	WorldEntityFlagUnk2
	WorldEntityFlagUnk3
	WorldEntityFlagUnk4
	WorldEntityFlagUnk5
	WorldEntityFlagUnk6
	WorldEntityFlagUnk7
	WorldEntityFlagUnk8
	WorldEntityFlagIsSynched_mb
	WorldEntityFlagBreaksEverything
	WorldEntityFlagMakesCharacterInvisible
	WorldEntityFlagMakesMovementVeryJumpy
	WorldEntityFlagUnk10
	WorldEntityFlagHasAttributes_mb
	WorldEntityFlagUnk12
)

type WorldEntityInitFlags uint8

const (
	WorldEntityInitFlagUnk0 WorldEntityInitFlags = 1 << iota
	WorldEntityInitFlagUnk1
	WorldEntityInitFlagUnk2
	WorldEntityInitFlagCustomAnimationSpeed
	WorldEntityInitFlagUnk4
	WorldEntityInitFlagUnk5
	WorldEntityInitFlagUnk6
	WorldEntityInitFlagUnk7
)

//go:generate go run ../../scripts/generatelua -type=WorldEntity -extends=Entity
type WorldEntity struct {
	*Entity

	HP drfloat.DRFloat
	MP drfloat.DRFloat

	CollisionRadius int

	AnimationsList *AnimationsList

	WorldPosition        datatypes.Vector3Float32
	Heading              float32
	WorldEntityFlags     uint32
	WorldEntityInitFlags byte
	Label                string

	// If Unk1Case is > 0 then it stops animations without any error
	Unk1Case uint16
	Unk2Case byte
	Unk4Case uint32

	UseCustomAnimationSpeed bool
	AnimationSpeed          float32

	luaScript      script.IEntityScript
	CanBeActivated bool

	Spawned bool

	/*
		PropertyWorldEntityPosition
		PropertyWorldEntityHeading
		PropertyWorldEntityBlocking
		PropertyWorldEntityCanBeActivated
		PropertyWorldEntityDescActivationOffset
		PropertyWorldEntityDescActivationRange
		PropertyWorldEntityDescPolyPick
		PropertyWorldEntityDescBlocking
		PropertyWorldEntityDescDynamicBlocking
		PropertyWorldEntityDescActivationRange
		PropertyWorldEntityDescPolyPick
		PropertyWorldEntityDescBlocking
	*/
}

func (g *WorldEntity) WriteSynch(b *byter.Byter) {
	flag := 0x02
	// TODO consider checking the zone to see if it's a town, as it is 0x02 will work in town
	//b.WriteByte(0x00) // 0x00 If in town
	b.WriteByte(byte(flag)) // 0x02 If in dungeon

	if flag == 0x02 {
		b.WriteUInt32(
			g.GetSynch(),
		) // HP - EntitySynchInfo::ReadFromStream
	}
}

func (g *WorldEntity) AddChild(child drobjecttypes.DRObject) {
	g.Entity.AddChild(child)
	child.SetParent(g)
}

func (g *WorldEntity) Tick() {
	if g.luaScript != nil {
		err := g.luaScript.Tick()

		if err != nil {
			log.Errorf("Error in entity script __tick: %s", err)
		}
	}

	for _, object := range g.Children() {
		object.Tick()
	}
}

func (g *WorldEntity) Init() {
	if g.luaScript != nil {
		err := g.luaScript.Init(g)
		if err != nil {
			log.Errorf("Error in entity script __init: %s", err)
		}
	}

	for _, object := range g.Children() {
		object.Init()
	}
}

func (g *WorldEntity) SetScript(script script.IEntityScript) {
	g.luaScript = script

	err := g.luaScript.Load()

	if err != nil {
		panic(err)
	}
}

func (g *WorldEntity) Animations() []*Animation {
	if g.AnimationsList == nil {
		return nil
	}

	return g.AnimationsList.Animations
}

func (g *WorldEntity) Activate(player *RRPlayer, u *UnitBehavior, id byte, sessionID byte) {
	CEWriter := NewClientEntityWriterWithByter()

	CEWriter.BeginComponentUpdate(u)
	CEWriter.CreateActionResponse(actions2.BehaviourActionActivate, id, sessionID)

	activateAction := actions2.ActionActivate{
		TargetEntityID: uint16(g.EntityProperties.ID),
	}

	activateAction.InitWithoutOpCode(CEWriter.Body)
	CEWriter.WriteSynch(u)

	player.MessageQueue.Enqueue(
		message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeBehaviourAction,
	)

	player.CurrentCharacter.GetChildByGCNativeType("Avatar").(*Avatar).SendFollowClient()
}

func (n *WorldEntity) SetPosition(position datatypes.Vector3Float32) {
	n.WorldPosition = position
}

func (n *WorldEntity) SetRotation(degrees float32) {
	n.Heading = degrees
}

func (n *WorldEntity) Type() drobjecttypes.DRObjectType {
	return drobjecttypes.DRObjectEntity
}

func (n *WorldEntity) WriteInit(b *byter.Byter) {
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
	// 0x07 is the least required to get NPCs working
	b.WriteUInt32(
		n.worldEntityFlags(), // With this one alone it was working
	)
	// These positions stopped working at some point
	b.WriteInt32(int32(n.WorldPosition.X * 256)) // Pos X
	b.WriteInt32(int32(n.WorldPosition.Y * 256)) // Pos Y
	b.WriteInt32(int32(n.WorldPosition.Z * 256)) // Pos Z
	b.WriteInt32(int32(n.Heading * 256))

	// Flags
	// Each flag adds one more section of data to read sequentially
	// 0x01 Has Parent?
	// 0x02 Unk
	// 0x04 Makes movement smoother, interpolated position?
	// 0x08 Unk
	//b.WriteByte(1 | 2 | 4 | 8)
	// When this is set to 0 the character is slightly less broken
	// With 1 | 2 | 4 | 8 it was causing the character to have no animations and
	// eventually collapse into itself
	//n.WorldEntityInitFlags := 0x04 | 0x08

	initFlags := n.WorldEntityInitFlags

	if n.UseCustomAnimationSpeed {
		initFlags |= byte(WorldEntityInitFlagCustomAnimationSpeed)
	} else {
		initFlags &= ^byte(WorldEntityInitFlagCustomAnimationSpeed)
	}

	b.WriteByte(initFlags)

	// Has parent/owner?
	if n.WorldEntityInitFlags&0x01 > 0 {
		if n.GCParent == nil {
			b.WriteUInt16(0)
		} else {
			b.WriteUInt16(uint16(n.GCParent.(IRREntityPropertiesHaver).GetRREntityProperties().ID))
		}
	}

	if n.WorldEntityInitFlags&0x02 > 0 {
		// Ox02
		b.WriteByte(n.Unk2Case)
	}

	if n.WorldEntityInitFlags&0x04 > 0 {
		// 0x04
		b.WriteUInt32(n.Unk4Case)
	}

	if n.UseCustomAnimationSpeed {
		// 0x08
		b.WriteUInt32(uint32(n.AnimationSpeed * 256))
	}
}

func (n *WorldEntity) worldEntityFlags() uint32 {
	flags := n.WorldEntityFlags

	if n.CanBeActivated {
		flags |= uint32(WorldEntityFlagCanBeActivated)
	} else {
		flags &= ^uint32(WorldEntityFlagCanBeActivated)
	}

	return flags
}

func (w *WorldEntity) GetSynch() uint32 {
	synch := w.HP.ToWire() & 0xFFFFFF00

	if w.Spawned {
		//Ref: void_Unit::computeAttributes_void_1
		//HP regen needs to be in here
		// HPRegen = ((0x200 * 0x100) >> 8) >> 8 == 0x02
		//synch |= 0xF0
	}

	return synch
}

func NewWorldEntity(gcType string) *WorldEntity {
	entity := NewEntity(gcType)
	entity.GCType = gcType

	return &WorldEntity{
		Entity:               entity,
		HP:                   drfloat.FromInt32(100),
		MP:                   drfloat.FromInt32(100),
		WorldEntityFlags:     0x04,
		WorldEntityInitFlags: 0xFF,
	}
}
