package actions

import "RainbowRunner/pkg/byter"

type MoveInDirectionAction struct {
}

func (d MoveInDirectionAction) OpCode() BehaviourAction {
	return BehaviourActionMoveInDirectionAction
}

func (d MoveInDirectionAction) Init(body *byter.Byter) {
	panic("implement me")
}
