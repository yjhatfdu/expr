package functions

import (
	"fmt"
	"github.com/yjhatfdu/expr/types"
	"strings"
)

type typeRule struct {
	input  []types.BaseType
	output types.BaseType
}

type Handler func([]types.INullableVector) (types.INullableVector, error)

func (h Handler) Handle(v []types.INullableVector) (types.INullableVector, error) {
	return h(v)
}

type handlerFunction struct {
	OutputType types.BaseType
	Handler    IHandler
	Argc       int
}
type IHandler interface {
	Handle([]types.INullableVector) (types.INullableVector, error)
}

func types2Names(typeInput []types.BaseType) []string {
	var typeNames []string
	for _, t := range typeInput {
		typeNames = append(typeNames, types.GetTypeName(t))
	}
	return typeNames
}

var functions = map[string]*Function{}

func GetFunction(name string, inputTypes []types.BaseType) (*handlerFunction, error) {
	if f, ok := functions[name]; ok {
		h, err := f.Match(inputTypes)
		if err != nil {
			return nil, err
		}
		if h == nil {
			return nil, fmt.Errorf("function overload %s(%s) not found, avaliables:\n%s", name, strings.Join(types2Names(inputTypes), ", "), f.Print())
		} else {
			return h, nil
		}
	} else {
		return nil, fmt.Errorf("function with name '%s' not defined", name)
	}
}

type Function struct {
	name             string
	typeRules        []typeRule
	handlers         []IHandler
	genericValidator func([]types.BaseType) (types.BaseType, error)
	genericHandler   IHandler
	isAggregation    bool
	comment          string
}

func NewFunction(name string) (*Function, error) {
	if functions[name] != nil {
		return nil, fmt.Errorf("function with name '%s' already exists", name)
	}
	f := &Function{
		name:      name,
		typeRules: make([]typeRule, 0),
		handlers:  make([]IHandler, 0),
	}
	functions[name] = f
	return f, nil
}

func (f *Function) Comment(c string) {
	f.comment = c
}

func (f *Function) Print() string {
	output := ""
	if f.comment != "" {
		output = "//" + f.comment
	}
	for _, tr := range f.typeRules {
		var typeNames = types2Names(tr.input)
		output += fmt.Sprintf("%s(%s):%s\n", f.name, strings.Join(typeNames, ", "), types.GetTypeName(tr.output))
	}
	if f.genericValidator != nil {
		output += fmt.Sprintf("generic function %s(any...):any", f.name)
	}
	return output
}

func (f *Function) Overload(inputTypes []types.BaseType, output types.BaseType, implementation Handler) {
	tr := typeRule{
		input:  inputTypes,
		output: output,
	}
	f.typeRules = append(f.typeRules, tr)
	f.handlers = append(f.handlers, implementation)
}

func (f *Function) OverloadHandler(inputTypes []types.BaseType, output types.BaseType, implementation IHandler) {
	tr := typeRule{
		input:  inputTypes,
		output: output,
	}
	f.typeRules = append(f.typeRules, tr)
	f.handlers = append(f.handlers, implementation)
}

func (f *Function) Generic(typeValidator func([]types.BaseType) (types.BaseType, error), implementation Handler) {
	if f.genericValidator != nil {
		panic("redeclare generic function " + f.name)
	}
	f.genericValidator = typeValidator
	f.genericHandler = implementation
}

func (f *Function) GenericHandler(typeValidator func([]types.BaseType) (types.BaseType, error), implementation IHandler) {
	if f.genericValidator != nil {
		panic("redeclare generic function " + f.name)
	}
	f.genericValidator = typeValidator
	f.genericHandler = implementation
}

func (f *Function) Match(inputTypes []types.BaseType) (*handlerFunction, error) {
	for i, tr := range f.typeRules {
		if len(tr.input) != len(inputTypes) {
			continue
		}
		isAllScala := true
		for j := range tr.input {
			isAllScala = isAllScala && inputTypes[j] > types.ScalaOffset
			if tr.input[j] != types.Any && tr.input[j] != inputTypes[j] && tr.input[j]+types.ScalaOffset != inputTypes[j] {
				goto next
			}
		}
		if isAllScala && tr.output < types.ScalaTypes {
			return &handlerFunction{
				OutputType: tr.output + types.ScalaOffset,
				Handler:    f.handlers[i],
				Argc:       len(tr.input),
			}, nil
		} else {
			return &handlerFunction{
				OutputType: tr.output,
				Handler:    f.handlers[i],
				Argc:       len(tr.input),
			}, nil
		}
	next:
	}
	if f.genericValidator != nil {
		t, err := f.genericValidator(inputTypes)
		if err != nil {
			return nil, err
		}
		return &handlerFunction{
			OutputType: t,
			Handler:    f.genericHandler,
			Argc:       len(inputTypes),
		}, nil
	}
	return nil, nil
}

//func BroadCast2Bool(left, right types.INullableVector, f func(i, j bool) (bool, error)) (types.INullableVector, error) {
//	var ll = left.Length()
//	var rl = right.Length()
//	var lIsScala = left.IsScala()
//	var rIsScala = right.IsScala()
//	if ll != rl && !lIsScala && !rIsScala {
//		return nil, fmt.Errorf(`invalid BroadCast2 between two vectors, first length is %d, second length is %d`, left.Length(), right.Length())
//	}
//	var length = ll
//	if rl > ll {
//		length = rl
//	}
//	output := &types.NullableBool{}
//	output.Init(length)
//	//values := make([]bool, length)
//	//isNull := make([]bool, length)
//	leftTruty := left.TruthyArr()
//	rightTruty := right.TruthyArr()
//	for i := 0; i < length; i++ {
//		li := i
//		ri := i
//		if lIsScala {
//			li = 0
//		}
//		if rIsScala {
//			ri = 0
//		}
//		leftv := leftTruty[li]
//		rightv := rightTruty[ri]
//		out, err := f(leftv, rightv)
//		if err != nil {
//			output.SetNull(i, true)
//			output.AddError(&types.VectorError{
//				Index: i,
//				Error: err,
//			})
//		}
//		output.Values[i] = out
//	}
//	return output, nil
//}

func BroadCast2(left, right, output types.INullableVector, f func(index, i, j int) error) (types.INullableVector, error) {
	var ll = left.Length()
	var rl = right.Length()
	var lIsScala = left.IsScala()
	var rIsScala = right.IsScala()
	if ll != rl && !lIsScala && !rIsScala {
		return nil, fmt.Errorf(`invalid BroadCast2 between two vectors, first length is %d, second length is %d`, left.Length(), right.Length())
	}
	var length = ll
	if rl > ll {
		length = rl
	}
	output.Init(length)
	if lIsScala && rIsScala {
		output.SetScala(true)
	}
	oNullArr := output.GetIsNullArr()
	lNullArr := left.GetIsNullArr()
	rNullArr := right.GetIsNullArr()
	for i := 0; i < length; i++ {
		li := i
		ri := i
		if lIsScala {
			li = 0
		}
		if rIsScala {
			ri = 0
		}
		if lNullArr[li] || rNullArr[ri] {
			oNullArr[i] = true
		} else {
			err := f(i, li, ri)
			if err != nil {
				output.SetNull(i, true)
				output.AddError(&types.VectorError{
					Index: i,
					Error: err,
				})
			}
		}
	}
	output.SetFilterArr(CalFilterMask([][]bool{left.GetFilterArr(), right.GetFilterArr()}))
	return output, nil
}

func BroadCast1(in types.INullableVector, output types.INullableVector, handler func(i int) error) (types.INullableVector, error) {
	l := in.Length()
	output.Init(l)
	if in.IsScala() {
		output.SetScala(true)
	}
	isNullArr := in.GetIsNullArr()
	outIsNullArr := output.GetIsNullArr()
	for i := 0; i < l; i++ {
		if isNullArr[i] == true {
			outIsNullArr[i] = true
		} else {
			err := handler(i)
			if err != nil {
				output.SetNull(i, true)
				output.AddError(&types.VectorError{
					Index: i,
					Error: err,
				})
			}
		}
	}
	output.SetFilterArr(in.GetFilterArr())
	return output, nil
}

func BroadCastMultiGeneric(input []types.INullableVector, outputType types.BaseType, handler func(values []interface{}, index int) (interface{}, error)) (types.INullableVector, error) {
	var output types.INullableVector
	switch outputType {
	case types.Int:
		output = &types.NullableInt{}
	case types.Float:
		output = &types.NullableFloat{}
	case types.Bool:
		output = &types.NullableBool{}
	case types.Text:
		output = &types.NullableText{}
	case types.Timestamp, types.Time, types.Date:
		output = &types.NullableTimestamp{}
	default:
		panic("should not happend")
	}
	maxLength := 0
	allIsScala := true
	for _, in := range input {
		if l := in.Length(); l > maxLength {
			maxLength = l
		}
		allIsScala = allIsScala && in.IsScala()
	}
	output.Init(maxLength)
	row := make([]interface{}, len(input))
	for i := 0; i < maxLength; i++ {
		for j := 0; j < len(input); j++ {
			row[j] = input[j].Index(i)
		}
		out, err := handler(row, i)
		if err != nil {
			output.SetNull(i, true)
			output.AddError(&types.VectorError{
				Index: i,
				Error: err,
			})
		} else {
			output.Seti(i, out)
		}
	}
	return output, nil
}

func CalFilterMask(mask [][]bool) []bool {
	var result []bool
	for _, m := range mask {
		if m != nil {
			if result == nil {
				result = make([]bool, len(m), cap(m))
			}
			orBool(result, m, result)
		}
	}
	return result
}

func PrintAllFunctions() string {
	s := make([]string, 0)
	for _, f := range functions {
		s = append(s, f.Print())
	}
	return strings.Join(s, "\n")
}
