package main

import (
	"RainbowRunner/internal/gosucks"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type FieldDef struct {
	Name  string
	field *ast.Field
}

type FuncDef struct {
	Name     string
	funcType *ast.FuncType
}

type StructDef struct {
	Name       string
	structType *ast.StructType
	Fields     []*FieldDef
	Methods    []*FuncDef
}

func main() {
	structs := make(map[string]*StructDef)

	err := getAllStructDefinitions(structs)

	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(structs, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func getAllStructDefinitions(structs map[string]*StructDef) error {
	// Parse the package using the Go parser
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "internal/objects", nil, parser.AllErrors)
	if err != nil {
		return err
	}

	// Iterate over the packages and their files
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
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
									fmt.Printf("%s=========\n", typeSpec.Name)

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

											structs[structName].Fields = append(structs[structName].Fields, &FieldDef{
												Name:  field.Names[0].String(),
												field: field,
											})

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
					if funcDecl.Recv == nil {
						continue
					}

					structName := ""

					if starExpr, ok := funcDecl.Recv.List[0].Type.(*ast.StarExpr); ok {
						structName = starExpr.X.(*ast.Ident).Name
					} else {
						continue
					}

					if _, ok := structs[structName]; !ok {
						structs[structName] = &StructDef{
							Name: structName,
						}
					}

					if structs[structName].Methods == nil {
						structs[structName].Methods = make([]*FuncDef, 0)
					}

					structs[structName].Methods = append(structs[structName].Methods, &FuncDef{
						Name:     funcDecl.Name.String(),
						funcType: funcDecl.Type,
					})

					gosucks.VAR(funcDecl)
				}
			}
		}
	}
	return nil
}
