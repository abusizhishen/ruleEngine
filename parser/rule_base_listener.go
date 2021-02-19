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

// EnterIDEN is called when production IDEN is entered.
func (s *BaseRuleListener) EnterIDEN(ctx *IDENContext) {}

// ExitIDEN is called when production IDEN is exited.
func (s *BaseRuleListener) ExitIDEN(ctx *IDENContext) {}

// EnterNUM is called when production NUM is entered.
func (s *BaseRuleListener) EnterNUM(ctx *NUMContext) {}

// ExitNUM is called when production NUM is exited.
func (s *BaseRuleListener) ExitNUM(ctx *NUMContext) {}

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

// EnterStatement is called when production statement is entered.
func (s *BaseRuleListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseRuleListener) ExitStatement(ctx *StatementContext) {}

// EnterInit is called when production init is entered.
func (s *BaseRuleListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseRuleListener) ExitInit(ctx *InitContext) {}
