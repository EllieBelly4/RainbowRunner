package objects

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/objects/actions"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
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
	WorldEntityFlagInteractable
	WorldEntityFlagVisible
	WorldEntityFlagUnk2
	WorldEntityFlagUnk3
	WorldEntityFlagUnk4
	WorldEntityFlagUnk5
	WorldEntityFlagUnk6
	WorldEntityFlagUnk7
	WorldEntityFlagUnk8
	WorldEntityFlagUnk9
	WorldEntityFlagBreaksEverything
	WorldEntityFlagMakesCharacterInvisible
	WorldEntityFlagMakesMovementVeryJumpy
	WorldEntityFlagUnk10
	WorldEntityFlagUnk11
	WorldEntityFlagUnk12

	WorldEntityFlagInteractableNPC = WorldEntityFlagStatic | WorldEntityFlagInteractable | WorldEntityFlagVisible
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
	WorldPosition        datatypes.Vector3Float32
	Rotation             float32
	WorldEntityFlags     uint32
	WorldEntityInitFlags byte
	Label                string

	// If Unk1Case is > 0 then it stops animations without any error
	Unk1Case uint16
	Unk2Case byte
	Unk4Case uint32

	UseCustomAnimationSpeed bool
	AnimationSpeed          float32
}

func (g *WorldEntity) Activate(player *RRPlayer, u *UnitBehavior, id byte, sessionID byte) {
	CEWriter := NewClientEntityWriterWithByter()

	CEWriter.BeginComponentUpdate(u)
	CEWriter.CreateActionResponse(actions.BehaviourActionActivate, id)

	activateAction := actions.Activate{
		TargetEntityID: uint16(g.EntityProperties.ID),
	}

	activateAction.InitWithoutOpCode(CEWriter.Body, sessionID)
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
	n.Rotation = degrees
}

func (n *WorldEntity) Type() DRObjectType {
	return DRObjectEntity
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
		n.WorldEntityFlags, // With this one alone it was working
	)
	// These positions stopped working at some point
	b.WriteInt32(int32(n.WorldPosition.X * 256)) // Pos X
	b.WriteInt32(int32(n.WorldPosition.Y * 256)) // Pos Y
	b.WriteInt32(int32(n.WorldPosition.Z * 256)) // Pos Z
	b.WriteInt32(int32(n.Rotation * 256))

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

	if n.WorldEntityInitFlags&0x01 > 0 {
		// 0x01
		b.WriteUInt16(n.Unk1Case)
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

func NewWorldEntity(gcType string) *WorldEntity {
	entity := NewEntity(gcType)
	entity.GCType = gcType

	return &WorldEntity{
		Entity:               entity,
		WorldEntityFlags:     0x04,
		WorldEntityInitFlags: 0xFF,
	}
}
