package main

import (
	"RainbowRunner/cmd/configparser/parser"
	"RainbowRunner/internal/database"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"
)

var baseMaps = make(map[string]*database.DRClass)

type DRArmourParser struct {
	*parser.BaseDRConfigListener
	DRClass *database.DRClass

	filePath      string
	fileClassName string

	classStack          *DRClassStack
	currentClass        string
	currentPropertyName string
	basePath            string

	depth          int
	IsGenericClass bool
}

func NewDRConfigListener(filePath string) *DRArmourParser {
	split := strings.Split(filePath, "\\")
	fileName := strings.Split(split[len(split)-1], ".")[0]

	return &DRArmourParser{
		filePath:      filePath,
		fileClassName: fileName,
		DRClass:       nil,
		classStack:    NewDRClassStack(),
		basePath:      filepath.Dir(filePath),
	}
}

func (t *DRArmourParser) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classEnter {
		t.depth++
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classLeave {
		current := t.classStack.Pop()
		t.DRClass = t.classStack.Current()

		if t.DRClass != current {
			mergeChildInto(t.DRClass, current, true)
		}

		t.depth--
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_propertyName {
		t.currentPropertyName = ctx.GetText()
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_propertyValue {
		t.DRClass.Properties[t.currentPropertyName] = ctx.GetText()
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classIdentifier {
		className := ctx.GetText()

		if className == "*" {
			t.IsGenericClass = true
		} else {
			t.pushNewClass(className)
		}

		//t.DRClass.Children = append(t.DRClass.Children, t.DRClass)

		//file := make(map[string]interface{})
		//
		////if _, ok := t.File[className]; !ok {
		////	t.File[className] = file
		////}
		//
		//t.Files.Push(file)
		//
		////t.currentMap = t.File[className].(map[string]interface{})
		//
		//if className != t.fileClassName {
		//	//fmt.Println(strings.Join(t.currentHierarchy.Stack[:t.currentHierarchy.Index], "."))
		//}
		//
		//ci := t.siblingIndex.Current()
		//t.siblingIndex.SetCurrent(ci + 1)
		//t.currentMap["index"] = ci
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_parentClass {
		fqParent := ctx.GetText()

		if t.IsGenericClass {
			t.pushNewClass(fqParent)
		}

		split := strings.SplitN(fqParent, ".", 2)
		parentFile := split[0]
		parentClass := ""

		if len(split) > 1 {
			parentClass = split[1]
		}

		var parentDRClass *database.DRClass

		if _, ok := baseMaps[parentFile]; !ok {
			parentFilePath := t.basePath + "/" + parentFile + ".txt"
			if finfo, err := os.Stat(parentFilePath); err == nil && !finfo.IsDir() {
				baseMaps[parentFile] = parseFile(parentFilePath)
			}
		}

		parentDRClass = baseMaps[parentFile]

		if parentDRClass != nil {
			foundClass := parentDRClass

			if parentClass != "" {
				foundClass = parentDRClass.Find(strings.Split(parentClass, "."))
			}

			if foundClass != nil {
				mergeProperties(foundClass, t.classStack.Current())

				if foundClass.Children != nil {
					for _, child := range foundClass.Children {
						mergeChildInto(t.classStack.Current(), child, false)
					}
				}
			}
		}

		//switch bm := baseMaps[parentFile].(type) {
		//case map[string]interface{}:
		//	if parentFile == parentClass {
		//		t.currentMap[parentClass] = bm
		//		//fmt.Printf("%x\n",bm)
		//	} else {
		//		t.currentMap[parentClass] = bm[parentClass]
		//		//fmt.Printf("%x\n",bm[parentClass])
		//	}
		//}
	}
}

func (t *DRArmourParser) pushNewClass(className string) {
	newClass := database.NewDRClass(className)

	t.classStack.Push(newClass)

	if t.DRClass != nil {
		mergeChildInto(t.DRClass, newClass, true)
		//t.DRClass.Children = append(t.DRClass.Children, newClass)
	} else {
		t.classStack.Push(newClass)
	}

	t.DRClass = newClass
}

func mergeChildInto(base *database.DRClass, newChild *database.DRClass, rightPriority bool) {
	foundChild := base.Find([]string{newChild.Name})

	if foundChild == nil {
		base.Children = append(base.Children, newChild)
	} else {
		if rightPriority {
			mergeProperties(foundChild, newChild)
		} else {
			mergeProperties(newChild, foundChild)
		}
	}
}

func mergeProperties(base *database.DRClass, child *database.DRClass) {
	for propKey, propVal := range child.Properties {
		//if _, currentHasProp := child.Properties[propKey]; !currentHasProp {
		base.Properties[propKey] = propVal
		//}
	}
}

func (t *DRArmourParser) ExitEveryRule(ctx antlr.ParserRuleContext) {
	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classDef {
	}
}

//func (this *TreeShapeListener) EnterClassDef(ctx *parser.ClassDefContext) {
//}
