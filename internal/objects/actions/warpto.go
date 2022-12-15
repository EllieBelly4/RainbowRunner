package actions

import "RainbowRunner/pkg/byter"

type WarpTo struct {
	PosX uint32
	PosY uint32
}

func (a *WarpTo) OpCode() BehaviourAction {
	return 17
}

func (a *WarpTo) Init(body *byter.Byter, sessionID byte) {
	// WarpTo::readData
	body.WriteByte(sessionID)
	body.WriteUInt32(a.PosX)
	body.WriteUInt32(a.PosY)
}
