package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=MoveInDirectionAction
type MoveInDirectionAction struct {
}

func (d MoveInDirectionAction) OpCode() BehaviourAction {
	return BehaviourActionMoveInDirectionAction
}

func (d MoveInDirectionAction) Init(body *byter.Byter) {
	panic("implement me")
}
