package actions

import (
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes/drfloat"
)

//go:generate go run ../../scripts/generatelua -type=ActionTurnAction
type ActionTurnAction struct {
	Speed drfloat.DRFloat
}

func (a ActionTurnAction) OpCode() BehaviourAction {
	return BehaviourActionTurnAction
}

func (a ActionTurnAction) Init(body *byter.Byter) {
	body.WriteUInt32(a.Speed.ToWire())
}

func NewActionTurnAction() *ActionTurnAction {
	return &ActionTurnAction{}
}
