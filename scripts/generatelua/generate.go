package main

import (
	"RainbowRunner/internal/gosucks"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"
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
	Name        string
	structType  *ast.StructType
	Fields      []*FieldDef
	Methods     []*FuncDef
	Constructor *FuncDef
}

var (
	typeName = flag.String("type", "", "Comma separated list of types to generate lua wrappers for")
)

func main() {
	flag.Parse()
	structs := make(map[string]*StructDef)

	fileName := os.Getenv("GOFILE")
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	extensionlessFileName := strings.Split(fileName, ".go")[0]

	outputFile := filepath.Join(cwd, fmt.Sprintf("lua_%s_generated.go", extensionlessFileName))
	filePath := filepath.Join(cwd, fileName)

	//err := getAllStructDefinitions(structs)
	err = parseFileStructDefinitionsFromString(structs, filePath)

	if err != nil {
		panic(err)
	}

	typeNames := strings.Split(*typeName, ",")

	err = executeGenerate(structs, typeNames, outputFile)

	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(structs, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func executeGenerate(structs map[string]*StructDef, typeNames []string, outputFile string) error {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))

	for _, name := range typeNames {
		if _, ok := structs[name]; !ok {
			return errors.New(fmt.Sprintf("could not find type %s in file", name))
		}

		data, err := generateWrapper(structs[name])

		if err != nil {
			return err
		}

		err = os.WriteFile(outputFile, data, 0755)

		if err != nil {
			return err
		}
	}

	return nil
}

func generateWrapper(def *StructDef) ([]byte, error) {
	t, err := template.New("wrapper").Parse(templateString)

	if err != nil {
		return nil, err
	}

	data := &TemplateData{
		Struct: def,
	}

	buf := &bytes.Buffer{}

	err = t.Execute(buf, data)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
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
			err := parseFileStructDefinitions(structs, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func parseFileStructDefinitionsFromString(structs map[string]*StructDef, file string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseFile(fset, file, nil, parser.AllErrors)
	if err != nil {
		return err
	}

	return parseFileStructDefinitions(structs, pkgs)
}

func parseFileStructDefinitions(structs map[string]*StructDef, file *ast.File) error {
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
			funcName := funcDecl.Name.String()

			funcDef := &FuncDef{
				Name:     funcName,
				funcType: funcDecl.Type,
			}

			if strings.HasPrefix(funcName, "New") {
				structName := strings.Replace(funcName, "New", "", 1)

				if _, ok := structs[structName]; !ok {
					structs[structName] = &StructDef{
						Name: structName,
					}
				}

				structs[structName].Constructor = funcDef
			}

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

			structs[structName].Methods = append(structs[structName].Methods, funcDef)
			gosucks.VAR(funcDecl)
		}
	}
	return nil
}
