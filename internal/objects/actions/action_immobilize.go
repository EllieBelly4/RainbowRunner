package actions

import "RainbowRunner/pkg/byter"

type Immobilize struct {
}

func (d Immobilize) OpCode() BehaviourAction {
	return BehaviourActionImmobilize
}

func (d Immobilize) Init(body *byter.Byter) {
	panic("implement me")
}
