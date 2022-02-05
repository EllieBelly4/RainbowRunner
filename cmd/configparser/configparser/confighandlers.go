package configparser

import (
	"RainbowRunner/cmd/configparser/parser"
	"RainbowRunner/internal/database"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

var baseMaps = make(map[string]*database.DRClass)

type DRConfigParser struct {
	*parser.BaseDRConfigListener
	//DRClass *database.DRClass

	filePath      string
	fileClassName string

	classStack          *DRClassStack
	currentClass        string
	currentPropertyName string
	basePath            string

	depth          int
	IsGenericClass bool
	gcBaseType     string
	drConfig       *DRConfig
}

func NewDRConfigListener(filePath string, rootPath string, config *DRConfig) *DRConfigParser {
	split := strings.Split(filePath, "\\")
	fileName := strings.Split(split[len(split)-1], ".")[0]
	extensionlessPath := strings.SplitN(filePath, ".", 2)[0]
	gcBaseType := strings.ReplaceAll(strings.Replace(extensionlessPath, rootPath+"\\", "", 1), "\\", ".")

	splitBaseType := strings.Split(gcBaseType, ".")
	curMap := config.Classes

	for i := 0; i < len(splitBaseType)-1; i++ {
		if _, ok := curMap.Children[splitBaseType[i]]; !ok {
			curMap.Children[splitBaseType[i]] = database.NewDRClassChildGroup("")
			curMap.Children[splitBaseType[i]].Entities = make([]*database.DRClass, 0)
			curMap.Children[splitBaseType[i]].Entities = append(curMap.Children[splitBaseType[i]].Entities, database.NewDRClass(""))
		}

		curMap = curMap.Children[splitBaseType[i]].Entities[0]
	}

	classStack := NewDRClassStack()

	classStack.Push(curMap)

	return &DRConfigParser{
		filePath:      filePath,
		gcBaseType:    gcBaseType,
		fileClassName: fileName,
		drConfig:      config,
		classStack:    classStack,
		basePath:      rootPath,
	}
}

func (t *DRConfigParser) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(t.classStack.Index)

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classEnter {
		t.IsGenericClass = false
		t.depth++
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classLeave {
		t.classStack.Pop()
		//
		//if t.DRClass != current {
		//	mergeChildInto(t.DRClass, current, true)
		//}

		t.depth--
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_propertyName {
		t.currentPropertyName = ctx.GetText()
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_propertyValue {
		t.classStack.Current().Properties[t.currentPropertyName] = ctx.GetText()
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classIdentifier {
		className := ctx.GetText()
		currentClass := t.classStack.Current()

		if className == "*" {
			t.IsGenericClass = true
			className = t.fileClassName

			if t.depth == 0 {
				t.classStack.Push(currentClass)
			}
		}

		if !t.IsGenericClass {
			newClass := database.NewDRClass("")

			if _, ok := currentClass.Children[className]; !ok {
				currentClass.Children[className] = database.NewDRClassChildGroup("")
			}

			currentClass.Children[className].Entities = append(currentClass.Children[className].Entities, newClass)

			t.classStack.Push(newClass)
		}
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_parentClass {
		currentClass := t.classStack.Current()
		parentClass := ctx.GetText()

		if t.IsGenericClass {
			if t.depth == 0 {
				currentClass.Extends = parentClass
			} else {
				splitParentClass := strings.Split(parentClass, ".")
				childParentClass := splitParentClass[len(splitParentClass)-1]

				newClass := database.NewDRClass("*")
				newClass.Extends = parentClass

				if _, ok := currentClass.Children[childParentClass]; !ok {
					currentClass.Children[childParentClass] = database.NewDRClassChildGroup("")
				}

				currentClass.Children[childParentClass].Entities = append(currentClass.Children[childParentClass].Entities, newClass)

				t.classStack.Push(newClass)
			}

		} else {
			currentClass.Extends = parentClass
		}
	}
}

//func mergeChildInto(base *database.DRClass, newChild *database.DRClass, rightPriority bool) {
//	foundChild := base.Find([]string{newChild.Name})
//
//	if foundChild == nil {
//		base.Children = append(base.Children, newChild)
//	} else {
//		if rightPriority {
//			mergeProperties(foundChild, newChild)
//		} else {
//			mergeProperties(newChild, foundChild)
//		}
//	}
//}
//
//func mergeProperties(base *database.DRClass, child *database.DRClass) {
//	for propKey, propVal := range child.Properties {
//		//if _, currentHasProp := child.Properties[propKey]; !currentHasProp {
//		base.Properties[propKey] = propVal
//		//}
//	}
//}

func (t *DRConfigParser) ExitEveryRule(ctx antlr.ParserRuleContext) {
	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classDef {
	}
}

func parseFile(path string, rootPath string, config *DRConfig) {
	input, _ := antlr.NewFileStream(path)
	lexer := parser.NewDRConfigLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewDRConfigParser(stream)
	p.AddErrorListener(NewErrorListener(path))
	p.BuildParseTrees = true
	tree := p.ClassDef()
	listener := NewDRConfigListener(path, rootPath, config)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

//func ParseAllFiles(files []string) ([]*database.DRClass, error) {
//	all := make([]*database.DRClass, 0)
//
//	for _, path := range files {
//		//split := strings.Split(path, "\\")
//		//fileName := strings.Split(split[len(split)-1], ".")[0]
//
//		file := parseFile(path, filepath.Dir(path), nil)
//
//		all = append(all, file)
//	}
//
//	return all, nil
//}

func ParseAllFilesToDRConfig(files []string, rootPath string) (*DRConfig, error) {
	drConfig := NewDRConfig()

	for _, path := range files {
		//split := strings.Split(path, "\\")
		//fileName := strings.Split(split[len(split)-1], ".")[0]

		fmt.Println("Parsing " + path)

		parseFile(path, rootPath, drConfig)
	}

	return drConfig, nil
}

type ErrorListener struct {
	*antlr.DiagnosticErrorListener
	Path string
}

func (d *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Printf("FAILED: %s\n", d.Path)
	d.DiagnosticErrorListener.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

//func (d *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//}
//
//func (d *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//}
//
//func (d *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//}

func NewErrorListener(path string) *ErrorListener {
	originalListener := antlr.NewDiagnosticErrorListener(true)

	listener := &ErrorListener{
		DiagnosticErrorListener: originalListener,
		Path:                    path,
	}

	return listener
}

func listContains(parses []string, path string) bool {
	for _, pars := range parses {
		if pars == path {
			return true
		}
	}

	return false
}

//func (this *TreeShapeListener) EnterClassDef(ctx *parser.ClassDefContext) {
//}
