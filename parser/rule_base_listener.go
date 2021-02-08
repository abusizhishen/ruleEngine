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

// EnterCompareStatement is called when production compareStatement is entered.
func (s *BaseRuleListener) EnterCompareStatement(ctx *CompareStatementContext) {}

// ExitCompareStatement is called when production compareStatement is exited.
func (s *BaseRuleListener) ExitCompareStatement(ctx *CompareStatementContext) {}

// EnterCalculateStatement is called when production calculateStatement is entered.
func (s *BaseRuleListener) EnterCalculateStatement(ctx *CalculateStatementContext) {}

// ExitCalculateStatement is called when production calculateStatement is exited.
func (s *BaseRuleListener) ExitCalculateStatement(ctx *CalculateStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseRuleListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseRuleListener) ExitStatement(ctx *StatementContext) {}

// EnterInit is called when production init is entered.
func (s *BaseRuleListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseRuleListener) ExitInit(ctx *InitContext) {}
