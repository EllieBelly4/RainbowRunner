package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TemplateData struct {
	Struct  *StructDef
	Imports []*ImportDef
	Extends []*StructDef
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

func (f *FieldDef) NameCamelcase() string {
	return camelcaseVarName(f.Name)
}

func camelcaseVarName(name string) string {
	var s strings.Builder

	lowered := false

	for _, c := range name {
		if !lowered && unicode.IsUpper(c) {
			s.WriteRune(unicode.ToLower(c))
			continue
		} else {
			lowered = true
		}

		s.WriteRune(c)
	}

	return s.String()
}

func (f *FuncDef) NameCamelcase() string {
	return camelcaseVarName(f.Name)
}

func (i *ImportDef) ImportString() string {
	s := ""

	if i.Name != nil {
		s += *i.Name + " "
	}

	s += "\"" + i.Path + "\""

	return s
}

//func (t *TemplateData) ExtendsString() string {
//	var s strings.Builder
//
//	for i, f := range t.Extends {
//		if i > 0 {
//			s.WriteString(", ")
//		}
//
//		s.WriteString(f.Name)
//	}
//
//	return s.String()
//}
