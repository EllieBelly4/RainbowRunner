package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionUse
type ActionUse struct {
}

func (a ActionUse) OpCode() BehaviourAction {
	return BehaviourActionUse
}

func (a ActionUse) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionUse() *ActionUse {
	return &ActionUse{}
}
