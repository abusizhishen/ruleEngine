// Code generated from Rule.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 30, 155,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9,
	80, 10, 9, 13, 9, 14, 9, 81, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 17, 6, 17, 113, 10, 17, 13, 17, 14, 17, 114, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 6, 27, 139,
	10, 27, 13, 27, 14, 27, 140, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 2, 2, 30, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14,
	27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23,
	45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 3, 2, 5, 3, 2,
	50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	157, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2,
	2, 2, 2, 57, 3, 2, 2, 2, 3, 59, 3, 2, 2, 2, 5, 61, 3, 2, 2, 2, 7, 63, 3,
	2, 2, 2, 9, 65, 3, 2, 2, 2, 11, 67, 3, 2, 2, 2, 13, 69, 3, 2, 2, 2, 15,
	71, 3, 2, 2, 2, 17, 79, 3, 2, 2, 2, 19, 83, 3, 2, 2, 2, 21, 89, 3, 2, 2,
	2, 23, 92, 3, 2, 2, 2, 25, 95, 3, 2, 2, 2, 27, 97, 3, 2, 2, 2, 29, 100,
	3, 2, 2, 2, 31, 106, 3, 2, 2, 2, 33, 112, 3, 2, 2, 2, 35, 116, 3, 2, 2,
	2, 37, 118, 3, 2, 2, 2, 39, 121, 3, 2, 2, 2, 41, 124, 3, 2, 2, 2, 43, 127,
	3, 2, 2, 2, 45, 129, 3, 2, 2, 2, 47, 131, 3, 2, 2, 2, 49, 133, 3, 2, 2,
	2, 51, 135, 3, 2, 2, 2, 53, 138, 3, 2, 2, 2, 55, 144, 3, 2, 2, 2, 57, 149,
	3, 2, 2, 2, 59, 60, 7, 42, 2, 2, 60, 4, 3, 2, 2, 2, 61, 62, 7, 43, 2, 2,
	62, 6, 3, 2, 2, 2, 63, 64, 7, 63, 2, 2, 64, 8, 3, 2, 2, 2, 65, 66, 7, 61,
	2, 2, 66, 10, 3, 2, 2, 2, 67, 68, 7, 125, 2, 2, 68, 12, 3, 2, 2, 2, 69,
	70, 7, 127, 2, 2, 70, 14, 3, 2, 2, 2, 71, 72, 7, 116, 2, 2, 72, 73, 7,
	103, 2, 2, 73, 74, 7, 118, 2, 2, 74, 75, 7, 119, 2, 2, 75, 76, 7, 116,
	2, 2, 76, 77, 7, 112, 2, 2, 77, 16, 3, 2, 2, 2, 78, 80, 9, 2, 2, 2, 79,
	78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2,
	2, 82, 18, 3, 2, 2, 2, 83, 84, 7, 112, 2, 2, 84, 85, 7, 113, 2, 2, 85,
	86, 7, 121, 2, 2, 86, 87, 7, 42, 2, 2, 87, 88, 7, 43, 2, 2, 88, 20, 3,
	2, 2, 2, 89, 90, 7, 40, 2, 2, 90, 91, 7, 40, 2, 2, 91, 22, 3, 2, 2, 2,
	92, 93, 7, 126, 2, 2, 93, 94, 7, 126, 2, 2, 94, 24, 3, 2, 2, 2, 95, 96,
	7, 35, 2, 2, 96, 26, 3, 2, 2, 2, 97, 98, 7, 107, 2, 2, 98, 99, 7, 104,
	2, 2, 99, 28, 3, 2, 2, 2, 100, 101, 7, 103, 2, 2, 101, 102, 7, 110, 2,
	2, 102, 103, 7, 117, 2, 2, 103, 104, 7, 107, 2, 2, 104, 105, 7, 104, 2,
	2, 105, 30, 3, 2, 2, 2, 106, 107, 7, 103, 2, 2, 107, 108, 7, 110, 2, 2,
	108, 109, 7, 117, 2, 2, 109, 110, 7, 103, 2, 2, 110, 32, 3, 2, 2, 2, 111,
	113, 9, 3, 2, 2, 112, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 112,
	3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 34, 3, 2, 2, 2, 116, 117, 7, 64,
	2, 2, 117, 36, 3, 2, 2, 2, 118, 119, 7, 64, 2, 2, 119, 120, 7, 63, 2, 2,
	120, 38, 3, 2, 2, 2, 121, 122, 7, 63, 2, 2, 122, 123, 7, 63, 2, 2, 123,
	40, 3, 2, 2, 2, 124, 125, 7, 62, 2, 2, 125, 126, 7, 63, 2, 2, 126, 42,
	3, 2, 2, 2, 127, 128, 7, 62, 2, 2, 128, 44, 3, 2, 2, 2, 129, 130, 7, 45,
	2, 2, 130, 46, 3, 2, 2, 2, 131, 132, 7, 47, 2, 2, 132, 48, 3, 2, 2, 2,
	133, 134, 7, 44, 2, 2, 134, 50, 3, 2, 2, 2, 135, 136, 7, 49, 2, 2, 136,
	52, 3, 2, 2, 2, 137, 139, 9, 4, 2, 2, 138, 137, 3, 2, 2, 2, 139, 140, 3,
	2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 142, 3, 2, 2,
	2, 142, 143, 8, 27, 2, 2, 143, 54, 3, 2, 2, 2, 144, 145, 7, 118, 2, 2,
	145, 146, 7, 116, 2, 2, 146, 147, 7, 119, 2, 2, 147, 148, 7, 103, 2, 2,
	148, 56, 3, 2, 2, 2, 149, 150, 7, 104, 2, 2, 150, 151, 7, 99, 2, 2, 151,
	152, 7, 110, 2, 2, 152, 153, 7, 117, 2, 2, 153, 154, 7, 103, 2, 2, 154,
	58, 3, 2, 2, 2, 6, 2, 81, 114, 140, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'='", "';'", "'{'", "'}'", "'return'", "", "'now()'",
	"'&&'", "'||'", "'!'", "'if'", "'elsif'", "'else'", "", "'>'", "'>='",
	"'=='", "'<='", "'<'", "'+'", "'-'", "'*'", "'/'", "", "'true'", "'false'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "NUM", "NOW", "AND", "OR", "NOT", "IF",
	"ELSIF", "ELSE", "IDENTIFY", "GT", "GTE", "EQ", "LTE", "LT", "ADD", "SUB",
	"MUL", "DIV", "WS", "TRUE", "FALSE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "NUM", "NOW", "AND",
	"OR", "NOT", "IF", "ELSIF", "ELSE", "IDENTIFY", "GT", "GTE", "EQ", "LTE",
	"LT", "ADD", "SUB", "MUL", "DIV", "WS", "TRUE", "FALSE",
}

type RuleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewRuleLexer(input antlr.CharStream) *RuleLexer {

	l := new(RuleLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Rule.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RuleLexer tokens.
const (
	RuleLexerT__0     = 1
	RuleLexerT__1     = 2
	RuleLexerT__2     = 3
	RuleLexerT__3     = 4
	RuleLexerT__4     = 5
	RuleLexerT__5     = 6
	RuleLexerT__6     = 7
	RuleLexerNUM      = 8
	RuleLexerNOW      = 9
	RuleLexerAND      = 10
	RuleLexerOR       = 11
	RuleLexerNOT      = 12
	RuleLexerIF       = 13
	RuleLexerELSIF    = 14
	RuleLexerELSE     = 15
	RuleLexerIDENTIFY = 16
	RuleLexerGT       = 17
	RuleLexerGTE      = 18
	RuleLexerEQ       = 19
	RuleLexerLTE      = 20
	RuleLexerLT       = 21
	RuleLexerADD      = 22
	RuleLexerSUB      = 23
	RuleLexerMUL      = 24
	RuleLexerDIV      = 25
	RuleLexerWS       = 26
	RuleLexerTRUE     = 27
	RuleLexerFALSE    = 28
)
