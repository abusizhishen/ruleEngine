package src

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
	"reflect"
)

func (v *RuleEngineVisitor) VisitMapValue(ctx *parser.MapValueContext) interface{} {
	fmt.Println("VisitMapValue:", ctx.GetText())
	var hash = make(map[string]interface{}, ctx.GetChildCount())
	for _, element := range ctx.AllPair() {
		element.Accept(v)
		v, k := v.pop(), v.pop()
		hash[k.(string)] = v
	}

	v.push(hash)
	return nil
}

func (v *RuleEngineVisitor) VisitPair(ctx *parser.PairContext) interface{} {
	fmt.Println("VisitPair:", ctx.GetText())

	origin := ctx.GetMapKey().GetText()
	key := origin[1 : len(origin)-1]
	v.push(key)
	ctx.ValueType().Accept(v)

	return nil
}

func (v *RuleEngineVisitor) VisitGetMapValue(ctx *parser.GetMapValueContext) interface{} {
	fmt.Println("VisitGetMapValue:", ctx.GetText())

	ctx.Identify().Accept(v)
	value := v.pop()
	mapValue, ok := value.(map[string]interface{})
	if !ok {
		t := reflect.TypeOf(value)
		v.setError(fmt.Errorf("无效的操作 on variable:%s, type:%s", ctx.Identify().GetText(), t.String()))
	}
	origin := ctx.Str().GetText()
	key := origin[1 : len(origin)-1]

	vv, ok := mapValue[key]
	if !ok {
		v.setError(fmt.Errorf("map %s has no key :%s", ctx.Identify().GetText(), key))
	}

	v.push(vv)
	return nil
}
