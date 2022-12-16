package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionConvertItemsToGold
type ActionConvertItemsToGold struct {
}

func (a ActionConvertItemsToGold) OpCode() BehaviourAction {
	return BehaviourActionConvertItemsToGold
}

func (a ActionConvertItemsToGold) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionConvertItemsToGold() *ActionConvertItemsToGold {
	return &ActionConvertItemsToGold{}
}
