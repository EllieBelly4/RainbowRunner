package main

import "text/template"

var templateFuncMap = template.FuncMap{
	"add":                        Add,
	"isNumberType":               IsNumberType,
	"isStringType":               IsStringType,
	"generateCallString":         GenerateCallString,
	"generateCallMemberFunction": GenerateCallMemberFunction,
	"isLuaConvertible":           IsLuaConvertible,
}

const (
	// language=gotemplate
	templateString string = `// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua2 "github.com/yuin/gopher-lua"
	lua "RainbowRunner/internal/lua"
	{{- range .Imports }}
	{{ .ImportString }}
	{{- end }}
)

type I{{ .Struct.Name }} interface {
	Get{{ .Struct.Name }}() *{{ .Struct.Name }}
}

func ({{ .Struct.MemberInitial }} *{{ .Struct.Name }}) Get{{ .Struct.Name }}() *{{ .Struct.Name }} {
	return {{ .Struct.MemberInitial }}
}

func registerLua{{ .Struct.Name }}(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("{{ .Struct.Name }}")
	state.SetGlobal("{{ .Struct.Name }}", mt)
{{- if .Struct.Constructor }}
	state.SetField(mt, "new", state.NewFunction(newLua{{ .Struct.Name }}))
{{- end }}
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethods{{ .Struct.Name }}(),
	))
}

func luaMethods{{ .Struct.Name }}() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
{{- if .Struct.Fields }}
	{{- $struct := .Struct }}
	{{- range $i, $field := .Struct.Fields }}
		{{- if not $field.IsExported }}
		{{- continue }}
		{{- end }}

		{{- if isStringType $field }}
		"{{ $field.NameCamelcase }}": luaGenericGetSetString[I{{ $struct.Name }}](func(v I{{ $struct.Name }}) *string { return &v.Get{{ $struct.Name }}().{{ $field.Name }} }),
		{{- else if isNumberType $field }}
		"{{ $field.NameCamelcase }}": luaGenericGetSetNumber[I{{ $struct.Name }}](func(v I{{ $struct.Name }}) *{{ $field.FullTypeString }} { return &v.Get{{ $struct.Name }}().{{ $field.Name }} }),
		{{- else if isLuaConvertible $field }}
		"{{ $field.NameCamelcase }}": luaGenericGetSetValue[I{{ $struct.Name }}, {{ $field.FullTypeStringWithPtr }}](func(v I{{ $struct.Name }}) *{{ $field.FullTypeStringWithPtr }} { return &v.Get{{ $struct.Name }}().{{ $field.Name }} }),
		{{- end }}
	{{- end }}
{{- end }}

{{- $struct := .Struct }}
{{- range $i, $method := .Struct.Methods }}
		{{- if not $method.IsExported }}
		{{- continue }}
		{{- end }}
		"{{ $method.NameCamelcase }}": {{ generateCallMemberFunction $struct $method }},
{{- end }}

	}, {{ .ExtendsString }})
}

{{- if .Struct.Constructor }}
func newLua{{ .Struct.Name }}(l *lua2.LState) int {
    obj := {{ generateCallString .Struct.Constructor 0 }}
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("{{ .Struct.Name }}"))
	l.Push(ud)
	return 1
}
{{- end }}

func ({{ .Struct.MemberInitial }} *{{ .Struct.Name }}) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = {{ .Struct.MemberInitial }}

	l.SetMetatable(ud, l.GetTypeMetatable("{{ .Struct.Name }}"))
	return ud
}
`
)
