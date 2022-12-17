package main

import (
	"strings"
	"text/template"
)

func GenerateCallMemberFunction(s *StructDef, def *FuncDef) string {
	t := template.New("callMemberFunctionTemplate")

	t.Funcs(mergeFuncMaps(template.FuncMap{
		"generateCallString":       GenerateCallString,
		"add":                      Add,
		"generateResultPushString": generateResultPushString,
	}, typeCheckFunctions))

	t, err := t.Parse(callMemberFunctionTemplate)

	if err != nil {
		panic(err)
	}

	var b strings.Builder

	err = t.Execute(&b, &CallMemberFunctionTemplateData{
		Struct: s,
		Method: def,
	})

	if err != nil {
		panic(err)
	}

	return b.String()
}
