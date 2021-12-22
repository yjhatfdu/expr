package functions

import (
	"github.com/yjhatfdu/expr/types"
	"time"
)

func init() {
	now, _ := NewFunction("now")
	now.Overload([]types.BaseType{}, types.TimestampS, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		out := &types.NullableTimestamp{
			TsType: types.Timestamp,
		}
		out.Init(1)
		out.SetScala(true)
		out.Set(0, time.Now().UnixNano(), false)
		return out, nil
	})
	getYear, _ := NewFunction("getYear")
	getYear.Overload([]types.BaseType{types.Timestamp}, types.Int, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{}
		return BroadCast1(input, output, func(i int) error {
			t := time.Unix(0, input.Index(i).(int64)).In(time.Local)
			output.Set(i, int64(t.Year()), false)
			return nil
		})
	})
	getMonth, _ := NewFunction("getMonth")
	getMonth.Overload([]types.BaseType{types.Timestamp}, types.Int, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{}
		return BroadCast1(input, output, func(i int) error {
			t := time.Unix(0, input.Index(i).(int64)).In(time.Local)
			output.Set(i, int64(t.Month()), false)
			return nil
		})
	})
	getDay, _ := NewFunction("getDay")
	getDay.Overload([]types.BaseType{types.Timestamp}, types.Int, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{}
		return BroadCast1(input, output, func(i int) error {
			t := time.Unix(0, input.Index(i).(int64)).In(time.Local)
			output.Set(i, int64(t.Day()), false)
			return nil
		})
	})
	getWeekDay, _ := NewFunction("getWeekDay")
	getWeekDay.Overload([]types.BaseType{types.Timestamp}, types.Int, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{}
		return BroadCast1(input, output, func(i int) error {
			t := time.Unix(0, input.Index(i).(int64)).In(time.Local)
			output.Set(i, int64(t.Weekday()), false)
			return nil
		})
	})
}
