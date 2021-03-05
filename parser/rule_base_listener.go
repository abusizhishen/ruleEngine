// Code generated from Rule.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Rule

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRuleListener is a complete listener for a parse tree produced by RuleParser.
type BaseRuleListener struct{}

var _ RuleListener = &BaseRuleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRuleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRuleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRuleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRuleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBoolOperate is called when production boolOperate is entered.
func (s *BaseRuleListener) EnterBoolOperate(ctx *BoolOperateContext) {}

// ExitBoolOperate is called when production boolOperate is exited.
func (s *BaseRuleListener) ExitBoolOperate(ctx *BoolOperateContext) {}

// EnterCalculate is called when production calculate is entered.
func (s *BaseRuleListener) EnterCalculate(ctx *CalculateContext) {}

// ExitCalculate is called when production calculate is exited.
func (s *BaseRuleListener) ExitCalculate(ctx *CalculateContext) {}

// EnterCompare is called when production compare is entered.
func (s *BaseRuleListener) EnterCompare(ctx *CompareContext) {}

// ExitCompare is called when production compare is exited.
func (s *BaseRuleListener) ExitCompare(ctx *CompareContext) {}

// EnterLogical is called when production logical is entered.
func (s *BaseRuleListener) EnterLogical(ctx *LogicalContext) {}

// ExitLogical is called when production logical is exited.
func (s *BaseRuleListener) ExitLogical(ctx *LogicalContext) {}

// EnterCOMPAREX is called when production COMPAREX is entered.
func (s *BaseRuleListener) EnterCOMPAREX(ctx *COMPAREXContext) {}

// ExitCOMPAREX is called when production COMPAREX is exited.
func (s *BaseRuleListener) ExitCOMPAREX(ctx *COMPAREXContext) {}

// EnterCOMPARE is called when production COMPARE is entered.
func (s *BaseRuleListener) EnterCOMPARE(ctx *COMPAREContext) {}

// ExitCOMPARE is called when production COMPARE is exited.
func (s *BaseRuleListener) ExitCOMPARE(ctx *COMPAREContext) {}

// EnterCalcu is called when production Calcu is entered.
func (s *BaseRuleListener) EnterCalcu(ctx *CalcuContext) {}

// ExitCalcu is called when production Calcu is exited.
func (s *BaseRuleListener) ExitCalcu(ctx *CalcuContext) {}

// EnterITEMCOMP is called when production ITEMCOMP is entered.
func (s *BaseRuleListener) EnterITEMCOMP(ctx *ITEMCOMPContext) {}

// ExitITEMCOMP is called when production ITEMCOMP is exited.
func (s *BaseRuleListener) ExitITEMCOMP(ctx *ITEMCOMPContext) {}

// EnterNUM is called when production NUM is entered.
func (s *BaseRuleListener) EnterNUM(ctx *NUMContext) {}

// ExitNUM is called when production NUM is exited.
func (s *BaseRuleListener) ExitNUM(ctx *NUMContext) {}

// EnterBoolValue is called when production boolValue is entered.
func (s *BaseRuleListener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production boolValue is exited.
func (s *BaseRuleListener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterIDENTIFY is called when production IDENTIFY is entered.
func (s *BaseRuleListener) EnterIDENTIFY(ctx *IDENTIFYContext) {}

// ExitIDENTIFY is called when production IDENTIFY is exited.
func (s *BaseRuleListener) ExitIDENTIFY(ctx *IDENTIFYContext) {}

// EnterSTRING is called when production STRING is entered.
func (s *BaseRuleListener) EnterSTRING(ctx *STRINGContext) {}

// ExitSTRING is called when production STRING is exited.
func (s *BaseRuleListener) ExitSTRING(ctx *STRINGContext) {}

// EnterCalculateValue is called when production calculateValue is entered.
func (s *BaseRuleListener) EnterCalculateValue(ctx *CalculateValueContext) {}

// ExitCalculateValue is called when production calculateValue is exited.
func (s *BaseRuleListener) ExitCalculateValue(ctx *CalculateValueContext) {}

// EnterITEMCALCU is called when production ITEMCALCU is entered.
func (s *BaseRuleListener) EnterITEMCALCU(ctx *ITEMCALCUContext) {}

// ExitITEMCALCU is called when production ITEMCALCU is exited.
func (s *BaseRuleListener) ExitITEMCALCU(ctx *ITEMCALCUContext) {}

// EnterADDSUB is called when production ADDSUB is entered.
func (s *BaseRuleListener) EnterADDSUB(ctx *ADDSUBContext) {}

// ExitADDSUB is called when production ADDSUB is exited.
func (s *BaseRuleListener) ExitADDSUB(ctx *ADDSUBContext) {}

// EnterMULDIV is called when production MULDIV is entered.
func (s *BaseRuleListener) EnterMULDIV(ctx *MULDIVContext) {}

// ExitMULDIV is called when production MULDIV is exited.
func (s *BaseRuleListener) ExitMULDIV(ctx *MULDIVContext) {}

// EnterCALCULATEX is called when production CALCULATEX is entered.
func (s *BaseRuleListener) EnterCALCULATEX(ctx *CALCULATEXContext) {}

// ExitCALCULATEX is called when production CALCULATEX is exited.
func (s *BaseRuleListener) ExitCALCULATEX(ctx *CALCULATEXContext) {}

// EnterBOOL is called when production BOOL is entered.
func (s *BaseRuleListener) EnterBOOL(ctx *BOOLContext) {}

// ExitBOOL is called when production BOOL is exited.
func (s *BaseRuleListener) ExitBOOL(ctx *BOOLContext) {}

// EnterBOOLOP is called when production BOOLOP is entered.
func (s *BaseRuleListener) EnterBOOLOP(ctx *BOOLOPContext) {}

// ExitBOOLOP is called when production BOOLOP is exited.
func (s *BaseRuleListener) ExitBOOLOP(ctx *BOOLOPContext) {}

// EnterIDENBOOL is called when production IDENBOOL is entered.
func (s *BaseRuleListener) EnterIDENBOOL(ctx *IDENBOOLContext) {}

// ExitIDENBOOL is called when production IDENBOOL is exited.
func (s *BaseRuleListener) ExitIDENBOOL(ctx *IDENBOOLContext) {}

// EnterCOMPAREBOOL is called when production COMPAREBOOL is entered.
func (s *BaseRuleListener) EnterCOMPAREBOOL(ctx *COMPAREBOOLContext) {}

// ExitCOMPAREBOOL is called when production COMPAREBOOL is exited.
func (s *BaseRuleListener) ExitCOMPAREBOOL(ctx *COMPAREBOOLContext) {}

// EnterBOOLOPX is called when production BOOLOPX is entered.
func (s *BaseRuleListener) EnterBOOLOPX(ctx *BOOLOPXContext) {}

// ExitBOOLOPX is called when production BOOLOPX is exited.
func (s *BaseRuleListener) ExitBOOLOPX(ctx *BOOLOPXContext) {}

// EnterValueType is called when production valueType is entered.
func (s *BaseRuleListener) EnterValueType(ctx *ValueTypeContext) {}

// ExitValueType is called when production valueType is exited.
func (s *BaseRuleListener) ExitValueType(ctx *ValueTypeContext) {}

// EnterSetValueStatement is called when production setValueStatement is entered.
func (s *BaseRuleListener) EnterSetValueStatement(ctx *SetValueStatementContext) {}

// ExitSetValueStatement is called when production setValueStatement is exited.
func (s *BaseRuleListener) ExitSetValueStatement(ctx *SetValueStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseRuleListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseRuleListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterElseIfStatement is called when production elseIfStatement is entered.
func (s *BaseRuleListener) EnterElseIfStatement(ctx *ElseIfStatementContext) {}

// ExitElseIfStatement is called when production elseIfStatement is exited.
func (s *BaseRuleListener) ExitElseIfStatement(ctx *ElseIfStatementContext) {}

// EnterElseStatement is called when production elseStatement is entered.
func (s *BaseRuleListener) EnterElseStatement(ctx *ElseStatementContext) {}

// ExitElseStatement is called when production elseStatement is exited.
func (s *BaseRuleListener) ExitElseStatement(ctx *ElseStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseRuleListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseRuleListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseRuleListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseRuleListener) ExitPair(ctx *PairContext) {}

// EnterMapValue is called when production mapValue is entered.
func (s *BaseRuleListener) EnterMapValue(ctx *MapValueContext) {}

// ExitMapValue is called when production mapValue is exited.
func (s *BaseRuleListener) ExitMapValue(ctx *MapValueContext) {}

// EnterGetMapValue is called when production getMapValue is entered.
func (s *BaseRuleListener) EnterGetMapValue(ctx *GetMapValueContext) {}

// ExitGetMapValue is called when production getMapValue is exited.
func (s *BaseRuleListener) ExitGetMapValue(ctx *GetMapValueContext) {}

// EnterArray is called when production array is entered.
func (s *BaseRuleListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseRuleListener) ExitArray(ctx *ArrayContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseRuleListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseRuleListener) ExitStatement(ctx *StatementContext) {}

// EnterInit is called when production init is entered.
func (s *BaseRuleListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseRuleListener) ExitInit(ctx *InitContext) {}
