package actions

import "RainbowRunner/pkg/byter"

type Follow struct {
}

func (d Follow) OpCode() BehaviourAction {
	return BehaviourActionFollow
}

func (d Follow) Init(body *byter.Byter) {
	panic("implement me")
}
