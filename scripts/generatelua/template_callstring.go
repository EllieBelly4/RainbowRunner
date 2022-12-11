package main

import (
	"fmt"
	"strings"
	"text/template"
)

const (
	//language=gotemplate
	callStringTemplate = `{{ .Func.Name }}(
{{- range $i, $param := .Func.Params -}}
	{{- $ii := add $i $.StartOffset -}}
	{{- if isNumberType $param }}
	{{- generateCheckNumber $param $ii}},
	{{- else if isStringType $param }}
	{{- generateCheckString $param $ii}},
	{{- else if $param.IsPointer }}
	lua.CheckReferenceValue[{{ $param.FullTypeString }}](l, {{ add $ii 1 }}),
	{{- else }}
	lua.CheckValue[{{ $param.FullTypeString }}](l, {{ add $ii 1 }}),
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

func GenerateCallString(def FuncDef, startOffset int) string {
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

	err = t.Execute(&b, struct {
		Func        FuncDef
		StartOffset int
	}{
		Func:        def,
		StartOffset: startOffset,
	})

	if err != nil {
		panic(err)
	}

	return b.String()
}
