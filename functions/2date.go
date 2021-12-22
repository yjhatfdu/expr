package functions

import (
	"github.com/yjhatfdu/expr/types"
	"time"
)

func init() {
	toDate, _ := NewFunction("toDate")
	toDate.Overload([]types.BaseType{types.Int}, types.Date, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableInt)
		output := &types.NullableTimestamp{
			TsType: types.Date,
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
				FilterArr: input.FilterArr,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toDate.Overload([]types.BaseType{types.Date}, types.Date, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toDate.Overload([]types.BaseType{types.Timestamp}, types.Date, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Date}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			//t := input.Values[i] - types.LocalOffsetNano
			_time := input.Index(i).(int64) % 24 * 3600 * 1e9
			output.Set(i, input.Index(i).(int64)-_time, false)
			return nil
		})
	})
	toDate.Overload([]types.BaseType{types.Text}, types.Date, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Date}
		input := vectors[0].(*types.NullableText)
		return BroadCast1(vectors[0], output, func(i int) error {
			s := input.Index(i).(string)
			ts, err := time.Parse("2006-01-02", s)
			if err != nil {
				return err
			}
			t := ts.UnixNano()
			output.Set(i, t, false)
			return nil
		})
	})
	toDate.Overload([]types.BaseType{types.Text, types.TextS}, types.Date, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Date}
		input := vectors[0].(*types.NullableText)
		standard := vectors[1].(*types.NullableText).Index(0).(string)
		gostyle := convert2GoTimeFormatStyle(standard)
		return BroadCast1(vectors[0], output, func(i int) error {
			s := input.Index(i).(string)
			ts, err := time.Parse(gostyle, s)
			if err != nil {
				return err
			}
			t := ts.UnixNano()
			output.Set(i, t, false)
			return nil
		})
	})
}
