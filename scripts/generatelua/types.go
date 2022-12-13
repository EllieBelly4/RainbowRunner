package main

import (
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
		if !method.IsExported {
			continue
		}

		for _, param := range method.Params {
			if param.Package != "" {
				imports[param.Package] = true
			}
		}
	}

	if f.Constructor != nil {
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

		// These types must be mirrored in the template or the imports will be wrong
		if !IsStringType(field) &&
			!IsNumberType(field) &&
			!IsFieldLuaConvertible(field) {
			continue
		}

		if field.Package != "" {
			imports[field.Package] = true
		}
	}

	res := make([]*ImportDef, 0)

	for name, _ := range imports {
		res = append(res, importDefs[name])
	}

	return res
}

func (v *ValueType) FullTypeString() string {
	s := ""

	if v.Package != "" {
		s = v.Package + "."
	}

	s += v.ParamType

	return s
}

func (v *ValueType) FullTypeStringWithPtr() string {
	typeString := v.FullTypeString()

	if v.IsPointer {
		typeString = "*" + typeString
	}

	return typeString
}

func (s *StructDef) MemberInitial() string {
	return strings.ToLower(s.Name[:1])
}
