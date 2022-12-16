package actions

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

//go:generate go run ../../../scripts/generatelua -type=ActionWarpTo
type ActionWarpTo struct {
	Position datatypes.Vector3Float32
}

func (a *ActionWarpTo) OpCode() BehaviourAction {
	return BehaviourActionWarpTo
}

func (a *ActionWarpTo) Init(body *byter.Byter) {
	body.WriteInt32(int32(a.Position.X * 256))
	body.WriteInt32(int32(a.Position.Y * 256))
	body.WriteInt32(int32(a.Position.Z * 256))

	// TODO verify this is still needed
	// This was used for the "embedded" actions where the action was
	// written along with the behavior
	//// WarpTo::readData
	//body.WriteByte(sessionID)
	//body.WriteUInt32(a.PosX)
	//body.WriteUInt32(a.PosY)
}
