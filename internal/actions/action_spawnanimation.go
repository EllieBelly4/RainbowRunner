package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionSpawnAnimation
type ActionSpawnAnimation struct {
	Unk0 byte
	Unk1 byte

	DataUnk0 byte
	DataUnk1 uint16
	DataUnk2 uint16
}

func (a ActionSpawnAnimation) OpCode() BehaviourAction {
	return BehaviourActionSpawnAnimation
}

func (a ActionSpawnAnimation) Init(body *byter.Byter) {
	// readInit
	//body.WriteByte(a.Unk0)
	//body.WriteByte(a.Unk1)

	// readData
	body.WriteByte(a.DataUnk0)
	body.WriteUInt16(a.DataUnk1)
	body.WriteUInt16(a.DataUnk2)

	//panic("implement me")
}

func NewActionSpawnAnimation() *ActionSpawnAnimation {
	return &ActionSpawnAnimation{}
}
