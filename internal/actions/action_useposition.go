package actions

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

//go:generate go run ../../scripts/generatelua -type=ActionUsePosition
type ActionUsePosition struct {
	ActionID byte

	Position datatypes.Vector3Float32
}

func (a ActionUsePosition) OpCode() BehaviourAction {
	return BehaviourActionUsePosition
}

func (a ActionUsePosition) Init(body *byter.Byter) {
	body.WriteByte(a.ActionID)

	body.WriteUInt32(uint32(a.Position.X * 256))
	body.WriteUInt32(uint32(a.Position.Y * 256))
	body.WriteUInt32(uint32(a.Position.Z * 256))
}

func NewActionUsePosition() *ActionUsePosition {
	return &ActionUsePosition{}
}
