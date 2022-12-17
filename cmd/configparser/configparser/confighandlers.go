package configparser

import (
	"RainbowRunner/cmd/configparser/parser"
	"RainbowRunner/internal/types/configtypes"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strings"
)

var baseMaps = make(map[string]*configtypes.DRClass)

type DRConfigParser struct {
	*parser.BaseDRConfigListener

	filePath      string
	fileClassName string

	classStack          *DRClassStack
	currentClass        string
	currentPropertyName string
	basePath            string

	depth          int
	IsGenericClass bool
	gcBaseType     string
	drConfig       *configtypes.DRConfig
}

func NewDRConfigListener(filePath string, rootPath string, config *configtypes.DRConfig) *DRConfigParser {
	split := strings.Split(filePath, "\\")
	fileName := strings.Split(split[len(split)-1], ".")[0]
	extensionlessPath := strings.SplitN(filePath, ".", 2)[0]
	gcBaseType := strings.ReplaceAll(strings.Replace(extensionlessPath, rootPath+"\\", "", 1), "\\", ".")

	splitBaseType := strings.Split(gcBaseType, ".")
	curMap := config.Classes

	for i := 0; i < len(splitBaseType)-1; i++ {
		gcTypeNameLowercase := strings.ToLower(splitBaseType[i])

		if _, ok := curMap.Children[gcTypeNameLowercase]; !ok {
			curMap.Children[gcTypeNameLowercase] = configtypes.NewDRClassChildGroup("")
			curMap.Children[gcTypeNameLowercase].Entities = make([]*configtypes.DRClass, 0)
			curMap.Children[gcTypeNameLowercase].Entities = append(curMap.Children[gcTypeNameLowercase].Entities, configtypes.NewDRClass(""))
		}

		curMap = curMap.Children[gcTypeNameLowercase].Entities[0]
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
	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classEnter {
		t.IsGenericClass = false
		t.depth++

		currentClassName := strings.ToLower(t.classStack.Current().Name)
		currentClassParent := t.classStack.Previous()

		parentChild, ok := currentClassParent.Children[currentClassName]

		if !ok {
			currentClassParent.Children[currentClassName] = configtypes.NewDRClassChildGroup("")
			parentChild = currentClassParent.Children[currentClassName]
		}

		parentChild.Entities =
			append(
				parentChild.Entities,
				t.classStack.Current(),
			)
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_parentClass {
		parent := ctx.GetText()

		if parent != "" {
			t.classStack.Current().Extends = parent
		}
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classLeave {
		t.classStack.Pop()

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

		if className == "*" {
			t.IsGenericClass = true
			if t.depth == 0 {
				className = t.fileClassName
			} else {
				className = "UNKNOWN"
			}
		}

		newClass := configtypes.NewDRClass(className)

		t.classStack.Push(newClass)
	}

	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_parentClass {
		currentClass := t.classStack.Current()
		parentClass := ctx.GetText()
		parentClass = strings.ToLower(parentClass)

		if t.IsGenericClass && currentClass.Name == "UNKNOWN" {
			currentClass.Name = parentClass
		}
	}
}

func (t *DRConfigParser) ExitEveryRule(ctx antlr.ParserRuleContext) {
	if ctx.GetRuleIndex() == parser.DRConfigParserRULE_classDef {
	}
}

func parseFile(path string, rootPath string, config *configtypes.DRConfig) {
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

func ParseAllFilesToDRConfig(files []string, rootPath string) (*configtypes.DRConfig, error) {
	drConfig := configtypes.NewDRConfig()

	for _, path := range files {
		_, err := os.Stat(path)

		if err != nil {
			return nil, err
		}

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

func NewErrorListener(path string) *ErrorListener {
	originalListener := antlr.NewDiagnosticErrorListener(true)

	listener := &ErrorListener{
		DiagnosticErrorListener: originalListener,
		Path:                    path,
	}

	return listener
}
