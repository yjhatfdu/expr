package functions

import (
	"bytes"
	"errors"
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
		return vectors[0], nil
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
			output.Set(i, input.Values[i].StringScale(input.Scale), false)
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
		var utf8Flag bool
		var gbkFlag bool
		sample := input.Index(0).([]byte)
		if utf8.Valid(sample) {
			utf8Flag = true
		} else if isGBK(sample) {
			gbkFlag = true
		}
		return BroadCast1(vectors[0], output, func(i int) error {
			if utf8Flag {
				output.Seti(i, string(input.Index(i).([]byte)))
			} else if gbkFlag {
				output.Set(i, ConvertToNewString(string(input.Index(i).([]byte)), "gbk", "utf8"), false)
			} else {
				return errors.New("未知的原始编码，请确定编码后使用双参数toText函数尝试转换")
			}
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Blob, types.TextS}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableBlob)
		oldEncoder := strings.ToLower(vectors[1].(*types.NullableText).Index(0).(string))
		return BroadCast1(vectors[0], output, func(i int) error {
			//if oldEncoder == "gbk" {
			//	bs, err := GbkToUtf8(input.Index(i).([]byte))
			//	if err != nil {
			//		return err
			//	}
			//	output.Seti(i, string(bs))
			//} else if oldEncoder == "gb18030" {
			//	bs, err := Gb18030ToUtf8(input.Index(i).([]byte))
			//	if err != nil {
			//		return err
			//	}
			//	output.Seti(i, string(bs))
			//} else {
			output.Seti(i, ConvertToNewString(string(input.Index(i).([]byte)), oldEncoder, "utf8"))
			//}
			return nil
		})
	})
}

func isGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0x7f {
			//编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		} else {
			//大于127的使用双字节编码，落在gbk编码范围内的字符
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

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Gb18030ToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGb18030(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func HZGB2312ToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToHZGB2312(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func ConvertToNewString(src string, oldDecoder string, newEncoder string) string {
	srcDecoder := mahonia.NewDecoder(oldDecoder)
	desDecoder := mahonia.NewDecoder(newEncoder)
	resStr := srcDecoder.ConvertString(src)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}
