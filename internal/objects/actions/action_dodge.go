package actions

import "RainbowRunner/pkg/byter"

type Dodge struct {
}

func (d Dodge) OpCode() BehaviourAction {
	return BehaviourActionDodge
}

func (d Dodge) Init(body *byter.Byter) {
	panic("implement me")
}
