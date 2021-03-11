package src

import (
	"github.com/abusizhishen/ruleEngine/parser"
)

func (v *RuleEngineVisitor) VisitGetMapOrArrayValue(ctx *parser.GetMapOrArrayValueContext) interface{} {
	//fmt.Println("VisitGetArrayValue:", ctx.GetText())
	//
	//ctx.Identify().Accept(v)
	//for _, key := range ctx.AllKey() {
	//	switch key.(type) {
	//	case parser.INumContext:
	//		v.VisitArray(key.(parser.INumContext))
	//	case parser.IStringValueContext:
	//		v.VisitIStringValue(key.(parser.IStringValueContext))
	//	}
	//	//value := v.pop()
	//	//key.Accept(v)
	//	//valueType := v.pop()
	//	//switch valueType.(type) {
	//	//case float64:
	//	//	v.VisitArray()
	//	//case string:
	//	//	v.VisitMap()
	//	//}
	//}

	return nil
}
