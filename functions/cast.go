package functions

import (
	"github.com/yjhatfdu/expr/types"
	"strconv"
	"time"
)

func cvtInt2Float([]int64, []float64)
func cvtFloat2Int([]float64, []int64)
func cvtInt2Numeric([]int64, []int64, int64)

func init() {
	toInt, _ := NewFunction("toInt")
	toInt.Overload([]types.BaseType{types.Int}, types.Int, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toInt.Overload([]types.BaseType{types.Float}, types.Int, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableInt{}
		input := vectors[0].(*types.NullableFloat)
		output.Init(input.Length())
		output.IsScalaV = input.IsScalaV
		cvtFloat2Int(input.Values, output.Values)
		output.IsNullArr = input.IsNullArr
		return output, nil
	})
	toInt.Overload([]types.BaseType{types.Text}, types.Int, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableInt{}
		input := vectors[0].(*types.NullableText)
		return BroadCast1(vectors[0], output, func(i int) error {
			ret, err := strconv.ParseInt(input.Values[i], 10, 64)
			if err != nil {
				return err
			}
			output.Set(i, ret, false)
			return nil
		})
	})
	toInt.Overload([]types.BaseType{types.Timestamp}, types.Int, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {

		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toInt.Overload([]types.BaseType{types.Time}, types.Int, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toInt.Overload([]types.BaseType{types.Date}, types.Int, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		input := vectors[0].(*types.NullableTimestamp)
		output := &types.NullableInt{
			NullableVector: types.NullableVector{
				IsNullArr: input.IsNullArr,
				IsScalaV:  input.IsScalaV,
			},
			Values: input.Values,
		}
		return output, nil
	})
	toInt.Overload([]types.BaseType{types.Numeric}, types.Int, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableInt{}
		input := vectors[0].(*types.NullableNumeric)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Numeric2Int(input.Values[i], input.Scale), false)
			return nil
		})
	})
}

func init() {
	toFloat, _ := NewFunction("toFloat")
	toFloat.Overload([]types.BaseType{types.Float}, types.Float, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toFloat.Overload([]types.BaseType{types.Int}, types.Float, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableFloat{}
		input := vectors[0].(*types.NullableInt)
		output.Init(input.Length())
		output.IsScalaV = input.IsScalaV
		cvtInt2Float(input.Values, output.Values)
		output.IsNullArr = input.IsNullArr
		return output, nil
	})
	toFloat.Overload([]types.BaseType{types.Numeric}, types.Float, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableFloat{}
		input := vectors[0].(*types.NullableNumeric)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Numeric2Float(input.Values[i], input.Scale), false)
			return nil
		})
	})
	toFloat.Overload([]types.BaseType{types.Text}, types.Float, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableFloat{}
		input := vectors[0].(*types.NullableText)
		return BroadCast1(vectors[0], output, func(i int) error {
			f, err := strconv.ParseFloat(input.Values[i], 64)
			if err != nil {
				return err
			}
			output.Set(i, f, false)
			return nil
		})
	})
}

func init() {
	toText, _ := NewFunction("toText")
	toText.Overload([]types.BaseType{types.Text}, types.Text, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toText.Overload([]types.BaseType{types.Int}, types.Text, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableInt)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, strconv.FormatInt(input.Values[i], 10), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Float}, types.Text, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableFloat)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, strconv.FormatFloat(input.Values[i], 'f', -1, 64), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Numeric}, types.Text, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableNumeric)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.Numeric2Text(input.Values[i], input.Scale), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Timestamp}, types.Text, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, time.Unix(0, input.Values[i]).In(time.Local).Format(time.RFC3339), false)
			return nil
		})
	})
	toText.Overload([]types.BaseType{types.Date}, types.Text, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableText{}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, time.Unix(0, input.Values[i]).In(time.Local).Format("2006-01-02"), false)
			return nil
		})
	})
}

func init() {
	toDate, _ := NewFunction("toDate")
	toDate.Overload([]types.BaseType{types.Date}, types.Date, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toDate.Overload([]types.BaseType{types.Timestamp}, types.Date, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Date}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			t := input.Values[i] + types.LocalOffsetNano
			dt := t - t%(24*3600*1e9)
			output.Set(i, dt, false)
			return nil
		})
	})
}

func init() {
	toTime, _ := NewFunction("toTime")
	toTime.Overload([]types.BaseType{types.Time}, types.Time, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toTime.Overload([]types.BaseType{types.Timestamp}, types.Time, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableTimestamp{TsType: types.Time}
		input := vectors[0].(*types.NullableTimestamp)
		return BroadCast1(vectors[0], output, func(i int) error {
			t := input.Values[i] + types.LocalOffsetNano
			dt := t % (24 * 3600 * 1e9)
			output.Set(i, dt, false)
			return nil
		})
	})
}

func init() {
	// default format is Numeric(12,4), use toNumeric(_,int) to specify scale
	toNumeric, _ := NewFunction("toNumeric")
	toNumeric.Overload([]types.BaseType{types.Numeric}, types.Numeric, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		return vectors[0], nil
	})
	toNumeric.Overload([]types.BaseType{types.Numeric, types.IntS}, types.Numeric, func(vectors []types.INullableVector) (vector types.INullableVector, e error) {
		output := &types.NullableNumeric{}
		input := vectors[0].(*types.NullableNumeric)
		scale := int(vectors[1].(*types.NullableInt).Values[0])
		output.Scale = scale
		return BroadCast1(vectors[0], output, func(i int) error {
			output.Set(i, types.NormalizeNumeric(input.Values[i], input.Scale, scale), false)
			return nil
		})
	})
}
