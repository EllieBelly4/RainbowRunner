package main

var (
	//language=gotemplate
	registerTemplate = `package objects


import lua2 "github.com/yuin/gopher-lua"

func registerAllLuaFunctions(state *lua2.LState) {
	{{- range . }}
		{{ . }}(state)
	{{- end }}
}
`
)
