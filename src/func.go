package src

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
	"reflect"
)

func (v *RuleEngineVisitor)SetFun(name string, fun interface{})  {
	v.data[name] = fun
}

func (v *RuleEngineVisitor)RemoveFun(name string)  {
	delete(v.data, name)
}

func (v *RuleEngineVisitor)VisitFunCall(ctx *parser.FunCallContext) interface{} {
	fmt.Println("VisitFunCall: ",ctx.GetText())

	ctx.Identify().Accept(v)
	f := v.pop()


	var args []reflect.Value
	for _, ars := range ctx.AllValueType(){
		ars.Accept(v)
		args  = append(args, reflect.ValueOf(v.pop()))
	}

	result := reflect.ValueOf(f).Call(args)[0]

	v.push(result.Interface())
	return nil
}