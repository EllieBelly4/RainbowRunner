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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 65, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 5, 2, 20, 10, 2, 3, 2, 3, 2, 3, 2, 5, 2, 25,
	10, 2, 3, 2, 3, 2, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 5, 8, 51, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 5, 8, 61, 10, 8, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10,
	12, 14, 16, 2, 3, 4, 2, 16, 16, 22, 22, 2, 66, 2, 19, 3, 2, 2, 2, 4, 36,
	3, 2, 2, 2, 6, 38, 3, 2, 2, 2, 8, 40, 3, 2, 2, 2, 10, 42, 3, 2, 2, 2, 12,
	44, 3, 2, 2, 2, 14, 60, 3, 2, 2, 2, 16, 62, 3, 2, 2, 2, 18, 20, 7, 7, 2,
	2, 19, 18, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 24,
	5, 8, 5, 2, 22, 23, 7, 6, 2, 2, 23, 25, 5, 10, 6, 2, 24, 22, 3, 2, 2, 2,
	24, 25, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 31, 5, 4, 3, 2, 27, 30, 5,
	2, 2, 2, 28, 30, 5, 12, 7, 2, 29, 27, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30,
	33, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2,
	2, 33, 31, 3, 2, 2, 2, 34, 35, 5, 6, 4, 2, 35, 3, 3, 2, 2, 2, 36, 37, 7,
	9, 2, 2, 37, 5, 3, 2, 2, 2, 38, 39, 7, 10, 2, 2, 39, 7, 3, 2, 2, 2, 40,
	41, 9, 2, 2, 2, 41, 9, 3, 2, 2, 2, 42, 43, 7, 22, 2, 2, 43, 11, 3, 2, 2,
	2, 44, 45, 5, 16, 9, 2, 45, 46, 7, 8, 2, 2, 46, 47, 5, 14, 8, 2, 47, 48,
	7, 12, 2, 2, 48, 13, 3, 2, 2, 2, 49, 51, 7, 11, 2, 2, 50, 49, 3, 2, 2,
	2, 50, 51, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 61, 7, 22, 2, 2, 53, 54,
	7, 22, 2, 2, 54, 55, 7, 15, 2, 2, 55, 61, 7, 22, 2, 2, 56, 61, 7, 23, 2,
	2, 57, 61, 7, 20, 2, 2, 58, 61, 7, 21, 2, 2, 59, 61, 7, 5, 2, 2, 60, 50,
	3, 2, 2, 2, 60, 53, 3, 2, 2, 2, 60, 56, 3, 2, 2, 2, 60, 57, 3, 2, 2, 2,
	60, 58, 3, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 15, 3, 2, 2, 2, 62, 63, 7,
	22, 2, 2, 63, 17, 3, 2, 2, 2, 8, 19, 24, 29, 31, 50, 60,
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
	"classDef", "classEnter", "classLeave", "classIdentifier", "parentClass",
	"property", "propertyValue", "propertyName",
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
	DRConfigParserRULE_classEnter      = 1
	DRConfigParserRULE_classLeave      = 2
	DRConfigParserRULE_classIdentifier = 3
	DRConfigParserRULE_parentClass     = 4
	DRConfigParserRULE_property        = 5
	DRConfigParserRULE_propertyValue   = 6
	DRConfigParserRULE_propertyName    = 7
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

func (s *ClassDefContext) ClassEnter() IClassEnterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassEnterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassEnterContext)
}

func (s *ClassDefContext) ClassLeave() IClassLeaveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassLeaveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassLeaveContext)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DRConfigParserSTATIC {
		{
			p.SetState(16)
			p.Match(DRConfigParserSTATIC)
		}

	}

	{
		p.SetState(19)
		p.ClassIdentifier()
	}

	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DRConfigParserEXTENDS {
		{
			p.SetState(20)
			p.Match(DRConfigParserEXTENDS)
		}
		{
			p.SetState(21)
			p.ParentClass()
		}

	}
	{
		p.SetState(24)
		p.ClassEnter()
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DRConfigParserSTATIC)|(1<<DRConfigParserASTERISK)|(1<<DRConfigParserIDENTIFIER))) != 0 {
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(25)
				p.ClassDef()
			}

		case 2:
			{
				p.SetState(26)
				p.Property()
			}

		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.ClassLeave()
	}

	return localctx
}

// IClassEnterContext is an interface to support dynamic dispatch.
type IClassEnterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassEnterContext differentiates from other interfaces.
	IsClassEnterContext()
}

type ClassEnterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassEnterContext() *ClassEnterContext {
	var p = new(ClassEnterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DRConfigParserRULE_classEnter
	return p
}

func (*ClassEnterContext) IsClassEnterContext() {}

func NewClassEnterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassEnterContext {
	var p = new(ClassEnterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DRConfigParserRULE_classEnter

	return p
}

func (s *ClassEnterContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassEnterContext) PARENL() antlr.TerminalNode {
	return s.GetToken(DRConfigParserPARENL, 0)
}

func (s *ClassEnterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassEnterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassEnterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.EnterClassEnter(s)
	}
}

func (s *ClassEnterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.ExitClassEnter(s)
	}
}

func (s *ClassEnterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DRConfigVisitor:
		return t.VisitClassEnter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DRConfigParser) ClassEnter() (localctx IClassEnterContext) {
	localctx = NewClassEnterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DRConfigParserRULE_classEnter)

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
		p.SetState(34)
		p.Match(DRConfigParserPARENL)
	}

	return localctx
}

// IClassLeaveContext is an interface to support dynamic dispatch.
type IClassLeaveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassLeaveContext differentiates from other interfaces.
	IsClassLeaveContext()
}

type ClassLeaveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassLeaveContext() *ClassLeaveContext {
	var p = new(ClassLeaveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DRConfigParserRULE_classLeave
	return p
}

func (*ClassLeaveContext) IsClassLeaveContext() {}

func NewClassLeaveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassLeaveContext {
	var p = new(ClassLeaveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DRConfigParserRULE_classLeave

	return p
}

func (s *ClassLeaveContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassLeaveContext) PARENR() antlr.TerminalNode {
	return s.GetToken(DRConfigParserPARENR, 0)
}

func (s *ClassLeaveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassLeaveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassLeaveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.EnterClassLeave(s)
	}
}

func (s *ClassLeaveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.ExitClassLeave(s)
	}
}

func (s *ClassLeaveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DRConfigVisitor:
		return t.VisitClassLeave(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DRConfigParser) ClassLeave() (localctx IClassLeaveContext) {
	localctx = NewClassLeaveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DRConfigParserRULE_classLeave)

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
		p.SetState(36)
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
	p.EnterRule(localctx, 6, DRConfigParserRULE_classIdentifier)
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
		p.SetState(38)
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
	p.EnterRule(localctx, 8, DRConfigParserRULE_parentClass)

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
		p.SetState(40)
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

func (s *PropertyContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *PropertyContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(DRConfigParserASSIGN, 0)
}

func (s *PropertyContext) PropertyValue() IPropertyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyValueContext)
}

func (s *PropertyContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DRConfigParserSEMI, 0)
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
	p.EnterRule(localctx, 10, DRConfigParserRULE_property)

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
		p.SetState(42)
		p.PropertyName()
	}
	{
		p.SetState(43)
		p.Match(DRConfigParserASSIGN)
	}
	{
		p.SetState(44)
		p.PropertyValue()
	}
	{
		p.SetState(45)
		p.Match(DRConfigParserSEMI)
	}

	return localctx
}

// IPropertyValueContext is an interface to support dynamic dispatch.
type IPropertyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyValueContext differentiates from other interfaces.
	IsPropertyValueContext()
}

type PropertyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyValueContext() *PropertyValueContext {
	var p = new(PropertyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DRConfigParserRULE_propertyValue
	return p
}

func (*PropertyValueContext) IsPropertyValueContext() {}

func NewPropertyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyValueContext {
	var p = new(PropertyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DRConfigParserRULE_propertyValue

	return p
}

func (s *PropertyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyValueContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(DRConfigParserIDENTIFIER)
}

func (s *PropertyValueContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(DRConfigParserIDENTIFIER, i)
}

func (s *PropertyValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DRConfigParserNUMBER, 0)
}

func (s *PropertyValueContext) SINGLESTR() antlr.TerminalNode {
	return s.GetToken(DRConfigParserSINGLESTR, 0)
}

func (s *PropertyValueContext) DOUBLESTR() antlr.TerminalNode {
	return s.GetToken(DRConfigParserDOUBLESTR, 0)
}

func (s *PropertyValueContext) VECTOR3() antlr.TerminalNode {
	return s.GetToken(DRConfigParserVECTOR3, 0)
}

func (s *PropertyValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(DRConfigParserCOLON, 0)
}

func (s *PropertyValueContext) EXCL() antlr.TerminalNode {
	return s.GetToken(DRConfigParserEXCL, 0)
}

func (s *PropertyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.EnterPropertyValue(s)
	}
}

func (s *PropertyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.ExitPropertyValue(s)
	}
}

func (s *PropertyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DRConfigVisitor:
		return t.VisitPropertyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DRConfigParser) PropertyValue() (localctx IPropertyValueContext) {
	localctx = NewPropertyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DRConfigParserRULE_propertyValue)
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
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DRConfigParserEXCL {
			{
				p.SetState(47)
				p.Match(DRConfigParserEXCL)
			}

		}
		{
			p.SetState(50)
			p.Match(DRConfigParserIDENTIFIER)
		}

	case 2:
		{
			p.SetState(51)
			p.Match(DRConfigParserIDENTIFIER)
		}
		{
			p.SetState(52)
			p.Match(DRConfigParserCOLON)
		}
		{
			p.SetState(53)
			p.Match(DRConfigParserIDENTIFIER)
		}

	case 3:
		{
			p.SetState(54)
			p.Match(DRConfigParserNUMBER)
		}

	case 4:
		{
			p.SetState(55)
			p.Match(DRConfigParserSINGLESTR)
		}

	case 5:
		{
			p.SetState(56)
			p.Match(DRConfigParserDOUBLESTR)
		}

	case 6:
		{
			p.SetState(57)
			p.Match(DRConfigParserVECTOR3)
		}

	}

	return localctx
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DRConfigParserRULE_propertyName
	return p
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DRConfigParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DRConfigParserIDENTIFIER, 0)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DRConfigListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (s *PropertyNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DRConfigVisitor:
		return t.VisitPropertyName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DRConfigParser) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DRConfigParserRULE_propertyName)

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
		p.SetState(60)
		p.Match(DRConfigParserIDENTIFIER)
	}

	return localctx
}
