package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=TurnAction
type TurnAction struct {
}

func (d TurnAction) OpCode() BehaviourAction {
	return BehaviourActionTurnAction
}

func (d TurnAction) Init(body *byter.Byter) {
	panic("implement me")
}
