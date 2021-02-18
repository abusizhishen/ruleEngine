// Code generated from Rule.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Rule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RuleListener is a complete listener for a parse tree produced by RuleParser.
type RuleListener interface {
	antlr.ParseTreeListener

	// EnterBoolOperate is called when entering the boolOperate production.
	EnterBoolOperate(c *BoolOperateContext)

	// EnterCalculate is called when entering the calculate production.
	EnterCalculate(c *CalculateContext)

	// EnterCompare is called when entering the compare production.
	EnterCompare(c *CompareContext)

	// EnterLogical is called when entering the logical production.
	EnterLogical(c *LogicalContext)

	// EnterCOMPARE is called when entering the COMPARE production.
	EnterCOMPARE(c *COMPAREContext)

	// EnterCOMPAREX is called when entering the COMPAREX production.
	EnterCOMPAREX(c *COMPAREXContext)

	// EnterIDEN is called when entering the IDEN production.
	EnterIDEN(c *IDENContext)

	// EnterNUM is called when entering the NUM production.
	EnterNUM(c *NUMContext)

	// EnterCalculateStatement is called when entering the calculateStatement production.
	EnterCalculateStatement(c *CalculateStatementContext)

	// EnterBOOL is called when entering the BOOL production.
	EnterBOOL(c *BOOLContext)

	// EnterIDENBOOL is called when entering the IDENBOOL production.
	EnterIDENBOOL(c *IDENBOOLContext)

	// EnterCOMPAREVALUE is called when entering the COMPAREVALUE production.
	EnterCOMPAREVALUE(c *COMPAREVALUEContext)

	// EnterBoolStatement is called when entering the boolStatement production.
	EnterBoolStatement(c *BoolStatementContext)

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// ExitBoolOperate is called when exiting the boolOperate production.
	ExitBoolOperate(c *BoolOperateContext)

	// ExitCalculate is called when exiting the calculate production.
	ExitCalculate(c *CalculateContext)

	// ExitCompare is called when exiting the compare production.
	ExitCompare(c *CompareContext)

	// ExitLogical is called when exiting the logical production.
	ExitLogical(c *LogicalContext)

	// ExitCOMPARE is called when exiting the COMPARE production.
	ExitCOMPARE(c *COMPAREContext)

	// ExitCOMPAREX is called when exiting the COMPAREX production.
	ExitCOMPAREX(c *COMPAREXContext)

	// ExitIDEN is called when exiting the IDEN production.
	ExitIDEN(c *IDENContext)

	// ExitNUM is called when exiting the NUM production.
	ExitNUM(c *NUMContext)

	// ExitCalculateStatement is called when exiting the calculateStatement production.
	ExitCalculateStatement(c *CalculateStatementContext)

	// ExitBOOL is called when exiting the BOOL production.
	ExitBOOL(c *BOOLContext)

	// ExitIDENBOOL is called when exiting the IDENBOOL production.
	ExitIDENBOOL(c *IDENBOOLContext)

	// ExitCOMPAREVALUE is called when exiting the COMPAREVALUE production.
	ExitCOMPAREVALUE(c *COMPAREVALUEContext)

	// ExitBoolStatement is called when exiting the boolStatement production.
	ExitBoolStatement(c *BoolStatementContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)
}
