package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go/ast"
	"strings"
)

// IValueType interface for ValueType
type IValueType interface {
	GetParamType() string
	GetPackage() string
}

type ValueType struct {
	Name      string
	ParamType string
	Package   string
	IsPointer bool
	IsArray   bool
}

type FuncParamDef struct {
	ValueType
}

type FuncResultDef struct {
	ValueType
}

type FieldDef struct {
	ValueType
	IsExported bool
	field      *ast.Field
}

func (v ValueType) IsScalar() bool {
	return v.IsStringType() || v.IsNumberType() || v.IsBoolType()
}

func (f *FuncDef) IsCallSupported() bool {
	if f.Params != nil {
		for _, param := range f.Params {
			if param.IsArray {
				log.Warn("Array parameters are not currently supported ", f.Name)
				return false
			}
		}
	}

	return true
}

type FuncDef struct {
	Name     string
	funcType *ast.FuncType

	Params     []*FuncParamDef
	IsExported bool
	Results    []*FuncResultDef
}

type StructDef struct {
	Name        string
	structType  *ast.StructType
	Fields      []*FieldDef
	Methods     []*FuncDef
	Constructor *FuncDef
}

type ImportDef struct {
	Name *string
	Path string
}

// GetParamType method for ValueType that returns the ParamType
func (v *ValueType) GetParamType() string {
	return v.ParamType
}

// GetPackage method for ValueType that returns the Package
func (v *ValueType) GetPackage() string {
	return v.Package
}

func (f *StructDef) FullTypeString() string {
	return f.Name
}

func (f *StructDef) GetRequiredImports(importDefs map[string]*ImportDef) []*ImportDef {
	imports := make(map[string]bool)

	for _, method := range f.Methods {
		if !method.IsExported || !method.IsCallSupported() {
			continue
		}

		for _, param := range method.Params {
			if param.Package != "" {
				imports[param.Package] = true
			}
		}
	}

	if f.Constructor != nil && f.Constructor.IsCallSupported() {
		for _, param := range f.Constructor.Params {
			if param.Package != "" {
				imports[param.Package] = true
			}
		}
	}

	for _, field := range f.Fields {
		if !field.IsExported {
			continue
		}

		if field.Package != "" {
			imports[field.Package] = true
		}
	}

	res := make([]*ImportDef, 0)

	for name, _ := range imports {
		if name == "lua" || name == "lua2" {
			continue
		}

		res = append(res, importDefs[name])
	}

	return res
}

func (v *ValueType) FullTypeString() string {
	s := ""

	if v.IsArray {
		s += "[]"
	}

	if v.Package != "" {
		s = v.Package + "."
	}

	s += v.ParamType

	return s
}

func (v *ValueType) FullTypeStringWithPtr() string {
	typeString := ""

	if v.Package != "" {
		typeString = v.Package + "."
	}

	if v.IsPointer {
		typeString = "*" + typeString
	}

	if v.IsArray {
		typeString = "[]" + typeString
	}

	typeString += v.ParamType

	return typeString
}

func (s *StructDef) MemberInitial() string {
	return strings.ToLower(s.Name[:1])
}

func (v ValueType) IsStringType() bool {
	return v.GetParamType() == "string"
}

func (v ValueType) IsNumberType() bool {
	return IsNumberType(&v)
}

func (v ValueType) IsBoolType() bool {
	return IsBoolType(&v)
}

func (f *FuncDef) ResultAssignmentString() string {
	if len(f.Results) == 0 {
		return ""
	}

	resNames := make([]string, len(f.Results))

	for i := 0; i < len(f.Results); i++ {
		resNames[i] = fmt.Sprintf("res%d", i)
	}

	s := strings.Join(resNames, ", ")

	return s
}
