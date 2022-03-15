package functions

import (
	"github.com/yjhatfdu/expr/types"
	"strconv"
)

func cvtInt2Float(in []int64, out []float64) {
	for i := range in {
		out[i] = float64(in[i])
	}
}

func init() {
	toFloat, _ := NewFunction("toFloat")
	toFloat.Overload([]types.BaseType{types.Float}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toFloat.Overload([]types.BaseType{types.Int}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableFloat{}
		input := vectors[0].(*types.NullableInt)
		output.Init(input.Length())
		output.IsScalaV = input.IsScalaV
		cvtInt2Float(input.Values, output.Values)
		output.IsNullArr = input.IsNullArr
		output.FilterArr = input.FilterArr
		return output, nil
	})
	toFloat.Overload([]types.BaseType{types.Numeric}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableFloat{}
		input := vectors[0].(*types.NullableNumeric)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, input.Index(i).(types.Decimal).ToFloat(), false)
			return nil
		})
	})
	toFloat.Overload([]types.BaseType{types.Text}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableFloat{}
		input := vectors[0].(*types.NullableText)
		return BroadCast1(vectors[0], output, func(i int) error {
			f, err := strconv.ParseFloat(input.Index(i).(string), 64)
			if err != nil {
				return err
			}
			output.Set(i, f, false)
			return nil
		})
	})
}
