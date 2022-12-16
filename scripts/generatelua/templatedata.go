package main

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
	"unicode"
)

type TemplateData struct {
	Struct  *StructDef
	Imports []*ImportDef
	Extends []*FuncDef
}

var packageCache = make(map[string]*packages.Package)

func (t *TemplateData) StructTypeNameVar() string {
	return fmt.Sprintf("lua%sTypeName", t.Struct.Name)
}

func Add(a, b int) int {
	return a + b
}

func IsNumberType(t IValueType) bool {
	tString := t.GetParamType()

	if tString == "byte" ||
		tString == "int" ||
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

func IsBoolType(v IValueType) bool {
	return v.GetParamType() == "bool"
}

func IsStringType(t IValueType) bool {
	return t.GetParamType() == "string"
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

func (t *TemplateData) ExtendsString() string {
	var s strings.Builder

	for i, f := range t.Extends {
		if i > 0 {
			s.WriteString(", ")
		}

		s.WriteString(f.Name)
	}

	return s.String()
}

func IsFieldLuaConvertible(t *FieldDef) bool {
	return isLuaConvertible(t.ValueType)
}

func IsResultLuaConvertible(t *FuncResultDef) bool {
	return isLuaConvertible(t.ValueType)
}

func isLuaConvertible(t ValueType) bool {
	pkg := getPackage(t.Package)

	fullTypeName := pkg.PkgPath + "." + t.ParamType

	for _, f := range pkg.TypesInfo.Types {
		if f.Type.String() == fullTypeName {
			if named, ok := f.Type.(*types.Named); ok {
				if iface, ok := named.Underlying().(*types.Interface); ok {
					for i := 0; i < iface.NumMethods(); i++ {
						m := iface.Method(i)

						if m.Name() == "ToLua" {
							return true
						}
					}
					continue
				}

				for i := 0; i < named.NumMethods(); i++ {
					m := named.Method(i)

					if m.Name() == "ToLua" {
						return true
					}
				}
			}
		}
	}

	return false
}

func isInterface(t ValueType) bool {
	pkg := getPackage(t.Package)

	fullTypeName := pkg.PkgPath + "." + t.ParamType

	for _, f := range pkg.TypesInfo.Types {
		if f.Type.String() == fullTypeName {
			if named, ok := f.Type.(*types.Named); ok {
				if _, ok := named.Underlying().(*types.Interface); ok {
					return true
				}
			}
		}
	}

	return false
}

func getPackage(p string) *packages.Package {
	if p == "" {
		return currentPkg
	}

	pkg, ok := packageCache[p]

	if ok {
		return pkg
	}

	allPkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedTypesInfo | packages.NeedSyntax,
	}, imports[p].Path)

	if err != nil {
		panic(err)
	}

	packageCache[p] = allPkgs[0]

	return allPkgs[0]
}
