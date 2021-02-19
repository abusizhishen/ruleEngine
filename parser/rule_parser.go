// Code generated from Rule.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 113,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 38, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 44, 10, 6, 12, 6, 14, 6, 47,
	11, 6, 3, 7, 3, 7, 5, 7, 51, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	5, 8, 59, 10, 8, 3, 8, 3, 8, 3, 8, 6, 8, 64, 10, 8, 13, 8, 14, 8, 65, 3,
	8, 3, 8, 3, 8, 6, 8, 71, 10, 8, 13, 8, 14, 8, 72, 7, 8, 75, 10, 8, 12,
	8, 14, 8, 78, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 88, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 94, 10, 9, 12, 9, 14, 9, 97,
	11, 9, 3, 10, 3, 10, 3, 10, 5, 10, 102, 10, 10, 3, 11, 7, 11, 105, 10,
	11, 12, 11, 14, 11, 108, 11, 11, 3, 11, 5, 11, 111, 10, 11, 3, 11, 2, 5,
	10, 14, 16, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 8, 3, 2, 7, 9, 3,
	2, 16, 19, 3, 2, 11, 15, 3, 2, 18, 19, 3, 2, 16, 17, 3, 2, 21, 22, 2, 119,
	2, 22, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 28, 3, 2, 2,
	2, 10, 37, 3, 2, 2, 2, 12, 50, 3, 2, 2, 2, 14, 58, 3, 2, 2, 2, 16, 87,
	3, 2, 2, 2, 18, 101, 3, 2, 2, 2, 20, 106, 3, 2, 2, 2, 22, 23, 9, 2, 2,
	2, 23, 3, 3, 2, 2, 2, 24, 25, 9, 3, 2, 2, 25, 5, 3, 2, 2, 2, 26, 27, 9,
	4, 2, 2, 27, 7, 3, 2, 2, 2, 28, 29, 9, 2, 2, 2, 29, 9, 3, 2, 2, 2, 30,
	31, 8, 6, 1, 2, 31, 38, 5, 12, 7, 2, 32, 38, 5, 14, 8, 2, 33, 34, 7, 3,
	2, 2, 34, 35, 5, 10, 6, 2, 35, 36, 7, 4, 2, 2, 36, 38, 3, 2, 2, 2, 37,
	30, 3, 2, 2, 2, 37, 32, 3, 2, 2, 2, 37, 33, 3, 2, 2, 2, 38, 45, 3, 2, 2,
	2, 39, 40, 12, 6, 2, 2, 40, 41, 5, 6, 4, 2, 41, 42, 5, 10, 6, 7, 42, 44,
	3, 2, 2, 2, 43, 39, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2,
	45, 46, 3, 2, 2, 2, 46, 11, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 51, 7,
	10, 2, 2, 49, 51, 7, 5, 2, 2, 50, 48, 3, 2, 2, 2, 50, 49, 3, 2, 2, 2, 51,
	13, 3, 2, 2, 2, 52, 53, 8, 8, 1, 2, 53, 59, 5, 12, 7, 2, 54, 55, 7, 3,
	2, 2, 55, 56, 5, 14, 8, 2, 56, 57, 7, 4, 2, 2, 57, 59, 3, 2, 2, 2, 58,
	52, 3, 2, 2, 2, 58, 54, 3, 2, 2, 2, 59, 76, 3, 2, 2, 2, 60, 63, 12, 6,
	2, 2, 61, 62, 9, 5, 2, 2, 62, 64, 5, 14, 8, 2, 63, 61, 3, 2, 2, 2, 64,
	65, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 75, 3, 2, 2,
	2, 67, 70, 12, 5, 2, 2, 68, 69, 9, 6, 2, 2, 69, 71, 5, 14, 8, 2, 70, 68,
	3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2,
	73, 75, 3, 2, 2, 2, 74, 60, 3, 2, 2, 2, 74, 67, 3, 2, 2, 2, 75, 78, 3,
	2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 15, 3, 2, 2, 2, 78,
	76, 3, 2, 2, 2, 79, 80, 8, 9, 1, 2, 80, 88, 9, 7, 2, 2, 81, 88, 7, 10,
	2, 2, 82, 88, 5, 10, 6, 2, 83, 84, 7, 3, 2, 2, 84, 85, 5, 16, 9, 2, 85,
	86, 7, 4, 2, 2, 86, 88, 3, 2, 2, 2, 87, 79, 3, 2, 2, 2, 87, 81, 3, 2, 2,
	2, 87, 82, 3, 2, 2, 2, 87, 83, 3, 2, 2, 2, 88, 95, 3, 2, 2, 2, 89, 90,
	12, 7, 2, 2, 90, 91, 5, 2, 2, 2, 91, 92, 5, 16, 9, 8, 92, 94, 3, 2, 2,
	2, 93, 89, 3, 2, 2, 2, 94, 97, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96,
	3, 2, 2, 2, 96, 17, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 102, 5, 16, 9,
	2, 99, 102, 5, 10, 6, 2, 100, 102, 5, 14, 8, 2, 101, 98, 3, 2, 2, 2, 101,
	99, 3, 2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 19, 3, 2, 2, 2, 103, 105, 5,
	18, 10, 2, 104, 103, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2,
	2, 2, 106, 107, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2,
	109, 111, 7, 2, 2, 3, 110, 109, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111,
	21, 3, 2, 2, 2, 15, 37, 45, 50, 58, 65, 72, 74, 76, 87, 95, 101, 106, 110,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "", "'now()'", "'and'", "'or'", "'not'", "", "'>'", "'>='",
	"'=='", "'<='", "'<'", "'+'", "'-'", "'*'", "'/'", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "NUM", "NOW", "AND", "OR", "NOT", "IDENTIFY", "GT", "GTE",
	"EQ", "LTE", "LT", "ADD", "SUB", "MUL", "DIV", "WS", "TRUE", "FALSE",
}

var ruleNames = []string{
	"boolOperate", "calculate", "compare", "logical", "compareStatement", "calculateValue",
	"calculateStatement", "boolStatement", "statement", "init",
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
	RuleParserNUM      = 3
	RuleParserNOW      = 4
	RuleParserAND      = 5
	RuleParserOR       = 6
	RuleParserNOT      = 7
	RuleParserIDENTIFY = 8
	RuleParserGT       = 9
	RuleParserGTE      = 10
	RuleParserEQ       = 11
	RuleParserLTE      = 12
	RuleParserLT       = 13
	RuleParserADD      = 14
	RuleParserSUB      = 15
	RuleParserMUL      = 16
	RuleParserDIV      = 17
	RuleParserWS       = 18
	RuleParserTRUE     = 19
	RuleParserFALSE    = 20
)

// RuleParser rules.
const (
	RuleParserRULE_boolOperate        = 0
	RuleParserRULE_calculate          = 1
	RuleParserRULE_compare            = 2
	RuleParserRULE_logical            = 3
	RuleParserRULE_compareStatement   = 4
	RuleParserRULE_calculateValue     = 5
	RuleParserRULE_calculateStatement = 6
	RuleParserRULE_boolStatement      = 7
	RuleParserRULE_statement          = 8
	RuleParserRULE_init               = 9
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

func (s *BoolOperateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBoolOperate(s)
	}
}

func (s *BoolOperateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBoolOperate(s)
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
		p.SetState(20)
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

func (s *CalculateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCalculate(s)
	}
}

func (s *CalculateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCalculate(s)
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
		p.SetState(22)
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

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCompare(s)
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
		p.SetState(24)
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

func (s *LogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterLogical(s)
	}
}

func (s *LogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitLogical(s)
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
		p.SetState(26)
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

func (s *COMPAREXContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCOMPAREX(s)
	}
}

func (s *COMPAREXContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCOMPAREX(s)
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

func (s *COMPAREContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCOMPARE(s)
	}
}

func (s *COMPAREContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCOMPARE(s)
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

func (s *CalcuContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCalcu(s)
	}
}

func (s *CalcuContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCalcu(s)
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

func (s *ITEMCOMPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterITEMCOMP(s)
	}
}

func (s *ITEMCOMPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitITEMCOMP(s)
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewITEMCOMPContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(29)
			p.CalculateValue()
		}

	case 2:
		localctx = NewCalcuContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.calculateStatement(0)
		}

	case 3:
		localctx = NewCOMPAREXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(32)
			p.compareStatement(0)
		}
		{
			p.SetState(33)
			p.Match(RuleParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(43)
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
			p.SetState(37)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(38)

				var _x = p.Compare()

				localctx.(*COMPAREContext).op = _x
			}
			{
				p.SetState(39)
				p.compareStatement(5)
			}

		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
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

func (s *CalculateValueContext) CopyFrom(ctx *CalculateValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CalculateValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IDENContext struct {
	*CalculateValueContext
}

func NewIDENContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDENContext {
	var p = new(IDENContext)

	p.CalculateValueContext = NewEmptyCalculateValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateValueContext))

	return p
}

func (s *IDENContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDENContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *IDENContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterIDEN(s)
	}
}

func (s *IDENContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitIDEN(s)
	}
}

type NUMContext struct {
	*CalculateValueContext
}

func NewNUMContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NUMContext {
	var p = new(NUMContext)

	p.CalculateValueContext = NewEmptyCalculateValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CalculateValueContext))

	return p
}

func (s *NUMContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NUMContext) NUM() antlr.TerminalNode {
	return s.GetToken(RuleParserNUM, 0)
}

func (s *NUMContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterNUM(s)
	}
}

func (s *NUMContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitNUM(s)
	}
}

func (p *RuleParser) CalculateValue() (localctx ICalculateValueContext) {
	localctx = NewCalculateValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RuleParserRULE_calculateValue)

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

	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserIDENTIFY:
		localctx = NewIDENContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Match(RuleParserIDENTIFY)
		}

	case RuleParserNUM:
		localctx = NewNUMContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Match(RuleParserNUM)
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

func (s *ITEMCALCUContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterITEMCALCU(s)
	}
}

func (s *ITEMCALCUContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitITEMCALCU(s)
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

func (s *ADDSUBContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterADDSUB(s)
	}
}

func (s *ADDSUBContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitADDSUB(s)
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

func (s *MULDIVContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterMULDIV(s)
	}
}

func (s *MULDIVContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitMULDIV(s)
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

func (s *CALCULATEXContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCALCULATEX(s)
	}
}

func (s *CALCULATEXContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCALCULATEX(s)
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
	_startState := 12
	p.EnterRecursionRule(localctx, 12, RuleParserRULE_calculateStatement, _p)
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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserNUM, RuleParserIDENTIFY:
		localctx = NewITEMCALCUContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(51)
			p.CalculateValue()
		}

	case RuleParserT__0:
		localctx = NewCALCULATEXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(53)
			p.calculateStatement(0)
		}
		{
			p.SetState(54)
			p.Match(RuleParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMULDIVContext(p, NewCalculateStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_calculateStatement)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(61)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(59)

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
							p.SetState(60)
							p.calculateStatement(0)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(63)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
				}

			case 2:
				localctx = NewADDSUBContext(p, NewCalculateStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_calculateStatement)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(68)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(66)

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
							p.SetState(67)
							p.calculateStatement(0)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(70)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
				}

			}

		}
		p.SetState(76)
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

func (s *BOOLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBOOL(s)
	}
}

func (s *BOOLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBOOL(s)
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

func (s *BOOLOPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBOOLOP(s)
	}
}

func (s *BOOLOPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBOOLOP(s)
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

func (s *IDENBOOLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterIDENBOOL(s)
	}
}

func (s *IDENBOOLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitIDENBOOL(s)
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

func (s *COMPAREBOOLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCOMPAREBOOL(s)
	}
}

func (s *COMPAREBOOLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCOMPAREBOOL(s)
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

func (s *BOOLOPXContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBOOLOPX(s)
	}
}

func (s *BOOLOPXContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBOOLOPX(s)
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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, RuleParserRULE_boolStatement, _p)
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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(78)
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
			p.SetState(79)
			p.Match(RuleParserIDENTIFY)
		}

	case 3:
		localctx = NewCOMPAREBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)
			p.compareStatement(0)
		}

	case 4:
		localctx = NewBOOLOPXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(81)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(82)
			p.boolStatement(0)
		}
		{
			p.SetState(83)
			p.Match(RuleParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(93)
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
			p.SetState(87)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(88)

				var _x = p.BoolOperate()

				localctx.(*BOOLOPContext).op = _x
			}
			{
				p.SetState(89)
				p.boolStatement(6)
			}

		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
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

func (s *StatementContext) CalculateStatement() ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *RuleParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RuleParserRULE_statement)

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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.boolStatement(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.compareStatement(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.calculateStatement(0)
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

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *RuleParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RuleParserRULE_init)
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
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserNUM)|(1<<RuleParserIDENTIFY)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE))) != 0 {
		{
			p.SetState(101)
			p.Statement()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(107)
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

	case 6:
		var t *CalculateStatementContext = nil
		if localctx != nil {
			t = localctx.(*CalculateStatementContext)
		}
		return p.CalculateStatement_Sempred(t, predIndex)

	case 7:
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
