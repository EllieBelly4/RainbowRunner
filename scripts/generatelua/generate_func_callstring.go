package main

import (
	"fmt"
	"strings"
	"text/template"
)

func GenerateCallString(def FuncDef, startOffset int) string {
	t := template.New("callStringTemplate")

	t.Funcs(mergeFuncMaps(template.FuncMap{
		"add":                 Add,
		"generateCheckNumber": generateCheckNumber,
		"generateCheckString": generateCheckString,
		"generateCheckBool":   generateCheckBool,
	}, typeCheckFunctions))

	t, err := t.Parse(callStringTemplate)

	if err != nil {
		panic(err)
	}

	var b strings.Builder

	err = t.Execute(&b, struct {
		Func        FuncDef
		StartOffset int
	}{
		Func:        def,
		StartOffset: startOffset,
	})

	if err != nil {
		panic(err)
	}

	return b.String()
}

func generateCheckNumber(param FuncParamDef, index int) string {
	s := fmt.Sprintf("%s(l.CheckNumber(%d))", param.FullTypeString(), index+1)

	if param.IsPointer {
		s = pointerWrapValue(param, s)
	}

	return s
}

func generateCheckString(param FuncParamDef, index int) string {
	s := fmt.Sprintf("%s(l.CheckString(%d))", param.FullTypeString(), index+1)

	if param.IsPointer {
		s = pointerWrapValue(param, s)
	}

	return s
}

func generateCheckBool(param FuncParamDef, index int) string {
	s := fmt.Sprintf("%s(l.CheckBool(%d))", param.FullTypeString(), index+1)

	if param.IsPointer {
		s = pointerWrapValue(param, s)
	}

	return s
}

func pointerWrapValue(param FuncParamDef, s string) string {
	return fmt.Sprintf("func(v %s) *%s {return &v}(%s)", param.FullTypeString(), param.FullTypeString(), s)
}
