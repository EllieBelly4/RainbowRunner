package main

import (
	"RainbowRunner/internal/gosucks"
	"encoding/json"
	"flag"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"strings"
)

var (
	typeName = flag.String("type", "", "Comma separated list of types to generate lua wrappers for")
	extends  = flag.String("extends", "", "Comma separated list of types to extend")
)

var (
	fileStructs = make(map[string]*StructDef)
	imports     = make(map[string]*ImportDef)
	allStructs  = make(map[string]*StructDef)
	funcDefs    = make(map[string]*FuncDef)
	currentPkg  *packages.Package
)

// TODO refactor to allow other packages to be generated
func main() {
	flag.Parse()

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

	currentPkg = pkg[0]

	addAllImports(imports, currentPkg)

	err = parseFileStructDefinitionsFromString(currentPkg, fileStructs, filePath)

	if err != nil {
		panic(err)
	}

	addAllMemberFunctions(fileStructs, funcDefs, currentPkg)

	err = getAllStructDefinitions(allStructs, cwd)

	if err != nil {
		panic(err)
	}

	extendFuncs, err := getExtendFuncs(splitExtends, funcDefs)

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
