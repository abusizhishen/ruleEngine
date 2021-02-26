// Code generated from Rule.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Rule

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 198,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 56, 10, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 7, 6, 62, 10, 6, 12, 6, 14, 6, 65, 11, 6, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 75, 10, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 83, 10, 11, 3, 11, 3, 11, 3, 11, 6,
	11, 88, 10, 11, 13, 11, 14, 11, 89, 3, 11, 3, 11, 3, 11, 6, 11, 95, 10,
	11, 13, 11, 14, 11, 96, 7, 11, 99, 10, 11, 12, 11, 14, 11, 102, 11, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 112, 10,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 118, 10, 12, 12, 12, 14, 12, 121,
	11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 127, 10, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 133, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 139,
	10, 15, 12, 15, 14, 15, 142, 11, 15, 3, 15, 3, 15, 7, 15, 146, 10, 15,
	12, 15, 14, 15, 149, 11, 15, 3, 15, 5, 15, 152, 10, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 7, 16, 158, 10, 16, 12, 16, 14, 16, 161, 11, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 7, 17, 168, 10, 17, 12, 17, 14, 17, 171, 11, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 179, 10, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 187, 10, 19, 3, 20, 7, 20, 190, 10,
	20, 12, 20, 14, 20, 193, 11, 20, 3, 20, 5, 20, 196, 10, 20, 3, 20, 2, 5,
	10, 20, 22, 21, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 2, 8, 3, 2, 12, 14, 3, 2, 26, 29, 3, 2, 21, 25, 3, 2, 18,
	19, 3, 2, 28, 29, 3, 2, 26, 27, 2, 209, 2, 40, 3, 2, 2, 2, 4, 42, 3, 2,
	2, 2, 6, 44, 3, 2, 2, 2, 8, 46, 3, 2, 2, 2, 10, 55, 3, 2, 2, 2, 12, 66,
	3, 2, 2, 2, 14, 68, 3, 2, 2, 2, 16, 70, 3, 2, 2, 2, 18, 74, 3, 2, 2, 2,
	20, 82, 3, 2, 2, 2, 22, 111, 3, 2, 2, 2, 24, 126, 3, 2, 2, 2, 26, 128,
	3, 2, 2, 2, 28, 134, 3, 2, 2, 2, 30, 153, 3, 2, 2, 2, 32, 164, 3, 2, 2,
	2, 34, 178, 3, 2, 2, 2, 36, 186, 3, 2, 2, 2, 38, 191, 3, 2, 2, 2, 40, 41,
	9, 2, 2, 2, 41, 3, 3, 2, 2, 2, 42, 43, 9, 3, 2, 2, 43, 5, 3, 2, 2, 2, 44,
	45, 9, 4, 2, 2, 45, 7, 3, 2, 2, 2, 46, 47, 9, 2, 2, 2, 47, 9, 3, 2, 2,
	2, 48, 49, 8, 6, 1, 2, 49, 56, 5, 18, 10, 2, 50, 56, 5, 20, 11, 2, 51,
	52, 7, 3, 2, 2, 52, 53, 5, 10, 6, 2, 53, 54, 7, 4, 2, 2, 54, 56, 3, 2,
	2, 2, 55, 48, 3, 2, 2, 2, 55, 50, 3, 2, 2, 2, 55, 51, 3, 2, 2, 2, 56, 63,
	3, 2, 2, 2, 57, 58, 12, 6, 2, 2, 58, 59, 5, 6, 4, 2, 59, 60, 5, 10, 6,
	7, 60, 62, 3, 2, 2, 2, 61, 57, 3, 2, 2, 2, 62, 65, 3, 2, 2, 2, 63, 61,
	3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 11, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2,
	66, 67, 7, 10, 2, 2, 67, 13, 3, 2, 2, 2, 68, 69, 9, 5, 2, 2, 69, 15, 3,
	2, 2, 2, 70, 71, 7, 20, 2, 2, 71, 17, 3, 2, 2, 2, 72, 75, 5, 16, 9, 2,
	73, 75, 5, 12, 7, 2, 74, 72, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 19, 3,
	2, 2, 2, 76, 77, 8, 11, 1, 2, 77, 83, 5, 18, 10, 2, 78, 79, 7, 3, 2, 2,
	79, 80, 5, 20, 11, 2, 80, 81, 7, 4, 2, 2, 81, 83, 3, 2, 2, 2, 82, 76, 3,
	2, 2, 2, 82, 78, 3, 2, 2, 2, 83, 100, 3, 2, 2, 2, 84, 87, 12, 6, 2, 2,
	85, 86, 9, 6, 2, 2, 86, 88, 5, 20, 11, 2, 87, 85, 3, 2, 2, 2, 88, 89, 3,
	2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 99, 3, 2, 2, 2, 91,
	94, 12, 5, 2, 2, 92, 93, 9, 7, 2, 2, 93, 95, 5, 20, 11, 2, 94, 92, 3, 2,
	2, 2, 95, 96, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99,
	3, 2, 2, 2, 98, 84, 3, 2, 2, 2, 98, 91, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2,
	100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 21, 3, 2, 2, 2, 102, 100,
	3, 2, 2, 2, 103, 104, 8, 12, 1, 2, 104, 112, 9, 5, 2, 2, 105, 112, 7, 20,
	2, 2, 106, 112, 5, 10, 6, 2, 107, 108, 7, 3, 2, 2, 108, 109, 5, 22, 12,
	2, 109, 110, 7, 4, 2, 2, 110, 112, 3, 2, 2, 2, 111, 103, 3, 2, 2, 2, 111,
	105, 3, 2, 2, 2, 111, 106, 3, 2, 2, 2, 111, 107, 3, 2, 2, 2, 112, 119,
	3, 2, 2, 2, 113, 114, 12, 7, 2, 2, 114, 115, 5, 2, 2, 2, 115, 116, 5, 22,
	12, 8, 116, 118, 3, 2, 2, 2, 117, 113, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2,
	119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 23, 3, 2, 2, 2, 121, 119,
	3, 2, 2, 2, 122, 127, 5, 14, 8, 2, 123, 127, 5, 16, 9, 2, 124, 127, 5,
	12, 7, 2, 125, 127, 5, 20, 11, 2, 126, 122, 3, 2, 2, 2, 126, 123, 3, 2,
	2, 2, 126, 124, 3, 2, 2, 2, 126, 125, 3, 2, 2, 2, 127, 25, 3, 2, 2, 2,
	128, 129, 7, 20, 2, 2, 129, 130, 7, 5, 2, 2, 130, 132, 5, 24, 13, 2, 131,
	133, 7, 6, 2, 2, 132, 131, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 27, 3,
	2, 2, 2, 134, 135, 7, 15, 2, 2, 135, 136, 5, 22, 12, 2, 136, 140, 7, 7,
	2, 2, 137, 139, 5, 36, 19, 2, 138, 137, 3, 2, 2, 2, 139, 142, 3, 2, 2,
	2, 140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 143, 3, 2, 2, 2, 142,
	140, 3, 2, 2, 2, 143, 147, 7, 8, 2, 2, 144, 146, 5, 30, 16, 2, 145, 144,
	3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2,
	2, 2, 148, 151, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 152, 5, 32, 17,
	2, 151, 150, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 29, 3, 2, 2, 2, 153,
	154, 7, 16, 2, 2, 154, 155, 5, 22, 12, 2, 155, 159, 7, 7, 2, 2, 156, 158,
	5, 36, 19, 2, 157, 156, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3,
	2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 162, 3, 2, 2, 2, 161, 159, 3, 2, 2,
	2, 162, 163, 7, 8, 2, 2, 163, 31, 3, 2, 2, 2, 164, 165, 7, 17, 2, 2, 165,
	169, 7, 7, 2, 2, 166, 168, 5, 36, 19, 2, 167, 166, 3, 2, 2, 2, 168, 171,
	3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 172, 3, 2,
	2, 2, 171, 169, 3, 2, 2, 2, 172, 173, 7, 8, 2, 2, 173, 33, 3, 2, 2, 2,
	174, 175, 7, 9, 2, 2, 175, 179, 5, 24, 13, 2, 176, 179, 5, 20, 11, 2, 177,
	179, 5, 22, 12, 2, 178, 174, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 177,
	3, 2, 2, 2, 179, 35, 3, 2, 2, 2, 180, 187, 5, 20, 11, 2, 181, 187, 5, 22,
	12, 2, 182, 187, 5, 10, 6, 2, 183, 187, 5, 28, 15, 2, 184, 187, 5, 26,
	14, 2, 185, 187, 5, 34, 18, 2, 186, 180, 3, 2, 2, 2, 186, 181, 3, 2, 2,
	2, 186, 182, 3, 2, 2, 2, 186, 183, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186,
	185, 3, 2, 2, 2, 187, 37, 3, 2, 2, 2, 188, 190, 5, 36, 19, 2, 189, 188,
	3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 192, 3, 2,
	2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 196, 7, 2, 2, 3,
	195, 194, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 39, 3, 2, 2, 2, 23, 55,
	63, 74, 82, 89, 96, 98, 100, 111, 119, 126, 132, 140, 147, 151, 159, 169,
	178, 186, 191, 195,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'='", "';'", "'{'", "'}'", "'return'", "", "'now()'",
	"'&&'", "'||'", "'!'", "'if'", "'elsif'", "'else'", "'true'", "'false'",
	"", "'>'", "'>='", "'=='", "'<='", "'<'", "'+'", "'-'", "'*'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "NUM", "NOW", "AND", "OR", "NOT", "IF",
	"ELSIF", "ELSE", "TRUE", "FALSE", "IDENTIFY", "GT", "GTE", "EQ", "LTE",
	"LT", "ADD", "SUB", "MUL", "DIV", "WS",
}

var ruleNames = []string{
	"boolOperate", "calculate", "compare", "logical", "compareStatement", "num",
	"boolValue", "identify", "calculateValue", "calculateStatement", "boolStatement",
	"valueType", "setValueStatement", "ifStatement", "elseIfStatement", "elseStatement",
	"returnStatement", "statement", "init",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RuleParser struct {
	*antlr.BaseParser
}

func NewRuleParser(input antlr.TokenStream) *RuleParser {
	this := new(RuleParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Rule.g4"

	return this
}

// RuleParser tokens.
const (
	RuleParserEOF      = antlr.TokenEOF
	RuleParserT__0     = 1
	RuleParserT__1     = 2
	RuleParserT__2     = 3
	RuleParserT__3     = 4
	RuleParserT__4     = 5
	RuleParserT__5     = 6
	RuleParserT__6     = 7
	RuleParserNUM      = 8
	RuleParserNOW      = 9
	RuleParserAND      = 10
	RuleParserOR       = 11
	RuleParserNOT      = 12
	RuleParserIF       = 13
	RuleParserELSIF    = 14
	RuleParserELSE     = 15
	RuleParserTRUE     = 16
	RuleParserFALSE    = 17
	RuleParserIDENTIFY = 18
	RuleParserGT       = 19
	RuleParserGTE      = 20
	RuleParserEQ       = 21
	RuleParserLTE      = 22
	RuleParserLT       = 23
	RuleParserADD      = 24
	RuleParserSUB      = 25
	RuleParserMUL      = 26
	RuleParserDIV      = 27
	RuleParserWS       = 28
)

// RuleParser rules.
const (
	RuleParserRULE_boolOperate        = 0
	RuleParserRULE_calculate          = 1
	RuleParserRULE_compare            = 2
	RuleParserRULE_logical            = 3
	RuleParserRULE_compareStatement   = 4
	RuleParserRULE_num                = 5
	RuleParserRULE_boolValue          = 6
	RuleParserRULE_identify           = 7
	RuleParserRULE_calculateValue     = 8
	RuleParserRULE_calculateStatement = 9
	RuleParserRULE_boolStatement      = 10
	RuleParserRULE_valueType          = 11
	RuleParserRULE_setValueStatement  = 12
	RuleParserRULE_ifStatement        = 13
	RuleParserRULE_elseIfStatement    = 14
	RuleParserRULE_elseStatement      = 15
	RuleParserRULE_returnStatement    = 16
	RuleParserRULE_statement          = 17
	RuleParserRULE_init               = 18
)

// IBoolOperateContext is an interface to support dynamic dispatch.
type IBoolOperateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolOperateContext differentiates from other interfaces.
	IsBoolOperateContext()
}

type BoolOperateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolOperateContext() *BoolOperateContext {
	var p = new(BoolOperateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_boolOperate
	return p
}

func (*BoolOperateContext) IsBoolOperateContext() {}

func NewBoolOperateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolOperateContext {
	var p = new(BoolOperateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_boolOperate

	return p
}

func (s *BoolOperateContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolOperateContext) AND() antlr.TerminalNode {
	return s.GetToken(RuleParserAND, 0)
}

func (s *BoolOperateContext) OR() antlr.TerminalNode {
	return s.GetToken(RuleParserOR, 0)
}

func (s *BoolOperateContext) NOT() antlr.TerminalNode {
	return s.GetToken(RuleParserNOT, 0)
}

func (s *BoolOperateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolOperateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolOperateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitBoolOperate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) BoolOperate() (localctx IBoolOperateContext) {
	localctx = NewBoolOperateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RuleParserRULE_boolOperate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserAND)|(1<<RuleParserOR)|(1<<RuleParserNOT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICalculateContext is an interface to support dynamic dispatch.
type ICalculateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalculateContext differentiates from other interfaces.
	IsCalculateContext()
}

type CalculateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalculateContext() *CalculateContext {
	var p = new(CalculateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_calculate
	return p
}

func (*CalculateContext) IsCalculateContext() {}

func NewCalculateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculateContext {
	var p = new(CalculateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_calculate

	return p
}

func (s *CalculateContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculateContext) ADD() antlr.TerminalNode {
	return s.GetToken(RuleParserADD, 0)
}

func (s *CalculateContext) SUB() antlr.TerminalNode {
	return s.GetToken(RuleParserSUB, 0)
}

func (s *CalculateContext) MUL() antlr.TerminalNode {
	return s.GetToken(RuleParserMUL, 0)
}

func (s *CalculateContext) DIV() antlr.TerminalNode {
	return s.GetToken(RuleParserDIV, 0)
}

func (s *CalculateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalculateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitCalculate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Calculate() (localctx ICalculateContext) {
	localctx = NewCalculateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RuleParserRULE_calculate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserADD)|(1<<RuleParserSUB)|(1<<RuleParserMUL)|(1<<RuleParserDIV))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICompareContext is an interface to support dynamic dispatch.
type ICompareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareContext differentiates from other interfaces.
	IsCompareContext()
}

type CompareContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareContext() *CompareContext {
	var p = new(CompareContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_compare
	return p
}

func (*CompareContext) IsCompareContext() {}

func NewCompareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareContext {
	var p = new(CompareContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_compare

	return p
}

func (s *CompareContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareContext) GT() antlr.TerminalNode {
	return s.GetToken(RuleParserGT, 0)
}

func (s *CompareContext) GTE() antlr.TerminalNode {
	return s.GetToken(RuleParserGTE, 0)
}

func (s *CompareContext) EQ() antlr.TerminalNode {
	return s.GetToken(RuleParserEQ, 0)
}

func (s *CompareContext) LTE() antlr.TerminalNode {
	return s.GetToken(RuleParserLTE, 0)
}

func (s *CompareContext) LT() antlr.TerminalNode {
	return s.GetToken(RuleParserLT, 0)
}

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Compare() (localctx ICompareContext) {
	localctx = NewCompareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RuleParserRULE_compare)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserGT)|(1<<RuleParserGTE)|(1<<RuleParserEQ)|(1<<RuleParserLTE)|(1<<RuleParserLT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogicalContext is an interface to support dynamic dispatch.
type ILogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalContext differentiates from other interfaces.
	IsLogicalContext()
}

type LogicalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalContext() *LogicalContext {
	var p = new(LogicalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_logical
	return p
}

func (*LogicalContext) IsLogicalContext() {}

func NewLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalContext {
	var p = new(LogicalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_logical

	return p
}

func (s *LogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalContext) AND() antlr.TerminalNode {
	return s.GetToken(RuleParserAND, 0)
}

func (s *LogicalContext) OR() antlr.TerminalNode {
	return s.GetToken(RuleParserOR, 0)
}

func (s *LogicalContext) NOT() antlr.TerminalNode {
	return s.GetToken(RuleParserNOT, 0)
}

func (s *LogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitLogical(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Logical() (localctx ILogicalContext) {
	localctx = NewLogicalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RuleParserRULE_logical)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserAND)|(1<<RuleParserOR)|(1<<RuleParserNOT))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICompareStatementContext is an interface to support dynamic dispatch.
type ICompareStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareStatementContext differentiates from other interfaces.
	IsCompareStatementContext()
}

type CompareStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareStatementContext() *CompareStatementContext {
	var p = new(CompareStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_compareStatement
	return p
}

func (*CompareStatementContext) IsCompareStatementContext() {}

func NewCompareStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareStatementContext {
	var p = new(CompareStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_compareStatement

	return p
}

func (s *CompareStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareStatementContext) CopyFrom(ctx *CompareStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CompareStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type COMPAREXContext struct {
	*CompareStatementContext
}

func NewCOMPAREXContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *COMPAREXContext {
	var p = new(COMPAREXContext)

	p.CompareStatementContext = NewEmptyCompareStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareStatementContext))

	return p
}

func (s *COMPAREXContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *COMPAREXContext) CompareStatement() ICompareStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *COMPAREXContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitCOMPAREX(s)

	default:
		return t.VisitChildren(s)
	}
}

type COMPAREContext struct {
	*CompareStatementContext
	op ICompareContext
}

func NewCOMPAREContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *COMPAREContext {
	var p = new(COMPAREContext)

	p.CompareStatementContext = NewEmptyCompareStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareStatementContext))

	return p
}

func (s *COMPAREContext) GetOp() ICompareContext { return s.op }

func (s *COMPAREContext) SetOp(v ICompareContext) { s.op = v }

func (s *COMPAREContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *COMPAREContext) AllCompareStatement() []ICompareStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICompareStatementContext)(nil)).Elem())
	var tst = make([]ICompareStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICompareStatementContext)
		}
	}

	return tst
}

func (s *COMPAREContext) CompareStatement(i int) ICompareStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *COMPAREContext) Compare() ICompareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareContext)
}

func (s *COMPAREContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitCOMPARE(s)

	default:
		return t.VisitChildren(s)
	}
}

type CalcuContext struct {
	*CompareStatementContext
}

func NewCalcuContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CalcuContext {
	var p = new(CalcuContext)

	p.CompareStatementContext = NewEmptyCompareStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareStatementContext))

	return p
}

func (s *CalcuContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcuContext) CalculateStatement() ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *CalcuContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitCalcu(s)

	default:
		return t.VisitChildren(s)
	}
}

type ITEMCOMPContext struct {
	*CompareStatementContext
}

func NewITEMCOMPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ITEMCOMPContext {
	var p = new(ITEMCOMPContext)

	p.CompareStatementContext = NewEmptyCompareStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CompareStatementContext))

	return p
}

func (s *ITEMCOMPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ITEMCOMPContext) CalculateValue() ICalculateValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateValueContext)
}

func (s *ITEMCOMPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitITEMCOMP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) CompareStatement() (localctx ICompareStatementContext) {
	return p.compareStatement(0)
}

func (p *RuleParser) compareStatement(_p int) (localctx ICompareStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCompareStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICompareStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, RuleParserRULE_compareStatement, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewITEMCOMPContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(47)
			p.CalculateValue()
		}

	case 2:
		localctx = NewCalcuContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.calculateStatement(0)
		}

	case 3:
		localctx = NewCOMPAREXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(50)
			p.compareStatement(0)
		}
		{
			p.SetState(51)
			p.Match(RuleParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCOMPAREContext(p, NewCompareStatementContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_compareStatement)
			p.SetState(55)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(56)

				var _x = p.Compare()

				localctx.(*COMPAREContext).op = _x
			}
			{
				p.SetState(57)
				p.compareStatement(5)
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// INumContext is an interface to support dynamic dispatch.
type INumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumContext differentiates from other interfaces.
	IsNumContext()
}

type NumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumContext() *NumContext {
	var p = new(NumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_num
	return p
}

func (*NumContext) IsNumContext() {}

func NewNumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumContext {
	var p = new(NumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_num

	return p
}

func (s *NumContext) GetParser() antlr.Parser { return s.parser }

func (s *NumContext) CopyFrom(ctx *NumContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NUMContext struct {
	*NumContext
}

func NewNUMContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NUMContext {
	var p = new(NUMContext)

	p.NumContext = NewEmptyNumContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumContext))

	return p
}

func (s *NUMContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NUMContext) NUM() antlr.TerminalNode {
	return s.GetToken(RuleParserNUM, 0)
}

func (s *NUMContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitNUM(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Num() (localctx INumContext) {
	localctx = NewNumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RuleParserRULE_num)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewNUMContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(RuleParserNUM)
	}

	return localctx
}

// IBoolValueContext is an interface to support dynamic dispatch.
type IBoolValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolValueContext differentiates from other interfaces.
	IsBoolValueContext()
}

type BoolValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolValueContext() *BoolValueContext {
	var p = new(BoolValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_boolValue
	return p
}

func (*BoolValueContext) IsBoolValueContext() {}

func NewBoolValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolValueContext {
	var p = new(BoolValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_boolValue

	return p
}

func (s *BoolValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RuleParserTRUE, 0)
}

func (s *BoolValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RuleParserFALSE, 0)
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitBoolValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) BoolValue() (localctx IBoolValueContext) {
	localctx = NewBoolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RuleParserRULE_boolValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RuleParserTRUE || _la == RuleParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifyContext is an interface to support dynamic dispatch.
type IIdentifyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifyContext differentiates from other interfaces.
	IsIdentifyContext()
}

type IdentifyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifyContext() *IdentifyContext {
	var p = new(IdentifyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_identify
	return p
}

func (*IdentifyContext) IsIdentifyContext() {}

func NewIdentifyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifyContext {
	var p = new(IdentifyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_identify

	return p
}

func (s *IdentifyContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifyContext) CopyFrom(ctx *IdentifyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IdentifyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IDENTIFYContext struct {
	*IdentifyContext
}

func NewIDENTIFYContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDENTIFYContext {
	var p = new(IDENTIFYContext)

	p.IdentifyContext = NewEmptyIdentifyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IdentifyContext))

	return p
}

func (s *IDENTIFYContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDENTIFYContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *IDENTIFYContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitIDENTIFY(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Identify() (localctx IIdentifyContext) {
	localctx = NewIdentifyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RuleParserRULE_identify)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewIDENTIFYContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(RuleParserIDENTIFY)
	}

	return localctx
}

// ICalculateValueContext is an interface to support dynamic dispatch.
type ICalculateValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalculateValueContext differentiates from other interfaces.
	IsCalculateValueContext()
}

type CalculateValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalculateValueContext() *CalculateValueContext {
	var p = new(CalculateValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_calculateValue
	return p
}

func (*CalculateValueContext) IsCalculateValueContext() {}

func NewCalculateValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculateValueContext {
	var p = new(CalculateValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_calculateValue

	return p
}

func (s *CalculateValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculateValueContext) Identify() IIdentifyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifyContext)
}

func (s *CalculateValueContext) Num() INumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *CalculateValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalculateValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitCalculateValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) CalculateValue() (localctx ICalculateValueContext) {
	localctx = NewCalculateValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RuleParserRULE_calculateValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(72)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserIDENTIFY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Identify()
		}

	case RuleParserNUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Num()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICalculateStatementContext is an interface to support dynamic dispatch.
type ICalculateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalculateStatementContext differentiates from other interfaces.
	IsCalculateStatementContext()
}

type CalculateStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalculateStatementContext() *CalculateStatementContext {
	var p = new(CalculateStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_calculateStatement
	return p
}

func (*CalculateStatementContext) IsCalculateStatementContext() {}

func NewCalculateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculateStatementContext {
	var p = new(CalculateStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_calculateStatement

	return p
}

func (s *CalculateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculateStatementContext) CopyFrom(ctx *CalculateStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CalculateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ITEMCALCUContext struct {
	*CalculateStatementContext
}

func NewITEMCALCUContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ITEMCALCUContext {
	var p = new(ITEMCALCUContext)

	p.CalculateStatementContext = NewEmptyCalculateStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateStatementContext))

	return p
}

func (s *ITEMCALCUContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ITEMCALCUContext) CalculateValue() ICalculateValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateValueContext)
}

func (s *ITEMCALCUContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitITEMCALCU(s)

	default:
		return t.VisitChildren(s)
	}
}

type ADDSUBContext struct {
	*CalculateStatementContext
	op antlr.Token
}

func NewADDSUBContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ADDSUBContext {
	var p = new(ADDSUBContext)

	p.CalculateStatementContext = NewEmptyCalculateStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateStatementContext))

	return p
}

func (s *ADDSUBContext) GetOp() antlr.Token { return s.op }

func (s *ADDSUBContext) SetOp(v antlr.Token) { s.op = v }

func (s *ADDSUBContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ADDSUBContext) AllCalculateStatement() []ICalculateStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem())
	var tst = make([]ICalculateStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICalculateStatementContext)
		}
	}

	return tst
}

func (s *ADDSUBContext) CalculateStatement(i int) ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *ADDSUBContext) AllADD() []antlr.TerminalNode {
	return s.GetTokens(RuleParserADD)
}

func (s *ADDSUBContext) ADD(i int) antlr.TerminalNode {
	return s.GetToken(RuleParserADD, i)
}

func (s *ADDSUBContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(RuleParserSUB)
}

func (s *ADDSUBContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(RuleParserSUB, i)
}

func (s *ADDSUBContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitADDSUB(s)

	default:
		return t.VisitChildren(s)
	}
}

type MULDIVContext struct {
	*CalculateStatementContext
	op antlr.Token
}

func NewMULDIVContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MULDIVContext {
	var p = new(MULDIVContext)

	p.CalculateStatementContext = NewEmptyCalculateStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateStatementContext))

	return p
}

func (s *MULDIVContext) GetOp() antlr.Token { return s.op }

func (s *MULDIVContext) SetOp(v antlr.Token) { s.op = v }

func (s *MULDIVContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MULDIVContext) AllCalculateStatement() []ICalculateStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem())
	var tst = make([]ICalculateStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICalculateStatementContext)
		}
	}

	return tst
}

func (s *MULDIVContext) CalculateStatement(i int) ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *MULDIVContext) AllMUL() []antlr.TerminalNode {
	return s.GetTokens(RuleParserMUL)
}

func (s *MULDIVContext) MUL(i int) antlr.TerminalNode {
	return s.GetToken(RuleParserMUL, i)
}

func (s *MULDIVContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(RuleParserDIV)
}

func (s *MULDIVContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(RuleParserDIV, i)
}

func (s *MULDIVContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitMULDIV(s)

	default:
		return t.VisitChildren(s)
	}
}

type CALCULATEXContext struct {
	*CalculateStatementContext
}

func NewCALCULATEXContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CALCULATEXContext {
	var p = new(CALCULATEXContext)

	p.CalculateStatementContext = NewEmptyCalculateStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateStatementContext))

	return p
}

func (s *CALCULATEXContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CALCULATEXContext) CalculateStatement() ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *CALCULATEXContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitCALCULATEX(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) CalculateStatement() (localctx ICalculateStatementContext) {
	return p.calculateStatement(0)
}

func (p *RuleParser) calculateStatement(_p int) (localctx ICalculateStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCalculateStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICalculateStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, RuleParserRULE_calculateStatement, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserNUM, RuleParserIDENTIFY:
		localctx = NewITEMCALCUContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(75)
			p.CalculateValue()
		}

	case RuleParserT__0:
		localctx = NewCALCULATEXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(76)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(77)
			p.calculateStatement(0)
		}
		{
			p.SetState(78)
			p.Match(RuleParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMULDIVContext(p, NewCalculateStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_calculateStatement)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(85)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(83)

							var _lt = p.GetTokenStream().LT(1)

							localctx.(*MULDIVContext).op = _lt

							_la = p.GetTokenStream().LA(1)

							if !(_la == RuleParserMUL || _la == RuleParserDIV) {
								var _ri = p.GetErrorHandler().RecoverInline(p)

								localctx.(*MULDIVContext).op = _ri
							} else {
								p.GetErrorHandler().ReportMatch(p)
								p.Consume()
							}
						}
						{
							p.SetState(84)
							p.calculateStatement(0)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(87)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
				}

			case 2:
				localctx = NewADDSUBContext(p, NewCalculateStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_calculateStatement)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(92)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(90)

							var _lt = p.GetTokenStream().LT(1)

							localctx.(*ADDSUBContext).op = _lt

							_la = p.GetTokenStream().LA(1)

							if !(_la == RuleParserADD || _la == RuleParserSUB) {
								var _ri = p.GetErrorHandler().RecoverInline(p)

								localctx.(*ADDSUBContext).op = _ri
							} else {
								p.GetErrorHandler().ReportMatch(p)
								p.Consume()
							}
						}
						{
							p.SetState(91)
							p.calculateStatement(0)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(94)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
				}

			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IBoolStatementContext is an interface to support dynamic dispatch.
type IBoolStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolStatementContext differentiates from other interfaces.
	IsBoolStatementContext()
}

type BoolStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolStatementContext() *BoolStatementContext {
	var p = new(BoolStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_boolStatement
	return p
}

func (*BoolStatementContext) IsBoolStatementContext() {}

func NewBoolStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolStatementContext {
	var p = new(BoolStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_boolStatement

	return p
}

func (s *BoolStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolStatementContext) CopyFrom(ctx *BoolStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BoolStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BOOLContext struct {
	*BoolStatementContext
}

func NewBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BOOLContext {
	var p = new(BOOLContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *BOOLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BOOLContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RuleParserTRUE, 0)
}

func (s *BOOLContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RuleParserFALSE, 0)
}

func (s *BOOLContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitBOOL(s)

	default:
		return t.VisitChildren(s)
	}
}

type BOOLOPContext struct {
	*BoolStatementContext
	op IBoolOperateContext
}

func NewBOOLOPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BOOLOPContext {
	var p = new(BOOLOPContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *BOOLOPContext) GetOp() IBoolOperateContext { return s.op }

func (s *BOOLOPContext) SetOp(v IBoolOperateContext) { s.op = v }

func (s *BOOLOPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BOOLOPContext) AllBoolStatement() []IBoolStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem())
	var tst = make([]IBoolStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolStatementContext)
		}
	}

	return tst
}

func (s *BOOLOPContext) BoolStatement(i int) IBoolStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *BOOLOPContext) BoolOperate() IBoolOperateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolOperateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolOperateContext)
}

func (s *BOOLOPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitBOOLOP(s)

	default:
		return t.VisitChildren(s)
	}
}

type IDENBOOLContext struct {
	*BoolStatementContext
}

func NewIDENBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDENBOOLContext {
	var p = new(IDENBOOLContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *IDENBOOLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDENBOOLContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *IDENBOOLContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitIDENBOOL(s)

	default:
		return t.VisitChildren(s)
	}
}

type COMPAREBOOLContext struct {
	*BoolStatementContext
}

func NewCOMPAREBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *COMPAREBOOLContext {
	var p = new(COMPAREBOOLContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *COMPAREBOOLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *COMPAREBOOLContext) CompareStatement() ICompareStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *COMPAREBOOLContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitCOMPAREBOOL(s)

	default:
		return t.VisitChildren(s)
	}
}

type BOOLOPXContext struct {
	*BoolStatementContext
}

func NewBOOLOPXContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BOOLOPXContext {
	var p = new(BOOLOPXContext)

	p.BoolStatementContext = NewEmptyBoolStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolStatementContext))

	return p
}

func (s *BOOLOPXContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BOOLOPXContext) BoolStatement() IBoolStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *BOOLOPXContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitBOOLOPX(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) BoolStatement() (localctx IBoolStatementContext) {
	return p.boolStatement(0)
}

func (p *RuleParser) boolStatement(_p int) (localctx IBoolStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoolStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, RuleParserRULE_boolStatement, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(102)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RuleParserTRUE || _la == RuleParserFALSE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		localctx = NewIDENBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(103)
			p.Match(RuleParserIDENTIFY)
		}

	case 3:
		localctx = NewCOMPAREBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(104)
			p.compareStatement(0)
		}

	case 4:
		localctx = NewBOOLOPXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(105)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(106)
			p.boolStatement(0)
		}
		{
			p.SetState(107)
			p.Match(RuleParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBOOLOPContext(p, NewBoolStatementContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_boolStatement)
			p.SetState(111)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(112)

				var _x = p.BoolOperate()

				localctx.(*BOOLOPContext).op = _x
			}
			{
				p.SetState(113)
				p.boolStatement(6)
			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IValueTypeContext is an interface to support dynamic dispatch.
type IValueTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueTypeContext differentiates from other interfaces.
	IsValueTypeContext()
}

type ValueTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueTypeContext() *ValueTypeContext {
	var p = new(ValueTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_valueType
	return p
}

func (*ValueTypeContext) IsValueTypeContext() {}

func NewValueTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueTypeContext {
	var p = new(ValueTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_valueType

	return p
}

func (s *ValueTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueTypeContext) BoolValue() IBoolValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolValueContext)
}

func (s *ValueTypeContext) Identify() IIdentifyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifyContext)
}

func (s *ValueTypeContext) Num() INumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *ValueTypeContext) CalculateStatement() ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *ValueTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitValueType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) ValueType() (localctx IValueTypeContext) {
	localctx = NewValueTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RuleParserRULE_valueType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.BoolValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Identify()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)
			p.Num()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			p.calculateStatement(0)
		}

	}

	return localctx
}

// ISetValueStatementContext is an interface to support dynamic dispatch.
type ISetValueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetValueStatementContext differentiates from other interfaces.
	IsSetValueStatementContext()
}

type SetValueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetValueStatementContext() *SetValueStatementContext {
	var p = new(SetValueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_setValueStatement
	return p
}

func (*SetValueStatementContext) IsSetValueStatementContext() {}

func NewSetValueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetValueStatementContext {
	var p = new(SetValueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_setValueStatement

	return p
}

func (s *SetValueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SetValueStatementContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *SetValueStatementContext) ValueType() IValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *SetValueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetValueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetValueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitSetValueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) SetValueStatement() (localctx ISetValueStatementContext) {
	localctx = NewSetValueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RuleParserRULE_setValueStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(RuleParserIDENTIFY)
	}
	{
		p.SetState(127)
		p.Match(RuleParserT__2)
	}
	{
		p.SetState(128)
		p.ValueType()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuleParserT__3 {
		{
			p.SetState(129)
			p.Match(RuleParserT__3)
		}

	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(RuleParserIF, 0)
}

func (s *IfStatementContext) BoolStatement() IBoolStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) AllElseIfStatement() []IElseIfStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfStatementContext)(nil)).Elem())
	var tst = make([]IElseIfStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) ElseIfStatement(i int) IElseIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfStatementContext)
}

func (s *IfStatementContext) ElseStatement() IElseStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RuleParserRULE_ifStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(RuleParserIF)
	}
	{
		p.SetState(133)
		p.boolStatement(0)
	}
	{
		p.SetState(134)
		p.Match(RuleParserT__4)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(135)
			p.Statement()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
		p.Match(RuleParserT__5)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RuleParserELSIF {
		{
			p.SetState(142)
			p.ElseIfStatement()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuleParserELSE {
		{
			p.SetState(148)
			p.ElseStatement()
		}

	}

	return localctx
}

// IElseIfStatementContext is an interface to support dynamic dispatch.
type IElseIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStatementContext differentiates from other interfaces.
	IsElseIfStatementContext()
}

type ElseIfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStatementContext() *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_elseIfStatement
	return p
}

func (*ElseIfStatementContext) IsElseIfStatementContext() {}

func NewElseIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_elseIfStatement

	return p
}

func (s *ElseIfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStatementContext) ELSIF() antlr.TerminalNode {
	return s.GetToken(RuleParserELSIF, 0)
}

func (s *ElseIfStatementContext) BoolStatement() IBoolStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *ElseIfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ElseIfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ElseIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitElseIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) ElseIfStatement() (localctx IElseIfStatementContext) {
	localctx = NewElseIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RuleParserRULE_elseIfStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(RuleParserELSIF)
	}
	{
		p.SetState(152)
		p.boolStatement(0)
	}
	{
		p.SetState(153)
		p.Match(RuleParserT__4)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(154)
			p.Statement()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(160)
		p.Match(RuleParserT__5)
	}

	return localctx
}

// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_elseStatement
	return p
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(RuleParserELSE, 0)
}

func (s *ElseStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ElseStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) ElseStatement() (localctx IElseStatementContext) {
	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RuleParserRULE_elseStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(RuleParserELSE)
	}
	{
		p.SetState(163)
		p.Match(RuleParserT__4)
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(164)
			p.Statement()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(170)
		p.Match(RuleParserT__5)
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value rule contexts.
	GetValue() IValueTypeContext

	// SetValue sets the value rule contexts.
	SetValue(IValueTypeContext)

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  IValueTypeContext
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) GetValue() IValueTypeContext { return s.value }

func (s *ReturnStatementContext) SetValue(v IValueTypeContext) { s.value = v }

func (s *ReturnStatementContext) ValueType() IValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *ReturnStatementContext) CalculateStatement() ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *ReturnStatementContext) BoolStatement() IBoolStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RuleParserRULE_returnStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Match(RuleParserT__6)
		}
		{
			p.SetState(173)

			var _x = p.ValueType()

			localctx.(*ReturnStatementContext).value = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.calculateStatement(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.boolStatement(0)
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CalculateStatement() ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *StatementContext) BoolStatement() IBoolStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *StatementContext) CompareStatement() ICompareStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) SetValueStatement() ISetValueStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetValueStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetValueStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RuleParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.calculateStatement(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.boolStatement(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(180)
			p.compareStatement(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(181)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(182)
			p.SetValueStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(183)
			p.ReturnStatement()
		}

	}

	return localctx
}

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_init
	return p
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *InitContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *InitContext) EOF() antlr.TerminalNode {
	return s.GetToken(RuleParserEOF, 0)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RuleParserRULE_init)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(186)
			p.Statement()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(192)
			p.Match(RuleParserEOF)
		}

	}

	return localctx
}

func (p *RuleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *CompareStatementContext = nil
		if localctx != nil {
			t = localctx.(*CompareStatementContext)
		}
		return p.CompareStatement_Sempred(t, predIndex)

	case 9:
		var t *CalculateStatementContext = nil
		if localctx != nil {
			t = localctx.(*CalculateStatementContext)
		}
		return p.CalculateStatement_Sempred(t, predIndex)

	case 10:
		var t *BoolStatementContext = nil
		if localctx != nil {
			t = localctx.(*BoolStatementContext)
		}
		return p.BoolStatement_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RuleParser) CompareStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleParser) CalculateStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleParser) BoolStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
