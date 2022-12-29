package actions

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes/drfloat"
)

//go:generate go run ../../scripts/generatelua -type=ActionMoveTo
type ActionMoveTo struct {
	PosX float32
	PosY float32
}

func (a *ActionMoveTo) OpCode() BehaviourAction {
	return BehaviourActionMoveTo
}

func (a *ActionMoveTo) Init(body *byter.Byter) {
	// MoveTo::readData
	x := drfloat.FromFloat32(a.PosX)
	y := drfloat.FromFloat32(a.PosY)
	body.WriteUInt32(x.ToWire())
	body.WriteUInt32(y.ToWire())

	//// MoveTo::readInit
	//// Not used when embedding in Behavior
	//// ^ doesn't seem to always be true, MonsterBehavior2 seems to require it
	//body.WriteByte(0x05)
}

func NewActionMoveTo() *ActionMoveTo {
	return &ActionMoveTo{}
}
