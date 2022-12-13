package main

import (
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/scripts/common"
	"flag"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"strings"
	template2 "text/template"
)

func main() {
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//err := getAllStructDefinitions(structs)
	pkg, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesSizes | packages.NeedTypesInfo | packages.NeedSyntax,
	}, cwd)

	if err != nil {
		panic(err)
	}

	gosucks.VAR(pkg)

	registerFuncs, err := getAllRegisterFuncs(pkg[0])

	template := template2.New("registerLuaFunctions")

	template, err = template.Parse(registerTemplate)

	if err != nil {
		panic(err)
	}

	var s strings.Builder

	err = template.Execute(&s, registerFuncs)

	if err != nil {
		panic(err)
	}

	data := common.FormatScript([]byte(s.String()))

	outFilePath := filepath.Join(cwd, "lua_generated__register.go")

	err = os.WriteFile(outFilePath, data, 0644)

	if err != nil {
		panic(err)
	}
}

// getAllRegisterFuncs returns a list of all functions that have a prefix of registerLua
func getAllRegisterFuncs(p *packages.Package) ([]string, error) {
	var registerFuncs []string

	for _, f := range p.Syntax {
		for _, d := range f.Decls {
			switch decl := d.(type) {
			case *ast.FuncDecl:
				if strings.HasPrefix(decl.Name.Name, "registerLua") {
					registerFuncs = append(registerFuncs, decl.Name.Name)
				}
			}
		}
	}

	return registerFuncs, nil
}
