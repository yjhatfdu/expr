package functions

import (
	"github.com/yjhatfdu/expr/types"
)

func init() {
	// default format is Numeric(12,4), use toNumeric(_,int) to specify scale
	toNumeric, _ := NewFunction("toNumeric")
	toNumeric.Overload([]types.BaseType{types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toNumeric.Overload([]types.BaseType{types.Int}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableInt)
		scale := 4
		output.Scale = 4
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.NormalizeNumeric(input.Index(i).(int64), 0, scale), false)
			return nil
		})
	})
	toNumeric.Overload([]types.BaseType{types.Int, types.IntS}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableInt)
		scale := int(vectors[1].(*types.NullableInt).Values[0])
		output.Scale = scale
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.NormalizeNumeric(input.Index(i).(int64), 0, scale), false)
			return nil
		})
	})
	toNumeric.Overload([]types.BaseType{types.Numeric, types.IntS}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableNumeric)
		scale := int(vectors[1].(*types.NullableInt).Values[0])
		output.Scale = scale
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.NormalizeNumeric(input.Index(i).(int64), input.Scale, scale), false)
			return nil
		})
	})
	toNumeric.Overload([]types.BaseType{types.Float}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableFloat)
		scale := 4
		output.Scale = 4
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Float2numeric(input.Index(i).(float64), scale), false)
			return nil
		})
	})
	toNumeric.Overload([]types.BaseType{types.Float, types.IntS}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableFloat)
		scale := int(vectors[1].(*types.NullableInt).Values[0])
		output.Scale = scale
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Float2numeric(input.Index(i).(float64), scale), false)
			return nil
		})
	})
}
