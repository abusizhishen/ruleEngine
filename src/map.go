package src

import (
	"github.com/abusizhishen/ruleEngine/parser"
)

func (v *RuleEngineVisitor) VisitMapValue(ctx *parser.MapValueContext) interface{} {
	var hash = make(map[string]interface{}, ctx.GetChildCount())
	for _, element := range ctx.AllPair() {
		element.Accept(v)
		v, k := v.pop(), v.pop()
		hash[k.(string)] = v
	}

	v.push(hash)
	return nil
}

func (v *RuleEngineVisitor) VisitPair(ctx *parser.PairContext) interface{} {
	origin := ctx.MapKey().GetText()
	key := origin[1 : len(origin)-1]
	v.push(key)
	ctx.ValueType().Accept(v)

	return nil
}
