package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=ActionAttackTarget2
type ActionAttackTarget2 struct {
	Unk0     byte
	TargetID uint16
}

func (a ActionAttackTarget2) OpCode() BehaviourAction {
	return BehaviourActionAttackTarget2
}

func (a ActionAttackTarget2) Init(body *byter.Byter) {
	body.WriteByte(a.Unk0)
	body.WriteUInt16(a.TargetID)
}

func NewActionAttackTarget2() *ActionAttackTarget2 {
	return &ActionAttackTarget2{}
}
