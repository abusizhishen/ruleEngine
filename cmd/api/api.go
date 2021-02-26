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

var ruleMap = make(map[string]string)

func http() {
	var g = gin.Default()
	//添加规则
	g.POST("/add", addRule)
	//运行规则
	g.POST("/run", runRule)
	//展示规则
	g.GET("/show", showRule)

	err := g.Run(":80")
	if err != nil {
		panic(err)
	}
}

//创建规则参数
type Create struct {
	Name string `json:"name" binding:"required"`
	Rule string `json:"rule" binding:"required"`
}

//运行规则参数
type RunParams struct {
	Name string                 `json:"name" binding:"required"`
	Data map[string]interface{} `json:"data" binding:"required"`
}

//添加规则
func addRule(ctx *gin.Context) {
	var create Create
	err := ctx.ShouldBindJSON(&create)
	if err != nil {
		ctx.String(200, err.Error())
		return
	}

	ruleMap[create.Name] = create.Rule
	ctx.JSON(200, Resp{
		Code: 0,
		Data: []int{},
		Msg:  "success",
	})
}

//运行规则
func runRule(ctx *gin.Context) {
	var param RunParams
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		ctx.String(200, err.Error())
		return
	}

	r, ok := ruleMap[param.Name]
	if !ok {
		ctx.JSON(200, Resp{
			Code: 0,
			Data: nil,
			Msg:  "无效的规则" + param.Name,
		})
		return
	}
	v := src.NewVisitor(param.Data)
	p := getParser(r)
	ctx.JSON(200, Resp{
		Code: 0,
		Data: p.Init().Accept(v),
		Msg:  "",
	})
}

//展示规则
func showRule(ctx *gin.Context) {
	ctx.JSON(200, ruleMap)
}

func getParser(rule string) *parser.RuleParser {
	input := antlr.NewInputStream(rule)
	lex := parser.NewRuleLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	p := parser.NewRuleParser(tokens)
	return p
}

//响应数据
type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
