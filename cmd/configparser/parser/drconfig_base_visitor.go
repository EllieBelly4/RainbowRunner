// Code generated from C:/Users/Sophie/go/src/RainbowRunner/antlr\DRConfig.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // DRConfig
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDRConfigVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDRConfigVisitor) VisitClassDef(ctx *ClassDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDRConfigVisitor) VisitClassEnter(ctx *ClassEnterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDRConfigVisitor) VisitClassLeave(ctx *ClassLeaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDRConfigVisitor) VisitClassIdentifier(ctx *ClassIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDRConfigVisitor) VisitParentClass(ctx *ParentClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDRConfigVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDRConfigVisitor) VisitPropertyValue(ctx *PropertyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDRConfigVisitor) VisitPropertyName(ctx *PropertyNameContext) interface{} {
	return v.VisitChildren(ctx)
}
