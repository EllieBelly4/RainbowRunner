package actions

import "RainbowRunner/pkg/byter"

type Activate struct {
	TargetEntityID uint16
}

func (a *Activate) OpCode() BehaviourAction {
	return BehaviourActionActivate
}

func (a *Activate) Init(body *byter.Byter) {
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

func (a *Activate) InitWithoutOpCode(body *byter.Byter) {
	// Used to be 0x02
	body.WriteUInt16(a.TargetEntityID)
}
