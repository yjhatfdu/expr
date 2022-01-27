package functions

import (
	"github.com/yjhatfdu/expr/types"
	"strings"
)

func init() {
	mulFunc, _ := NewFunction("mul")
	mulFunc.Overload([]types.BaseType{types.Int, types.Int}, types.Int, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableInt{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(int64)*right.Index(j).(int64), false)
			return nil
		})
	})
	mulFunc.Overload([]types.BaseType{types.Int, types.Float}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, float64(left.Index(i).(int64))*right.Index(j).(float64), false)
			return nil
		})
	})
	mulFunc.Overload([]types.BaseType{types.Float, types.Int}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(float64)*float64(right.Index(j).(int64)), false)
			return nil
		})
	})
	mulFunc.Overload([]types.BaseType{types.Float, types.Float}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(float64)*right.Index(j).(float64), false)
			return nil
		})
	})
	mulFunc.Overload([]types.BaseType{types.Numeric, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = left.Scale + right.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(int64)*right.Index(j).(int64), false)
			return nil
		})
	})
	mulFunc.Overload([]types.BaseType{types.Numeric, types.Int}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableInt)
		output.Scale = left.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(int64)*types.Int2numeric(right.Index(j).(int64), left.Scale), false)
			return nil
		})
	})
	mulFunc.Overload([]types.BaseType{types.Int, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = right.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.Int2numeric(left.Index(i).(int64), right.Scale)*right.Index(j).(int64), false)
			return nil
		})
	})
	mulFunc.Overload([]types.BaseType{types.Numeric, types.Float}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableFloat)
		output.Scale = left.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Index(i).(int64)*types.Float2numeric(right.Index(j).(float64), left.Scale), false)
			return nil
		})
	})
	mulFunc.Overload([]types.BaseType{types.Float, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = right.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.Float2numeric(left.Index(i).(float64), right.Scale)*right.Index(j).(int64), false)
			return nil
		})
	})
	mulFunc.Overload([]types.BaseType{types.Text, types.IntS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableText{}
		left := vectors[0].(*types.NullableText)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, strings.Repeat(left.Index(i).(string), int(right.Index(j).(int64))), false)
			return nil
		})
	})
}
