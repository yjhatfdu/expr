package functions

import "expr/types"

func not(arr []bool) {
	for i := 0; i < len(arr); i++ {
		arr[i] = !arr[i]
	}
}

func init() {
	addFunc, _ := NewFunction("not")
	addFunc.Overload([]types.BaseType{types.Any}, types.Bool, func(vectors []types.INullableVector) (types.INullableVector, error) {
		v := vectors[0]
		length := v.Length()
		out := &types.NullableBool{
			NullableVector: types.NullableVector{
				IsNullArr: make([]bool, length, 32*(length/32+1)),
				IsScalaV:  v.IsScala(),
			},
			Values: nil,
		}
		truthArr := v.TruthyArr()
		not(truthArr)
		out.Values = truthArr
		return out, nil
	})
}
