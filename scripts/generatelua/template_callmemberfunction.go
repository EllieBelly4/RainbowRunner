package main

const (
	//language=gotemplate
	callMemberFunctionTemplate string = `func (l *lua2.LState) int {
	objInterface := lua.CheckInterfaceValue[I{{ .Struct.FullTypeString }}](l, 1)
	obj := objInterface.Get{{ .Struct.Name }}()
	
	{{- $resultsLen := len .Method.Results -}}

	{{- if eq $resultsLen 0 }}
	obj.{{ generateCallString .Method 1 }}
	{{- else }}
	{{ .Method.ResultAssignmentString }} := obj.{{ generateCallString .Method 1 -}}
		{{- range $i, $result := .Method.Results }}
			{{- $resVarName := printf "res%d" $i }}
			{{- generateResultPushString $result $resVarName }}
		{{- end }}
	{{- end }}

	return {{ len .Method.Results }}
}`
)

type CallMemberFunctionTemplateData struct {
	Struct *StructDef
	Method *FuncDef
}
