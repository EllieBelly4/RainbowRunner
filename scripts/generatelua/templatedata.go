package main

import (
	"fmt"
)

type TemplateData struct {
	Struct  *StructDef
	Imports []string
}

func (t *TemplateData) StructTypeNameVar() string {
	return fmt.Sprintf("lua%sTypeName", t.Struct.Name)
}

func Add(a, b int) int {
	return a + b
}

func IsNumberType(t IValueType) bool {
	tString := t.GetParamType()

	if tString == "int" ||
		tString == "int8" ||
		tString == "int16" ||
		tString == "int32" ||
		tString == "int64" ||
		tString == "uint" ||
		tString == "uint8" ||
		tString == "uint16" ||
		tString == "uint32" ||
		tString == "uint64" ||
		tString == "uintptr" ||
		tString == "float32" ||
		tString == "float64" {
		return true
	}

	return false
}

func IsStringType(t IValueType) bool {
	if t.GetParamType() == "string" {
		return true
	}

	return false
}
