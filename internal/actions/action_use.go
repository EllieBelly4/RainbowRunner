package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionUse
type ActionUse struct {
	SlotID byte
}

func (a ActionUse) OpCode() BehaviourAction {
	return BehaviourActionUse
}

func (a ActionUse) Init(body *byter.Byter) {
	body.WriteByte(a.SlotID)
}

func NewActionUse() *ActionUse {
	return &ActionUse{}
}
