package functions

import "github.com/yjhatfdu/expr/types"

func init() {
	minusFunc, _ := NewFunction("minus")
	minusFunc.Overload([]types.BaseType{types.Int, types.Int}, types.Int, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableInt{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]-right.Values[j], false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Int, types.Float}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, float64(left.Values[i])-right.Values[j], false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Float, types.Int}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]-float64(right.Values[j]), false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Float, types.Float}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+right.Values[j], false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Numeric, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableNumeric)
		s := types.NumericScale(left.Scale, right.Scale)
		output.Scale = s
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.NormalizeNumeric(left.Values[i], left.Scale, s)-types.NormalizeNumeric(right.Values[j], left.Scale, s), false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Numeric, types.Int}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableInt)
		output.Scale = left.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]-types.Int2numeric(right.Values[j], left.Scale), false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Int, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = right.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.Int2numeric(left.Values[i], right.Scale)-right.Values[j], false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Numeric, types.Float}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableFloat)
		output.Scale = left.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]-types.Float2numeric(right.Values[j], left.Scale), false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Float, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = right.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.Float2numeric(left.Values[i], right.Scale)-right.Values[j], false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Timestamp, types.Timestamp}, types.Interval, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableTimestamp{TsType: types.Interval}
		output.TsType = types.Interval
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]-right.Values[j], false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Timestamp, types.Interval}, types.Timestamp, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableTimestamp{TsType: types.Timestamp}
		output.TsType = types.Timestamp
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]-right.Values[j], false)
			return nil
		})
	})
	minusFunc.Overload([]types.BaseType{types.Interval, types.Interval}, types.Interval, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableTimestamp{TsType: types.Interval}
		output.TsType = types.Interval
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]-right.Values[j], false)
			return nil
		})
	})
}
