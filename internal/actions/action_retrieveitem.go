package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionRetrieveItem
type ActionRetrieveItem struct {
}

func (a ActionRetrieveItem) OpCode() BehaviourAction {
	return BehaviourActionRetrieveItem
}

func (a ActionRetrieveItem) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionRetrieveItem() *ActionRetrieveItem {
	return &ActionRetrieveItem{}
}
