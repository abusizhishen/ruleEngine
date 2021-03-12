package main

import (
	"fmt"
	"github.com/abusizhishen/ruleEngine/parser"
	"github.com/abusizhishen/ruleEngine/src"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/c-bata/go-prompt"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		run(antlr.NewInputStream(input()))
	}
}

var data map[string]interface{}
func run(input antlr.CharStream)  {
	defer func() {
		if err := recover(); err != nil{
			log.Printf("exception: %v",err)
		}
	}()
	lexer := parser.NewRuleLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer,antlr.LexerDefaultTokenChannel)
	p := parser.NewRuleParser(tokens)

	v := src.NewVisitor(data)
	defer func() {data = v.Data()}()
	result := p.Init().Accept(v)

	fmt.Println("result: ", result)
}


func input() string{
	var buf strings.Builder
	//循环读取 遇到 ';' 结束
	for {
		s := prompt.Input("> ", completer)
		buf.WriteString(s)

		switch s {
		case "exit","e","quit","q":
			os.Exit(0)
		default:
			if len(s) > 0 && s[len(s)-1] == ';'{
				s = buf.String()
				buf.Reset()
				return s
			}
		}
	}

	return ""
}

func completer(d prompt.Document) []prompt.Suggest {
	var s  []prompt.Suggest
	return prompt.FilterHasPrefix(s, ";", true)
}