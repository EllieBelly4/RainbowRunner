package main

import (
	"strings"
	"text/template"
)

const callStringTemplate = `{{ .Name }}(
{{- range $i, $param := .Params }}
	{{- if isNumberType $param }}
	{{ $param.FullTypeString }}(l.CheckNumber({{ add $i 1 }})),
	{{- else if isStringType $param }}
	l.CheckString({{ add $i 1 }}),
	{{- else }}
	lua.CheckReferenceValue[{{ $param.FullTypeString }}](l, {{ add $i 1 }}),
	{{- end }}
{{- end }}
)`

func GenerateCallString(def FuncDef) string {
	t := template.New("callStringTemplate")

	t.Funcs(template.FuncMap{
		"add":          Add,
		"isNumberType": IsNumberType,
		"isStringType": IsStringType,
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
