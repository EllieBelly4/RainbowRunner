package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionUseItemPosition
type ActionUseItemPosition struct {
}

func (a ActionUseItemPosition) OpCode() BehaviourAction {
	return BehaviourActionUseItemPosition
}

func (a ActionUseItemPosition) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionUseItemPosition() *ActionUseItemPosition {
	return &ActionUseItemPosition{}
}
