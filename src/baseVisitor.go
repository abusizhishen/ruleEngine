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

	Info  bool
	Error bool
}

func (v *RuleEngineVisitor) pop() interface{} {
	if len(v.stack) == 0 {
		panic("stack is empty")
	}

	l := len(v.stack)
	vv := v.stack[l-1]
	v.stack = v.stack[:len(v.stack)-1]
	return vv
}

func (v *RuleEngineVisitor) push(i interface{}) {
	v.stack = append(v.stack, i)
}

func (v *RuleEngineVisitor) setError(e error) {
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

func (v *RuleEngineVisitor) VisitInit(ctx *parser.InitContext) interface{} {
	fmt.Println("VisitInit:", ctx.GetText())

	for _, statement := range ctx.AllStatement() {
		statement.Accept(v)
	}

	return v.stack
}

func (v *RuleEngineVisitor) VisitBOOLOP(ctx *parser.BOOLOPContext) interface{} {
	fmt.Println("VisitBOOLOP", ctx.GetText())
	right, left := v.pop(), v.pop()

	switch ctx.GetOp().GetText() {
	case "&&":
		v.push(left.(bool) && right.(bool))
	case "||":
		v.push(left.(bool) || right.(bool))
	case "!":
		v.push(left.(bool) && (!right.(bool)))
	}

	return nil
}
func (v *RuleEngineVisitor) VisitCOMPAREX(ctx *parser.COMPAREXContext) interface{} {
	fmt.Println("VisitCOMPAREX", ctx.GetText())
	ctx.CompareStatement().Accept(v)
	return nil
}

func (v *RuleEngineVisitor) VisitCalcu(ctx *parser.CalcuContext) interface{} {
	fmt.Println("VisitCalcu", ctx.GetText())
	ctx.CalculateStatement().Accept(v)
	return nil
}

func (v *RuleEngineVisitor) VisitCOMPARE(ctx *parser.COMPAREContext) interface{} {
	fmt.Println("VisitCOMPARE: ", ctx.GetText())
	ctx.CompareStatement(0).Accept(v)
	ctx.CompareStatement(1).Accept(v)
	right, left := v.pop().(int), v.pop().(int)

	switch ctx.GetOp().GetText() {
	case ">=":
		v.push(left >= right)
	case ">":
		v.push(left > right)
	case "==":
		v.push(left == right)
	case "<=":
		v.push(left <= right)
	case "<":
		v.push(left < right)
	}
	return nil
}

func (v *RuleEngineVisitor) VisitMULDIV(ctx *parser.MULDIVContext) interface{} {
	fmt.Println("VisitMULDIV: ", ctx.GetText())
	ctx.CalculateStatement(0).Accept(v)
	ctx.CalculateStatement(1).Accept(v)
	right, left := v.pop().(int), v.pop().(int)

	switch ctx.GetOp().GetText() {
	case "*":
		v.push(left * right)
	case "/":
		if right == 0 {
			v.setError(fmt.Errorf("被除数不能为0"))
		}
		v.push(left / right)
	default:
		panic("invalid op: " + ctx.GetOp().GetText())
	}
	return nil
}

func (v *RuleEngineVisitor) VisitADDSUB(ctx *parser.ADDSUBContext) interface{} {
	fmt.Println("VisitADDSUB: ", ctx.GetText())
	ctx.CalculateStatement(0).Accept(v)
	ctx.CalculateStatement(1).Accept(v)

	right, left := v.pop().(int), v.pop().(int)

	switch ctx.GetOp().GetText() {
	case "+":
		v.push(left + right)
	case "-":
		v.push(left - right)
	default:
		panic("invalid op: " + ctx.GetOp().GetText())
	}

	return nil
}

func (v *RuleEngineVisitor) VisitIDENTIFY(ctx *parser.IDENTIFYContext) interface{} {
	fmt.Println("VisitIDENTIFY: ", ctx.GetText())
	vv, ok := v.data[ctx.GetText()]
	if !ok {
		err := fmt.Errorf("invalid variable %s", ctx.GetText())
		v.setError(err)
		return nil
	}

	v.push(vv)

	return nil
}

func (v *RuleEngineVisitor) VisitNUM(ctx *parser.NUMContext) interface{} {
	fmt.Println("VisitNUM:", ctx.GetText())

	str := ctx.GetText()
	va, err := strconv.Atoi(str)
	if err == nil {
		v.push(va)
		return nil
	}

	err = fmt.Errorf("invalid num %s, err:%v", str, err)
	v.setError(err)
	return nil
}

func (v *RuleEngineVisitor) VisitBOOL(ctx *parser.BOOLContext) interface{} {
	fmt.Println("VisitBOOL:", ctx.GetText())

	switch ctx.GetText() {
	case "true":
		v.push(true)
	case "false":
		v.push(false)
	}
	return nil

}

func (v *RuleEngineVisitor) VisitIDENBOOL(ctx *parser.IDENBOOLContext) interface{} {
	fmt.Println("VisitIDENBOOL:", ctx.GetText())

	vv, ok := v.data[ctx.GetText()]
	if !ok {
		err := fmt.Errorf("invalid variable %s", ctx.GetText())
		v.setError(err)
		return nil
	}

	va, ok := vv.(bool)
	if ok {
		v.push(va)
		return nil
	}

	err := fmt.Errorf("invalid variable %s, value:%v, expect bool value", ctx.GetText(), v)
	v.setError(err)
	return nil
}

func (v *RuleEngineVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	fmt.Println("VisitIfStatement:", ctx.GetText())
	ctx.BoolStatement().Accept(v)

	if v.pop().(bool) {
		for _, statement := range ctx.AllStatement() {
			statement.Accept(v)
		}

	} else {
		ctx.ElseStatement().Accept(v)
	}

	return nil
}

func (v *RuleEngineVisitor) VisitSetValueStatement(ctx *parser.SetValueStatementContext) interface{} {
	fmt.Println("VisitSetValue:", ctx.GetText())
	ctx.ValueType().Accept(v)
	key := ctx.IDENTIFY().GetText()
	v.data[key] = v.pop()
	fmt.Println(fmt.Sprintf("%s=%v", key, v.data[key]))

	return nil
}

func (v *RuleEngineVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Println("VisitStatement:", ctx.GetText())

	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChildOfType(i, nil)

		//ty := reflect.TypeOf(child)
		//fmt.Println(ty.String(), ty.Name())
		child.Accept(v)
	}

	return nil
}

func (v *RuleEngineVisitor) VisitElseStatement(ctx *parser.ElseStatementContext) interface{} {
	fmt.Println("VisitElseStatement:", ctx.GetText())

	for _, statement := range ctx.AllStatement() {
		statement.Accept(v)
	}

	return nil
}

func (v *RuleEngineVisitor) VisitValueType(ctx *parser.ValueTypeContext) interface{} {
	fmt.Println("VisitCalculate:", ctx.GetText())

	for i := 0; i < ctx.GetChildCount(); i++ {
		ctx.GetChildOfType(i, nil).Accept(v)
	}

	return nil
}

func (v *RuleEngineVisitor) VisitCOMPAREBOOL(ctx *parser.COMPAREBOOLContext) interface{} {
	fmt.Println("VisitCOMPAREBOOL:", ctx.GetText())
	ctx.CompareStatement().Accept(v)
	return nil
}

func (v *RuleEngineVisitor) VisitCalculate(ctx *parser.CalculateContext) interface{} {
	fmt.Println("VisitCalculate:", ctx.GetText())

	for i := 0; i < ctx.GetChildCount(); i++ {
		ctx.GetChildOfType(i, nil).Accept(v)
	}

	return nil
}

func (v *RuleEngineVisitor) VisitITEMCALCU(ctx *parser.ITEMCALCUContext) interface{} {
	fmt.Println("VisitITEMCALCU:", ctx.GetText())

	for i := 0; i < ctx.GetChildCount(); i++ {
		//fmt.Println(reflect.TypeOf(ctx.GetChild(i)).String())
		ctx.GetChildOfType(i, nil).Accept(v)

	}

	return nil
}

func (v *RuleEngineVisitor) VisitITEMCOMP(ctx *parser.ITEMCOMPContext) interface{} {
	fmt.Println("VisitITEMCOMP:", ctx.GetText())

	for i := 0; i < ctx.GetChildCount(); i++ {
		//fmt.Println(reflect.TypeOf(ctx.GetChild(i)).String())
		ctx.GetChildOfType(i, nil).Accept(v)

	}

	return nil
}

func (v *RuleEngineVisitor) VisitCalculateValue(ctx *parser.CalculateValueContext) interface{} {
	fmt.Println("VisitCalculateValue:", ctx.GetText())

	for i := 0; i < ctx.GetChildCount(); i++ {
		//fmt.Println(reflect.TypeOf(ctx.GetChild(i)).String())
		ctx.GetChildOfType(i, nil).Accept(v)

	}

	return nil
}

func (v *RuleEngineVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	fmt.Println("VisitReturnStatement:", ctx.GetText())
	ctx.GetValue().Accept(v)

	return nil
}
