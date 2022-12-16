package main

const (
	ActionTemplate = `package actions

import "RainbowRunner/pkg/byter"

//go:generate go run ../../../scripts/generatelua -type=Action{{ .ActionName }}
type Action{{ .ActionName }} struct {
}

func (a Action{{ .ActionName }}) OpCode() BehaviourAction {
	return BehaviourAction{{ .ActionName }}
}

func (a Action{{ .ActionName }}) Init(body *byter.Byter) {
	panic("implement me")
}

func NewAction{{ .ActionName }}() *Action{{ .ActionName }} {
	return &Action{{ .ActionName }}{}
}
`
)
