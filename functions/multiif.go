package functions

import (
	"fmt"
	"github.com/yjhatfdu/expr/types"
)

func init() {
	// multiIf(if,then,else)
	// multiIf(case1,then1,case2,then2....)
	// multiIf(case1,then1,case2,then2....,default)
	multiIf, _ := NewFunction("multiIf")
	multiIf.Generic(func(inputTypes []types.BaseType) (baseType types.BaseType, e error) {
		if len(inputTypes) < 2 {
			return 0, fmt.Errorf("multiIf should have at least 2 arguments got (%d) argument", len(inputTypes))
		}
		outType := inputTypes[1]

		for i := 1; i < len(inputTypes); i++ {
			if inputTypes[i] != types.Any {
				outType = inputTypes[i]
				break
			}
		}

		for i := 0; i < len(inputTypes)/2; i++ {
			t := inputTypes[2*i+1]
			if t != outType && t+types.ScalaOffset != outType && t-types.ScalaOffset != outType && t != types.Any {
				return 0, fmt.Errorf("argument #%d should be %s, got %s", 2*i+1, types.GetTypeName(outType), types.GetTypeName(t))
			}
		}

		//todo 使用函数进行类型比较
		lastType := inputTypes[len(inputTypes)-1]
		if lastType != outType && lastType+types.ScalaOffset != outType && lastType-types.ScalaOffset != outType && lastType != types.Any {
			return 0, fmt.Errorf("argument #%d should be %s, got %s", len(inputTypes)-1, types.GetTypeName(outType), types.GetTypeName(lastType))
		}
		return outType, nil
	}, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		var hasDefault bool
		if len(vectors)%2 == 1 {
			hasDefault = true
		}
		l := len(vectors)
		filterMasks := make([][]bool, l)
		for i := range vectors {
			filterMasks[i] = vectors[i].GetFilterArr()
		}

		var output types.INullableVector
		var outType = types.Any

		for i := 1; i < len(vectors); i++ {
			if vectors[i].Type() != types.Any {
				outType = vectors[i].Type()
				break
			}
		}
		t := outType
		switch t {
		case types.Int:
			output = &types.NullableInt{}
		case types.Float:
			output = &types.NullableFloat{}
		case types.Bool:
			output = &types.NullableBool{}
		case types.Text:
			output = &types.NullableText{}
		case types.Timestamp, types.Time, types.Date:
			output = &types.NullableTimestamp{
				TsType: t,
			}
		case types.Numeric:
			output = &types.NullableNumeric{Scale: vectors[1].(*types.NullableNumeric).Scale}
		case types.TextA:
			output=&types.NullableTextArray{}
		default:
			panic("should not happend")
		}

		out, err := BroadCastMultiGeneric(vectors, output, func(values []interface{}, index int) (i interface{}, e error) {
			for i := 0; i < l/2; i++ {
				v := values[2*i]
				var truthy bool
				if v == nil {
					continue
				}
				switch vi := v.(type) {
				case int64:
					truthy = vi != 0
				case float64:
					truthy = vi != 0
				case string:
					truthy = vi != ""
				case bool:
					truthy = vi
				}
				if truthy {
					return values[2*i+1], nil
				}
			}
			if hasDefault {
				return values[len(values)-1], nil
			} else {
				return nil, nil
			}
		})
		if err != nil {
			return nil, err
		}
		out.SetFilterArr(CalFilterMask(filterMasks))
		return out, nil
	})
}
