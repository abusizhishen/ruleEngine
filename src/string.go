package src

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
)

func (v *RuleEngineVisitor) VisitSTRING(ctx *parser.STRINGContext) interface{} {
	fmt.Println("VisitSTRING", ctx.GetText())
	v.push(ctx.Str().GetText())
	return nil
}

func (v *RuleEngineVisitor) VisitStringValue(ctx *parser.StringValueContext) interface{} {

	return nil
}
