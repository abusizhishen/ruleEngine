package main

import (
	"github.com/abusizhishen/ruleEngine/parser"
	"github.com/abusizhishen/ruleEngine/src"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input,err := antlr.NewFileStream("test.rule")
	if err != nil{
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
	antlr.ParseTreeWalkerDefault.Walk(src.New(map[string]interface{}{"a":false, "b": 1,"c":0}), p.Init())
}
