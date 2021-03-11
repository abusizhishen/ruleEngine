// Code generated from Rule.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Rule

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRuleVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRuleVisitor) VisitBoolOperate(ctx *BoolOperateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitCalculate(ctx *CalculateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitCompare(ctx *CompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitLogical(ctx *LogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitCOMPAREX(ctx *COMPAREXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitCOMPARE(ctx *COMPAREContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitCalcu(ctx *CalcuContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitITEMCOMP(ctx *ITEMCOMPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitNUM(ctx *NUMContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitBoolValue(ctx *BoolValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitIDENTIFY(ctx *IDENTIFYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitSTRING(ctx *STRINGContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitCalculateValue(ctx *CalculateValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitITEMCALCU(ctx *ITEMCALCUContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitADDSUB(ctx *ADDSUBContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitMULDIV(ctx *MULDIVContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitCALCULATEX(ctx *CALCULATEXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitBOOL(ctx *BOOLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitBOOLOP(ctx *BOOLOPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitIDENBOOL(ctx *IDENBOOLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitCOMPAREBOOL(ctx *COMPAREBOOLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitBOOLOPX(ctx *BOOLOPXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitValueType(ctx *ValueTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitSetValueStatement(ctx *SetValueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitElseIfStatement(ctx *ElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitElseStatement(ctx *ElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitMapValue(ctx *MapValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitGetMapValue(ctx *GetMapValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitGetArrayValue(ctx *GetArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuleVisitor) VisitInit(ctx *InitContext) interface{} {
	return v.VisitChildren(ctx)
}
