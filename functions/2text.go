package functions

import (
	"github.com/axgle/mahonia"
	"github.com/yjhatfdu/expr/types"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
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
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableText)
		return BroadCast1(vectors[0], output, func(i int) error {
			s := input.Index(i).(string)
			if utf8.ValidString(s) {
				output.Seti(i, s)
				return nil
			} else if ValidGBKString(s) {
				reader := transform.NewReader(strings.NewReader(s), simplifiedchinese.GB18030.NewDecoder())
				d, err := ioutil.ReadAll(reader)
				if err != nil {
					return err
				}
				output.Seti(i, string(d))
			} else {
				output.Seti(i, s)
			}

			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Bool}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableBool)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, strconv.FormatBool(input.Index(i).(bool)), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Int}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableInt)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, strconv.FormatInt(input.Index(i).(int64), 10), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Float}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableFloat)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, strconv.FormatFloat(input.Index(i).(float64), 'f', -1, 64), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Numeric}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableNumeric)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Numeric2Text(input.Index(i).(int64), input.Scale), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Timestamp}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, time.Unix(0, input.Index(i).(int64)).Format(time.RFC3339), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Date}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, time.Unix(0, input.Index(i).(int64)).Format("2006-01-02"), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Time}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, time.Unix(0, input.Index(i).(int64)).Format("15:04:05"), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Timestamp, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		standard := vectors[1].(*types.NullableText).Index(0).(string)
		gostyle := convert2GoTimeFormatStyle(standard)
		return BroadCast1(vectors[0], output, func(i int) error {
			t := time.Unix(0, input.Index(i).(int64))
			output.Set(i, t.Format(gostyle), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Date, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		standard := vectors[1].(*types.NullableText).Index(0).(string)
		gostyle := convert2GoTimeFormatStyle(standard)
		return BroadCast1(vectors[0], output, func(i int) error {
			t := time.Unix(0, input.Index(i).(int64))
			output.Set(i, t.Format(gostyle), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Time, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		standard := vectors[1].(*types.NullableText).Index(0).(string)
		gostyle := convert2GoTimeFormatStyle(standard)
		return BroadCast1(vectors[0], output, func(i int) error {
			t := time.Unix(0, input.Index(i).(int64))
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
		oldEncoder := vectors[1].(*types.NullableText).Index(0).(string)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, ConvertToNewString(string(input.Index(i).([]byte)), oldEncoder, "utf8"), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Blob, types.TextS, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableBlob)
		oldEncoder := vectors[1].(*types.NullableText).Index(0).(string)
		newEncoder := vectors[2].(*types.NullableText).Index(0).(string)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, ConvertToNewString(string(input.Index(i).([]byte)), oldEncoder, newEncoder), false)
			return nil
		})
	})
}

func ValidGBKString(s string) bool {
	data := []byte(s)
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0xff {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func ConvertToNewString(src string, oldEncoder string, newEncoder string) string {
	srcDecoder := mahonia.NewDecoder(oldEncoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}
