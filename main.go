package main

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
	"github.com/abusizhishen/ruleEngine/src"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input, err := antlr.NewFileStream("test.rule")
	if err != nil {
		panic(err)
	}

	lex := parser.NewRuleLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	p := parser.NewRuleParser(tokens)
	v := src.NewVisitor(data)
	v.SetFun("add", func(a,b float64) float64{return a+b})
	result := p.Init().Accept(v)

	fmt.Println("result: ", result)

}

var data = map[string]interface{}{}
