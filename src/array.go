package src

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
)

func (v *RuleEngineVisitor) VisitArray(ctx *parser.ArrayContext) interface{} {
	fmt.Println("VisitArray:", ctx.GetText())

	var array []interface{}
	for _, valueType := range ctx.AllValueType() {
		valueType.Accept(v)
		value := v.pop()
		array = append(array, value)
	}

	v.push(array)

	return nil
}
