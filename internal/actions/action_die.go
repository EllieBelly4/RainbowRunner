package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionDie
type ActionDie struct {
}

func (d ActionDie) OpCode() BehaviourAction {
	return BehaviourActionDie
}

func (d ActionDie) Init(body *byter.Byter) {
	// FaceTarget::readInit
	body.WriteByte(0x00) // Unk

	// UnSpawn::readInit
	body.WriteByte(0x01) // Unk
	body.WriteByte(0x01) // Unk
}
