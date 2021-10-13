package functions

import (
	"errors"
	"fmt"
	"github.com/yjhatfdu/expr/types"
	"math"
)

func init() {
	in, _ := NewFunction("in")
	in.Generic(func(inputTypes []types.BaseType) (baseType types.BaseType, e error) {
		if len(inputTypes) < 1 {
			return 0, errors.New("require at least 2 argument")
		}
		t := inputTypes[0]
		if t > types.ScalaTypes {
			t = t - types.ScalaOffset
		}
		for i := 1; i < len(inputTypes); i++ {
			if inputTypes[i] != t+types.ScalaOffset {
				return 0, errors.New(fmt.Sprintf("in函数的输入值参数和匹配值参数必须为相同类型，且匹配值必须为常量，实际输入值类型 %s, 期望类型 %s", types.GetTypeName(t), types.GetTypeName(inputTypes[i])))
			}
		}
		return types.Bool, nil
	}, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		var output types.NullableBool
		vec := vectors[0]
		output.Init(vec.Length())
		output.SetScala(vec.IsScala())
		output.IsNullArr = vec.GetIsNullArr()
		switch vectors[0].Type() {
		case types.Int:
			var inArr = make([]int64, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableInt).Values[0])
			}
			v := vec.(*types.NullableInt)
			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				for j := 0; j < len(inArr); j++ {
					if v.Values[i] == inArr[j] {
						output.Values[i] = true
						break
					}
				}
			}
		case types.Float:
			var inArr = make([]float64, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableFloat).Values[0])
			}
			v := vec.(*types.NullableFloat)
			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				for j := 0; j < len(inArr); j++ {
					if math.Abs(v.Values[i]-inArr[j]) <= math.SmallestNonzeroFloat64 {
						output.Values[i] = true
						break
					}
				}
			}
		case types.Bool:
			var inArr = make([]bool, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableBool).Values[0])
			}
			v := vec.(*types.NullableBool)
			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				for j := 0; j < len(inArr); j++ {
					if v.Values[i] == inArr[j] {
						output.Values[i] = true
						break
					}
				}
			}
		case types.Text:
			var inArr = make([]string, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableText).Values[0])
			}
			v := vec.(*types.NullableText)
			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				for j := 0; j < len(inArr); j++ {
					if v.Values[i] == inArr[j] {
						output.Values[i] = true
						break
					}
				}
			}
		case types.Numeric:
			var inArr = make([]int64, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableNumeric).Values[0])
			}
			v := vec.(*types.NullableNumeric)
			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				for j := 0; j < len(inArr); j++ {
					if v.Values[i] == inArr[j] {
						output.Values[i] = true
						break
					}
				}
			}
		default:
			panic("should not happend")
		}
		(&output).SetFilterArr(vectors[0].GetFilterArr())
		return &output, nil
	})
}

func init() {
	notin, _ := NewFunction("notIn")
	notin.Generic(func(inputTypes []types.BaseType) (baseType types.BaseType, e error) {
		if len(inputTypes) < 1 {
			return 0, errors.New("require at least 2 argument")
		}
		t := inputTypes[0]
		if t > types.ScalaTypes {
			t = t - types.ScalaOffset
		}
		for i := 1; i < len(inputTypes); i++ {
			if inputTypes[i] != t+types.ScalaOffset {
				return 0, errors.New("参数必须为同样类型，且为常量")
			}
		}
		return types.Bool, nil
	}, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		var output types.NullableBool
		vec := vectors[0]
		output.Init(vec.Length())
		output.SetScala(vec.IsScala())
		output.IsNullArr = vec.GetIsNullArr()
		switch vectors[0].Type() {
		case types.Int:
			var inArr = make([]int64, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableInt).Values[0])
			}
			v := vec.(*types.NullableInt)
			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				output.Values[i] = true
				for j := 0; j < len(inArr); j++ {
					if v.Values[i] == inArr[j] {
						output.Values[i] = false
						break
					}
				}
			}
		case types.Float:
			var inArr = make([]float64, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableFloat).Values[0])
			}
			v := vec.(*types.NullableFloat)
			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				output.Values[i] = true
				for j := 0; j < len(inArr); j++ {
					if math.Abs(v.Values[i]-inArr[j]) <= math.SmallestNonzeroFloat64 {
						output.Values[i] = false
						break
					}
				}
			}
		case types.Bool:
			var inArr = make([]bool, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableBool).Values[0])
			}
			v := vec.(*types.NullableBool)

			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				output.Values[i] = true
				for j := 0; j < len(inArr); j++ {
					if v.Values[i] == inArr[j] {
						output.Values[i] = false
						break
					}
				}
			}
		case types.Text:
			var inArr = make([]string, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableText).Values[0])
			}
			v := vec.(*types.NullableText)
			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				output.Values[i] = true
				for j := 0; j < len(inArr); j++ {
					if v.Values[i] == inArr[j] {
						output.Values[i] = false
						break
					}
				}
			}
		case types.Numeric:
			var inArr = make([]int64, 0, len(vectors)-1)
			for i := 1; i < len(vectors); i++ {
				inArr = append(inArr, vectors[i].(*types.NullableNumeric).Values[0])
			}
			v := vec.(*types.NullableNumeric)
			for i := 0; i < len(v.Values); i++ {
				if output.IsNullArr[i] {
					continue
				}
				output.Values[i] = true
				for j := 0; j < len(inArr); j++ {
					if v.Values[i] == inArr[j] {
						output.Values[i] = false
						break
					}
				}
			}
		default:
			panic("should not happend")
		}
		(&output).SetFilterArr(vectors[0].GetFilterArr())
		return &output, nil
	})
}
