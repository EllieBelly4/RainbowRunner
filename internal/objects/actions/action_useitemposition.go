package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=UseItemPosition
type UseItemPosition struct {
}

func (d UseItemPosition) OpCode() BehaviourAction {
	return BehaviourActionUseItemPosition
}

func (d UseItemPosition) Init(body *byter.Byter) {
	panic("implement me")
}
