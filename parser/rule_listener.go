// Code generated from Rule.g4 by ANTLR 4.7.2. DO NOT EDIT.

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

	// EnterCOMPAREX is called when entering the COMPAREX production.
	EnterCOMPAREX(c *COMPAREXContext)

	// EnterCOMPARE is called when entering the COMPARE production.
	EnterCOMPARE(c *COMPAREContext)

	// EnterCalcu is called when entering the Calcu production.
	EnterCalcu(c *CalcuContext)

	// EnterITEMCOMP is called when entering the ITEMCOMP production.
	EnterITEMCOMP(c *ITEMCOMPContext)

	// EnterNUM is called when entering the NUM production.
	EnterNUM(c *NUMContext)

	// EnterIDENTIFY is called when entering the IDENTIFY production.
	EnterIDENTIFY(c *IDENTIFYContext)

	// EnterCalculateValue is called when entering the calculateValue production.
	EnterCalculateValue(c *CalculateValueContext)

	// EnterITEMCALCU is called when entering the ITEMCALCU production.
	EnterITEMCALCU(c *ITEMCALCUContext)

	// EnterADDSUB is called when entering the ADDSUB production.
	EnterADDSUB(c *ADDSUBContext)

	// EnterMULDIV is called when entering the MULDIV production.
	EnterMULDIV(c *MULDIVContext)

	// EnterCALCULATEX is called when entering the CALCULATEX production.
	EnterCALCULATEX(c *CALCULATEXContext)

	// EnterBOOL is called when entering the BOOL production.
	EnterBOOL(c *BOOLContext)

	// EnterBOOLOP is called when entering the BOOLOP production.
	EnterBOOLOP(c *BOOLOPContext)

	// EnterIDENBOOL is called when entering the IDENBOOL production.
	EnterIDENBOOL(c *IDENBOOLContext)

	// EnterCOMPAREBOOL is called when entering the COMPAREBOOL production.
	EnterCOMPAREBOOL(c *COMPAREBOOLContext)

	// EnterBOOLOPX is called when entering the BOOLOPX production.
	EnterBOOLOPX(c *BOOLOPXContext)

	// EnterValueType is called when entering the valueType production.
	EnterValueType(c *ValueTypeContext)

	// EnterSetValueStatement is called when entering the setValueStatement production.
	EnterSetValueStatement(c *SetValueStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterElseIfStatement is called when entering the elseIfStatement production.
	EnterElseIfStatement(c *ElseIfStatementContext)

	// EnterElseStatement is called when entering the elseStatement production.
	EnterElseStatement(c *ElseStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

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

	// ExitCOMPAREX is called when exiting the COMPAREX production.
	ExitCOMPAREX(c *COMPAREXContext)

	// ExitCOMPARE is called when exiting the COMPARE production.
	ExitCOMPARE(c *COMPAREContext)

	// ExitCalcu is called when exiting the Calcu production.
	ExitCalcu(c *CalcuContext)

	// ExitITEMCOMP is called when exiting the ITEMCOMP production.
	ExitITEMCOMP(c *ITEMCOMPContext)

	// ExitNUM is called when exiting the NUM production.
	ExitNUM(c *NUMContext)

	// ExitIDENTIFY is called when exiting the IDENTIFY production.
	ExitIDENTIFY(c *IDENTIFYContext)

	// ExitCalculateValue is called when exiting the calculateValue production.
	ExitCalculateValue(c *CalculateValueContext)

	// ExitITEMCALCU is called when exiting the ITEMCALCU production.
	ExitITEMCALCU(c *ITEMCALCUContext)

	// ExitADDSUB is called when exiting the ADDSUB production.
	ExitADDSUB(c *ADDSUBContext)

	// ExitMULDIV is called when exiting the MULDIV production.
	ExitMULDIV(c *MULDIVContext)

	// ExitCALCULATEX is called when exiting the CALCULATEX production.
	ExitCALCULATEX(c *CALCULATEXContext)

	// ExitBOOL is called when exiting the BOOL production.
	ExitBOOL(c *BOOLContext)

	// ExitBOOLOP is called when exiting the BOOLOP production.
	ExitBOOLOP(c *BOOLOPContext)

	// ExitIDENBOOL is called when exiting the IDENBOOL production.
	ExitIDENBOOL(c *IDENBOOLContext)

	// ExitCOMPAREBOOL is called when exiting the COMPAREBOOL production.
	ExitCOMPAREBOOL(c *COMPAREBOOLContext)

	// ExitBOOLOPX is called when exiting the BOOLOPX production.
	ExitBOOLOPX(c *BOOLOPXContext)

	// ExitValueType is called when exiting the valueType production.
	ExitValueType(c *ValueTypeContext)

	// ExitSetValueStatement is called when exiting the setValueStatement production.
	ExitSetValueStatement(c *SetValueStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitElseIfStatement is called when exiting the elseIfStatement production.
	ExitElseIfStatement(c *ElseIfStatementContext)

	// ExitElseStatement is called when exiting the elseStatement production.
	ExitElseStatement(c *ElseStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)
}
