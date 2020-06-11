package functions

import "expr/types"

func init() {
	addFunc, _ := NewFunction("and")
	addFunc.Overload([]types.BaseType{types.Any, types.Any}, types.Bool, func(vectors []types.INullableVector) (types.INullableVector, error) {
		return BroadCast2Bool(vectors[0], vectors[1], func(i, j bool) (i2 bool, e error) {
			return i && j, nil
		})
	})
}
