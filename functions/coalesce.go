package functions

import (
	"errors"
	"github.com/yjhatfdu/expr/types"
)

func init() {
	coalesce, _ := NewFunction("coalesce")
	coalesce.Generic(func(types []types.BaseType) (baseType types.BaseType, e error) {
		if len(types) == 0 {
			return 0, errors.New("require at least 1 argument")
		}
		t := types[0]
		for _, ti := range types {
			if ti != t {
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
