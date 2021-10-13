package functions

import (
	"errors"
	"github.com/yjhatfdu/expr/types"
)

func init() {
	isNull, _ := NewFunction("isNull")
	isNull.Generic(func(inputTypes []types.BaseType) (baseType types.BaseType, e error) {
		if len(inputTypes) != 1 {
			return 0, errors.New("require 1 argument")
		}
		return types.Bool, nil
	}, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		var outFilterMask []bool
		outFilterMask = vectors[0].GetFilterArr()
		out := &types.NullableBool{}
		out.SetScala(vectors[0].IsScala())
		out.Values = vectors[0].GetIsNullArr()
		out.SetFilterArr(outFilterMask)
		return out, nil
	})
	isNotNull, _ := NewFunction("isNotNull")
	isNotNull.Generic(func(inputTypes []types.BaseType) (baseType types.BaseType, e error) {
		if len(inputTypes) != 1 {
			return 0, errors.New("require 1 argument")
		}
		return types.Bool, nil
	}, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		var outFilterMask []bool
		outFilterMask = vectors[0].GetFilterArr()
		out := &types.NullableBool{}
		out.SetScala(vectors[0].IsScala())
		out.Init(vectors[0].Length())
		copy(out.Values, vectors[0].GetIsNullArr())
		not(out.Values)
		out.SetFilterArr(outFilterMask)
		return out, nil
	})
}
