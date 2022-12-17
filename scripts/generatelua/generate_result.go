package main

import (
	"strings"
	"text/template"
)

type generateResultTemplateData struct {
	Result  *FuncResultDef
	VarName string
}

func (g generateResultTemplateData) PushString() string {
	pushString := "l.Push"

	if !g.Result.IsArray {
		return pushString
	}

	return g.VarName + "Array.Append"
}

func generateResultPushString(result *FuncResultDef, varName string) string {
	t := template.New("resultPushTemplate")

	t.Funcs(mergeFuncMaps(template.FuncMap{
		"generateCallString":       GenerateCallString,
		"add":                      Add,
		"generateResultPushString": generateResultPushString,
	}, typeCheckFunctions))

	t, err := t.Parse(resultPushTemplate)

	if err != nil {
		panic(err)
	}

	var b strings.Builder

	err = t.Execute(&b, generateResultTemplateData{
		Result:  result,
		VarName: varName,
	})

	if err != nil {
		panic(err)
	}

	return b.String()
}
