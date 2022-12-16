package actions

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

type WarpTo struct {
	Position datatypes.Vector3Float32
}

func (a *WarpTo) OpCode() BehaviourAction {
	return 17
}

func (a *WarpTo) Init(body *byter.Byter, sessionID byte) {

	// TODO verify this is correct
	// This was used for the "embedded" actions where the action was
	// written along with the behavior
	//// WarpTo::readData
	//body.WriteByte(sessionID)
	//body.WriteUInt32(a.PosX)
	//body.WriteUInt32(a.PosY)
}
