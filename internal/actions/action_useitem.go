package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionUseItem
type ActionUseItem struct {
}

func (a ActionUseItem) OpCode() BehaviourAction {
	return BehaviourActionUseItem
}

func (a ActionUseItem) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionUseItem() *ActionUseItem {
	return &ActionUseItem{}
}
