package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	types "go/types"
	"golang.org/x/tools/go/packages"
	"regexp"
	"strings"
)

func addAllMemberFunctions(structs map[string]*StructDef, defs map[string]*FuncDef, p *packages.Package) {
	for _, f := range p.Syntax {
		for _, decl := range f.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				funcDef := NewFuncDef(funcDecl)

				defs[funcDef.Name] = funcDef

				if funcDef.Name == "ToLua" {
					continue
				}

				if funcDecl.Recv != nil {
					for _, field := range funcDecl.Recv.List {
						if field.Names != nil {
							for _, fieldName := range field.Names {
								fieldRecvType := fieldName.Obj.Decl.(*ast.Field).Type

								if starExpr, ok := fieldRecvType.(*ast.StarExpr); ok {
									fieldRecvType = starExpr.X
								}

								structName := types.ExprString(fieldRecvType)

								if structDef, ok := structs[structName]; ok {
									structDef.Methods = append(structDef.Methods, funcDef)
								}
							}
						}
					}
				}
			}
		}
	}
}

func getAllStructDefinitions(structs map[string]*StructDef, cwd string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, cwd, nil, parser.AllErrors)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			err := parseFileStructDefinitions(nil, structs, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func parseFileStructDefinitionsFromString(p *packages.Package, structs map[string]*StructDef, file string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseFile(fset, file, nil, parser.AllErrors)
	if err != nil {
		return err
	}

	return parseFileStructDefinitions(p, structs, pkgs)
}

func parseFileStructDefinitions(p *packages.Package, structs map[string]*StructDef, file *ast.File) error {
	// Iterate over the top-level declarations in the file
	for _, decl := range file.Decls {
		// Check if the declaration is a GenDecl
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			// Check if the GenDecl is a type declaration
			if genDecl.Tok == token.TYPE {
				// Iterate over the GenDecl's specifications
				for _, spec := range genDecl.Specs {
					// Check if the specification is a type spec
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						// Check if the type is a struct
						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							// Print the name of the struct
							//fmt.Printf("%s=========\n", typeSpec.Name)

							structName := typeSpec.Name.String()
							if _, ok := structs[structName]; !ok {
								structs[structName] = &StructDef{
									Name:       structName,
									structType: structType,
								}
							}

							// Iterate over the fields in the struct
							for _, field := range structType.Fields.List {
								if field.Names == nil {
									//if star, ok := field.Type.(*ast.StarExpr); ok {
									//	if starIdent, ok := star.X.(*ast.Ident); ok {
									//		//fmt.Println(starIdent)
									//	}
									//}
								} else {
									if structs[structName].Fields == nil {
										structs[structName].Fields = make([]*FieldDef, 0)
									}

									fieldName := field.Names[0].String()
									fieldDef := &FieldDef{
										ValueType: ValueType{
											Name: fieldName,
										},
										field:      field,
										IsExported: strings.HasPrefix(fieldName, strings.ToUpper(fieldName[:1])),
									}

									addValueType(field, &fieldDef.ValueType)

									structs[structName].Fields = append(structs[structName].Fields, fieldDef)

									// Print the name and type of each field
									//fmt.Printf("%s: %s\n", field.Names[0], field.Type)
								}
							}
						}
					}
				}
			}
		}

		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			funcDef := NewFuncDef(funcDecl)

			if strings.HasPrefix(funcDef.Name, "New") {
				structName := strings.Replace(funcDef.Name, "New", "", 1)

				if _, ok := structs[structName]; !ok {
					structs[structName] = &StructDef{
						Name: structName,
					}
				}

				structs[structName].Constructor = funcDef
			}
		}
	}
	return nil
}

// NewFuncDef creates a new FuncDef from an ast.FuncDecl
func NewFuncDef(decl *ast.FuncDecl) *FuncDef {
	funcName := decl.Name.String()

	funcDef := &FuncDef{
		Name:       funcName,
		funcType:   decl.Type,
		IsExported: strings.HasPrefix(funcName, strings.ToUpper(funcName[:1])),
	}

	if decl.Type.Params != nil {
		funcDef.Params = make([]*FuncParamDef, 0)

		for _, field := range decl.Type.Params.List {
			if field.Names == nil {
				funcParamDef := &FuncParamDef{
					ValueType: ValueType{},
				}

				addValueType(field, &funcParamDef.ValueType)

				funcDef.Params = append(funcDef.Params, funcParamDef)
				continue
			}

			for _, fieldName := range field.Names {
				funcParamDef := &FuncParamDef{
					ValueType: ValueType{
						Name: fieldName.Name,
					},
				}

				addValueType(field, &funcParamDef.ValueType)

				funcDef.Params = append(funcDef.Params, funcParamDef)
			}
		}
	}

	if decl.Type.Results != nil {
		funcDef.Results = make([]*FuncResultDef, 0)

		for _, field := range decl.Type.Results.List {
			if field.Names == nil {
				funcResultDef := &FuncResultDef{
					ValueType: ValueType{},
				}

				addValueType(field, &funcResultDef.ValueType)

				funcDef.Results = append(funcDef.Results, funcResultDef)
				continue
			}

			for _, fieldName := range field.Names {
				funcResultDef := &FuncResultDef{
					ValueType: ValueType{
						Name: fieldName.Name,
					},
				}

				addValueType(field, &funcResultDef.ValueType)

				funcDef.Results = append(funcDef.Results, funcResultDef)
			}
		}
	}

	return funcDef
}

func addValueType(field *ast.Field, funcParamDef *ValueType) {
	if ident, ok := field.Type.(*ast.Ident); ok {
		funcParamDef.ParamType = ident.Name
	} else if star, ok := field.Type.(*ast.StarExpr); ok {
		funcParamDef.IsPointer = true

		if selectorExpr, ok := star.X.(*ast.SelectorExpr); ok {
			pkgString := types.ExprString(selectorExpr.X)
			funcParamDef.Package = pkgString
			funcParamDef.ParamType = selectorExpr.Sel.Name
		} else if ident2, ok := star.X.(*ast.Ident); ok {
			funcParamDef.ParamType = ident2.Name
		}
	} else if selector, ok := field.Type.(*ast.SelectorExpr); ok {
		funcParamDef.Package = types.ExprString(selector.X)
		funcParamDef.ParamType = types.ExprString(selector.Sel)

		if _, ok := selector.X.(*ast.StarExpr); ok {
			funcParamDef.IsPointer = true
		}
	} else {
		funcParamDef.ParamType = types.ExprString(field.Type)
	}
}

func addAllImports(imports map[string]*ImportDef, p *packages.Package) {
	for _, file := range p.Syntax {
		for _, i := range file.Imports {
			cleanPath := strings.ReplaceAll(i.Path.Value, "\"", "")
			splitPath := strings.Split(cleanPath, "/")
			splitPath = strings.Split(splitPath[len(splitPath)-1], "-")
			finalSeg := splitPath[len(splitPath)-1]

			key := finalSeg

			if i.Name != nil {
				key = i.Name.String()
			}

			i2 := &ImportDef{
				Path: cleanPath,
			}

			if i.Name != nil {
				i2.Name = &i.Name.Name
			}

			imports[key] = i2
		}
	}
}

func getExtendStructs(splitExtends []string, structs map[string]*StructDef) ([]*StructDef, error) {
	ret := make([]*StructDef, 0)

	for _, extend := range splitExtends {
		if _, ok := structs[extend]; !ok {
			return nil, errors.New(fmt.Sprintf("could not find type %s", extend))
		}

		ret = append(ret, structs[extend])
	}

	return ret, nil
}

func getExtendFuncs(splitExtends []string, funcDefs map[string]*FuncDef) ([]*FuncDef, error) {
	if splitExtends == nil || len(splitExtends) == 0 {
		return nil, nil
	}

	allFuncDefs := make([]*FuncDef, 0)

	lgFuncRegex := regexp.MustCompile(`map\[string]lua[0-9]?.LGFunction`)

	for _, extend := range splitExtends {
		funcName := fmt.Sprintf("luaMethods%s", extend)

		if _, ok := funcDefs[funcName]; !ok {
			return nil, errors.New(fmt.Sprintf("could not find methods function to extend %s", funcName))
		}

		fun := funcDefs[funcName]

		if len(fun.Results) != 1 {
			return nil, errors.New(fmt.Sprintf("extend function %s does not return a table", funcName))
		}

		// TODO maybe do some proper type checking here
		if !lgFuncRegex.Match([]byte(fun.Results[0].ParamType)) {
			return nil, errors.New(fmt.Sprintf("extend function %s does not return a table", funcName))
		}

		allFuncDefs = append(allFuncDefs, fun)
	}

	return allFuncDefs, nil
}
