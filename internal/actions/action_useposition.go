package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionUsePosition
type ActionUsePosition struct {
}

func (a ActionUsePosition) OpCode() BehaviourAction {
	return BehaviourActionUsePosition
}

func (a ActionUsePosition) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionUsePosition() *ActionUsePosition {
	return &ActionUsePosition{}
}
