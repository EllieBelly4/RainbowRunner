package main

import (
	"RainbowRunner/scripts/common"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

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

		data = common.FormatScript(data)

		outputFile := filepath.Join(cwd, fmt.Sprintf("lua_generated_%s.go", strings.ToLower(name)))

		err = os.WriteFile(outputFile, data, 0755)

		if err != nil {
			return err
		}
	}

	return nil
}

func generateWrapper(extends []*FuncDef, imports map[string]*ImportDef, def *StructDef) ([]byte, error) {
	t := template.New("wrapper")

	t = t.Funcs(mergeFuncMaps(templateFuncMap, typeCheckFunctions))

	t, err := t.Parse(templateString)

	if err != nil {
		return nil, err
	}

	requiredImports := def.GetRequiredImports(imports)

	data := &TemplateData{
		Struct:      def,
		Imports:     requiredImports,
		Extends:     extends,
		PackageName: currentPkg.Name,
	}

	buf := &bytes.Buffer{}

	err = t.Execute(buf, data)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
