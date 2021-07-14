// Code generated from C:/Users/Sophie/go/src/RainbowRunner/scripts/config-language\DRConfig.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // DRConfig

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DRConfigListener is a complete listener for a parse tree produced by DRConfigParser.
type DRConfigListener interface {
	antlr.ParseTreeListener

	// EnterClassDef is called when entering the classDef production.
	EnterClassDef(c *ClassDefContext)

	// EnterClassIdentifier is called when entering the classIdentifier production.
	EnterClassIdentifier(c *ClassIdentifierContext)

	// EnterParentClass is called when entering the parentClass production.
	EnterParentClass(c *ParentClassContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// ExitClassDef is called when exiting the classDef production.
	ExitClassDef(c *ClassDefContext)

	// ExitClassIdentifier is called when exiting the classIdentifier production.
	ExitClassIdentifier(c *ClassIdentifierContext)

	// ExitParentClass is called when exiting the parentClass production.
	ExitParentClass(c *ParentClassContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)
}
