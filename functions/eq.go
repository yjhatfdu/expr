package functions

import (
	"github.com/yjhatfdu/expr/types"
	"math"
)

func init() {
	eq, _ := NewFunction("eq")
	eq.Overload([]types.BaseType{types.Int, types.Int}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(int64) == right.Index(j).(int64), false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Float, types.Float}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, math.Abs(left.Index(i).(float64)-right.Index(j).(float64)) <= math.SmallestNonzeroFloat64, false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Text, types.Text}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableText)
		right := vectors[1].(*types.NullableText)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(string) == right.Index(j).(string), false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Numeric, types.Numeric}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableNumeric)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareDecimal(left.Values[i], right.Values[j]) == 0, false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Numeric, types.Int}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareDecimal(left.Values[i], types.Int2Decimal(right.Values[j], 0)) == 0, false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Int, types.Numeric}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableNumeric)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareDecimal(types.Int2Decimal(left.Values[i], 0), right.Values[j]) == 0, false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Numeric, types.Float}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareDecimal(left.Values[i], types.Float2Decimal(right.Values[j], left.Scale)) == 0, false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Float, types.Numeric}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableNumeric)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareDecimal(types.Float2Decimal(left.Values[i], right.Scale), right.Values[j]) == 0, false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Timestamp, types.Timestamp}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(int64) == right.Index(j).(int64), false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Date, types.Date}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(int64) == right.Index(j).(int64), false)
			return nil
		})
	})
	eq.Overload([]types.BaseType{types.Time, types.Time}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(int64) == right.Index(j).(int64), false)
			return nil
		})
	})
}
