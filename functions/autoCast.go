package functions

import (
	"fmt"
	"github.com/yjhatfdu/expr/types"
)

type castKey struct {
	source types.BaseType
	target types.BaseType
}

type CastFunc func(v types.INullableVector) types.INullableVector

var castRules map[castKey]func(v types.INullableVector) types.INullableVector

func init() {

	castRules = map[castKey]func(v types.INullableVector) types.INullableVector{
		castKey{types.Any, types.Int}: func(v types.INullableVector) types.INullableVector {
			ret := types.NullableInt{}
			ret.Init(1)
			ret.SetScala(true)
			ret.SetNull(0, true)
			return &ret
		}, castKey{types.Any, types.Time}: func(v types.INullableVector) types.INullableVector {
			ret := types.NullableTimestamp{}
			ret.TsType = types.Time
			ret.Init(1)
			ret.SetScala(true)
			ret.SetNull(0, true)
			return &ret
		}, castKey{types.Any, types.Bool}: func(v types.INullableVector) types.INullableVector {
			ret := types.NullableBool{}
			ret.Init(1)
			ret.SetScala(true)
			ret.SetNull(0, true)
			return &ret
		}, castKey{types.Any, types.Float}: func(v types.INullableVector) types.INullableVector {
			ret := types.NullableFloat{}
			ret.Init(1)
			ret.SetScala(true)
			ret.SetNull(0, true)
			return &ret
		}, castKey{types.Any, types.Numeric}: func(v types.INullableVector) types.INullableVector {
			ret := types.NullableNumeric{}
			ret.SetScala(true)
			ret.SetNull(0, true)
			return &ret
		}, castKey{types.Any, types.Date}: func(v types.INullableVector) types.INullableVector {
			ret := types.NullableTimestamp{}
			ret.Init(1)
			ret.TsType = types.Date
			ret.SetScala(true)
			ret.SetNull(0, true)
			return &ret
		}, castKey{types.Any, types.Interval}: func(v types.INullableVector) types.INullableVector {
			ret := types.NullableTimestamp{}
			ret.Init(1)
			ret.TsType = types.Interval
			ret.SetScala(true)
			ret.SetNull(0, true)
			return &ret
		}, castKey{types.Any, types.Text}: func(v types.INullableVector) types.INullableVector {
			ret := types.NullableText{}
			ret.Init(1)
			ret.SetScala(true)
			ret.SetNull(0, true)
			return &ret
		}, castKey{types.Int, types.Numeric}: func(v types.INullableVector) types.INullableVector {
			f, _ := GetFunction("toNumeric", []types.BaseType{types.Int})
			ret, _ := f.Handler.Handle([]types.INullableVector{v}, nil)
			return ret
		}, castKey{types.Float, types.Numeric}: func(v types.INullableVector) types.INullableVector {
			f, _ := GetFunction("toNumeric", []types.BaseType{types.Float})
			ret, _ := f.Handler.Handle([]types.INullableVector{v}, nil)
			return ret
		}, castKey{types.Int, types.Float}: func(v types.INullableVector) types.INullableVector {
			f, _ := GetFunction("toFloat", []types.BaseType{types.Int})
			ret, _ := f.Handler.Handle([]types.INullableVector{v}, nil)
			return ret
		}, castKey{types.Int, types.Text}: func(v types.INullableVector) types.INullableVector {
			f, _ := GetFunction("toText", []types.BaseType{types.Int})
			ret, _ := f.Handler.Handle([]types.INullableVector{v}, nil)
			return ret
		}, castKey{types.Numeric, types.Text}: func(v types.INullableVector) types.INullableVector {
			f, _ := GetFunction("toText", []types.BaseType{types.Numeric})
			ret, _ := f.Handler.Handle([]types.INullableVector{v}, nil)
			return ret
		},
	}
}
func AutoCast(v types.INullableVector, targetType types.BaseType) (types.INullableVector, error) {
	f, ok := castRules[castKey{v.Type(), targetType}]
	if !ok {
		return nil, fmt.Errorf("no auto cast for type %v to %v", v.Type(), targetType)
	}
	return f(v), nil
}

func matchAutoMatch(from types.BaseType, to types.BaseType) func(v types.INullableVector) types.INullableVector {
	return castRules[castKey{from, to}]
}
