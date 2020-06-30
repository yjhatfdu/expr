package tests

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/yjhatfdu/expr"
	"github.com/yjhatfdu/expr/types"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func loader(file string) {
	log.Printf("testing %s", file)
	fp, err := os.Open(file)
	defer fp.Close()
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fp)
	typesStr, err := r.ReadString('\n')
	typesStr = typesStr[0 : len(typesStr)-1]
	typesSegs := strings.Split(typesStr, ",")
	allTypes := make([]types.BaseType, len(typesSegs))
	for i := range typesSegs {
		t, ok := types.GetTypeByName(typesSegs[i])
		if !ok {
			panic(fmt.Errorf("invalid type %s", typesSegs[i]))
		}
		allTypes[i] = t
	}
	inputTypes := allTypes[0 : len(allTypes)-1]
	outputType := allTypes[len(allTypes)-1]
	code, err := r.ReadString('\n')
	program, err := expr.Compile(code, inputTypes)
	if err != nil {
		if strings.HasSuffix(file, ".cerr") {
			return
		} else {
			panic(err)
		}
	}
	if program.OutputType != outputType && program.OutputType-types.ScalaOffset != outputType {
		panic(fmt.Errorf("output type should be %s, was %s\n", types.GetTypeName(outputType), types.GetTypeName(program.OutputType)))
	}
	data, _ := ioutil.ReadAll(r)
	cr := csv.NewReader(bytes.NewReader(data))
	rows, err := cr.ReadAll()
	if err != nil {
		panic(err)
	}
	arrCount := len(allTypes)
	dataArr := make([]types.INullableVector, arrCount)
	for i := range dataArr {
		arr := make([]string, len(rows))
		for j := range rows {
			arr[j] = rows[j][i]
		}
		var vec types.INullableVector
		switch allTypes[i] {
		case types.Int, types.IntS:
			vec = buildIntVec(arr, allTypes[i] > types.ScalaTypes)
		case types.Float, types.FloatS:
			vec = buildFloatVec(arr, allTypes[i] > types.ScalaTypes)
		case types.Text, types.TextS:
			vec = buildTextVec(arr, allTypes[i] > types.ScalaTypes)
		case types.Bool, types.BoolS:
			vec = buildBoolVec(arr, allTypes[i] > types.ScalaTypes)
		case types.Timestamp, types.TimestampS:
			vec = buildTimestampVec(arr, allTypes[i] > types.ScalaTypes)
		case types.Numeric, types.NumericS:
			vec = buildNumericVec(arr, allTypes[i] > types.ScalaTypes)
		case types.Interval, types.IntervalS:
			vec = buildIntervalVec(arr, allTypes[i] > types.ScalaTypes)
		case types.Time, types.TimeS:
			vec = buildTimeVec(arr, allTypes[i] > types.ScalaTypes)
		case types.Date, types.DateS:
			vec = buildDateVec(arr, allTypes[i] > types.ScalaTypes)
		}
		dataArr[i] = vec
	}
	result, err := program.Run(dataArr[0 : len(dataArr)-1])
	if err != nil {
		if strings.HasSuffix(file, ".rerr") {
			return
		} else {
			panic(err)
		}
	}
	if len(result.GetErrors()) > 0 {
		if strings.HasSuffix(file, ".lerr") {
			return
		} else {
			for _, e := range result.GetErrors() {
				fmt.Println(e)
			}
			panic("line error")
		}
	}
	target := dataArr[len(dataArr)-1]
	if !compare(target, result) {
		panic(fmt.Errorf("result not match, \ntarget: %s \nresult: %s\n", types.ToString(target), types.ToString(result)))
	}
}

func buildIntVec(d []string, isScalar bool) types.INullableVector {
	val := types.NullableInt{}
	if isScalar {
		val.Init(1)
	} else {
		val.Init(len(d))
	}
	for i, s := range d {
		if s == `\N` || s == "" {
			val.Set(i, 0, true)
		} else {
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				panic(err)
			}
			val.Set(i, v, false)
		}
	}
	val.IsScalaV = isScalar
	return &val
}

func buildFloatVec(d []string, isScalar bool) types.INullableVector {
	val := types.NullableFloat{}
	if isScalar {
		val.Init(1)
	} else {
		val.Init(len(d))
	}
	for i, s := range d {
		if s == `\N` || s == "" {
			val.Set(i, 0, true)
		} else {
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				panic(err)
			}
			val.Set(i, v, false)
		}
	}
	val.IsScalaV = isScalar
	return &val
}

func buildTextVec(d []string, isScalar bool) types.INullableVector {
	val := types.NullableText{}
	if isScalar {
		val.Init(1)
	} else {
		val.Init(len(d))
	}
	for i, s := range d {
		if s == `\N` || s == "" {
			val.Set(i, "", true)
		} else {
			val.Set(i, s, false)
		}
	}
	val.IsScalaV = isScalar
	return &val
}

func buildNumericVec(d []string, isScalar bool) types.INullableVector {
	val := types.NullableNumeric{}
	val.Scale = 4
	if isScalar {
		val.Init(1)
	} else {
		val.Init(len(d))
	}
	for i, s := range d {
		if s == `\N` || s == "" {
			val.Set(i, 0, true)
		} else {
			v, err := types.Text2Numeric(s, 4)
			if err != nil {
				panic(err)
			}
			val.Set(i, v, false)
		}
	}
	val.IsScalaV = isScalar
	return &val
}

func buildBoolVec(d []string, isScalar bool) types.INullableVector {
	val := types.NullableBool{}
	if isScalar {
		val.Init(1)
	} else {
		val.Init(len(d))
	}
	for i, s := range d {
		if s == `\N` || s == "" {
			val.Set(i, false, true)
		} else {
			v, err := strconv.ParseBool(s)
			if err != nil {
				panic(err)
			}
			val.Set(i, v, false)
		}
	}
	val.IsScalaV = isScalar
	return &val
}

func buildTimestampVec(d []string, isScalar bool) types.INullableVector {
	val := types.NullableTimestamp{}
	if isScalar {
		val.Init(1)
	} else {
		val.Init(len(d))
	}
	for i, s := range d {
		if s == `\N` || s == "" {
			val.Set(i, 0, true)
		} else {
			v, err := time.Parse(time.RFC3339, s)
			if err != nil {
				panic(err)
			}
			val.Set(i, v.UnixNano(), false)
		}
	}
	val.IsScalaV = isScalar
	val.TsType = types.Timestamp
	return &val
}

func buildIntervalVec(d []string, isScalar bool) types.INullableVector {
	val := types.NullableTimestamp{}
	if isScalar {
		val.Init(1)
	} else {
		val.Init(len(d))
	}
	for i, s := range d {
		if s == `\N` || s == "" {
			val.Set(i, 0, true)
		} else {
			v, err := time.ParseDuration(s)
			if err != nil {
				panic(err)
			}
			val.Set(i, int64(v), false)
		}
	}
	val.IsScalaV = isScalar
	val.TsType = types.Interval
	return &val
}
func buildDateVec(d []string, isScalar bool) types.INullableVector {
	val := types.NullableTimestamp{}
	if isScalar {
		val.Init(1)
	} else {
		val.Init(len(d))
	}
	for i, s := range d {
		if s == `\N` || s == "" {
			val.Set(i, 0, true)
		} else {
			v, err := time.Parse("2006-01-02", s)
			if err != nil {
				panic(err)
			}
			val.Set(i, v.UnixNano(), false)
		}
	}
	val.IsScalaV = isScalar
	val.TsType = types.Date
	return &val
}

func buildTimeVec(d []string, isScalar bool) types.INullableVector {
	val := types.NullableTimestamp{}
	if isScalar {
		val.Init(1)
	} else {
		val.Init(len(d))
	}
	for i, s := range d {
		if s == `\N` || s == "" {
			val.Set(i, 0, true)
		} else {
			v, err := time.Parse(time.RFC3339, "1970-01-01T"+s+"Z")
			if err != nil {
				panic(err)
			}
			val.Set(i, v.UnixNano(), false)
		}
	}
	val.IsScalaV = isScalar
	val.TsType = types.Time
	return &val
}

func compare(vec1, vec2 types.INullableVector) bool {
	return types.ToString(vec1) == types.ToString(vec2)
}
