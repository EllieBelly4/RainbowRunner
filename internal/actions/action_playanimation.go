package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionPlayAnimation
type ActionPlayAnimation struct {
	Unk0 byte
	Unk1 uint32
	Unk2 uint32
	Unk3 uint32
	Unk4 uint32
}

func (d ActionPlayAnimation) OpCode() BehaviourAction {
	return BehaviourActionPlayAnimation
}

func (d ActionPlayAnimation) Init(body *byter.Byter) {
	body.WriteByte(d.Unk0)
	body.WriteUInt32(d.Unk1)
	body.WriteUInt32(d.Unk2)
	body.WriteUInt32(d.Unk3)
	body.WriteUInt32(d.Unk4)
}

func NewActionPlayAnimation() *ActionPlayAnimation {
	return &ActionPlayAnimation{}
}
