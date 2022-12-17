package main

const (
	//language=gotemplate
	resultPushTemplate string = `{{- if .Result.IsArray }}
			{{ .VarName }}Array := l.NewTable()

			for _, {{ .VarName }} := range {{ .VarName }} {
			{{- end }}

			{{- if isNumberType .Result }}
			{{ .PushString }}(lua2.LNumber({{ .VarName }}))
			{{- else if isStringType .Result }}
			{{ .PushString }}(lua2.LString({{ .VarName }}))
			{{ else if isBoolType .Result }}
			{{ .PushString }}(lua2.LBool({{ .VarName }}))
			{{- else if isResultLuaConvertible .Result }}
			if {{ .VarName }} != nil {
				{{ .PushString }}({{ .VarName }}.ToLua(l))
			} else {
				{{ .PushString }}(lua2.LNil)
			}
			{{- else }}
			ud := l.NewUserData()
			ud.Value = {{ .VarName }}
			l.SetMetatable(ud, l.GetTypeMetatable("{{ .Result.FullTypeString }}"))
			{{ .PushString }}(ud)
			{{- end }}

			{{- if .Result.IsArray }}
			}
			{{- end }}
`
)
