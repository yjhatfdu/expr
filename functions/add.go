package functions

import "github.com/yjhatfdu/expr/types"

func addIntInt(data1, data2, out []int64) {
	for i := range out {
		out[i] = data1[i] + data2[i]
	}
}
func addIntS(data1, out []int64, intS int64) {
	for i := range data1 {
		out[i] = data1[i] + intS
	}
}
func addIntFloat(data1 []int64, data2 []float64, out []float64) {
	for i := range out {
		out[i] = float64(data1[i]) + data2[i]
	}
}
func addIntSFloat(data1 []float64, out []float64, iv int64) {
	for i := range out {
		out[i] = data1[i] + float64(iv)
	}
}
func addIntFloatS(data1 []int64, out []float64, fv float64) {
	for i := range out {
		out[i] = float64(data1[i]) + fv
	}
}

func init() {
	addFunc, _ := NewFunction("add")
	addFunc.Overload([]types.BaseType{types.Int, types.Int}, types.Int, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableInt{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableInt)
		if left.IsScalaV && right.IsScalaV {
			output.Init(1)
			output.SetScala(true)
			output.IsNullArr[0] = left.IsNullArr[0] || right.IsNullArr[0]
			output.Values[0] = left.Values[0] + right.Values[0]
			return &output, nil
		}
		if len(left.Values) > len(right.Values) {
			output.Init(len(left.Values))
		} else {
			output.Init(len(right.Values))
		}
		if left.IsScalaV {
			addIntS(right.Values, output.Values, left.Values[0])
			orBoolS(right.IsNullArr, output.IsNullArr, left.IsNullArr[0])
		} else if right.IsScalaV {
			addIntS(left.Values, output.Values, right.Values[0])
			orBoolS(left.IsNullArr, output.IsNullArr, right.IsNullArr[0])
		} else {
			addIntInt(left.Values, right.Values, output.Values)
			orBool(left.IsNullArr, right.IsNullArr, output.IsNullArr)
		}
		output.FilterArr = CalFilterMask([][]bool{left.FilterArr, right.FilterArr})
		return &output, nil
	})
	addFunc.Overload([]types.BaseType{types.Int, types.Float}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableFloat)
		if left.IsScalaV && right.IsScalaV {
			output.Init(1)
			output.SetScala(true)
			output.IsNullArr[0] = left.IsNullArr[0] || right.IsNullArr[0]
			output.Values[0] = float64(left.Values[0]) + right.Values[0]
			return &output, nil
		}
		if len(left.Values) > len(right.Values) {
			output.Init(len(left.Values))
		} else {
			output.Init(len(right.Values))
		}
		if left.IsScalaV {
			addIntSFloat(right.Values, output.Values, left.Values[0])
			orBoolS(right.IsNullArr, output.IsNullArr, left.IsNullArr[0])
		} else if right.IsScalaV {
			addIntFloatS(left.Values, output.Values, right.Values[0])
			orBoolS(left.IsNullArr, output.IsNullArr, right.IsNullArr[0])
		} else {
			addIntFloat(left.Values, right.Values, output.Values)
			orBool(left.IsNullArr, right.IsNullArr, output.IsNullArr)
		}
		output.FilterArr = CalFilterMask([][]bool{left.FilterArr, right.FilterArr})
		return &output, nil
	})
	addFunc.Overload([]types.BaseType{types.Float, types.Int}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableInt)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+float64(right.Values[j]), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Float, types.Float}, types.Float, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableFloat{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableFloat)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+right.Values[j], false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Text, types.Text}, types.Text, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableText{}
		left := vectors[0].(*types.NullableText)
		right := vectors[1].(*types.NullableText)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+right.Values[j], false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Timestamp, types.Interval}, types.Timestamp, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableTimestamp{TsType: types.Timestamp}
		output.TsType = types.Timestamp
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+right.Values[j], false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Time, types.Interval}, types.Time, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableTimestamp{TsType: types.Time}
		output.TsType = types.Time
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+right.Values[j], false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Date, types.Interval}, types.Date, func(vectors []types.INullableVector, env map[string]string) (vector types.INullableVector, e error) {
		output := types.NullableTimestamp{TsType: types.Date}
		output.TsType = types.Date
		left := vectors[0].(*types.NullableTimestamp)
		right := vectors[1].(*types.NullableTimestamp)
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+right.Values[j], false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Numeric, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableNumeric)
		s := types.NumericScale(left.Scale, right.Scale)
		output.Scale = s
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.NormalizeNumeric(left.Values[i], left.Scale, s)+types.NormalizeNumeric(right.Values[j], left.Scale, s), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Numeric, types.Int}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableInt)
		output.Scale = left.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+types.Int2numeric(right.Values[j], left.Scale), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Int, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableInt)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = right.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.Int2numeric(left.Values[i], right.Scale)+right.Values[j], false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Numeric, types.Float}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableNumeric)
		right := vectors[1].(*types.NullableFloat)
		output.Scale = left.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, left.Values[i]+types.Float2numeric(right.Values[j], left.Scale), false)
			return nil
		})
	})
	addFunc.Overload([]types.BaseType{types.Float, types.Numeric}, types.Numeric, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		output := types.NullableNumeric{}
		left := vectors[0].(*types.NullableFloat)
		right := vectors[1].(*types.NullableNumeric)
		output.Scale = right.Scale
		return BroadCast2(vectors[0], vectors[1], &output, func(index, i, j int) error {
			output.Set(index, types.Float2numeric(left.Values[i], right.Scale)+right.Values[j], false)
			return nil
		})
	})
}
