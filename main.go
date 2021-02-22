package main

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
	"github.com/abusizhishen/ruleEngine/src"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input, err := antlr.NewFileStream("test.ru")
	if err != nil {
		panic(err)
	}

	lex := parser.NewRuleLexer(input)
	//tok := lex.NextToken()
	//fmt.Println(lex.GetAllTokens())
	//for _,token := range lex.GetAllTokens(){
	//	fmt.Println(token)
	//}
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	p := parser.NewRuleParser(tokens)
	v := src.NewVisitor(data)
	result := p.Init().Accept(v)
	fmt.Println("result: ", result)
}

var data = map[string]interface{}{"a": true, "b": 1, "c": 0}