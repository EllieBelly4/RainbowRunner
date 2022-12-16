package actions

import "RainbowRunner/pkg/byter"

type Die struct {
}

func (d Die) OpCode() BehaviourAction {
	return 0xFF
}

func (d Die) Init(body *byter.Byter, sessionID byte) {
	// FaceTarget::readInit
	body.WriteByte(0x00) // Unk

	// UnSpawn::readInit
	body.WriteByte(0x01) // Unk
	body.WriteByte(0x01) // Unk
}
