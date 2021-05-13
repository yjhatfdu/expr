package functions

import (
	"github.com/yjhatfdu/expr/types"
	"math"
)

func init() {
	neq, _ := NewFunction("neq")
	neq.Overload([]types.BaseType{types.Int, types.Int}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i] != right.Values[j], false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Float, types.Float}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, math.Abs(left.Values[i]-right.Values[j]) <= math.SmallestNonzeroFloat64, false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Text, types.Text}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableText)
		right := vectors[1].(*types.NullableText)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i] != right.Values[j], false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Numeric, types.Numeric}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableNumeric)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareNumeric(left.Values[i], left.Scale, right.Values[i], right.Scale) != 0, false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Numeric, types.Int}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareNumericInt(left.Values[i], left.Scale, right.Values[j]) != 0, false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Int, types.Numeric}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableNumeric)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareNumericInt(right.Values[i], right.Scale, left.Values[j]) != 0, false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Numeric, types.Float}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareNumericFloat(left.Values[i], left.Scale, right.Values[j]) != 0, false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Float, types.Numeric}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableNumeric)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.CompareNumericFloat(right.Values[i], right.Scale, left.Values[j]) != 0, false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Timestamp, types.Timestamp}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i] != right.Values[j], false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Date, types.Date}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i] != right.Values[j], false)
			return nil
		})
	})
	neq.Overload([]types.BaseType{types.Time, types.Time}, types.Bool, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableBool{}
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i] != right.Values[j], false)
			return nil
		})
	})
}
