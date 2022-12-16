package actions

import "RainbowRunner/pkg/byter"

type ConvertItemsToGold struct {
}

func (d ConvertItemsToGold) OpCode() BehaviourAction {
	return BehaviourActionConvertItemsToGold
}

func (d ConvertItemsToGold) Init(body *byter.Byter) {
	panic("implement me")
}
