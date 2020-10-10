package functions

import "github.com/yjhatfdu/expr/types"

func orBool(data1, data2, out []bool)
func orBoolS(data1, out []bool, bools bool)
func init() {
	addFunc, _ := NewFunction("or")
	addFunc.Overload([]types.BaseType{types.Any, types.Any}, types.Bool, func(vectors []types.INullableVector) (types.INullableVector, error) {
		truty1 := vectors[0].TruthyArr()
		truty2 := vectors[1].TruthyArr()
		length := len(truty1)
		if len(truty2) > length {
			length = len(truty2)
		}
		out := &types.NullableBool{}
		out.Init(length)
		leftIsScalar := vectors[0].IsScala()
		rightIsScalar := vectors[1].IsScala()
		if leftIsScalar && rightIsScalar {
			out.IsScalaV = true
			out.Values[0] = truty1[0] || truty2[0]
		} else if leftIsScalar {
			orBoolS(truty2, out.Values, truty1[0])
		} else if rightIsScalar {
			orBoolS(truty1, out.Values, truty2[0])
		} else {
			orBool(truty1, truty2, out.Values)
		}
		out.FilterArr = calFilterMask([][]bool{vectors[0].GetFilterArr(), vectors[1].GetFilterArr()})
		return out, nil
	})
}
