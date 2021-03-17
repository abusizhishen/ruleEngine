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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 35, 251,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 68, 10,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 74, 10, 6, 12, 6, 14, 6, 77, 11, 6, 3,
	7, 3, 7, 5, 7, 81, 10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 5, 12, 93, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 5, 13, 101, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	7, 13, 109, 10, 13, 12, 13, 14, 13, 112, 11, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 122, 10, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 7, 14, 128, 10, 14, 12, 14, 14, 14, 131, 11, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 141, 10, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 147, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 153,
	10, 17, 12, 17, 14, 17, 156, 11, 17, 3, 17, 3, 17, 7, 17, 160, 10, 17,
	12, 17, 14, 17, 163, 11, 17, 3, 17, 5, 17, 166, 10, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 7, 18, 172, 10, 18, 12, 18, 14, 18, 175, 11, 18, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 19, 7, 19, 182, 10, 19, 12, 19, 14, 19, 185, 11, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 193, 10, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 203, 10, 22, 12, 22, 14,
	22, 206, 11, 22, 5, 22, 208, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 6, 23, 217, 10, 23, 13, 23, 14, 23, 218, 3, 24, 3, 24, 3,
	24, 3, 24, 7, 24, 225, 10, 24, 12, 24, 14, 24, 228, 11, 24, 5, 24, 230,
	10, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25,
	240, 10, 25, 3, 26, 7, 26, 243, 10, 26, 12, 26, 14, 26, 246, 11, 26, 3,
	26, 5, 26, 249, 10, 26, 3, 26, 2, 5, 10, 24, 26, 27, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 2, 8, 3, 2, 16, 18, 3, 2, 30, 33, 3, 2, 25, 29, 3, 2, 22, 23, 3, 2,
	32, 33, 3, 2, 30, 31, 2, 264, 2, 52, 3, 2, 2, 2, 4, 54, 3, 2, 2, 2, 6,
	56, 3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 67, 3, 2, 2, 2, 12, 80, 3, 2, 2,
	2, 14, 82, 3, 2, 2, 2, 16, 84, 3, 2, 2, 2, 18, 86, 3, 2, 2, 2, 20, 88,
	3, 2, 2, 2, 22, 92, 3, 2, 2, 2, 24, 100, 3, 2, 2, 2, 26, 121, 3, 2, 2,
	2, 28, 140, 3, 2, 2, 2, 30, 142, 3, 2, 2, 2, 32, 148, 3, 2, 2, 2, 34, 167,
	3, 2, 2, 2, 36, 178, 3, 2, 2, 2, 38, 192, 3, 2, 2, 2, 40, 194, 3, 2, 2,
	2, 42, 198, 3, 2, 2, 2, 44, 211, 3, 2, 2, 2, 46, 220, 3, 2, 2, 2, 48, 239,
	3, 2, 2, 2, 50, 244, 3, 2, 2, 2, 52, 53, 9, 2, 2, 2, 53, 3, 3, 2, 2, 2,
	54, 55, 9, 3, 2, 2, 55, 5, 3, 2, 2, 2, 56, 57, 9, 4, 2, 2, 57, 7, 3, 2,
	2, 2, 58, 59, 9, 2, 2, 2, 59, 9, 3, 2, 2, 2, 60, 61, 8, 6, 1, 2, 61, 68,
	5, 22, 12, 2, 62, 68, 5, 24, 13, 2, 63, 64, 7, 3, 2, 2, 64, 65, 5, 10,
	6, 2, 65, 66, 7, 4, 2, 2, 66, 68, 3, 2, 2, 2, 67, 60, 3, 2, 2, 2, 67, 62,
	3, 2, 2, 2, 67, 63, 3, 2, 2, 2, 68, 75, 3, 2, 2, 2, 69, 70, 12, 6, 2, 2,
	70, 71, 5, 6, 4, 2, 71, 72, 5, 10, 6, 7, 72, 74, 3, 2, 2, 2, 73, 69, 3,
	2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76,
	11, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 81, 5, 20, 11, 2, 79, 81, 5, 14,
	8, 2, 80, 78, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 13, 3, 2, 2, 2, 82, 83,
	7, 14, 2, 2, 83, 15, 3, 2, 2, 2, 84, 85, 9, 5, 2, 2, 85, 17, 3, 2, 2, 2,
	86, 87, 7, 24, 2, 2, 87, 19, 3, 2, 2, 2, 88, 89, 7, 35, 2, 2, 89, 21, 3,
	2, 2, 2, 90, 93, 5, 18, 10, 2, 91, 93, 5, 14, 8, 2, 92, 90, 3, 2, 2, 2,
	92, 91, 3, 2, 2, 2, 93, 23, 3, 2, 2, 2, 94, 95, 8, 13, 1, 2, 95, 101, 5,
	22, 12, 2, 96, 97, 7, 3, 2, 2, 97, 98, 5, 24, 13, 2, 98, 99, 7, 4, 2, 2,
	99, 101, 3, 2, 2, 2, 100, 94, 3, 2, 2, 2, 100, 96, 3, 2, 2, 2, 101, 110,
	3, 2, 2, 2, 102, 103, 12, 6, 2, 2, 103, 104, 9, 6, 2, 2, 104, 109, 5, 24,
	13, 7, 105, 106, 12, 5, 2, 2, 106, 107, 9, 7, 2, 2, 107, 109, 5, 24, 13,
	6, 108, 102, 3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110,
	108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 25, 3, 2, 2, 2, 112, 110, 3,
	2, 2, 2, 113, 114, 8, 14, 1, 2, 114, 122, 9, 5, 2, 2, 115, 122, 7, 24,
	2, 2, 116, 122, 5, 10, 6, 2, 117, 118, 7, 3, 2, 2, 118, 119, 5, 26, 14,
	2, 119, 120, 7, 4, 2, 2, 120, 122, 3, 2, 2, 2, 121, 113, 3, 2, 2, 2, 121,
	115, 3, 2, 2, 2, 121, 116, 3, 2, 2, 2, 121, 117, 3, 2, 2, 2, 122, 129,
	3, 2, 2, 2, 123, 124, 12, 7, 2, 2, 124, 125, 5, 2, 2, 2, 125, 126, 5, 26,
	14, 8, 126, 128, 3, 2, 2, 2, 127, 123, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2,
	129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 27, 3, 2, 2, 2, 131, 129,
	3, 2, 2, 2, 132, 141, 5, 16, 9, 2, 133, 141, 5, 18, 10, 2, 134, 141, 5,
	14, 8, 2, 135, 141, 5, 24, 13, 2, 136, 141, 5, 42, 22, 2, 137, 141, 5,
	44, 23, 2, 138, 141, 5, 20, 11, 2, 139, 141, 5, 46, 24, 2, 140, 132, 3,
	2, 2, 2, 140, 133, 3, 2, 2, 2, 140, 134, 3, 2, 2, 2, 140, 135, 3, 2, 2,
	2, 140, 136, 3, 2, 2, 2, 140, 137, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140,
	139, 3, 2, 2, 2, 141, 29, 3, 2, 2, 2, 142, 143, 7, 24, 2, 2, 143, 144,
	7, 5, 2, 2, 144, 146, 5, 28, 15, 2, 145, 147, 7, 6, 2, 2, 146, 145, 3,
	2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 31, 3, 2, 2, 2, 148, 149, 7, 19, 2,
	2, 149, 150, 5, 26, 14, 2, 150, 154, 7, 7, 2, 2, 151, 153, 5, 48, 25, 2,
	152, 151, 3, 2, 2, 2, 153, 156, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154,
	155, 3, 2, 2, 2, 155, 157, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 157, 161,
	7, 8, 2, 2, 158, 160, 5, 34, 18, 2, 159, 158, 3, 2, 2, 2, 160, 163, 3,
	2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 165, 3, 2, 2,
	2, 163, 161, 3, 2, 2, 2, 164, 166, 5, 36, 19, 2, 165, 164, 3, 2, 2, 2,
	165, 166, 3, 2, 2, 2, 166, 33, 3, 2, 2, 2, 167, 168, 7, 20, 2, 2, 168,
	169, 5, 26, 14, 2, 169, 173, 7, 7, 2, 2, 170, 172, 5, 48, 25, 2, 171, 170,
	3, 2, 2, 2, 172, 175, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2,
	2, 2, 174, 176, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 176, 177, 7, 8, 2, 2,
	177, 35, 3, 2, 2, 2, 178, 179, 7, 21, 2, 2, 179, 183, 7, 7, 2, 2, 180,
	182, 5, 48, 25, 2, 181, 180, 3, 2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 181,
	3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 186, 3, 2, 2, 2, 185, 183, 3, 2,
	2, 2, 186, 187, 7, 8, 2, 2, 187, 37, 3, 2, 2, 2, 188, 189, 7, 9, 2, 2,
	189, 193, 5, 28, 15, 2, 190, 193, 5, 24, 13, 2, 191, 193, 5, 26, 14, 2,
	192, 188, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192, 191, 3, 2, 2, 2, 193,
	39, 3, 2, 2, 2, 194, 195, 7, 35, 2, 2, 195, 196, 7, 10, 2, 2, 196, 197,
	5, 28, 15, 2, 197, 41, 3, 2, 2, 2, 198, 207, 7, 7, 2, 2, 199, 204, 5, 40,
	21, 2, 200, 201, 7, 11, 2, 2, 201, 203, 5, 40, 21, 2, 202, 200, 3, 2, 2,
	2, 203, 206, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205,
	208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 207, 199, 3, 2, 2, 2, 207, 208,
	3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 210, 7, 8, 2, 2, 210, 43, 3, 2,
	2, 2, 211, 216, 5, 18, 10, 2, 212, 213, 7, 12, 2, 2, 213, 214, 5, 12, 7,
	2, 214, 215, 7, 13, 2, 2, 215, 217, 3, 2, 2, 2, 216, 212, 3, 2, 2, 2, 217,
	218, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 45, 3,
	2, 2, 2, 220, 229, 7, 12, 2, 2, 221, 226, 5, 28, 15, 2, 222, 223, 7, 11,
	2, 2, 223, 225, 5, 28, 15, 2, 224, 222, 3, 2, 2, 2, 225, 228, 3, 2, 2,
	2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228,
	226, 3, 2, 2, 2, 229, 221, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231,
	3, 2, 2, 2, 231, 232, 7, 13, 2, 2, 232, 47, 3, 2, 2, 2, 233, 240, 5, 24,
	13, 2, 234, 240, 5, 26, 14, 2, 235, 240, 5, 10, 6, 2, 236, 240, 5, 32,
	17, 2, 237, 240, 5, 30, 16, 2, 238, 240, 5, 38, 20, 2, 239, 233, 3, 2,
	2, 2, 239, 234, 3, 2, 2, 2, 239, 235, 3, 2, 2, 2, 239, 236, 3, 2, 2, 2,
	239, 237, 3, 2, 2, 2, 239, 238, 3, 2, 2, 2, 240, 49, 3, 2, 2, 2, 241, 243,
	5, 48, 25, 2, 242, 241, 3, 2, 2, 2, 243, 246, 3, 2, 2, 2, 244, 242, 3,
	2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 248, 3, 2, 2, 2, 246, 244, 3, 2, 2,
	2, 247, 249, 7, 2, 2, 3, 248, 247, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249,
	51, 3, 2, 2, 2, 27, 67, 75, 80, 92, 100, 108, 110, 121, 129, 140, 146,
	154, 161, 165, 173, 183, 192, 204, 207, 218, 226, 229, 239, 244, 248,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'='", "';'", "'{'", "'}'", "'return'", "':'", "','",
	"'['", "']'", "", "'now()'", "'&&'", "'||'", "'!'", "'if'", "'elsif'",
	"'else'", "'true'", "'false'", "", "'>'", "'>='", "'=='", "'<='", "'<'",
	"'+'", "'-'", "'*'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "NUM", "NOW", "AND", "OR",
	"NOT", "IF", "ELSIF", "ELSE", "TRUE", "FALSE", "IDENTIFY", "GT", "GTE",
	"EQ", "LTE", "LT", "ADD", "SUB", "MUL", "DIV", "WS", "Str",
}

var ruleNames = []string{
	"boolOperate", "calculate", "compare", "logical", "compareStatement", "key",
	"num", "boolValue", "identify", "stringValue", "calculateValue", "calculateStatement",
	"boolStatement", "valueType", "setValueStatement", "ifStatement", "elseIfStatement",
	"elseStatement", "returnStatement", "pair", "mapValue", "getMapOrArrayValue",
	"array", "statement", "init",
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
	RuleParserT__7     = 8
	RuleParserT__8     = 9
	RuleParserT__9     = 10
	RuleParserT__10    = 11
	RuleParserNUM      = 12
	RuleParserNOW      = 13
	RuleParserAND      = 14
	RuleParserOR       = 15
	RuleParserNOT      = 16
	RuleParserIF       = 17
	RuleParserELSIF    = 18
	RuleParserELSE     = 19
	RuleParserTRUE     = 20
	RuleParserFALSE    = 21
	RuleParserIDENTIFY = 22
	RuleParserGT       = 23
	RuleParserGTE      = 24
	RuleParserEQ       = 25
	RuleParserLTE      = 26
	RuleParserLT       = 27
	RuleParserADD      = 28
	RuleParserSUB      = 29
	RuleParserMUL      = 30
	RuleParserDIV      = 31
	RuleParserWS       = 32
	RuleParserStr      = 33
)

// RuleParser rules.
const (
	RuleParserRULE_boolOperate        = 0
	RuleParserRULE_calculate          = 1
	RuleParserRULE_compare            = 2
	RuleParserRULE_logical            = 3
	RuleParserRULE_compareStatement   = 4
	RuleParserRULE_key                = 5
	RuleParserRULE_num                = 6
	RuleParserRULE_boolValue          = 7
	RuleParserRULE_identify           = 8
	RuleParserRULE_stringValue        = 9
	RuleParserRULE_calculateValue     = 10
	RuleParserRULE_calculateStatement = 11
	RuleParserRULE_boolStatement      = 12
	RuleParserRULE_valueType          = 13
	RuleParserRULE_setValueStatement  = 14
	RuleParserRULE_ifStatement        = 15
	RuleParserRULE_elseIfStatement    = 16
	RuleParserRULE_elseStatement      = 17
	RuleParserRULE_returnStatement    = 18
	RuleParserRULE_pair               = 19
	RuleParserRULE_mapValue           = 20
	RuleParserRULE_getMapOrArrayValue = 21
	RuleParserRULE_array              = 22
	RuleParserRULE_statement          = 23
	RuleParserRULE_init               = 24
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
		p.SetState(50)
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
		p.SetState(52)
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
		p.SetState(54)
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
		p.SetState(56)
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewITEMCOMPContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(59)
			p.CalculateValue()
		}

	case 2:
		localctx = NewCalcuContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(60)
			p.calculateStatement(0)
		}

	case 3:
		localctx = NewCOMPAREXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(62)
			p.compareStatement(0)
		}
		{
			p.SetState(63)
			p.Match(RuleParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(73)
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
			p.SetState(67)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(68)

				var _x = p.Compare()

				localctx.(*COMPAREContext).op = _x
			}
			{
				p.SetState(69)
				p.compareStatement(5)
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *KeyContext) Num() INumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RuleParserRULE_key)

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

	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserStr:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.StringValue()
		}

	case RuleParserNUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Num()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 12, RuleParserRULE_num)

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
		p.SetState(80)
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
	p.EnterRule(localctx, 14, RuleParserRULE_boolValue)
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
		p.SetState(82)
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
	p.EnterRule(localctx, 16, RuleParserRULE_identify)

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
		p.SetState(84)
		p.Match(RuleParserIDENTIFY)
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) CopyFrom(ctx *StringValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type STRINGContext struct {
	*StringValueContext
}

func NewSTRINGContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *STRINGContext {
	var p = new(STRINGContext)

	p.StringValueContext = NewEmptyStringValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringValueContext))

	return p
}

func (s *STRINGContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *STRINGContext) Str() antlr.TerminalNode {
	return s.GetToken(RuleParserStr, 0)
}

func (s *STRINGContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitSTRING(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RuleParserRULE_stringValue)

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

	localctx = NewSTRINGContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(RuleParserStr)
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
	p.EnterRule(localctx, 20, RuleParserRULE_calculateValue)

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

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserIDENTIFY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Identify()
		}

	case RuleParserNUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
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

func (s *ADDSUBContext) ADD() antlr.TerminalNode {
	return s.GetToken(RuleParserADD, 0)
}

func (s *ADDSUBContext) SUB() antlr.TerminalNode {
	return s.GetToken(RuleParserSUB, 0)
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

func (s *MULDIVContext) MUL() antlr.TerminalNode {
	return s.GetToken(RuleParserMUL, 0)
}

func (s *MULDIVContext) DIV() antlr.TerminalNode {
	return s.GetToken(RuleParserDIV, 0)
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
	_startState := 22
	p.EnterRecursionRule(localctx, 22, RuleParserRULE_calculateStatement, _p)
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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuleParserNUM, RuleParserIDENTIFY:
		localctx = NewITEMCALCUContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(93)
			p.CalculateValue()
		}

	case RuleParserT__0:
		localctx = NewCALCULATEXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(94)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(95)
			p.calculateStatement(0)
		}
		{
			p.SetState(96)
			p.Match(RuleParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMULDIVContext(p, NewCalculateStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_calculateStatement)
				p.SetState(100)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(101)

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
					p.SetState(102)
					p.calculateStatement(5)
				}

			case 2:
				localctx = NewADDSUBContext(p, NewCalculateStatementContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_calculateStatement)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(104)

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
					p.SetState(105)
					p.calculateStatement(4)
				}

			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, RuleParserRULE_boolStatement, _p)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(112)
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
			p.SetState(113)
			p.Match(RuleParserIDENTIFY)
		}

	case 3:
		localctx = NewCOMPAREBOOLContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(114)
			p.compareStatement(0)
		}

	case 4:
		localctx = NewBOOLOPXContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(115)
			p.Match(RuleParserT__0)
		}
		{
			p.SetState(116)
			p.boolStatement(0)
		}
		{
			p.SetState(117)
			p.Match(RuleParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBOOLOPContext(p, NewBoolStatementContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, RuleParserRULE_boolStatement)
			p.SetState(121)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(122)

				var _x = p.BoolOperate()

				localctx.(*BOOLOPContext).op = _x
			}
			{
				p.SetState(123)
				p.boolStatement(6)
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
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

func (s *ValueTypeContext) MapValue() IMapValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapValueContext)
}

func (s *ValueTypeContext) GetMapOrArrayValue() IGetMapOrArrayValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetMapOrArrayValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetMapOrArrayValueContext)
}

func (s *ValueTypeContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *ValueTypeContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
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
	p.EnterRule(localctx, 26, RuleParserRULE_valueType)

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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.BoolValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Identify()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Num()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(133)
			p.calculateStatement(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(134)
			p.MapValue()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(135)
			p.GetMapOrArrayValue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(136)
			p.StringValue()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(137)
			p.Array()
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
	p.EnterRule(localctx, 28, RuleParserRULE_setValueStatement)
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
		p.SetState(140)
		p.Match(RuleParserIDENTIFY)
	}
	{
		p.SetState(141)
		p.Match(RuleParserT__2)
	}
	{
		p.SetState(142)
		p.ValueType()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuleParserT__3 {
		{
			p.SetState(143)
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
	p.EnterRule(localctx, 30, RuleParserRULE_ifStatement)
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
		p.SetState(146)
		p.Match(RuleParserIF)
	}
	{
		p.SetState(147)
		p.boolStatement(0)
	}
	{
		p.SetState(148)
		p.Match(RuleParserT__4)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(149)
			p.Statement()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(155)
		p.Match(RuleParserT__5)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RuleParserELSIF {
		{
			p.SetState(156)
			p.ElseIfStatement()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuleParserELSE {
		{
			p.SetState(162)
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
	p.EnterRule(localctx, 32, RuleParserRULE_elseIfStatement)
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
		p.SetState(165)
		p.Match(RuleParserELSIF)
	}
	{
		p.SetState(166)
		p.boolStatement(0)
	}
	{
		p.SetState(167)
		p.Match(RuleParserT__4)
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(168)
			p.Statement()
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(174)
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
	p.EnterRule(localctx, 34, RuleParserRULE_elseStatement)
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
		p.SetState(176)
		p.Match(RuleParserELSE)
	}
	{
		p.SetState(177)
		p.Match(RuleParserT__4)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(178)
			p.Statement()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
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
	p.EnterRule(localctx, 36, RuleParserRULE_returnStatement)

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

	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(RuleParserT__6)
		}
		{
			p.SetState(187)

			var _x = p.ValueType()

			localctx.(*ReturnStatementContext).value = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.calculateStatement(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(189)
			p.boolStatement(0)
		}

	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMapKey returns the mapKey token.
	GetMapKey() antlr.Token

	// SetMapKey sets the mapKey token.
	SetMapKey(antlr.Token)

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	mapKey antlr.Token
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) GetMapKey() antlr.Token { return s.mapKey }

func (s *PairContext) SetMapKey(v antlr.Token) { s.mapKey = v }

func (s *PairContext) ValueType() IValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *PairContext) Str() antlr.TerminalNode {
	return s.GetToken(RuleParserStr, 0)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RuleParserRULE_pair)

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
		p.SetState(192)

		var _m = p.Match(RuleParserStr)

		localctx.(*PairContext).mapKey = _m
	}
	{
		p.SetState(193)
		p.Match(RuleParserT__7)
	}
	{
		p.SetState(194)
		p.ValueType()
	}

	return localctx
}

// IMapValueContext is an interface to support dynamic dispatch.
type IMapValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapValueContext differentiates from other interfaces.
	IsMapValueContext()
}

type MapValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapValueContext() *MapValueContext {
	var p = new(MapValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_mapValue
	return p
}

func (*MapValueContext) IsMapValueContext() {}

func NewMapValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapValueContext {
	var p = new(MapValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_mapValue

	return p
}

func (s *MapValueContext) GetParser() antlr.Parser { return s.parser }

func (s *MapValueContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *MapValueContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *MapValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitMapValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) MapValue() (localctx IMapValueContext) {
	localctx = NewMapValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RuleParserRULE_mapValue)
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
		p.SetState(196)
		p.Match(RuleParserT__4)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuleParserStr {
		{
			p.SetState(197)
			p.Pair()
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RuleParserT__8 {
			{
				p.SetState(198)
				p.Match(RuleParserT__8)
			}
			{
				p.SetState(199)
				p.Pair()
			}

			p.SetState(204)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(207)
		p.Match(RuleParserT__5)
	}

	return localctx
}

// IGetMapOrArrayValueContext is an interface to support dynamic dispatch.
type IGetMapOrArrayValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetMapOrArrayValueContext differentiates from other interfaces.
	IsGetMapOrArrayValueContext()
}

type GetMapOrArrayValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetMapOrArrayValueContext() *GetMapOrArrayValueContext {
	var p = new(GetMapOrArrayValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_getMapOrArrayValue
	return p
}

func (*GetMapOrArrayValueContext) IsGetMapOrArrayValueContext() {}

func NewGetMapOrArrayValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetMapOrArrayValueContext {
	var p = new(GetMapOrArrayValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_getMapOrArrayValue

	return p
}

func (s *GetMapOrArrayValueContext) GetParser() antlr.Parser { return s.parser }

func (s *GetMapOrArrayValueContext) Identify() IIdentifyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifyContext)
}

func (s *GetMapOrArrayValueContext) AllKey() []IKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyContext)(nil)).Elem())
	var tst = make([]IKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyContext)
		}
	}

	return tst
}

func (s *GetMapOrArrayValueContext) Key(i int) IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *GetMapOrArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetMapOrArrayValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetMapOrArrayValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitGetMapOrArrayValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) GetMapOrArrayValue() (localctx IGetMapOrArrayValueContext) {
	localctx = NewGetMapOrArrayValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RuleParserRULE_getMapOrArrayValue)
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
		p.SetState(209)
		p.Identify()
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RuleParserT__9 {
		{
			p.SetState(210)
			p.Match(RuleParserT__9)
		}
		{
			p.SetState(211)
			p.Key()
		}
		{
			p.SetState(212)
			p.Match(RuleParserT__10)
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuleParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuleParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValueType() []IValueTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueTypeContext)(nil)).Elem())
	var tst = make([]IValueTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueTypeContext)
		}
	}

	return tst
}

func (s *ArrayContext) ValueType(i int) IValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuleVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuleParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RuleParserRULE_array)
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
		p.SetState(218)
		p.Match(RuleParserT__9)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__4)|(1<<RuleParserT__9)|(1<<RuleParserNUM)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0) || _la == RuleParserStr {
		{
			p.SetState(219)
			p.ValueType()
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RuleParserT__8 {
			{
				p.SetState(220)
				p.Match(RuleParserT__8)
			}
			{
				p.SetState(221)
				p.ValueType()
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(229)
		p.Match(RuleParserT__10)
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
	p.EnterRule(localctx, 46, RuleParserRULE_statement)

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

	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.calculateStatement(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
			p.boolStatement(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.compareStatement(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(234)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(235)
			p.SetValueStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(236)
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
	p.EnterRule(localctx, 48, RuleParserRULE_init)
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
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuleParserT__0)|(1<<RuleParserT__6)|(1<<RuleParserNUM)|(1<<RuleParserIF)|(1<<RuleParserTRUE)|(1<<RuleParserFALSE)|(1<<RuleParserIDENTIFY))) != 0 {
		{
			p.SetState(239)
			p.Statement()
		}

		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(245)
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

	case 11:
		var t *CalculateStatementContext = nil
		if localctx != nil {
			t = localctx.(*CalculateStatementContext)
		}
		return p.CalculateStatement_Sempred(t, predIndex)

	case 12:
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
