package functions

import (
	"expr/types"
	"fmt"
	"strings"
)

type typeRule struct {
	input  []types.BaseType
	output types.BaseType
}

type Handler func([]types.INullableVector) (types.INullableVector, error)
type handlerFunction struct {
	OutputType types.BaseType
	Handler    Handler
	Argc       int
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
		h := f.Match(inputTypes)
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
	handlers         []Handler
	genericValidator func([]types.BaseType) (types.BaseType, error)
	genericHandler   Handler
}

func NewFunction(name string) (*Function, error) {
	if functions[name] != nil {
		return nil, fmt.Errorf("function with name '%s' already exists", name)
	}
	f := &Function{
		name:      name,
		typeRules: make([]typeRule, 0),
		handlers:  make([]Handler, 0),
	}
	functions[name] = f
	return f, nil
}

func (f *Function) Print() string {
	output := ""
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

func (f *Function) Generic(typeValidator func([]types.BaseType) (types.BaseType, error), implementation Handler) {
	if f.genericValidator != nil {
		panic("redeclare generic function " + f.name)
	}
	f.genericValidator = typeValidator
	f.genericHandler = implementation
}

func (f *Function) Match(inputTypes []types.BaseType) *handlerFunction {
	for i, tr := range f.typeRules {
		if len(tr.input) != len(inputTypes) {
			continue
		}
		for j := range tr.input {
			if tr.input[j] != types.Any && tr.input[j] != inputTypes[j] {
				goto next
			}
		}
		return &handlerFunction{
			OutputType: tr.output,
			Handler:    f.handlers[i],
			Argc:       len(tr.input),
		}
	next:
	}
	if f.genericValidator != nil {
		t, err := f.genericValidator(inputTypes)
		if err != nil {
			return nil
		}
		return &handlerFunction{
			OutputType: t,
			Handler:    f.genericHandler,
			Argc:       len(inputTypes),
		}
	}
	return nil
}
func BroadCast2Bool(left, right types.INullableVector, f func(i, j bool) (bool, error)) (types.INullableVector, error) {
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
	values := make([]bool, length)
	isNull := make([]bool, length)
	leftTruty := left.TruthyArr()
	rightTruty := right.TruthyArr()
	for i := 0; i < length; i++ {
		li := i
		ri := i
		if lIsScala {
			li = 0
		}
		if rIsScala {
			ri = 0
		}
		leftv := leftTruty[li]
		rightv := rightTruty[ri]
		out, err := f(leftv, rightv)
		if err != nil {
			return nil, err
		}
		values[i] = out
	}
	return &types.NullableBool{
		Values: values,
		NullableVector: types.NullableVector{
			IsNullArr: isNull,
			IsScalaV:  rIsScala && lIsScala,
		},
	}, nil
}

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
				return nil, err
			}
		}
	}
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
				return nil, err
			}
		}
	}
	return output, nil
}

func BroadCastMultiGeneric(input []types.INullableVector, outputType types.BaseType, handler func(values []interface{}) (interface{}, error)) (types.INullableVector, error) {
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
		out, err := handler(row)
		if err != nil {
			return nil, err
		}
		output.Seti(i, out)
	}
	return output, nil
}
