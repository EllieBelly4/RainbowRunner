package main

import (
	"RainbowRunner/cmd/configparser/parser"
	"RainbowRunner/pkg/datatypes"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strings"
)

var baseMaps = make(map[string]interface{})

type TreeShapeListener struct {
	*parser.BaseDRConfigListener
	filePath         string
	fileClassName    string
	currentMap       map[string]interface{}
	currentClass     string
	currentHierarchy *datatypes.StringStack
	File             map[string]interface{}
}

func NewDRConfigListener(filePath string) *TreeShapeListener {
	split := strings.Split(filePath, "\\")
	fileName := strings.Split(split[len(split)-1], ".")[0]

	return &TreeShapeListener{
		filePath:         filePath,
		fileClassName:    fileName,
		currentMap:       make(map[string]interface{}),
		currentHierarchy: datatypes.NewStringStack(),
		File:             make(map[string]interface{}),
	}
}

func (t *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classIdentifier {
		className := ctx.GetText()

		if className == "*" {
			if t.currentHierarchy.Index > 0 {
				className = t.currentHierarchy.Stack[t.currentHierarchy.Index-1]
			} else {
				className = t.fileClassName
			}
		}

		t.currentClass = className

		t.currentHierarchy.Push(className)

		if _, ok := t.File[className]; !ok {
			t.File[className] = make(map[string]interface{})
		}

		t.currentMap = t.File[className].(map[string]interface{})

		if className != t.fileClassName {
			//fmt.Println(strings.Join(t.currentHierarchy.Stack[:t.currentHierarchy.Index], "."))
		}
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_parentClass {
		fqParent := ctx.GetText()

		split := strings.Split(fqParent, ".")
		parentFile := split[0]
		parentClass := parentFile

		if len(split) > 1 {
			parentClass = split[1]
		}

		if _, ok := baseMaps[parentFile]; !ok {
			parentFilePath := "D:\\Work\\dungeon-runners\\666 dumps\\" + parentFile + ".txt"
			if finfo, err := os.Stat(parentFilePath); err == nil && !finfo.IsDir() {
				parentFileFile := parseFile(parentFilePath)
				baseMaps[parentFile] = parentFileFile
			} else {
				t.File[t.currentClass].(map[string]interface{})[parentFile] = make(map[string]interface{})
			}
		}

		switch bm := baseMaps[parentFile].(type) {
		case map[string]interface{}:
			if parentFile == parentClass {
				t.File[t.currentClass].(map[string]interface{})[parentClass] = bm
				//fmt.Printf("%x\n",bm)
			} else {
				t.File[t.currentClass].(map[string]interface{})[parentClass] = bm[parentClass]
				//fmt.Printf("%x\n",bm[parentClass])
			}
		}
	}
}

func (t *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classDef {
		t.currentHierarchy.Pop()
	}
}

//func (this *TreeShapeListener) EnterClassDef(ctx *parser.ClassDefContext) {
//}
