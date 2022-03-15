package functions

import (
	"github.com/yjhatfdu/expr/types"
	"strconv"
)

func cvtFloat2Int(in []float64, out []int64) {
	for i := range in {
		out[i] = int64(in[i])
	}
}

func init() {
	toInt, _ := NewFunction("toInt")
	toInt.Overload([]types.BaseType{types.Int}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toInt.Overload([]types.BaseType{types.Float}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableInt{}
		input := vectors[0].(*types.NullableFloat)
		output.Init(input.Length())
		output.IsScalaV = input.IsScalaV
		cvtFloat2Int(input.Values, output.Values)
		output.IsNullArr = input.IsNullArr
		output.FilterArr = input.FilterArr
		return output, nil
	})
	toInt.Overload([]types.BaseType{types.Text}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableInt{}
		input := vectors[0].(*types.NullableText)
		//output.FilterArr = input.FilterArr
		return BroadCast1(vectors[0], output, func(i int) error {
			ret, err := strconv.ParseInt(input.Index(i).(string), 10, 64)
			if err != nil {
				return err
			}
			output.Set(i, ret, false)
			return nil
		})
	})
	toInt.Overload([]types.BaseType{types.Timestamp}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
				FilterArr: input.FilterArr,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toInt.Overload([]types.BaseType{types.Time}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
				FilterArr: input.FilterArr,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toInt.Overload([]types.BaseType{types.Date}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
				FilterArr: input.FilterArr,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toInt.Overload([]types.BaseType{types.Numeric}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableInt{}
		input := vectors[0].(*types.NullableNumeric)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, input.Index(i).(types.Decimal).ToInt(), false)
			return nil
		})
	})
	toInt.Overload([]types.BaseType{types.Bool}, types.Int, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableInt{}
		input := vectors[0].(*types.NullableBool)
		return BroadCast1(vectors[0], output, func(i int) error {
			if input.Index(i).(bool) {
				output.Set(i, 1, false)
			} else {
				output.Set(i, 0, false)
			}
			return nil
		})
	})
}
