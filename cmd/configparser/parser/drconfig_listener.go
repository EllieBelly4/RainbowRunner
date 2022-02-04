// Code generated from C:/Users/Sophie/go/src/RainbowRunner/antlr\DRConfig.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // DRConfig
import "github.com/antlr/antlr4/runtime/Go/antlr"

// DRConfigListener is a complete listener for a parse tree produced by DRConfigParser.
type DRConfigListener interface {
	antlr.ParseTreeListener

	// EnterClassDef is called when entering the classDef production.
	EnterClassDef(c *ClassDefContext)

	// EnterClassEnter is called when entering the classEnter production.
	EnterClassEnter(c *ClassEnterContext)

	// EnterClassLeave is called when entering the classLeave production.
	EnterClassLeave(c *ClassLeaveContext)

	// EnterClassIdentifier is called when entering the classIdentifier production.
	EnterClassIdentifier(c *ClassIdentifierContext)

	// EnterParentClass is called when entering the parentClass production.
	EnterParentClass(c *ParentClassContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterPropertyValue is called when entering the propertyValue production.
	EnterPropertyValue(c *PropertyValueContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// ExitClassDef is called when exiting the classDef production.
	ExitClassDef(c *ClassDefContext)

	// ExitClassEnter is called when exiting the classEnter production.
	ExitClassEnter(c *ClassEnterContext)

	// ExitClassLeave is called when exiting the classLeave production.
	ExitClassLeave(c *ClassLeaveContext)

	// ExitClassIdentifier is called when exiting the classIdentifier production.
	ExitClassIdentifier(c *ClassIdentifierContext)

	// ExitParentClass is called when exiting the parentClass production.
	ExitParentClass(c *ParentClassContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitPropertyValue is called when exiting the propertyValue production.
	ExitPropertyValue(c *PropertyValueContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)
}
