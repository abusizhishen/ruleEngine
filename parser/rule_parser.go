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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 22, 72, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 5, 5, 31, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 37,
	10, 5, 12, 5, 14, 5, 40, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 5, 6, 49, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 55, 10, 6, 12, 6, 14,
	6, 58, 11, 6, 3, 7, 3, 7, 5, 7, 62, 10, 7, 3, 8, 7, 8, 65, 10, 8, 12, 8,
	14, 8, 68, 11, 8, 3, 8, 3, 8, 3, 8, 2, 4, 8, 10, 9, 2, 4, 6, 8, 10, 12,
	14, 2, 5, 3, 2, 16, 19, 3, 2, 11, 15, 3, 2, 8, 10, 2, 73, 2, 16, 3, 2,
	2, 2, 4, 18, 3, 2, 2, 2, 6, 20, 3, 2, 2, 2, 8, 30, 3, 2, 2, 2, 10, 48,
	3, 2, 2, 2, 12, 61, 3, 2, 2, 2, 14, 66, 3, 2, 2, 2, 16, 17, 9, 2, 2, 2,
	17, 3, 3, 2, 2, 2, 18, 19, 9, 3, 2, 2, 19, 5, 3, 2, 2, 2, 20, 21, 9, 4,
	2, 2, 21, 7, 3, 2, 2, 2, 22, 23, 8, 5, 1, 2, 23, 31, 7, 5, 2, 2, 24, 31,
	7, 21, 2, 2, 25, 31, 7, 22, 2, 2, 26, 27, 7, 3, 2, 2, 27, 28, 5, 8, 5,
	2, 28, 29, 7, 4, 2, 2, 29, 31, 3, 2, 2, 2, 30, 22, 3, 2, 2, 2, 30, 24,
	3, 2, 2, 2, 30, 25, 3, 2, 2, 2, 30, 26, 3, 2, 2, 2, 31, 38, 3, 2, 2, 2,
	32, 33, 12, 7, 2, 2, 33, 34, 5, 4, 3, 2, 34, 35, 5, 8, 5, 8, 35, 37, 3,
	2, 2, 2, 36, 32, 3, 2, 2, 2, 37, 40, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38,
	39, 3, 2, 2, 2, 39, 9, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 41, 42, 8, 6, 1,
	2, 42, 49, 7, 5, 2, 2, 43, 49, 7, 6, 2, 2, 44, 45, 7, 3, 2, 2, 45, 46,
	5, 10, 6, 2, 46, 47, 7, 4, 2, 2, 47, 49, 3, 2, 2, 2, 48, 41, 3, 2, 2, 2,
	48, 43, 3, 2, 2, 2, 48, 44, 3, 2, 2, 2, 49, 56, 3, 2, 2, 2, 50, 51, 12,
	6, 2, 2, 51, 52, 5, 2, 2, 2, 52, 53, 5, 10, 6, 7, 53, 55, 3, 2, 2, 2, 54,
	50, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2,
	2, 57, 11, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 62, 5, 10, 6, 2, 60, 62,
	5, 8, 5, 2, 61, 59, 3, 2, 2, 2, 61, 60, 3, 2, 2, 2, 62, 13, 3, 2, 2, 2,
	63, 65, 5, 12, 7, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3,
	2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69,
	70, 7, 2, 2, 3, 70, 15, 3, 2, 2, 2, 8, 30, 38, 48, 56, 61, 66,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "", "", "'now()'", "'and'", "'or'", "'not'", "'>'", "'>='",
	"'=='", "'<='", "'<'", "'+'", "'-'", "'*'", "'/'", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "", "", "IDENTIFY", "NUM", "NOW", "AND", "OR", "NOT", "GT", "GTE",
	"EQ", "LTE", "LT", "ADD", "SUB", "MUL", "DIV", "WS", "TRUE", "FALSE",
}

var ruleNames = []string{
	"calculate", "compare", "logical", "compareStatement", "calculateStatement",
	"statement", "init",
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
	RuleParserIDENTIFY = 3
	RuleParserNUM      = 4
	RuleParserNOW      = 5
	RuleParserAND      = 6
	RuleParserOR       = 7
	RuleParserNOT      = 8
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
	RuleParserRULE_calculate          = 0
	RuleParserRULE_compare            = 1
	RuleParserRULE_logical            = 2
	RuleParserRULE_compareStatement   = 3
	RuleParserRULE_calculateStatement = 4
	RuleParserRULE_statement          = 5
	RuleParserRULE_init               = 6
)

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
	p.EnterRule(localctx, 0, RuleParserRULE_calculate)
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
		p.SetState(14)
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
	p.EnterRule(localctx, 2, RuleParserRULE_compare)
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
		p.SetState(16)
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
	p.EnterRule(localctx, 4, RuleParserRULE_logical)
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
		p.SetState(18)
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

	// GetOp returns the op rule contexts.
	GetOp() ICompareContext

	// SetOp sets the op rule contexts.
	SetOp(ICompareContext)

	// IsCompareStatementContext differentiates from other interfaces.
	IsCompareStatementContext()
}

type CompareStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     ICompareContext
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

func (s *CompareStatementContext) GetOp() ICompareContext { return s.op }

func (s *CompareStatementContext) SetOp(v ICompareContext) { s.op = v }

func (s *CompareStatementContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *CompareStatementContext) TRUE() antlr.TerminalNode {
	return s.GetToken(RuleParserTRUE, 0)
}

func (s *CompareStatementContext) FALSE() antlr.TerminalNode {
	return s.GetToken(RuleParserFALSE, 0)
}

func (s *CompareStatementContext) AllCompareStatement() []ICompareStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICompareStatementContext)(nil)).Elem())
	var tst = make([]ICompareStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICompareStatementContext)
		}
	}

	return tst
}

func (s *CompareStatementContext) CompareStatement(i int) ICompareStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
}

func (s *CompareStatementContext) Compare() ICompareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareContext)
}

func (s *CompareStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.EnterCompareStatement(s)
	}
}

func (s *CompareStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuleListener); ok {
		listenerT.ExitCompareStatement(s)
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
	_startState := 6
	p.EnterRecursionRule(localctx, 6, RuleParserRULE_compareStatement, _p)

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
	p.SetState(28)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserIDENTIFY:
		{
			p.SetState(21)
			p.Match(RuleParserIDENTIFY)
		}

	case RuleParserTRUE:
		{
			p.SetState(22)
			p.Match(RuleParserTRUE)
		}

	case RuleParserFALSE:
		{
			p.SetState(23)
			p.Match(RuleParserFALSE)
		}

	case RuleParserT__0:
		{
			p.SetState(24)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(25)
			p.compareStatement(0)
		}
		{
			p.SetState(26)
			p.Match(RuleParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCompareStatementContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_compareStatement)
			p.SetState(30)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(31)

				var _x = p.Compare()

				localctx.(*CompareStatementContext).op = _x
			}
			{
				p.SetState(32)
				p.compareStatement(6)
			}

		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
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

func (s *CalculateStatementContext) IDENTIFY() antlr.TerminalNode {
	return s.GetToken(RuleParserIDENTIFY, 0)
}

func (s *CalculateStatementContext) NUM() antlr.TerminalNode {
	return s.GetToken(RuleParserNUM, 0)
}

func (s *CalculateStatementContext) AllCalculateStatement() []ICalculateStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem())
	var tst = make([]ICalculateStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICalculateStatementContext)
		}
	}

	return tst
}

func (s *CalculateStatementContext) CalculateStatement(i int) ICalculateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICalculateStatementContext)
}

func (s *CalculateStatementContext) Calculate() ICalculateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalculateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalculateContext)
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
	return p.calculateStatement(0)
}

func (p *RuleParser) calculateStatement(_p int) (localctx ICalculateStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCalculateStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICalculateStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, RuleParserRULE_calculateStatement, _p)

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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserIDENTIFY:
		{
			p.SetState(40)
			p.Match(RuleParserIDENTIFY)
		}

	case RuleParserNUM:
		{
			p.SetState(41)
			p.Match(RuleParserNUM)
		}

	case RuleParserT__0:
		{
			p.SetState(42)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(43)
			p.calculateStatement(0)
		}
		{
			p.SetState(44)
			p.Match(RuleParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCalculateStatementContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_calculateStatement)
			p.SetState(48)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(49)

				var _x = p.Calculate()

				localctx.(*CalculateStatementContext).op = _x
			}
			{
				p.SetState(50)
				p.calculateStatement(5)
			}

		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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

func (s *StatementContext) CompareStatement() ICompareStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareStatementContext)
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
	p.EnterRule(localctx, 10, RuleParserRULE_statement)

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

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.calculateStatement(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.compareStatement(0)
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

func (s *InitContext) EOF() antlr.TerminalNode {
	return s.GetToken(RuleParserEOF, 0)
}

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
	p.EnterRule(localctx, 12, RuleParserRULE_init)
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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserIDENTIFY)|(1<<RuleParserNUM)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE))) != 0 {
		{
			p.SetState(61)
			p.Statement()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(RuleParserEOF)
	}

	return localctx
}

func (p *RuleParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *CompareStatementContext = nil
		if localctx != nil {
			t = localctx.(*CompareStatementContext)
		}
		return p.CompareStatement_Sempred(t, predIndex)

	case 4:
		var t *CalculateStatementContext = nil
		if localctx != nil {
			t = localctx.(*CalculateStatementContext)
		}
		return p.CalculateStatement_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RuleParser) CompareStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuleParser) CalculateStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
