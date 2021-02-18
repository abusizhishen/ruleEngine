// Code generated from Rule.g4 by ANTLR 4.8. DO NOT EDIT.

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

// EnterCOMPARE is called when production COMPARE is entered.
func (s *BaseRuleListener) EnterCOMPARE(ctx *COMPAREContext) {}

// ExitCOMPARE is called when production COMPARE is exited.
func (s *BaseRuleListener) ExitCOMPARE(ctx *COMPAREContext) {}

// EnterCOMPAREX is called when production COMPAREX is entered.
func (s *BaseRuleListener) EnterCOMPAREX(ctx *COMPAREXContext) {}

// ExitCOMPAREX is called when production COMPAREX is exited.
func (s *BaseRuleListener) ExitCOMPAREX(ctx *COMPAREXContext) {}

// EnterIDEN is called when production IDEN is entered.
func (s *BaseRuleListener) EnterIDEN(ctx *IDENContext) {}

// ExitIDEN is called when production IDEN is exited.
func (s *BaseRuleListener) ExitIDEN(ctx *IDENContext) {}

// EnterNUM is called when production NUM is entered.
func (s *BaseRuleListener) EnterNUM(ctx *NUMContext) {}

// ExitNUM is called when production NUM is exited.
func (s *BaseRuleListener) ExitNUM(ctx *NUMContext) {}

// EnterCalculateStatement is called when production calculateStatement is entered.
func (s *BaseRuleListener) EnterCalculateStatement(ctx *CalculateStatementContext) {}

// ExitCalculateStatement is called when production calculateStatement is exited.
func (s *BaseRuleListener) ExitCalculateStatement(ctx *CalculateStatementContext) {}

// EnterBOOL is called when production BOOL is entered.
func (s *BaseRuleListener) EnterBOOL(ctx *BOOLContext) {}

// ExitBOOL is called when production BOOL is exited.
func (s *BaseRuleListener) ExitBOOL(ctx *BOOLContext) {}

// EnterIDENBOOL is called when production IDENBOOL is entered.
func (s *BaseRuleListener) EnterIDENBOOL(ctx *IDENBOOLContext) {}

// ExitIDENBOOL is called when production IDENBOOL is exited.
func (s *BaseRuleListener) ExitIDENBOOL(ctx *IDENBOOLContext) {}

// EnterCOMPAREVALUE is called when production COMPAREVALUE is entered.
func (s *BaseRuleListener) EnterCOMPAREVALUE(ctx *COMPAREVALUEContext) {}

// ExitCOMPAREVALUE is called when production COMPAREVALUE is exited.
func (s *BaseRuleListener) ExitCOMPAREVALUE(ctx *COMPAREVALUEContext) {}

// EnterBoolStatement is called when production boolStatement is entered.
func (s *BaseRuleListener) EnterBoolStatement(ctx *BoolStatementContext) {}

// ExitBoolStatement is called when production boolStatement is exited.
func (s *BaseRuleListener) ExitBoolStatement(ctx *BoolStatementContext) {}

// EnterInit is called when production init is entered.
func (s *BaseRuleListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseRuleListener) ExitInit(ctx *InitContext) {}
