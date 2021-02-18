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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 81, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 5, 6, 40, 10, 6, 3, 7, 3, 7, 5, 7, 44, 10, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 54, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 60, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5,
	10, 70, 10, 10, 3, 11, 7, 11, 73, 10, 11, 12, 11, 14, 11, 76, 11, 11, 3,
	11, 5, 11, 79, 10, 11, 3, 11, 2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 2, 5, 3, 2, 7, 9, 3, 2, 16, 19, 3, 2, 11, 15, 2, 80, 2, 22, 3, 2, 2,
	2, 4, 24, 3, 2, 2, 2, 6, 26, 3, 2, 2, 2, 8, 28, 3, 2, 2, 2, 10, 39, 3,
	2, 2, 2, 12, 43, 3, 2, 2, 2, 14, 53, 3, 2, 2, 2, 16, 59, 3, 2, 2, 2, 18,
	69, 3, 2, 2, 2, 20, 74, 3, 2, 2, 2, 22, 23, 9, 2, 2, 2, 23, 3, 3, 2, 2,
	2, 24, 25, 9, 3, 2, 2, 25, 5, 3, 2, 2, 2, 26, 27, 9, 4, 2, 2, 27, 7, 3,
	2, 2, 2, 28, 29, 9, 2, 2, 2, 29, 9, 3, 2, 2, 2, 30, 31, 5, 12, 7, 2, 31,
	32, 5, 6, 4, 2, 32, 33, 5, 12, 7, 2, 33, 40, 3, 2, 2, 2, 34, 40, 5, 14,
	8, 2, 35, 36, 7, 3, 2, 2, 36, 37, 5, 10, 6, 2, 37, 38, 7, 4, 2, 2, 38,
	40, 3, 2, 2, 2, 39, 30, 3, 2, 2, 2, 39, 34, 3, 2, 2, 2, 39, 35, 3, 2, 2,
	2, 40, 11, 3, 2, 2, 2, 41, 44, 7, 10, 2, 2, 42, 44, 7, 5, 2, 2, 43, 41,
	3, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44, 13, 3, 2, 2, 2, 45, 46, 5, 12, 7, 2,
	46, 47, 5, 4, 3, 2, 47, 48, 5, 12, 7, 2, 48, 54, 3, 2, 2, 2, 49, 50, 7,
	3, 2, 2, 50, 51, 5, 14, 8, 2, 51, 52, 7, 4, 2, 2, 52, 54, 3, 2, 2, 2, 53,
	45, 3, 2, 2, 2, 53, 49, 3, 2, 2, 2, 54, 15, 3, 2, 2, 2, 55, 60, 7, 21,
	2, 2, 56, 60, 7, 22, 2, 2, 57, 60, 7, 10, 2, 2, 58, 60, 5, 10, 6, 2, 59,
	55, 3, 2, 2, 2, 59, 56, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 58, 3, 2, 2,
	2, 60, 17, 3, 2, 2, 2, 61, 62, 5, 16, 9, 2, 62, 63, 5, 2, 2, 2, 63, 64,
	5, 16, 9, 2, 64, 70, 3, 2, 2, 2, 65, 66, 7, 3, 2, 2, 66, 67, 5, 18, 10,
	2, 67, 68, 7, 4, 2, 2, 68, 70, 3, 2, 2, 2, 69, 61, 3, 2, 2, 2, 69, 65,
	3, 2, 2, 2, 70, 19, 3, 2, 2, 2, 71, 73, 5, 18, 10, 2, 72, 71, 3, 2, 2,
	2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 78,
	3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 79, 7, 2, 2, 3, 78, 77, 3, 2, 2, 2,
	78, 79, 3, 2, 2, 2, 79, 21, 3, 2, 2, 2, 9, 39, 43, 53, 59, 69, 74, 78,
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
	"calculateStatement", "boolValue", "boolStatement", "init",
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
	RuleParserRULE_boolValue          = 7
	RuleParserRULE_boolStatement      = 8
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

func (s *COMPAREXContext) CalculateStatement() ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
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

func (s *COMPAREContext) AllCalculateValue() []ICalculateValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICalculateValueContext)(nil)).Elem())
	var tst = make([]ICalculateValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICalculateValueContext)
		}
	}

	return tst
}

func (s *COMPAREContext) CalculateValue(i int) ICalculateValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICalculateValueContext)
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

func (p *RuleParser) CompareStatement() (localctx ICompareStatementContext) {
	localctx = NewCompareStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RuleParserRULE_compareStatement)

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

	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCOMPAREContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.CalculateValue()
		}
		{
			p.SetState(29)

			var _x = p.Compare()

			localctx.(*COMPAREContext).op = _x
		}
		{
			p.SetState(30)
			p.CalculateValue()
		}

	case 2:
		localctx = NewCOMPAREXContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.CalculateStatement()
		}

	case 3:
		localctx = NewCOMPAREXContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(34)
			p.CompareStatement()
		}
		{
			p.SetState(35)
			p.Match(RuleParserT__1)
		}

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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserIDENTIFY:
		localctx = NewIDENContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Match(RuleParserIDENTIFY)
		}

	case RuleParserNUM:
		localctx = NewNUMContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
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

	// GetOp returns the op rule contexts.
	GetOp() ICalculateContext

	// SetOp sets the op rule contexts.
	SetOp(ICalculateContext)

	// IsCalculateStatementContext differentiates from other interfaces.
	IsCalculateStatementContext()
}

type CalculateStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     ICalculateContext
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

func (s *CalculateStatementContext) GetOp() ICalculateContext { return s.op }

func (s *CalculateStatementContext) SetOp(v ICalculateContext) { s.op = v }

func (s *CalculateStatementContext) AllCalculateValue() []ICalculateValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICalculateValueContext)(nil)).Elem())
	var tst = make([]ICalculateValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICalculateValueContext)
		}
	}

	return tst
}

func (s *CalculateStatementContext) CalculateValue(i int) ICalculateValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICalculateValueContext)
}

func (s *CalculateStatementContext) Calculate() ICalculateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateContext)
}

func (s *CalculateStatementContext) CalculateStatement() ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *CalculateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalculateStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCalculateStatement(s)
	}
}

func (s *CalculateStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCalculateStatement(s)
	}
}

func (p *RuleParser) CalculateStatement() (localctx ICalculateStatementContext) {
	localctx = NewCalculateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RuleParserRULE_calculateStatement)

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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserNUM, RuleParserIDENTIFY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.CalculateValue()
		}
		{
			p.SetState(44)

			var _x = p.Calculate()

			localctx.(*CalculateStatementContext).op = _x
		}
		{
			p.SetState(45)
			p.CalculateValue()
		}

	case RuleParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(48)
			p.CalculateStatement()
		}
		{
			p.SetState(49)
			p.Match(RuleParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *BoolValueContext) CopyFrom(ctx *BoolValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type COMPAREVALUEContext struct {
	*BoolValueContext
}

func NewCOMPAREVALUEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *COMPAREVALUEContext {
	var p = new(COMPAREVALUEContext)

	p.BoolValueContext = NewEmptyBoolValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolValueContext))

	return p
}

func (s *COMPAREVALUEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *COMPAREVALUEContext) CompareStatement() ICompareStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *COMPAREVALUEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCOMPAREVALUE(s)
	}
}

func (s *COMPAREVALUEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCOMPAREVALUE(s)
	}
}

type BOOLContext struct {
	*BoolValueContext
}

func NewBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BOOLContext {
	var p = new(BOOLContext)

	p.BoolValueContext = NewEmptyBoolValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolValueContext))

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

type IDENBOOLContext struct {
	*BoolValueContext
}

func NewIDENBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IDENBOOLContext {
	var p = new(IDENBOOLContext)

	p.BoolValueContext = NewEmptyBoolValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolValueContext))

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

func (p *RuleParser) BoolValue() (localctx IBoolValueContext) {
	localctx = NewBoolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RuleParserRULE_boolValue)

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

	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBOOLContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Match(RuleParserTRUE)
		}

	case 2:
		localctx = NewBOOLContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(RuleParserFALSE)
		}

	case 3:
		localctx = NewIDENBOOLContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Match(RuleParserIDENTIFY)
		}

	case 4:
		localctx = NewCOMPAREVALUEContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.CompareStatement()
		}

	}

	return localctx
}

// IBoolStatementContext is an interface to support dynamic dispatch.
type IBoolStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op rule contexts.
	GetOp() IBoolOperateContext

	// SetOp sets the op rule contexts.
	SetOp(IBoolOperateContext)

	// IsBoolStatementContext differentiates from other interfaces.
	IsBoolStatementContext()
}

type BoolStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     IBoolOperateContext
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

func (s *BoolStatementContext) GetOp() IBoolOperateContext { return s.op }

func (s *BoolStatementContext) SetOp(v IBoolOperateContext) { s.op = v }

func (s *BoolStatementContext) AllBoolValue() []IBoolValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolValueContext)(nil)).Elem())
	var tst = make([]IBoolValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolValueContext)
		}
	}

	return tst
}

func (s *BoolStatementContext) BoolValue(i int) IBoolValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolValueContext)
}

func (s *BoolStatementContext) BoolOperate() IBoolOperateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolOperateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolOperateContext)
}

func (s *BoolStatementContext) BoolStatement() IBoolStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
}

func (s *BoolStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterBoolStatement(s)
	}
}

func (s *BoolStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitBoolStatement(s)
	}
}

func (p *RuleParser) BoolStatement() (localctx IBoolStatementContext) {
	localctx = NewBoolStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RuleParserRULE_boolStatement)

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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.BoolValue()
		}
		{
			p.SetState(60)

			var _x = p.BoolOperate()

			localctx.(*BoolStatementContext).op = _x
		}
		{
			p.SetState(61)
			p.BoolValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(64)
			p.BoolStatement()
		}
		{
			p.SetState(65)
			p.Match(RuleParserT__1)
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

func (s *InitContext) AllBoolStatement() []IBoolStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem())
	var tst = make([]IBoolStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolStatementContext)
		}
	}

	return tst
}

func (s *InitContext) BoolStatement(i int) IBoolStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolStatementContext)
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
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserNUM)|(1<<RuleParserIDENTIFY)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE))) != 0 {
		{
			p.SetState(69)
			p.BoolStatement()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(75)
			p.Match(RuleParserEOF)
		}

	}

	return localctx
}
