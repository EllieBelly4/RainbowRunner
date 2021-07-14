// Code generated from C:/Users/Sophie/go/src/RainbowRunner/scripts/config-language\DRConfig.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // DRConfig

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 50, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 5, 2, 12, 10, 2, 3,
	2, 3, 2, 3, 2, 5, 2, 17, 10, 2, 3, 2, 3, 2, 3, 2, 7, 2, 22, 10, 2, 12,
	2, 14, 2, 25, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 5, 5, 36, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5,
	5, 46, 10, 5, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 3, 4, 2, 16, 16,
	22, 22, 2, 55, 2, 11, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 30, 3, 2, 2, 2,
	8, 32, 3, 2, 2, 2, 10, 12, 7, 7, 2, 2, 11, 10, 3, 2, 2, 2, 11, 12, 3, 2,
	2, 2, 12, 13, 3, 2, 2, 2, 13, 16, 5, 4, 3, 2, 14, 15, 7, 6, 2, 2, 15, 17,
	5, 6, 4, 2, 16, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2,
	18, 23, 7, 9, 2, 2, 19, 22, 5, 2, 2, 2, 20, 22, 5, 8, 5, 2, 21, 19, 3,
	2, 2, 2, 21, 20, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23,
	24, 3, 2, 2, 2, 24, 26, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 26, 27, 7, 10,
	2, 2, 27, 3, 3, 2, 2, 2, 28, 29, 9, 2, 2, 2, 29, 5, 3, 2, 2, 2, 30, 31,
	7, 22, 2, 2, 31, 7, 3, 2, 2, 2, 32, 33, 7, 22, 2, 2, 33, 45, 7, 8, 2, 2,
	34, 36, 7, 11, 2, 2, 35, 34, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 37, 3,
	2, 2, 2, 37, 46, 7, 22, 2, 2, 38, 39, 7, 22, 2, 2, 39, 40, 7, 15, 2, 2,
	40, 46, 7, 22, 2, 2, 41, 46, 7, 23, 2, 2, 42, 46, 7, 20, 2, 2, 43, 46,
	7, 21, 2, 2, 44, 46, 7, 5, 2, 2, 45, 35, 3, 2, 2, 2, 45, 38, 3, 2, 2, 2,
	45, 41, 3, 2, 2, 2, 45, 42, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 44, 3,
	2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 48, 7, 12, 2, 2, 48, 9, 3, 2, 2, 2, 8,
	11, 16, 21, 23, 35, 45,
}
var literalNames = []string{
	"", "", "", "", "'extends'", "'static'", "'='", "'{'", "'}'", "'!'", "';'",
	"'.'", "','", "':'", "'*'", "'//'", "", "'\n'",
}
var symbolicNames = []string{
	"", "COMMENT", "MLCOMMENT", "VECTOR3", "EXTENDS", "STATIC", "ASSIGN", "PARENL",
	"PARENR", "EXCL", "SEMI", "DOT", "COMMA", "COLON", "ASTERISK", "SLASHSLASH",
	"WS", "EOL", "SINGLESTR", "DOUBLESTR", "IDENTIFIER", "NUMBER", "ANY",
}

var ruleNames = []string{
	"classDef", "classIdentifier", "parentClass", "property",
}

type DRConfigParser struct {
	*antlr.BaseParser
}

// NewDRConfigParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DRConfigParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDRConfigParser(input antlr.TokenStream) *DRConfigParser {
	this := new(DRConfigParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DRConfig.g4"

	return this
}

// DRConfigParser tokens.
const (
	DRConfigParserEOF        = antlr.TokenEOF
	DRConfigParserCOMMENT    = 1
	DRConfigParserMLCOMMENT  = 2
	DRConfigParserVECTOR3    = 3
	DRConfigParserEXTENDS    = 4
	DRConfigParserSTATIC     = 5
	DRConfigParserASSIGN     = 6
	DRConfigParserPARENL     = 7
	DRConfigParserPARENR     = 8
	DRConfigParserEXCL       = 9
	DRConfigParserSEMI       = 10
	DRConfigParserDOT        = 11
	DRConfigParserCOMMA      = 12
	DRConfigParserCOLON      = 13
	DRConfigParserASTERISK   = 14
	DRConfigParserSLASHSLASH = 15
	DRConfigParserWS         = 16
	DRConfigParserEOL        = 17
	DRConfigParserSINGLESTR  = 18
	DRConfigParserDOUBLESTR  = 19
	DRConfigParserIDENTIFIER = 20
	DRConfigParserNUMBER     = 21
	DRConfigParserANY        = 22
)

// DRConfigParser rules.
const (
	DRConfigParserRULE_classDef        = 0
	DRConfigParserRULE_classIdentifier = 1
	DRConfigParserRULE_parentClass     = 2
	DRConfigParserRULE_property        = 3
)

// IClassDefContext is an interface to support dynamic dispatch.
type IClassDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDefContext differentiates from other interfaces.
	IsClassDefContext()
}

type ClassDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDefContext() *ClassDefContext {
	var p = new(ClassDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DRConfigParserRULE_classDef
	return p
}

func (*ClassDefContext) IsClassDefContext() {}

func NewClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefContext {
	var p = new(ClassDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DRConfigParserRULE_classDef

	return p
}

func (s *ClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefContext) PARENL() antlr.TerminalNode {
	return s.GetToken(DRConfigParserPARENL, 0)
}

func (s *ClassDefContext) PARENR() antlr.TerminalNode {
	return s.GetToken(DRConfigParserPARENR, 0)
}

func (s *ClassDefContext) ClassIdentifier() IClassIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassIdentifierContext)
}

func (s *ClassDefContext) STATIC() antlr.TerminalNode {
	return s.GetToken(DRConfigParserSTATIC, 0)
}

func (s *ClassDefContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(DRConfigParserEXTENDS, 0)
}

func (s *ClassDefContext) ParentClass() IParentClassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParentClassContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParentClassContext)
}

func (s *ClassDefContext) AllClassDef() []IClassDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassDefContext)(nil)).Elem())
	var tst = make([]IClassDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassDefContext)
		}
	}

	return tst
}

func (s *ClassDefContext) ClassDef(i int) IClassDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassDefContext)
}

func (s *ClassDefContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *ClassDefContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *ClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.EnterClassDef(s)
	}
}

func (s *ClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.ExitClassDef(s)
	}
}

func (s *ClassDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DRConfigVisitor:
		return t.VisitClassDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DRConfigParser) ClassDef() (localctx IClassDefContext) {
	localctx = NewClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DRConfigParserRULE_classDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DRConfigParserSTATIC {
		{
			p.SetState(8)
			p.Match(DRConfigParserSTATIC)
		}

	}

	{
		p.SetState(11)
		p.ClassIdentifier()
	}

	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DRConfigParserEXTENDS {
		{
			p.SetState(12)
			p.Match(DRConfigParserEXTENDS)
		}
		{
			p.SetState(13)
			p.ParentClass()
		}

	}
	{
		p.SetState(16)
		p.Match(DRConfigParserPARENL)
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DRConfigParserSTATIC)|(1<<DRConfigParserASTERISK)|(1<<DRConfigParserIDENTIFIER))) != 0 {
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(17)
				p.ClassDef()
			}

		case 2:
			{
				p.SetState(18)
				p.Property()
			}

		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(24)
		p.Match(DRConfigParserPARENR)
	}

	return localctx
}

// IClassIdentifierContext is an interface to support dynamic dispatch.
type IClassIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassIdentifierContext differentiates from other interfaces.
	IsClassIdentifierContext()
}

type ClassIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassIdentifierContext() *ClassIdentifierContext {
	var p = new(ClassIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DRConfigParserRULE_classIdentifier
	return p
}

func (*ClassIdentifierContext) IsClassIdentifierContext() {}

func NewClassIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassIdentifierContext {
	var p = new(ClassIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DRConfigParserRULE_classIdentifier

	return p
}

func (s *ClassIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DRConfigParserIDENTIFIER, 0)
}

func (s *ClassIdentifierContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(DRConfigParserASTERISK, 0)
}

func (s *ClassIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.EnterClassIdentifier(s)
	}
}

func (s *ClassIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.ExitClassIdentifier(s)
	}
}

func (s *ClassIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DRConfigVisitor:
		return t.VisitClassIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DRConfigParser) ClassIdentifier() (localctx IClassIdentifierContext) {
	localctx = NewClassIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DRConfigParserRULE_classIdentifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DRConfigParserASTERISK || _la == DRConfigParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IParentClassContext is an interface to support dynamic dispatch.
type IParentClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParentClassContext differentiates from other interfaces.
	IsParentClassContext()
}

type ParentClassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParentClassContext() *ParentClassContext {
	var p = new(ParentClassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DRConfigParserRULE_parentClass
	return p
}

func (*ParentClassContext) IsParentClassContext() {}

func NewParentClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParentClassContext {
	var p = new(ParentClassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DRConfigParserRULE_parentClass

	return p
}

func (s *ParentClassContext) GetParser() antlr.Parser { return s.parser }

func (s *ParentClassContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DRConfigParserIDENTIFIER, 0)
}

func (s *ParentClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParentClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.EnterParentClass(s)
	}
}

func (s *ParentClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.ExitParentClass(s)
	}
}

func (s *ParentClassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DRConfigVisitor:
		return t.VisitParentClass(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DRConfigParser) ParentClass() (localctx IParentClassContext) {
	localctx = NewParentClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DRConfigParserRULE_parentClass)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Match(DRConfigParserIDENTIFIER)
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DRConfigParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DRConfigParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(DRConfigParserIDENTIFIER)
}

func (s *PropertyContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(DRConfigParserIDENTIFIER, i)
}

func (s *PropertyContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(DRConfigParserASSIGN, 0)
}

func (s *PropertyContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DRConfigParserSEMI, 0)
}

func (s *PropertyContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DRConfigParserNUMBER, 0)
}

func (s *PropertyContext) SINGLESTR() antlr.TerminalNode {
	return s.GetToken(DRConfigParserSINGLESTR, 0)
}

func (s *PropertyContext) DOUBLESTR() antlr.TerminalNode {
	return s.GetToken(DRConfigParserDOUBLESTR, 0)
}

func (s *PropertyContext) VECTOR3() antlr.TerminalNode {
	return s.GetToken(DRConfigParserVECTOR3, 0)
}

func (s *PropertyContext) COLON() antlr.TerminalNode {
	return s.GetToken(DRConfigParserCOLON, 0)
}

func (s *PropertyContext) EXCL() antlr.TerminalNode {
	return s.GetToken(DRConfigParserEXCL, 0)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DRConfigVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DRConfigParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DRConfigParserRULE_property)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(DRConfigParserIDENTIFIER)
	}
	{
		p.SetState(31)
		p.Match(DRConfigParserASSIGN)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DRConfigParserEXCL {
			{
				p.SetState(32)
				p.Match(DRConfigParserEXCL)
			}

		}
		{
			p.SetState(35)
			p.Match(DRConfigParserIDENTIFIER)
		}

	case 2:
		{
			p.SetState(36)
			p.Match(DRConfigParserIDENTIFIER)
		}
		{
			p.SetState(37)
			p.Match(DRConfigParserCOLON)
		}
		{
			p.SetState(38)
			p.Match(DRConfigParserIDENTIFIER)
		}

	case 3:
		{
			p.SetState(39)
			p.Match(DRConfigParserNUMBER)
		}

	case 4:
		{
			p.SetState(40)
			p.Match(DRConfigParserSINGLESTR)
		}

	case 5:
		{
			p.SetState(41)
			p.Match(DRConfigParserDOUBLESTR)
		}

	case 6:
		{
			p.SetState(42)
			p.Match(DRConfigParserVECTOR3)
		}

	}
	{
		p.SetState(45)
		p.Match(DRConfigParserSEMI)
	}

	return localctx
}
