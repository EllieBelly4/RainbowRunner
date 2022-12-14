package main

import "text/template"

var templateFuncMap = template.FuncMap{
	"add":                        Add,
	"generateCallString":         GenerateCallString,
	"generateCallMemberFunction": GenerateCallMemberFunction,
}

var typeCheckFunctions = template.FuncMap{
	"isStringType":           IsStringType,
	"isNumberType":           IsNumberType,
	"isBoolType":             IsBoolType,
	"isInterfaceType":        isInterface,
	"isFieldLuaConvertible":  IsFieldLuaConvertible,
	"isResultLuaConvertible": IsResultLuaConvertible,
	"isLuaConvertible":       isLuaConvertible,
	"isLuaConvertiblePtr":    isLuaConvertiblePtr,
	"isLuaConvertibleValue":  isLuaConvertibleValue,
}

const (
	// language=gotemplate
	templateString string = `// Code generated by scripts/generatelua DO NOT EDIT.
package {{ .PackageName }}

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
	{{- if .Struct.Constructor.IsCallSupported }}
	state.SetField(mt, "new", state.NewFunction(newLua{{ .Struct.Name }}))
	{{- end }}
{{- end }}
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethods{{ .Struct.Name }}(),
	))
}

func luaMethods{{ .Struct.Name }}() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
{{- if .Struct.Fields }}
	{{- $struct := .Struct }}
	{{- range $i, $field := .Struct.Fields }}
		{{- if not $field.IsExported }}
		{{- continue }}
		{{- end }}

		{{- if isStringType $field }}
		"{{ $field.NameCamelcase }}": lua.LuaGenericGetSetString[I{{ $struct.Name }}](func(v I{{ $struct.Name }}) *string { return &v.Get{{ $struct.Name }}().{{ $field.Name }} }),
		{{- else if isNumberType $field }}
		"{{ $field.NameCamelcase }}": lua.LuaGenericGetSetNumber[I{{ $struct.Name }}](func(v I{{ $struct.Name }}) *{{ $field.FullTypeString }} { return &v.Get{{ $struct.Name }}().{{ $field.Name }} }),
		{{- else if isBoolType $field }}
		"{{ $field.NameCamelcase }}": lua.LuaGenericGetSetBool[I{{ $struct.Name }}](func(v I{{ $struct.Name }}) *bool { return &v.Get{{ $struct.Name }}().{{ $field.Name }} }),
		{{- else }}
		"{{ $field.NameCamelcase }}": lua.LuaGenericGetSetValueAny[I{{ $struct.Name }}](func(v I{{ $struct.Name }}) *{{ $field.FullTypeStringWithPtr }} { return &v.Get{{ $struct.Name }}().{{ $field.Name }} }),
		{{- end }}
	{{- end }}
{{- end }}

{{- $struct := .Struct }}
{{- range $i, $method := .Struct.Methods }}
		{{- if not $method.IsExported }}
		{{- continue }}
		{{- end }}
		{{- if not $method.IsCallSupported }}
		{{- continue }}
		{{- end }}

		"{{ $method.NameCamelcase }}": {{ generateCallMemberFunction $struct $method }},
{{- end }}

	}, {{ .ExtendsString }})
}

{{- if .Struct.Constructor }}
{{- if not .Struct.Constructor.IsCallSupported }}
// -------------------------------------------------------------------------------------------------------------
// Unsupported constructor {{ .Struct.Constructor.Name }} is not supported
// -------------------------------------------------------------------------------------------------------------
{{- else }}
func newLua{{ .Struct.Name }}(l *lua2.LState) int {
    obj := {{ generateCallString .Struct.Constructor 0 }}
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("{{ .Struct.Name }}"))
	l.Push(ud)
	return 1
}
{{- end }}
{{- end }}

func ({{ .Struct.MemberInitial }} *{{ .Struct.Name }}) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = {{ .Struct.MemberInitial }}

	l.SetMetatable(ud, l.GetTypeMetatable("{{ .Struct.Name }}"))
	return ud
}
`
)
