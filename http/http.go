package main

import (
	"github.com/abusizhishen/ruleEngine/parser"
	"github.com/abusizhishen/ruleEngine/src"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gin-gonic/gin"
)

func main() {
	http()
}

type Rule struct {
	rule string
	p *parser.RuleParser
	v *src.RuleEngineVisitor
}

var ruleMap = make(map[string]string)

func http()  {
	var g = gin.Default()
	g.POST("/create", createRule)
	g.POST("/calculate", calculateRule)
	g.GET("/show", showRule)


	err := g.Run(":80")
	if err != nil{
		panic(err)
	}
}

type Create struct {
	Name string `json:"name" binding:"required"`
	Rule string `json:"rule" binding:"required"`
}

type Calculate struct {
	Name string `json:"name" binding:"required"`
	Data map[string]interface{} `json:"data" binding:"required"`
}

func createRule(ctx *gin.Context)  {
	var create Create
	err := ctx.ShouldBindJSON(&create)
	if err != nil{
		ctx.String(200,err.Error())
		return
	}

	ruleMap[create.Name] = create.Rule
	ctx.JSON(200,nil)
}

func calculateRule(ctx *gin.Context)  {
	var calculate Calculate
	err := ctx.ShouldBindJSON(&calculate)
	if err != nil{
		ctx.String(200,err.Error())
		return
	}

	r,ok := ruleMap[calculate.Name]
	if !ok{
		ctx.String(200,"无效的规则: "+ calculate.Name)
		return
	}
	v := src.NewVisitor(calculate.Data)
	p := getParser(r)
	ctx.JSON(200,map[string]interface{}{"data":p.Init().Accept(v)})
}

func showRule(ctx *gin.Context)  {
	ctx.JSON(200,ruleMap)
}

func getParser(rule string) *parser.RuleParser{
	input := antlr.NewInputStream(rule)
	lex := parser.NewRuleLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	p := parser.NewRuleParser(tokens)

	return p
}

func resp(ctx *gin.Context ,data interface{})  {
	if data == nil{
		data = map[string]string{}
	}
	ctx.JSON(200,data)
}