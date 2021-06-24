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
		var t = types.Any
		for _, ti := range inputTypes {
			if t == types.Any && ti != types.Any {
				t = ti
				if t > types.ScalaTypes {
					t = t - types.ScalaOffset
				}
				continue
			}
			if ti != t && ti != t+types.ScalaOffset && ti != types.Any {
				return 0, errors.New("require argument with same type")
			}
		}
		return t, nil
	}, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		filterMasks := make([][]bool, len(vectors))
		var filterMaskTemp []bool
		for i := range vectors {
			m := vectors[i].GetFilterArr()
			filterMasks[i] = m
			if m != nil {
				filterMaskTemp = m
			}
		}
		var outFilterMask []bool
		if filterMaskTemp != nil {
			outFilterMask = make([]bool, len(filterMaskTemp), cap(filterMaskTemp))
		}
		var t = types.Any
		for _, ti := range vectors {
			if t == types.Any && ti.Type() != types.Any {
				t = ti.Type()
				if t > types.ScalaTypes {
					t = t - types.ScalaOffset
				}
				continue
			}
		}
		out, err := BroadCastMultiGeneric(vectors, t, func(values []interface{}, index int) (vector interface{}, e error) {
			if outFilterMask != nil {
				isFiltered := true
				for j, v := range values {
					var f bool
					if filterMasks[j] == nil {
						f = false
					} else {
						f = filterMasks[j][index]
					}
					isFiltered = isFiltered && f
					if v != nil && !f {
						outFilterMask[index] = false
						return v, nil
					}
				}
				outFilterMask[index] = isFiltered
				return nil, nil
			} else {
				for _, v := range values {
					if v != nil {
						return v, nil
					}
				}
				return nil, nil
			}
		})
		if err != nil {
			return nil, err
		}
		out.SetFilterArr(outFilterMask)
		return out, nil
	})
}
