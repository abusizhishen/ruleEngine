// Code generated from Rule.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Rule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RuleListener is a complete listener for a parse tree produced by RuleParser.
type RuleListener interface {
	antlr.ParseTreeListener

	// EnterCalculate is called when entering the calculate production.
	EnterCalculate(c *CalculateContext)

	// EnterCompare is called when entering the compare production.
	EnterCompare(c *CompareContext)

	// EnterLogical is called when entering the logical production.
	EnterLogical(c *LogicalContext)

	// EnterCompareStatement is called when entering the compareStatement production.
	EnterCompareStatement(c *CompareStatementContext)

	// EnterCalculateStatement is called when entering the calculateStatement production.
	EnterCalculateStatement(c *CalculateStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// ExitCalculate is called when exiting the calculate production.
	ExitCalculate(c *CalculateContext)

	// ExitCompare is called when exiting the compare production.
	ExitCompare(c *CompareContext)

	// ExitLogical is called when exiting the logical production.
	ExitLogical(c *LogicalContext)

	// ExitCompareStatement is called when exiting the compareStatement production.
	ExitCompareStatement(c *CompareStatementContext)

	// ExitCalculateStatement is called when exiting the calculateStatement production.
	ExitCalculateStatement(c *CalculateStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)
}
