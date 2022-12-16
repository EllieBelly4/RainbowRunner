package actions

import "RainbowRunner/pkg/byter"

type RetrieveItem struct {
}

func (d RetrieveItem) OpCode() BehaviourAction {
	return BehaviourActionRetrieveItem
}

func (d RetrieveItem) Init(body *byter.Byter) {
	panic("implement me")
}
