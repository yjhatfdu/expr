package functions

import (
	"errors"
	"github.com/yjhatfdu/expr/types"
)

func init() {
	addFunc, _ := NewFunction("div")
	addFunc.Overload([]types.BaseType{types.Int, types.Int}, types.Int, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableInt{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			if right.Index(j).(int64) == 0 {
				return errors.New("divide zero")
			}
			output.Set(index, left.Index(i).(int64)/right.Index(j).(int64), false)
			return nil
		})

	})
	addFunc.Overload([]types.BaseType{types.Int, types.Float}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, float64(left.Index(i).(int64))/right.Index(j).(float64), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Float, types.Int}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(float64)/float64(right.Index(j).(int64)), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Float, types.Float}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(float64)/right.Index(j).(float64), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Numeric, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = left.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			if right.Values[j].IsZero() {
				return errors.New("divide zero")
			}
			output.Set(index, types.DivideDecimal(left.Values[i], right.Values[j]), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Numeric, types.Int}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableInt)
		output.Scale = left.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			if right.Index(j).(int64) == 0 {
				return errors.New("divide zero")
			}
			output.Set(index, types.DivideDecimal(left.Values[i], types.Int2Decimal(right.Values[j], 0)), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Int, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = right.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			if right.Values[j].IsZero() {
				return errors.New("divide zero")
			}
			output.Set(index, types.DivideDecimal(types.Int2Decimal(left.Values[i], 0), right.Values[j]), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Numeric, types.Float}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableFloat)
		output.Scale = left.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			if right.Index(j).(float64) == 0 {
				return errors.New("divide zero")
			}
			output.Set(index, types.DivideDecimal(left.Values[i], types.Float2Decimal(right.Values[j], 0)), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Float, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = right.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			if right.Values[j].IsZero() {
				return errors.New("divide zero")
			}
			output.Set(index, types.DivideDecimal(types.Float2Decimal(left.Values[i], 0), right.Values[j]), false)
			return nil
		})
	})
}
