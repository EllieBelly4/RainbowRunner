package actions

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

type MoveTo struct {
	PosX float32
	PosY float32
}

func (a *MoveTo) OpCode() BehaviourAction {
	return 1
}

func (a *MoveTo) Init(body *byter.Byter) {
	// MoveTo::readData
	body.WriteUInt32(datatypes.DRFloat(a.PosX).ToUInt())
	body.WriteUInt32(datatypes.DRFloat(a.PosY).ToUInt())

	//// MoveTo::readInit
	//// Not used when embedding in Behavior
	//// ^ doesn't seem to always be true, MonsterBehavior2 seems to require it
	//body.WriteByte(0x05)
}
