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
	"regexp"
	"strings"
	"text/template"
)

var (
	typeName = flag.String("type", "", "Comma separated list of types to generate lua wrappers for")
	extends  = flag.String("extends", "", "Comma separated list of types to extend")
)

func main() {
	flag.Parse()
	fileStructs := make(map[string]*StructDef)
	imports := make(map[string]*ImportDef)
	allStructs := make(map[string]*StructDef)

	splitExtends := strings.Split(*extends, ",")

	if splitExtends[0] == "" {
		splitExtends = nil
	}

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

	addAllImports(imports, pkg[0])

	err = parseFileStructDefinitionsFromString(pkg[0], fileStructs, filePath)

	if err != nil {
		panic(err)
	}

	addAllMemberFunctions(fileStructs, typeDefs, pkg[0])

	err = getAllStructDefinitions(allStructs, cwd)

	if err != nil {
		panic(err)
	}

	extendFuncs, err := getExtendFuncs(splitExtends, typeDefs)

	if err != nil {
		panic(err)
	}

	//extendStructs, err := getExtendStructs(splitExtends, allStructs)
	//
	//if err != nil {
	//	panic(err)
	//}

	typeNames := strings.Split(*typeName, ",")

	err = executeGenerate(extendFuncs, imports, fileStructs, typeNames, cwd)

	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(fileStructs, "", "  ")

	if err != nil {
		panic(err)
	}

	gosucks.VAR(data)
	//fmt.Println(string(data))
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

func executeGenerate(splitExtends []*FuncDef, imports map[string]*ImportDef, structs map[string]*StructDef, typeNames []string, cwd string) error {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))

	for _, name := range typeNames {
		if _, ok := structs[name]; !ok {
			return errors.New(fmt.Sprintf("could not find type %s in file", name))
		}

		data, err := generateWrapper(splitExtends, imports, structs[name])

		if err != nil {
			return err
		}

		data = formatScript(data)

		outputFile := filepath.Join(cwd, fmt.Sprintf("lua_generated_%s.go", strings.ToLower(name)))

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

func generateWrapper(extends []*FuncDef, imports map[string]*ImportDef, def *StructDef) ([]byte, error) {
	t := template.New("wrapper")

	t = t.Funcs(templateFuncMap)

	t, err := t.Parse(templateString)

	if err != nil {
		return nil, err
	}

	requiredImports := def.GetRequiredImports(imports)

	data := &TemplateData{
		Struct:  def,
		Imports: requiredImports,
		Extends: extends,
	}

	buf := &bytes.Buffer{}

	err = t.Execute(buf, data)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
