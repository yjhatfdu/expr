package functions

import (
	"errors"
	"fmt"
	"github.com/yjhatfdu/expr/types"
	"regexp"
	"time"
)

var tsfmt0 = regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{6}`)
var tsfmt1 = regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)
var tsfmt2 = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`)
var tsfmt3 = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

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

			var format string
			var ts time.Time
			if tsfmt0.MatchString(s) {
				format = "2006-01-02 15:04:05.000000"
			} else if tsfmt1.MatchString(s) {
				format = "2006-01-02 15:04:05"
			} else if tsfmt2.MatchString(s) {
				format = time.RFC3339
			} else if tsfmt3.MatchString(s) {
				format = "2006-01-02"
			} else {
				return errors.New("未支持的时间字符串 " + s)
			}
			var err error
			ts, err = time.ParseInLocation(format, s, time.Local)
			if err != nil {
				return errors.New(fmt.Sprintf("未能成功根据指定时间字符串格式 %s 解析目标字符串 %s", format, s))
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
