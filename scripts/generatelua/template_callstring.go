package main

const (
	//language=gotemplate
	callStringTemplate = `{{ .Func.Name }}(
{{- range $i, $param := .Func.Params -}}
	{{- $ii := add $i $.StartOffset -}}
	{{- if isNumberType $param }}
	{{- generateCheckNumber $param $ii}},
	{{- else if isStringType $param }}
	{{- generateCheckString $param $ii}},
	{{- else if isBoolType $param }}
	{{- generateCheckBool $param $ii}},
	{{- else if $param.IsPointer }}
	lua.CheckReferenceValue[{{ $param.FullTypeString }}](l, {{ add $ii 1 }}),
	{{- else }}
	lua.CheckValue[{{ $param.FullTypeString }}](l, {{ add $ii 1 }}),
	{{- end }}
{{- end }}
)`
)
