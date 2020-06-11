package expr

import (
	"errors"
	"expr/functions"
	"expr/types"
	"fmt"
	"strconv"
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

func (s *stack) popn(n int) []types.INullableVector {
	if s.n-n < 0 {
		panic("pop empty stack")
	}
	s.n -= n
	return s.list[s.n : s.n+n]
}
func (s *stack) push(v types.INullableVector) {
	if s.n >= len(s.list) {
		s.list = append(s.list, v)
	} else {
		s.list[s.n] = v
	}
	s.n++
}

func (p *Program) Run(input []types.INullableVector) (types.INullableVector, error) {
	if len(input) != len(p.InputTypes) {
		return nil, errors.New("input argument length not match")
	}
	//for i := range input {
	//	if input[i].Type()
	//}
	s := stack{
		list: make([]types.INullableVector, 8),
		n:    0,
	}
	for i := range p.opCode {
		op := p.opCode[i]
		switch op.op {
		case CONST:
			s.push(op.v)
		case VAR:
			s.push(input[op.varIndex])
		case FUNC:
			args := s.popn(op.argc)
			ret, err := op.handler(args)
			if err != nil {
				return nil, err
			}
			s.push(ret)
		}
	}
	return s.popn(1)[0], nil
}

type operation struct {
	op       int
	argc     int
	v        types.INullableVector
	varIndex int
	handler  functions.Handler
}

type context struct {
	ops []operation
}

func (ct *context) addOperation(op operation) {
	ct.ops = append(ct.ops, op)
}

func Compile(code string, inputType []types.BaseType) (p *Program, err error) {
	//defer func() {
	//	r := recover()
	//	if r != nil {
	//		err = errors.New(fmt.Sprintf("%v", r))
	//	}
	//}()
	l := NewLexer(code)
	yyErrorVerbose = true
	yyParse(l)
	ast := l.parseResult
	ctx := &context{ops: []operation{}}
	err = compile(ast, ctx, inputType)
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

func compile(an *AstNode, ctx *context, inputType []types.BaseType) error {
	for _, node := range an.Children {
		if err := compile(node, ctx, inputType); err != nil {
			return err
		}
	}
	switch an.NodeType {
	case CONST:
		switch an.ValueType {
		case types.Int:
			v, err := strconv.ParseInt(an.Value, 10, 64)
			if err != nil {
				return err
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
			an.OutType = types.Int
		case types.Float:
			v, err := strconv.ParseFloat(an.Value, 64)
			if err != nil {
				return err
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
			an.OutType = types.Float
		case types.Text:
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
			an.OutType = types.Float
		case types.Bool:
			v, err := strconv.ParseBool(an.Value)
			if err != nil {
				return err
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
			an.OutType = types.Float
		}
	case VAR:
		if len(inputType) == 0 {
			return fmt.Errorf("cannot reference '$%s' of zero input arguments expression", an.Value)
		}
		varIndex, err := strconv.Atoi(an.Value)
		if err != nil {
			return fmt.Errorf("invalid variable syntax '$%s'", an.Value)
		}
		if varIndex > len(inputType) {
			return fmt.Errorf("variable index '$%s' out of input argument range '$1-$%d'", an.Value, len(inputType))
		}
		ctx.addOperation(operation{
			op:       VAR,
			argc:     0,
			v:        nil,
			varIndex: varIndex - 1,
		})
		an.OutType = inputType[varIndex-1]
	case FUNC:
		inputTypes := make([]types.BaseType, len(an.Children))
		for i, c := range an.Children {
			inputTypes[i] = c.OutType
		}
		f, err := functions.GetFunction(an.Value, inputTypes)
		if err != nil {
			return err
		}
		ctx.addOperation(operation{
			op:      FUNC,
			argc:    f.Argc,
			handler: f.Handler,
		})
		an.OutType = f.OutputType
	default:
		panic("unknown operator")
	}
	return nil
}
