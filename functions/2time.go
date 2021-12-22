package functions

import (
	"github.com/yjhatfdu/expr/types"
	"time"
)

func init() {
	toTime, _ := NewFunction("toTime")
	toTime.Overload([]types.BaseType{types.Int}, types.Time, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableInt)
		output := &types.NullableTimestamp{
			TsType: types.Time,
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
				FilterArr: input.FilterArr,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toTime.Overload([]types.BaseType{types.Time}, types.Time, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toTime.Overload([]types.BaseType{types.Timestamp}, types.Time, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Time}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			//t := input.Values[i] - types.LocalOffsetNano
			dt := input.Index(i).(int64) % (24 * 3600 * 1e9)
			output.Set(i, dt, false)
			return nil
		})
	})
	toTime.Overload([]types.BaseType{types.Text}, types.Time, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Time}
		input := vectors[0].(*types.NullableText)
		return BroadCast1(vectors[0], output, func(i int) error {
			s := input.Index(i).(string)
			ts, err := time.Parse("15:04:05", s)
			if err != nil {
				return err
			}
			t := ts.UnixNano()
			dt := t % (24 * 3600 * 1e9)
			output.Set(i, dt, false)
			return nil
		})
	})
	toTime.Overload([]types.BaseType{types.Text, types.TextS}, types.Time, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Time}
		input := vectors[0].(*types.NullableText)
		standard := vectors[1].(*types.NullableText).Values[0]
		gostyle := convert2GoTimeFormatStyle(standard)
		return BroadCast1(vectors[0], output, func(i int) error {
			s := input.Index(i).(string)
			ts, err := time.Parse(gostyle, s)
			if err != nil {
				return err
			}
			t := ts.UnixNano()
			dt := t % (24 * 3600 * 1e9)
			output.Set(i, dt, false)
			return nil
		})
	})
}
