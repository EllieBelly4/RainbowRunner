package main

import (
	"fmt"
	"strings"
	"text/template"
)

const (
	//language=gotemplate
	callStringTemplate = `{{ .Name }}(
{{- range $i, $param := .Params -}}
	{{- if isNumberType $param }}
	{{- generateCheckNumber $param $i}},
	{{- else if isStringType $param }}
	{{- generateCheckString $param $i}},
	{{- else if $param.IsPointer }}
	lua.CheckReferenceValue[{{ $param.FullTypeString }}](l, {{ add $i 1 }}),
	{{- else }}
	lua.CheckValue[{{ $param.FullTypeString }}](l, {{ add $i 1 }}),
	{{- end }}
{{- end }}
)`
)

func generateCheckNumber(param FuncParamDef, index int) string {
	s := fmt.Sprintf("%s(l.CheckNumber(%d))", param.FullTypeString(), index+1)

	if param.IsPointer {
		s = pointerWrapValue(param, s)
	}

	return s
}

func generateCheckString(param FuncParamDef, index int) string {
	s := fmt.Sprintf("%s(l.CheckString(%d))", param.FullTypeString(), index+1)

	if param.IsPointer {
		s = pointerWrapValue(param, s)
	}

	return s
}

func pointerWrapValue(param FuncParamDef, s string) string {
	return fmt.Sprintf("func(v %s) *%s {return &v}(%s)", param.FullTypeString(), param.FullTypeString(), s)
}

func GenerateCallString(def FuncDef) string {
	t := template.New("callStringTemplate")

	t.Funcs(template.FuncMap{
		"add":                 Add,
		"isNumberType":        IsNumberType,
		"isStringType":        IsStringType,
		"generateCheckNumber": generateCheckNumber,
		"generateCheckString": generateCheckString,
	})

	t, err := t.Parse(callStringTemplate)

	if err != nil {
		panic(err)
	}

	var b strings.Builder

	err = t.Execute(&b, def)

	if err != nil {
		panic(err)
	}

	return b.String()
}
