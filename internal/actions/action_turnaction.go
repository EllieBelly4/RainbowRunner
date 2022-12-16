package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionTurnAction
type ActionTurnAction struct {
}

func (a ActionTurnAction) OpCode() BehaviourAction {
	return BehaviourActionTurnAction
}

func (a ActionTurnAction) Init(body *byter.Byter) {
	panic("implement me")
}

func NewActionTurnAction() *ActionTurnAction {
	return &ActionTurnAction{}
}
