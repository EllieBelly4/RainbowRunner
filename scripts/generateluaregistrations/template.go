package main

var (
	//language=gotemplate

	registerTemplate = `// Code generated by scripts/generateluaregistrations DO NOT EDIT.
package {{ .PackageName }}


import lua2 "github.com/yuin/gopher-lua"

func RegisterAllLuaFunctions(state *lua2.LState) {
	{{- range .RegisterFuncs }}
		{{ . }}(state)
	{{- end }}
}
`
)
