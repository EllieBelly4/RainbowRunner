package behavior

import (
	byter "RainbowRunner/pkg/byter"
)

//go:generate stringer -type=BehaviourAction
type BehaviourAction byte

const (
	BehaviourActionMoveTo                BehaviourAction = 1
	BehaviourActionSpawn                 BehaviourAction = 4
	BehaviourActionActivate              BehaviourAction = 6
	BehaviourActionKnockBack             BehaviourAction = 10
	BehaviourActionKnockDown             BehaviourAction = 11
	BehaviourActionStun                  BehaviourAction = 12
	BehaviourActionSearchForAttack       BehaviourAction = 13
	BehaviourActionSpawnAnimation        BehaviourAction = 14
	BehaviourActionUnSpawn               BehaviourAction = 15
	BehaviourActionDodge                 BehaviourAction = 16
	BehaviourActionWarpTo                BehaviourAction = 17
	BehaviourActionAmbush                BehaviourAction = 18
	BehaviourActionMoveInDirectionAction BehaviourAction = 19
	BehaviourActionTurnAction            BehaviourAction = 20
	BehaviourActionWander                BehaviourAction = 21
	BehaviourActionFollow                BehaviourAction = 22
	BehaviourActionPlayAnimation         BehaviourAction = 32
	BehaviourActionFaceTarget            BehaviourAction = 33
	BehaviourActionWait                  BehaviourAction = 34
	BehaviourActionImmobilize            BehaviourAction = 35
	BehaviourActionIdle                  BehaviourAction = 47
	BehaviourActionRessurect             BehaviourAction = 48
	BehaviourActionRemove                BehaviourAction = 64
	BehaviourActionHide                  BehaviourAction = 65
	BehaviourActionSetBlocking           BehaviourAction = 69
	BehaviourActionFlee                  BehaviourAction = 79
	BehaviourActionUseTarget             BehaviourAction = 80
	BehaviourActionUsePosition           BehaviourAction = 81
	BehaviourActionUse                   BehaviourAction = 82
	BehaviourActionUseItemTarget         BehaviourAction = 83
	BehaviourActionUseItemPosition       BehaviourAction = 84
	BehaviourActionUseItem               BehaviourAction = 85
	BehaviourActionDoEffect              BehaviourAction = 131
	BehaviourActionRetrieveItem          BehaviourAction = 160
	BehaviourActionConvertItemsToGold    BehaviourAction = 161
	BehaviourActionAttackTarget2         BehaviourAction = 240
	BehaviourActionKill                  BehaviourAction = 254
	BehaviourActionDie                   BehaviourAction = 255
)

// Action See [resources/Docs/v2/Behavior.md]
type Action interface {
	OpCode() uint8
	Init(body *byter.Byter)
}

type MoveTo struct {
	PosX uint32
	PosY uint32
}

func (a *MoveTo) OpCode() uint8 {
	return 1
}

func (a *MoveTo) Init(body *byter.Byter) {
	body.WriteByte(a.OpCode())

	// MoveTo::readData
	body.WriteByte(0xFF)
	body.WriteUInt32(a.PosX)
	body.WriteUInt32(a.PosY)

	// MoveTo::readInit
	// Not used when embedding in Behavior
	// ^ doesn't seem to always be true, MonsterBehavior2 seems to require it
	body.WriteByte(0x05)
}

type Activate struct {
	TargetEntityID uint16
}

func (a *Activate) OpCode() uint8 {
	return 6
}

func (a *Activate) Init(body *byter.Byter) {
	body.WriteByte(a.OpCode())

	// Activate::readData
	body.WriteByte(0xFF)
	// Used to be 0x02
	body.WriteUInt16(a.TargetEntityID)

	// StateMachine::ReadMessage
	// Flags
	// 0x02
	// 0x04
	// 0x08
	// 0x10 - Sub message? Chain message?
	// 0x20
	body.WriteByte(0x02 | 0x04 | 0x08 | 0x20)

	body.WriteUInt16(0x01)
	body.WriteUInt16(0x0003)
	body.WriteUInt16(0x01)
	body.WriteUInt32(0x00)
}

func (a *Activate) InitWithoutOpCode(body *byter.Byter) {
	// Activate::readData
	body.WriteByte(0xFF)
	// Used to be 0x02
	body.WriteUInt16(a.TargetEntityID)
}

type Die struct {
}

func (d Die) OpCode() uint8 {
	return 0xFF
}

func (d Die) Init(body *byter.Byter) {
	// FaceTarget::readInit
	body.WriteByte(0x00) // Unk

	// UnSpawn::readInit
	body.WriteByte(0x01) // Unk
	body.WriteByte(0x01) // Unk
}

type WarpTo struct {
	PosX uint32
	PosY uint32
}

func (a *WarpTo) OpCode() uint8 {
	return 17
}

func (a *WarpTo) Init(body *byter.Byter) {
	body.WriteByte(a.OpCode())

	// WarpTo::readData
	body.WriteByte(0xFF)
	body.WriteUInt32(a.PosX)
	body.WriteUInt32(a.PosY)
}
