package src

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
)

type RuleEngineListening struct {
	*parser.BaseRuleListener
	input map[string]interface{}
	output map[string]interface{}
}

func New(input map[string]interface{}) RuleEngineListening {
	var r = RuleEngineListening{
		BaseRuleListener: new(parser.BaseRuleListener),
		input:            make(map[string]interface{}),
		output:           make(map[string]interface{}),
	}

	for k,v := range input{
		r.input[k] = v
	}

	return r
}

func (r *RuleEngineListening)EnterStatement(ctx *parser.StatementContext)  {
	fmt.Println("EnterStatement: ",ctx.GetText())
}

func (r *RuleEngineListening)EnterInit(ctx *parser.InitContext)  {
	fmt.Println("init: ",ctx.GetText())
}

func (r *RuleEngineListening)EnterCalculateStatement(ctx *parser.CalculateStatementContext)  {
	fmt.Println("EnterCalculateStatement: ",ctx.GetText())
}

func (r *RuleEngineListening)EnterCalculateStatement(ctx *parser.id)  {
	fmt.Println("EnterCalculateStatement: ",ctx.GetText())
}