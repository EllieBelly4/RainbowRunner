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
			{{- if isNumberType $result }}
			l.Push(lua2.LNumber({{ $resVarName }}))
			{{- else if isStringType $result }}
			l.Push(lua2.LString({{ $resVarName }}))
			{{ else if isBoolType $result }}
			l.Push(lua2.LBool({{ $resVarName }}))
			{{- else if isResultLuaConvertible $result }}
			if {{ $resVarName }} != nil {
				l.Push({{ $resVarName }}.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}
			{{- else }}
			ud := l.NewUserData()
			ud.Value = {{ $resVarName }}
			l.SetMetatable(ud, l.GetTypeMetatable("{{ $result.FullTypeString }}"))
			l.Push(ud)
			{{- end }}
		{{- end }}
	{{- end }}

	return {{ len .Method.Results }}
}`
)

type CallMemberFunctionTemplateData struct {
	Struct *StructDef
	Method *FuncDef
}
