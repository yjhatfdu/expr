package expr

import (
	"expr/types"
)

type AstNode struct {
	Children  []*AstNode
	NodeType  int
	Value     string
	ValueType types.BaseType
	OutType   types.BaseType
}

func newAst(t int, v string, vt types.BaseType, children ...*AstNode) *AstNode {
	return &AstNode{
		Children:  children,
		NodeType:  t,
		Value:     v,
		ValueType: vt,
	}
}
