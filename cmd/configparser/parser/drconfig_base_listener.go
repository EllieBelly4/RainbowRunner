// Code generated from C:/Users/Sophie/go/src/RainbowRunner/scripts/config-language\DRConfig.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // DRConfig

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDRConfigListener is a complete listener for a parse tree produced by DRConfigParser.
type BaseDRConfigListener struct{}

var _ DRConfigListener = &BaseDRConfigListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDRConfigListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDRConfigListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDRConfigListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDRConfigListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseDRConfigListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseDRConfigListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterClassIdentifier is called when production classIdentifier is entered.
func (s *BaseDRConfigListener) EnterClassIdentifier(ctx *ClassIdentifierContext) {}

// ExitClassIdentifier is called when production classIdentifier is exited.
func (s *BaseDRConfigListener) ExitClassIdentifier(ctx *ClassIdentifierContext) {}

// EnterParentClass is called when production parentClass is entered.
func (s *BaseDRConfigListener) EnterParentClass(ctx *ParentClassContext) {}

// ExitParentClass is called when production parentClass is exited.
func (s *BaseDRConfigListener) ExitParentClass(ctx *ParentClassContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseDRConfigListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseDRConfigListener) ExitProperty(ctx *PropertyContext) {}
