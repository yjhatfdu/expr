package functions

import (
	"errors"
	"github.com/yjhatfdu/expr/types"
)

func init() {
	coalesce, _ := NewFunction("coalesce")
	coalesce.Generic(func(inputTypes []types.BaseType) (baseType types.BaseType, e error) {
		if len(inputTypes) == 0 {
			return 0, errors.New("require at least 1 argument")
		}
		t := inputTypes[0]
		if t > types.ScalaTypes {
			t = t - types.ScalaOffset
		}
		for _, ti := range inputTypes {
			if ti != t && ti != t+types.ScalaOffset {
				return 0, errors.New("require argument with same type")
			}
		}
		return t, nil
	}, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		return BroadCastMultiGeneric(vectors, vectors[0].Type(), func(values []interface{}) (vector interface{}, e error) {
			for _, v := range values {
				if v != nil {
					return v, nil
				}
			}
			return nil, nil
		})
	})
}
