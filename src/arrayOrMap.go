package src

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
	"reflect"
)

func (v *RuleEngineVisitor) VisitGetMapOrArrayValue(ctx *parser.GetMapOrArrayValueContext) interface{} {
	fmt.Println("VisitGetMapOrArrayValue:", ctx.GetText())

	ctx.Identify().Accept(v)
	for _, key := range ctx.AllIDENTIFY() {
		key.Accept(v)
		//keyType := v.pop()
		//switch keyType.(type) {
		//case string:
		//	v.pu
		//	//v.HandlerIStringValue()
		//}
		//switch key.(type) {
		//case parser.INumContext:
		//	v.handlerArray(key.(parser.INumContext))
		//case parser.IStringValueContext:
		//	v.HandlerIStringValue(key.(parser.IStringValueContext))
		//default:
		//	ty := reflect.TypeOf(key)
		//	fmt.Println(111, ty.String())
		//}

		//value := v.pop()
		//key.Accept(v)
		//valueType := v.pop()
		//switch valueType.(type) {
		//case float64:
		//	v.VisitArray()
		//case string:
		//	v.VisitMap()
		//}
	}

	return nil
}

func (v *RuleEngineVisitor) VisitKey(ctx *parser.KeyContext) interface{} {
	fmt.Println("VisitKey:", ctx.GetText())
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChildOfType(i, nil)

		//ty := reflect.TypeOf(child)
		//fmt.Println(111, ty.String(), ty.Name())
		child.Accept(v)

		switch child.(type) {
		case *parser.NUMContext:
			v.handlerArray()
		case *parser.STRINGContext:
			v.HandlerMap()
		default:
			ty := reflect.TypeOf(child)
			v.setError(fmt.Errorf("无效的类型:%s", ty.String()))
		}
	}
	v.VisitChildren(ctx)
	return nil
}
