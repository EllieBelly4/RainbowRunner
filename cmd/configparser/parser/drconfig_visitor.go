// Code generated from C:/Users/Sophie/go/src/RainbowRunner/antlr\DRConfig.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // DRConfig
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DRConfigParser.
type DRConfigVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DRConfigParser#classDef.
	VisitClassDef(ctx *ClassDefContext) interface{}

	// Visit a parse tree produced by DRConfigParser#classEnter.
	VisitClassEnter(ctx *ClassEnterContext) interface{}

	// Visit a parse tree produced by DRConfigParser#classLeave.
	VisitClassLeave(ctx *ClassLeaveContext) interface{}

	// Visit a parse tree produced by DRConfigParser#classIdentifier.
	VisitClassIdentifier(ctx *ClassIdentifierContext) interface{}

	// Visit a parse tree produced by DRConfigParser#parentClass.
	VisitParentClass(ctx *ParentClassContext) interface{}

	// Visit a parse tree produced by DRConfigParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by DRConfigParser#propertyValue.
	VisitPropertyValue(ctx *PropertyValueContext) interface{}

	// Visit a parse tree produced by DRConfigParser#propertyName.
	VisitPropertyName(ctx *PropertyNameContext) interface{}
}
