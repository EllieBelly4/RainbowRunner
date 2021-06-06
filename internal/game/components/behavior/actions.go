package behavior

import "RainbowRunner/internal/byter"

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
	//body.WriteByte(0x05)
}

type Activate struct {
}

func (a *Activate) OpCode() uint8 {
	return 6
}

func (a *Activate) Init(body *byter.Byter) {
	body.WriteByte(a.OpCode())

	// Activate::readData
	body.WriteByte(0xFF)
	body.WriteUInt16(0x02) // EntityID?

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
	body.WriteUInt32(0x0)
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
