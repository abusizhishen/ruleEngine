package src

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
	"reflect"
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

func (v *RuleEngineVisitor) VisitGetArrayValue(ctx *parser.GetArrayValueContext) interface{} {
	fmt.Println("VisitGetArrayValue:", ctx.GetText())

	ctx.Identify().Accept(v)
	value := v.pop()
	ctx.Num().Accept(v)
	index := v.pop().(float64)
	array, ok := value.([]interface{})
	if !ok {
		t := reflect.TypeOf(value)
		v.setError(fmt.Errorf("无效的操作 on variable:%s, type:%s", ctx.Identify().GetText(), t.String()))
	}


	if index< 0 || len(array) <= int(index) {
		v.setError(fmt.Errorf("access array %s err, index out of range :%v", ctx.Identify().GetText(), index))
	}
	vv := array[int(index)]
	v.push(vv)
	return nil
}