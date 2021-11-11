package functions

import (
	"github.com/axgle/mahonia"
	"github.com/yjhatfdu/expr/types"
	"strconv"
	"strings"
	"time"
)

func convert2GoTimeFormatStyle(standard string) (gostyle string) {
	standard = strings.Replace(standard, "yyyy", "2006", 1)
	standard = strings.Replace(standard, "yy", "06", 1)
	standard = strings.Replace(standard, "MM", "01", 1)
	standard = strings.Replace(standard, "dd", "02", 1)
	standard = strings.Replace(standard, "HH", "15", 1)
	standard = strings.Replace(standard, "mm", "04", 1)
	standard = strings.Replace(standard, "ss", "05", 1)
	standard = strings.Replace(standard, "SSS", "000", 1)

	// ±07:00 ±hh:mm
	standard = strings.Replace(standard, "hh", "08", 1)
	standard = strings.Replace(standard, "mm", "00", 1)

	return standard
}

func init() {
	toText, _ := NewFunction("toText")
	toText.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toText.Overload([]types.BaseType{types.Bool}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableBool)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, strconv.FormatBool(input.Values[i]), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Int}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableInt)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, strconv.FormatInt(input.Values[i], 10), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Float}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableFloat)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, strconv.FormatFloat(input.Values[i], 'f', -1, 64), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Numeric}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableNumeric)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Numeric2Text(input.Values[i], input.Scale), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Timestamp}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, time.Unix(0, input.Values[i]).Format(time.RFC3339), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Date}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, time.Unix(0, input.Values[i]).Format("2006-01-02"), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Time}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, time.Unix(0, input.Values[i]).Format("15:04:05"), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Timestamp, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		standard := vectors[1].(*types.NullableText).Values[0]
		gostyle := convert2GoTimeFormatStyle(standard)
		return BroadCast1(vectors[0], output, func(i int) error {
			t := time.Unix(0, input.Values[i])
			output.Set(i, t.Format(gostyle), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Date, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		standard := vectors[1].(*types.NullableText).Values[0]
		gostyle := convert2GoTimeFormatStyle(standard)
		return BroadCast1(vectors[0], output, func(i int) error {
			t := time.Unix(0, input.Values[i])
			output.Set(i, t.Format(gostyle), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Time, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		standard := vectors[1].(*types.NullableText).Values[0]
		gostyle := convert2GoTimeFormatStyle(standard)
		return BroadCast1(vectors[0], output, func(i int) error {
			t := time.Unix(0, input.Values[i])
			output.Set(i, t.Format(gostyle), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Blob}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableBlob)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, ConvertToNewString(string(input.Index(i).([]byte)), "gbk", "utf8"), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Blob, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableBlob)
		oldEncoder := vectors[1].(*types.NullableText).Values[0]
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, ConvertToNewString(string(input.Index(i).([]byte)), oldEncoder, "utf8"), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Blob, types.TextS, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableBlob)
		oldEncoder := vectors[1].(*types.NullableText).Values[0]
		newEncoder := vectors[2].(*types.NullableText).Values[0]
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, ConvertToNewString(string(input.Index(i).([]byte)), oldEncoder, newEncoder), false)
			return nil
		})
	})
	//toText.Overload([]types.BaseType{types.Blob, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
	//	output := &types.NullableText{}
	//	input := vectors[0].(*)
	//	return BroadCast1(vectors[0], output, func(i int) error {
	//		output.Set(i, strconv.FormatBool(input.Values[i]), false)
	//		return nil
	//	})
	//})
}

func ConvertToNewString(src string, oldEncoder string, newEncoder string) string {
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}
