// Code generated from Rule.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Rule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RuleParser.
type RuleVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RuleParser#boolOperate.
	VisitBoolOperate(ctx *BoolOperateContext) interface{}

	// Visit a parse tree produced by RuleParser#calculate.
	VisitCalculate(ctx *CalculateContext) interface{}

	// Visit a parse tree produced by RuleParser#compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by RuleParser#logical.
	VisitLogical(ctx *LogicalContext) interface{}

	// Visit a parse tree produced by RuleParser#COMPAREX.
	VisitCOMPAREX(ctx *COMPAREXContext) interface{}

	// Visit a parse tree produced by RuleParser#COMPARE.
	VisitCOMPARE(ctx *COMPAREContext) interface{}

	// Visit a parse tree produced by RuleParser#Calcu.
	VisitCalcu(ctx *CalcuContext) interface{}

	// Visit a parse tree produced by RuleParser#ITEMCOMP.
	VisitITEMCOMP(ctx *ITEMCOMPContext) interface{}

	// Visit a parse tree produced by RuleParser#NUM.
	VisitNUM(ctx *NUMContext) interface{}

	// Visit a parse tree produced by RuleParser#IDENTIFY.
	VisitIDENTIFY(ctx *IDENTIFYContext) interface{}

	// Visit a parse tree produced by RuleParser#calculateValue.
	VisitCalculateValue(ctx *CalculateValueContext) interface{}

	// Visit a parse tree produced by RuleParser#ITEMCALCU.
	VisitITEMCALCU(ctx *ITEMCALCUContext) interface{}

	// Visit a parse tree produced by RuleParser#ADDSUB.
	VisitADDSUB(ctx *ADDSUBContext) interface{}

	// Visit a parse tree produced by RuleParser#MULDIV.
	VisitMULDIV(ctx *MULDIVContext) interface{}

	// Visit a parse tree produced by RuleParser#CALCULATEX.
	VisitCALCULATEX(ctx *CALCULATEXContext) interface{}

	// Visit a parse tree produced by RuleParser#BOOL.
	VisitBOOL(ctx *BOOLContext) interface{}

	// Visit a parse tree produced by RuleParser#BOOLOP.
	VisitBOOLOP(ctx *BOOLOPContext) interface{}

	// Visit a parse tree produced by RuleParser#IDENBOOL.
	VisitIDENBOOL(ctx *IDENBOOLContext) interface{}

	// Visit a parse tree produced by RuleParser#COMPAREBOOL.
	VisitCOMPAREBOOL(ctx *COMPAREBOOLContext) interface{}

	// Visit a parse tree produced by RuleParser#BOOLOPX.
	VisitBOOLOPX(ctx *BOOLOPXContext) interface{}

	// Visit a parse tree produced by RuleParser#valueType.
	VisitValueType(ctx *ValueTypeContext) interface{}

	// Visit a parse tree produced by RuleParser#setValueStatement.
	VisitSetValueStatement(ctx *SetValueStatementContext) interface{}

	// Visit a parse tree produced by RuleParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by RuleParser#elseIfStatement.
	VisitElseIfStatement(ctx *ElseIfStatementContext) interface{}

	// Visit a parse tree produced by RuleParser#elseStatement.
	VisitElseStatement(ctx *ElseStatementContext) interface{}

	// Visit a parse tree produced by RuleParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by RuleParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by RuleParser#init.
	VisitInit(ctx *InitContext) interface{}
}
