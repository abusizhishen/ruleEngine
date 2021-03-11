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

func (v *RuleEngineVisitor) handlerArray() interface{} {
	fmt.Println("handlerArray:")

	index := v.pop().(float64)
	value := v.pop()

	array, ok := value.([]interface{})
	if !ok {
		t := reflect.TypeOf(value)
		v.setError(fmt.Errorf("无效的操作 on variable:%+v", t.String()))
	}

	if index < 0 || len(array) <= int(index) {
		v.setError(fmt.Errorf("access array %+v err, index out of range :%v", array, index))
	}
	vv := array[int(index)]
	v.push(vv)
	return nil
	//var array []interface{}
	//for _, valueType := range ctx.AllValueType() {
	//	valueType.Accept(v)
	//	value := v.pop()
	//	array = append(array, value)
	//}
	//
	//v.push(array)

	return nil
}

//func (v *RuleEngineVisitor) VisitGetArrayValue(ctx *parser.GetArrayValueContext) interface{} {
//	fmt.Println("VisitGetArrayValue:", ctx.GetText())
//
//	ctx.Identify().Accept(v)
//
//}
