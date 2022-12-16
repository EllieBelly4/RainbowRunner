package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=UseItem
type UseItem struct {
}

func (d UseItem) OpCode() BehaviourAction {
	return BehaviourActionUseItem
}

func (d UseItem) Init(body *byter.Byter) {
	panic("implement me")
}
