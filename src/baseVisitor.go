package src

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
	"strconv"
)

type RuleEngineVisitor struct {
	*parser.BaseRuleVisitor
	data   map[string]interface{}
	stack  []interface{}
	errors []error
}

func (r *RuleEngineVisitor) pop() interface{} {
	if len(r.stack) == 0 {
		panic("stack is empty")
	}

	l := len(r.stack)
	v := r.stack[l-1]
	r.stack = r.stack[:len(r.stack)-1]
	return v
}

func (r *RuleEngineVisitor) push(i interface{}) {
	r.stack = append(r.stack, i)
}

func (r *RuleEngineVisitor) setError(e error) {
	//r.errors = append(r.errors, e)
	panic(e.Error())
}

func NewVisitor(input map[string]interface{}) *RuleEngineVisitor {
	var r = RuleEngineVisitor{
		BaseRuleVisitor: new(parser.BaseRuleVisitor),
		data:            make(map[string]interface{}),
	}

	for k, v := range input {
		r.data[k] = v
	}

	return &r
}

func (r *RuleEngineVisitor) EnterInit(ctx *parser.InitContext) {
	fmt.Println("EnterInit: ", ctx.GetText())
}

func (r *RuleEngineVisitor) VisitInit(ctx *parser.InitContext) interface{} {
	if len(r.errors) != 0 {
		for _, e := range r.errors {
			fmt.Println(e)
		}
		return nil
	}

	fmt.Println("stack:", r.stack)
	fmt.Println("out", r.data)

	return nil
}

func (r *RuleEngineVisitor) VisitBOOLOP(ctx *parser.BOOLOPContext) interface{} {
	fmt.Println("VisitBOOLOP", ctx.GetText())
	right, left := r.pop(), r.pop()

	switch ctx.GetOp().GetText() {
	case "and":
		r.push(left.(bool) && right.(bool))
	case "or":
		r.push(left.(bool) || right.(bool))
	case "not":
		r.push(left.(bool) && (!right.(bool)))
	}

	return nil
}

func (r *RuleEngineVisitor) VisitCOMPARE(ctx *parser.COMPAREContext) interface{} {
	fmt.Println("VisitCompare: ", ctx.GetText())
	right, left := r.pop().(int), r.pop().(int)

	switch ctx.GetOp().GetText() {
	case ">=":
		r.push(left >= right)
	case ">":
		r.push(left > right)
	case "==":
		r.push(left == right)
	case "<=":
		r.push(left <= right)
	case "<":
		r.push(left < right)
	}
	return nil
}

func (r *RuleEngineVisitor) VisitMULDIV(ctx *parser.MULDIVContext) interface{} {
	fmt.Println("VisitMULDIV: ", ctx.GetText())
	right, left := r.pop().(int), r.pop().(int)

	switch ctx.GetOp().GetText() {
	case "*":
		r.push(left * right)
	case "/":
		if right == 0 {
			r.setError(fmt.Errorf("被除数不能为0"))
		}
		r.push(left / right)
	default:

	}
	return nil
}

func (r *RuleEngineVisitor) VisitADDSUB(ctx *parser.ADDSUBContext) interface{} {
	fmt.Println("VisitADDSUB: ", ctx.GetText())
	right, left := r.pop().(int), r.pop().(int)

	switch ctx.GetOp().GetText() {
	case "+":
		r.push(left + right)
	case "-":
		r.push(left - right)
	default:

	}

	return nil
}

func (r *RuleEngineVisitor) VisitIDENTIFY(ctx *parser.IDENTIFYContext) interface{} {
	fmt.Println("VisitIDEN: ", ctx.GetText())
	v, ok := r.data[ctx.GetText()]
	if !ok {
		err := fmt.Errorf("invalid variable %s", ctx.GetText())
		r.setError(err)
		return nil
	}

	r.push(v)

	return nil
}

func (r *RuleEngineVisitor) VisitNUM(ctx *parser.NUMContext) interface{} {
	str := ctx.GetText()
	va, err := strconv.Atoi(str)
	if err == nil {
		r.push(va)
		return nil
	}

	err = fmt.Errorf("invalid num %s, err:%v", str, err)
	r.setError(err)
	return nil
}

func (r *RuleEngineVisitor) VisitBOOL(ctx *parser.BOOLContext) interface{} {
	switch ctx.GetText() {
	case "true":
		r.push(true)
	case "false":
		r.push(false)
	}
	return nil

}

func (r *RuleEngineVisitor) VisitIDENBOOL(ctx *parser.IDENBOOLContext) interface{} {
	v, ok := r.data[ctx.GetText()]
	if !ok {
		err := fmt.Errorf("invalid variable %s", ctx.GetText())
		r.setError(err)
		return nil
	}

	va, ok := v.(bool)
	if ok {
		r.push(va)
		return nil
	}

	err := fmt.Errorf("invalid variable %s, value:%v, expect bool value", ctx.GetText(), v)
	r.setError(err)
	return nil
}

func (r *RuleEngineVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	fmt.Println("VisitIfStatement:", ctx.GetText())
	fmt.Println("condition:", r.pop())
	return nil
}

func (r *RuleEngineVisitor) VisitSetValueStatement(ctx *parser.SetValueStatementContext) interface{} {
	fmt.Println("VisitSetValue:", ctx.GetText())
	key := ctx.IDENTIFY().GetText()
	ctx.ValueType()
	r.data[key] = r.pop()
	r.push(r.data[key])
	fmt.Println(fmt.Sprintf("%s=%v", key, r.data[key]))

	return nil
}
