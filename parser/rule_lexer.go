// Code generated from Rule.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 113,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 6, 4, 49, 10, 4, 13, 4, 14, 4, 50, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 6, 9, 71, 10, 9, 13, 9, 14, 9, 72, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 6, 19, 97, 10,
	19, 13, 19, 14, 19, 98, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 2, 2, 22, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15,
	29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 3, 2, 5, 3, 2,
	50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	115, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 3, 43, 3, 2, 2, 2, 5, 45, 3, 2, 2, 2, 7, 48, 3, 2, 2,
	2, 9, 52, 3, 2, 2, 2, 11, 58, 3, 2, 2, 2, 13, 62, 3, 2, 2, 2, 15, 65, 3,
	2, 2, 2, 17, 70, 3, 2, 2, 2, 19, 74, 3, 2, 2, 2, 21, 76, 3, 2, 2, 2, 23,
	79, 3, 2, 2, 2, 25, 82, 3, 2, 2, 2, 27, 85, 3, 2, 2, 2, 29, 87, 3, 2, 2,
	2, 31, 89, 3, 2, 2, 2, 33, 91, 3, 2, 2, 2, 35, 93, 3, 2, 2, 2, 37, 96,
	3, 2, 2, 2, 39, 102, 3, 2, 2, 2, 41, 107, 3, 2, 2, 2, 43, 44, 7, 42, 2,
	2, 44, 4, 3, 2, 2, 2, 45, 46, 7, 43, 2, 2, 46, 6, 3, 2, 2, 2, 47, 49, 9,
	2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50,
	51, 3, 2, 2, 2, 51, 8, 3, 2, 2, 2, 52, 53, 7, 112, 2, 2, 53, 54, 7, 113,
	2, 2, 54, 55, 7, 121, 2, 2, 55, 56, 7, 42, 2, 2, 56, 57, 7, 43, 2, 2, 57,
	10, 3, 2, 2, 2, 58, 59, 7, 99, 2, 2, 59, 60, 7, 112, 2, 2, 60, 61, 7, 102,
	2, 2, 61, 12, 3, 2, 2, 2, 62, 63, 7, 113, 2, 2, 63, 64, 7, 116, 2, 2, 64,
	14, 3, 2, 2, 2, 65, 66, 7, 112, 2, 2, 66, 67, 7, 113, 2, 2, 67, 68, 7,
	118, 2, 2, 68, 16, 3, 2, 2, 2, 69, 71, 9, 3, 2, 2, 70, 69, 3, 2, 2, 2,
	71, 72, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 18, 3,
	2, 2, 2, 74, 75, 7, 64, 2, 2, 75, 20, 3, 2, 2, 2, 76, 77, 7, 64, 2, 2,
	77, 78, 7, 63, 2, 2, 78, 22, 3, 2, 2, 2, 79, 80, 7, 63, 2, 2, 80, 81, 7,
	63, 2, 2, 81, 24, 3, 2, 2, 2, 82, 83, 7, 62, 2, 2, 83, 84, 7, 63, 2, 2,
	84, 26, 3, 2, 2, 2, 85, 86, 7, 62, 2, 2, 86, 28, 3, 2, 2, 2, 87, 88, 7,
	45, 2, 2, 88, 30, 3, 2, 2, 2, 89, 90, 7, 47, 2, 2, 90, 32, 3, 2, 2, 2,
	91, 92, 7, 44, 2, 2, 92, 34, 3, 2, 2, 2, 93, 94, 7, 49, 2, 2, 94, 36, 3,
	2, 2, 2, 95, 97, 9, 4, 2, 2, 96, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98,
	96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 101, 8, 19,
	2, 2, 101, 38, 3, 2, 2, 2, 102, 103, 7, 118, 2, 2, 103, 104, 7, 116, 2,
	2, 104, 105, 7, 119, 2, 2, 105, 106, 7, 103, 2, 2, 106, 40, 3, 2, 2, 2,
	107, 108, 7, 104, 2, 2, 108, 109, 7, 99, 2, 2, 109, 110, 7, 110, 2, 2,
	110, 111, 7, 117, 2, 2, 111, 112, 7, 103, 2, 2, 112, 42, 3, 2, 2, 2, 6,
	2, 50, 72, 98, 3, 8, 2, 2,
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
	"", "'('", "')'", "", "'now()'", "'and'", "'or'", "'not'", "", "'>'", "'>='",
	"'=='", "'<='", "'<'", "'+'", "'-'", "'*'", "'/'", "", "'true'", "'false'",
}

var lexerSymbolicNames = []string{
	"", "", "", "NUM", "NOW", "AND", "OR", "NOT", "IDENTIFY", "GT", "GTE",
	"EQ", "LTE", "LT", "ADD", "SUB", "MUL", "DIV", "WS", "TRUE", "FALSE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "NUM", "NOW", "AND", "OR", "NOT", "IDENTIFY", "GT", "GTE",
	"EQ", "LTE", "LT", "ADD", "SUB", "MUL", "DIV", "WS", "TRUE", "FALSE",
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
	RuleLexerNUM      = 3
	RuleLexerNOW      = 4
	RuleLexerAND      = 5
	RuleLexerOR       = 6
	RuleLexerNOT      = 7
	RuleLexerIDENTIFY = 8
	RuleLexerGT       = 9
	RuleLexerGTE      = 10
	RuleLexerEQ       = 11
	RuleLexerLTE      = 12
	RuleLexerLT       = 13
	RuleLexerADD      = 14
	RuleLexerSUB      = 15
	RuleLexerMUL      = 16
	RuleLexerDIV      = 17
	RuleLexerWS       = 18
	RuleLexerTRUE     = 19
	RuleLexerFALSE    = 20
)
