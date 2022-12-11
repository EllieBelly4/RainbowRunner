package main

import (
	"RainbowRunner/internal/gosucks"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	typeName = flag.String("type", "", "Comma separated list of types to generate lua wrappers for")
)

func main() {
	flag.Parse()
	structs := make(map[string]*StructDef)
	imports := make(map[string]bool)

	fileName := os.Getenv("GOFILE")
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(cwd, fileName)

	//err := getAllStructDefinitions(structs)
	pkg, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedTypesInfo | packages.NeedSyntax,
	}, cwd)

	if err != nil {
		panic(err)
	}

	err = parseFileStructDefinitionsFromString(pkg[0], structs, filePath)

	if err != nil {
		panic(err)
	}

	addAllMemberFunctions(structs, typeDefs, pkg[0])
	addAllImports(imports, pkg[0])

	typeNames := strings.Split(*typeName, ",")

	err = executeGenerate(imports, structs, typeNames, cwd)

	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(structs, "", "  ")

	if err != nil {
		panic(err)
	}

	gosucks.VAR(data)
	//fmt.Println(string(data))
}

func addAllImports(imports map[string]bool, p *packages.Package) {
	for _, file := range p.Syntax {
		for _, i := range file.Imports {
			imports[i.Path.Value] = true
		}
	}
}

func executeGenerate(imports map[string]bool, structs map[string]*StructDef, typeNames []string, cwd string) error {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))

	for _, name := range typeNames {
		if _, ok := structs[name]; !ok {
			return errors.New(fmt.Sprintf("could not find type %s in file", name))
		}

		data, err := generateWrapper(imports, structs[name])

		if err != nil {
			return err
		}

		data = formatScript(data)

		outputFile := filepath.Join(cwd, fmt.Sprintf("lua_%s_generated.go", strings.ToLower(name)))

		err = os.WriteFile(outputFile, data, 0755)

		if err != nil {
			return err
		}
	}

	return nil
}

// formatScript formats a byte array of golang using go/format
func formatScript(data []byte) []byte {
	data, err := format.Source(data)

	if err != nil {
		panic(err)
	}

	return data
}

func generateWrapper(imports map[string]bool, def *StructDef) ([]byte, error) {
	t := template.New("wrapper")

	t = t.Funcs(templateFuncMap)

	t, err := t.Parse(templateString)

	if err != nil {
		return nil, err
	}

	importStrings := make([]string, 0)

	for k, _ := range imports {
		importStrings = append(importStrings, k)
	}

	data := &TemplateData{
		Struct:  def,
		Imports: importStrings,
	}

	buf := &bytes.Buffer{}

	err = t.Execute(buf, data)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
