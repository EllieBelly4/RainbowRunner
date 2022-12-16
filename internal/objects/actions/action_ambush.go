package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=Ambush
type Ambush struct {
}

func (d Ambush) OpCode() BehaviourAction {
	return BehaviourActionAmbush
}

func (d Ambush) Init(body *byter.Byter) {
	panic("implement me")
}
