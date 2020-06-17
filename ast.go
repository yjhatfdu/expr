package expr

import (
	"github.com/yjhatfdu/expr/types"
)

type AstNode struct {
	Children  []*AstNode
	NodeType  int
	Value     string
	ValueType types.BaseType
	OutType   types.BaseType
	Offset    int
	Length    int
}

func newAst(t int, v string, vt types.BaseType, pos int, children ...*AstNode) *AstNode {
	return &AstNode{
		Children:  children,
		NodeType:  t,
		Value:     v,
		ValueType: vt,
		Offset:    pos,
	}
}

func (an *AstNode) SetOffset(offset int) *AstNode {
	an.Offset = offset
	return an
}
