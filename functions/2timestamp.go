package functions

import (
	"github.com/yjhatfdu/expr/types"
	"time"
)

func init() {
	toTimestamp, _ := NewFunction("toTimestamp")
	toTimestamp.Overload([]types.BaseType{types.Int}, types.Timestamp, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableInt)
		output := &types.NullableTimestamp{
			TsType: types.Timestamp,
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
				FilterArr: input.FilterArr,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toTimestamp.Overload([]types.BaseType{types.Date}, types.Timestamp, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableTimestamp{
			TsType: types.Timestamp,
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
				FilterArr: input.FilterArr,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toTimestamp.Overload([]types.BaseType{types.Timestamp}, types.Timestamp, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toTimestamp.Overload([]types.BaseType{types.Text}, types.Timestamp, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Timestamp}
		input := vectors[0].(*types.NullableText)
		return BroadCast1(vectors[0], output, func(i int) error {
			s := input.Index(i).(string)
			ts, err := time.ParseInLocation(time.RFC3339, s, time.Local)
			if err != nil {
				return err
			}
			t := ts.UnixNano()
			output.Set(i, t, false)
			return nil
		})
	})
	toTimestamp.Overload([]types.BaseType{types.Text, types.TextS}, types.Timestamp, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Timestamp}
		input := vectors[0].(*types.NullableText)
		standard := vectors[1].(*types.NullableText).Index(0).(string)
		gostyle := convert2GoTimeFormatStyle(standard)
		return BroadCast1(vectors[0], output, func(i int) error {
			s := input.Index(i).(string)
			ts, err := time.ParseInLocation(gostyle, s, time.Local)
			if err != nil {
				return err
			}
			t := ts.UnixNano()
			output.Set(i, t, false)
			return nil
		})
	})
}
