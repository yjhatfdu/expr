package functions

import "github.com/yjhatfdu/expr/types"

func andBool(data1, data2, out []bool)
func andBoolS(data1, out []bool, bools bool)

func init() {
	addFunc, _ := NewFunction("and")
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
			out.Values[0] = truty1[0] && truty2[0]
		} else if leftIsScalar {
			andBoolS(truty2, out.Values, truty1[0])
		} else if rightIsScalar {
			andBoolS(truty1, out.Values, truty2[0])
		} else {
			andBool(truty1, truty2, out.Values)
		}
		out.FilterArr = CalFilterMask([][]bool{vectors[0].GetFilterArr(), vectors[1].GetFilterArr()})
		return out, nil
	})
}
