package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=ActionMoveInDirectionAction
type ActionMoveInDirectionAction struct {
}

func (a ActionMoveInDirectionAction) OpCode() BehaviourAction {
	return BehaviourActionMoveInDirectionAction
}

func (a ActionMoveInDirectionAction) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionMoveInDirectionAction() *ActionMoveInDirectionAction {
	return &ActionMoveInDirectionAction{}
}
