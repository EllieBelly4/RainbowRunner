package actions

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"RainbowRunner/pkg/datatypes/drfloat"
)

//go:generate go run ../../scripts/generatelua -type=ActionSpawn
type ActionSpawn struct {
	Pos     datatypes.Vector3Float32
	Heading drfloat.DRFloat

	SomeUnitID uint16
}

func (a ActionSpawn) OpCode() BehaviourAction {
	return BehaviourActionSpawn
}

func (a ActionSpawn) Init(body *byter.Byter) {
	pos := a.Pos.ToVector3DRFloat()

	body.WriteUInt32(pos.X.ToWire())
	body.WriteUInt32(pos.Y.ToWire())
	body.WriteUInt32(pos.Z.ToWire())
	body.WriteUInt16(a.SomeUnitID)
}

func NewActionSpawn() *ActionSpawn {
	return &ActionSpawn{}
}
