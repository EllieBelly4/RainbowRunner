package main

const (
	ActionTemplate = `package actions

import "RainbowRunner/pkg/byter"

type {{ .ActionName }} struct {
}

func (d {{ .ActionName }}) OpCode() BehaviourAction {
	return BehaviourAction{{ .ActionName }}
}

func (d {{ .ActionName }}) Init(body *byter.Byter) {
	panic("implement me")
}
`
)
