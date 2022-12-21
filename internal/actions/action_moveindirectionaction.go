package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionMoveInDirectionAction
type ActionMoveInDirectionAction struct {
	Unk0 byte
}

func (a ActionMoveInDirectionAction) OpCode() BehaviourAction {
	return BehaviourActionMoveInDirectionAction
}

func (a ActionMoveInDirectionAction) Init(body *byter.Byter) {
	body.WriteByte(a.Unk0)
}

func NewActionMoveInDirectionAction() *ActionMoveInDirectionAction {
	return &ActionMoveInDirectionAction{}
}
