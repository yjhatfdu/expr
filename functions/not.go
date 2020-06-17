package functions

import "github.com/yjhatfdu/expr/types"

func init() {
	addFunc, _ := NewFunction("not")
	addFunc.Overload([]types.BaseType{types.Any}, types.Bool, func(vectors []types.INullableVector) (types.INullableVector, error) {
		v := vectors[0]
		length := v.Length()
		values := make([]bool, length)
		isNull := make([]bool, length)
		truthArr := v.TruthyArr()
		for i := 0; i < length; i++ {
			values[i] = !truthArr[i]
		}
		return &types.NullableBool{
			NullableVector: types.NullableVector{
				IsNullArr: isNull,
				IsScalaV:  v.IsScala(),
			},
			Values: values,
		}, nil
	})
}
