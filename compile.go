package expr

import (
	"errors"
	"fmt"
	"github.com/yjhatfdu/expr/functions"
	"github.com/yjhatfdu/expr/types"
	"strconv"
	"strings"
	"sync"
)

type Program struct {
	InputTypes []types.BaseType
	OutputType types.BaseType
	opCode     []operation
}
type stack struct {
	list []types.INullableVector
	n    int
}

func (s *stack) popn(n int) ([]types.INullableVector, error) {
	if s.n-n < 0 {
		return nil, errors.New("pop empty stack")
	}
	s.n -= n
	return s.list[s.n : s.n+n], nil
}
func (s *stack) push(v types.INullableVector) {
	if s.n >= len(s.list) {
		s.list = append(s.list, v)
	} else {
		s.list[s.n] = v
	}
	s.n++
}

var stackPool = sync.Pool{
	New: func() interface{} {
		return make([]types.INullableVector, 8)
	},
}

func (p *Program) Run(input []types.INullableVector, env map[string]string) (vector types.INullableVector, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
				return
			}
			err = errors.New(fmt.Sprintf("%v", err))
		}
	}()

	if len(input) != len(p.InputTypes) {
		return nil, errors.New("input argument length not match")
	}
	s := stack{
		list: stackPool.Get().([]types.INullableVector),
		n:    0,
	}
	defer stackPool.Put(s.list)
	for i := range p.opCode {
		op := p.opCode[i]
		switch op.op {
		case CONST:
			s.push(op.v)
		case VAR:
			s.push(input[op.varIndex])
		case FUNC:
			args, err := s.popn(op.argc)
			if err != nil {
				return nil, err
			}
			ret, err := op.handler.Handle(args, env)
			if err != nil {
				return nil, err
			}
			s.push(ret)
		}
	}
	popn, err := s.popn(1)
	if err != nil {
		return nil, err
	}
	return popn[0], nil
}

type operation struct {
	op       int
	argc     int
	v        types.INullableVector
	varIndex int
	handler  functions.IHandler
}

type context struct {
	ops  []operation
	code string
}

func (ct *context) addOperation(op operation) {
	ct.ops = append(ct.ops, op)
}

func Compile(code string, inputType []types.BaseType, env map[string]string) (p *Program, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	l := NewLexer(code)
	yyErrorVerbose = true
	yyParse(l)
	ast := l.parseResult
	ctx := &context{ops: []operation{}, code: code}
	err = compile(ast, ctx, inputType, env)
	if err != nil {
		return
	}
	p = &Program{
		InputTypes: inputType,
		OutputType: ast.OutType,
		opCode:     ctx.ops,
	}
	return
}

func buildErrInfo(node *AstNode, code string) string {
	return fmt.Sprintf("\n%s\n%s", code, strings.Repeat(" ", node.Offset)+"^")
}

func compile(an *AstNode, ctx *context, inputType []types.BaseType, env map[string]string) error {
	for _, node := range an.Children {
		if err := compile(node, ctx, inputType, env); err != nil {
			return err
		}
	}
	switch an.NodeType {
	case CONST:
		switch an.ValueType {
		case types.Null:
			ctx.addOperation(operation{
				op:   CONST,
				argc: 0,
				v: &types.NullVector{
					IsNullArr: []bool{true},
					FilterArr: []bool{false},
					IsScalaV:  true,
					Values:    []string{"NULL"},
				},
			})
			an.OutType = types.Any
		case types.Int:
			v, err := strconv.ParseInt(an.Value, 10, 64)
			if err != nil {
				return fmt.Errorf("compile error:%s\ncaused by:%v", buildErrInfo(an, ctx.code), err)
			}
			ctx.addOperation(operation{
				op:   CONST,
				argc: 0,
				v: &types.NullableInt{
					NullableVector: types.NullableVector{
						IsScalaV:  true,
						IsNullArr: []bool{false},
					},
					Values: []int64{v},
				},
			})
			an.OutType = types.IntS
		case types.Float:
			v, err := strconv.ParseFloat(an.Value, 64)
			if err != nil {
				return fmt.Errorf("compile error:%s\ncaused by:%v", buildErrInfo(an, ctx.code), err)
			}
			ctx.addOperation(operation{
				op:   CONST,
				argc: 0,
				v: &types.NullableFloat{
					NullableVector: types.NullableVector{
						IsScalaV:  true,
						IsNullArr: []bool{false},
					},
					Values: []float64{v},
				},
			})
			an.OutType = types.FloatS
		case types.Text:
			var str string
			if strings.HasPrefix(an.Value, `"`) && strings.HasSuffix(an.Value, `"`) {
				str = an.Value[1 : len(an.Value)-1]
			} else if strings.HasPrefix(an.Value, `'`) && strings.HasSuffix(an.Value, `'`) {
				str = an.Value[1 : len(an.Value)-1]
			}
			str = strings.ReplaceAll(str, "\\\\", "\\")
			ctx.addOperation(operation{
				op:   CONST,
				argc: 0,
				v: &types.NullableText{
					NullableVector: types.NullableVector{
						IsScalaV:  true,
						IsNullArr: []bool{false},
					},
					Values: []string{str},
				},
			})
			an.OutType = types.TextS
		case RAWSTR:
			ctx.addOperation(operation{
				op:   CONST,
				argc: 0,
				v: &types.NullableText{
					NullableVector: types.NullableVector{
						IsScalaV:  true,
						IsNullArr: []bool{false},
					},
					Values: []string{an.Value},
				},
			})
			an.OutType = types.TextS
		case types.Bool:
			v, err := strconv.ParseBool(an.Value)
			if err != nil {
				return fmt.Errorf("compile error:%s\ncaused by:%v", buildErrInfo(an, ctx.code), err)
			}
			ctx.addOperation(operation{
				op:   CONST,
				argc: 0,
				v: &types.NullableBool{
					NullableVector: types.NullableVector{
						IsScalaV:  true,
						IsNullArr: []bool{false},
					},
					Values: []bool{v},
				},
			})
			an.OutType = types.BoolS
		}
	case VAR:
		if len(inputType) == 0 {
			err := fmt.Errorf("cannot reference '$%s' of zero input arguments expression", an.Value)
			return fmt.Errorf("compile error:%s\ncaused by:%v", buildErrInfo(an, ctx.code), err)
		}
		if an.Value == "ALL" {
			for i := range inputType {
				ctx.addOperation(operation{
					op:       VAR,
					argc:     0,
					v:        nil,
					varIndex: i,
				})
			}
		} else {
			varIndex, err := strconv.Atoi(an.Value)
			if err != nil {
				err := fmt.Errorf("invalid variable syntax '$%s'", an.Value)
				return fmt.Errorf("compile error:%s\ncaused by:%v", buildErrInfo(an, ctx.code), err)
			}
			if varIndex > len(inputType) {
				err := fmt.Errorf("variable index '$%s' out of input argument range '$1-$%d'", an.Value, len(inputType))
				return fmt.Errorf("compile error:%s\ncaused by:%v", buildErrInfo(an, ctx.code), err)
			}
			ctx.addOperation(operation{
				op:       VAR,
				argc:     0,
				v:        nil,
				varIndex: varIndex - 1,
			})
			an.OutType = inputType[varIndex-1]
		}
	case FUNC:
		inputTypes := make([]types.BaseType, len(an.Children))
		if len(an.Children) == 1 && an.Children[0].NodeType == VAR && an.Children[0].Value == "ALL" {
			inputTypes = inputType
		} else {
			for i, c := range an.Children {
				inputTypes[i] = c.OutType
			}
		}
		consts := make([]string, len(an.Children))
		for i, c := range an.Children {
			if c.NodeType == CONST {
				consts[i] = c.Value
			}
		}
		f, err := functions.GetFunction(an.Value, inputTypes)
		if err != nil {
			return fmt.Errorf("compile error:%s\ncaused by:%v", buildErrInfo(an, ctx.code), err)
		}
		err = f.Handler.Init(consts, env)
		if err != nil {
			return err
		}
		ctx.addOperation(operation{
			op:      FUNC,
			argc:    len(inputTypes),
			handler: f.Handler,
		})
		an.OutType = f.OutputType
	default:
		panic("unknown operator")
	}
	return nil
}
