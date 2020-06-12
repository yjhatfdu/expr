package functions

import "expr/types"

func init() {
	addFunc, _ := NewFunction("minus")
	addFunc.Overload([]types.BaseType{types.Int, types.Int}, types.Int, func(vectors []types.INullableVector) (types.INullableVector, error) {
		output := types.NullableInt{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]-right.Values[j], false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Int, types.Float}, types.Float, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, float64(left.Values[i])-right.Values[j], false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Float, types.Int}, types.Float, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]-float64(right.Values[j]), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Float, types.Float}, types.Float, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+right.Values[j], false)
			return nil
		})
	})
}
