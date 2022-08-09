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
		output.Scale = 0
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Int2Decimal(input.Values[i], 0), false)
			return nil
		})
	})
	toNumeric.Overload([]types.BaseType{types.Int, types.IntS}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableInt)
		scale := int(vectors[1].(*types.NullableInt).Index(0).(int64))
		output.Scale = scale
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Int2Decimal(input.Values[i], scale), false)
			return nil
		})
	})
	toNumeric.Overload([]types.BaseType{types.Numeric, types.IntS}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableNumeric)
		scale := int(vectors[1].(*types.NullableInt).Index(0).(int64))
		output.Scale = scale
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, input.Values[i], false)
			return nil
		})
	})
	toNumeric.Overload([]types.BaseType{types.Float}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableFloat)
		scale := 0
		output.Scale = 0
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Float2Decimal(input.Values[i], scale), false)
			return nil
		})
	})
	toNumeric.Overload([]types.BaseType{types.Text}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableText)
		return BroadCast1(vectors[0], output, func(i int) error {
			t, err := types.Text2Decimal(input.Values[i])
			if err != nil {
				return err
			}
			output.Set(i, t, false)
			return nil
		})
	})
	toNumeric.Overload([]types.BaseType{types.Float, types.IntS}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableFloat)
		scale := int(vectors[1].(*types.NullableInt).Index(0).(int64))
		output.Scale = scale
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Float2Decimal(input.Values[i], scale), false)
			return nil
		})
	})
}
