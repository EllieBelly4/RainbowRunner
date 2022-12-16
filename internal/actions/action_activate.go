package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionActivate
type ActionActivate struct {
	TargetEntityID uint16
}

func (a *ActionActivate) OpCode() BehaviourAction {
	return BehaviourActionActivate
}

func (a *ActionActivate) Init(body *byter.Byter) {
	// Activate::readData
	//body.WriteByte(sessionID)
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

func (a *ActionActivate) InitWithoutOpCode(body *byter.Byter) {
	// Used to be 0x02
	body.WriteUInt16(a.TargetEntityID)
}

func NewActionActivate() *ActionActivate {
	return &ActionActivate{}
}
