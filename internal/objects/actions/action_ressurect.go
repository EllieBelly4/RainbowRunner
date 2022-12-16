package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=Ressurect
type Ressurect struct {
}

func (d Ressurect) OpCode() BehaviourAction {
	return BehaviourActionRessurect
}

func (d Ressurect) Init(body *byter.Byter) {
	panic("implement me")
}
