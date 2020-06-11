package types

import (
	"fmt"
)

type BaseType int

const (
	Any BaseType = iota + 1000
	Int
	Float
	Numeric //only support Numeric(12,4)
	Text
	Bool
	Timestamp
	Date
	Time
	Interval
	Blob
)

var typeNames = map[BaseType]string{
	Any:       "Any",
	Int:       "Int",
	Float:     "Float",
	Numeric:   "Numeric",
	Text:      "Text",
	Bool:      "Bool",
	Timestamp: "Timestamp",
	Date:      "Date",
	Time:      "Time",
	Interval:  "Interval",
	Blob:      "Blob",
}

func GetTypeName(t BaseType) string {
	if name, ok := typeNames[t]; ok {
		return name
	} else {
		return "undefined type"
	}
}

type INullableVector interface {
	IsNull(i int) bool
	GetIsNullArr() []bool
	IsScala() bool
	Length() int
	Index(i int) interface{}
	Truthy(i int) bool
	TruthyArr() []bool
	Type() BaseType
	SetNull(i int, isNull bool)
	Init(length int)
	SetScala(isScala bool)
	Seti(i int, v interface{})
}

type NullableVector struct {
	IsNullArr []bool
	IsScalaV  bool
}

func (v *NullableVector) SetNull(i int, isNull bool) {
	v.IsNullArr[i] = isNull
}

func (v *NullableVector) SetScala(isScala bool) {
	v.IsScalaV = isScala
}

func (v NullableVector) IsNull(i int) bool {
	if v.IsScalaV {
		return v.IsNullArr[0]
	}
	return v.IsNullArr[i]
}

func (v NullableVector) GetIsNullArr() []bool {
	return v.IsNullArr
}

func (v NullableVector) IsScala() bool {
	return v.IsScalaV
}

func (v NullableVector) Length() int {
	return len(v.IsNullArr)
}

func ToString(v INullableVector) string {
	ret := ""
	for i := 0; i < v.Length(); i++ {
		if v.IsNull(i) == true {
			ret += "null "
		} else {
			ret += fmt.Sprintf("%v ", v.Index(i))
		}
	}
	return "[" + ret + "]"
}

type NullableInt struct {
	NullableVector
	Values []int64
}

func (v *NullableInt) Init(length int) {
	v.IsNullArr = make([]bool, length)
	v.Values = make([]int64, length)
}
func (v NullableInt) Set(i int, val int64, isNull bool) {
	v.Values[i] = val
	v.IsNullArr[i] = isNull
}

func (v NullableInt) Seti(i int, val interface{}) {
	if vval, ok := val.(int64); ok {
		v.Set(i, vval, false)
	} else {
		v.SetNull(i, true)
	}
}

func (v NullableInt) Type() BaseType {
	return Int
}

func (v NullableInt) Truthy(i int) bool {
	return v.IsNullArr[i] == false && v.Values[i] != 0
}
func (v NullableInt) TruthyArr() []bool {
	arr := make([]bool, len(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false && v.Values[i] != 0
	}
	return arr
}

func (v NullableInt) Index(i int) interface{} {

	if v.IsScalaV {
		if v.IsNullArr[0] {
			return nil
		} else {
			return v.Values[0]
		}
	} else {
		if v.IsNullArr[i] {
			return nil
		} else {
			return v.Values[i]
		}
	}
}

type NullableFloat struct {
	NullableVector
	Values []float64
}

func (v *NullableFloat) Init(length int) {
	v.IsNullArr = make([]bool, length)
	v.Values = make([]float64, length)
}

func (v NullableFloat) Set(i int, val float64, isNull bool) {
	v.Values[i] = val
	v.IsNullArr[i] = isNull
}
func (v NullableFloat) Seti(i int, val interface{}) {
	if vval, ok := val.(float64); ok {
		v.Set(i, vval, false)
	} else {
		v.SetNull(i, true)
	}
}

func (v NullableFloat) Type() BaseType {
	return Float
}

func (v NullableFloat) Truthy(i int) bool {
	return v.IsNullArr[i] == false && v.Values[i] != 0
}

func (v NullableFloat) TruthyArr() []bool {
	arr := make([]bool, len(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false && v.Values[i] != 0
	}
	return arr
}

func (v NullableFloat) Index(i int) interface{} {
	if v.IsScalaV {
		if v.IsNullArr[0] {
			return nil
		} else {
			return v.Values[0]
		}
	} else {
		if v.IsNullArr[i] {
			return nil
		} else {
			return v.Values[i]
		}
	}
}

type NullableBool struct {
	NullableVector
	Values []bool
}

func (v NullableBool) Set(i int, val bool, isNull bool) {
	v.Values[i] = val
	v.IsNullArr[i] = isNull
}

func (v NullableBool) Seti(i int, val interface{}) {
	if vval, ok := val.(bool); ok {
		v.Set(i, vval, false)
	} else {
		v.SetNull(i, true)
	}
}

func (v *NullableBool) Init(length int) {
	v.IsNullArr = make([]bool, length)
	v.Values = make([]bool, length)
}

func (v NullableBool) Type() BaseType {
	return Bool
}

func (v NullableBool) Truthy(i int) bool {
	return v.IsNullArr[i] == false && v.Values[i]
}
func (v NullableBool) TruthyArr() []bool {
	arr := make([]bool, len(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false && v.Values[i]
	}
	return arr
}

func (v NullableBool) Index(i int) interface{} {
	if v.IsScalaV {
		if v.IsNullArr[0] {
			return nil
		} else {
			return v.Values[0]
		}
	} else {
		if v.IsNullArr[i] {
			return nil
		} else {
			return v.Values[i]
		}
	}
}

type NullableNumeric struct {
	NullableVector
	Values []int64
}

func (v NullableNumeric) Set(i int, val int64, isNull bool) {
	v.Values[i] = val
	v.IsNullArr[i] = isNull
}

func (v NullableNumeric) Seti(i int, val interface{}) {
	if vval, ok := val.(int64); ok {
		v.Set(i, vval, false)
	} else {
		v.SetNull(i, true)
	}
}

func (v *NullableNumeric) Init(length int) {
	v.IsNullArr = make([]bool, length)
	v.Values = make([]int64, length)
}

func (v NullableNumeric) Type() BaseType {
	return Numeric
}

func (v NullableNumeric) Index(i int) interface{} {
	if v.IsScalaV {
		if v.IsNullArr[0] {
			return nil
		} else {
			return v.Values[0]
		}
	} else {
		if v.IsNullArr[i] {
			return nil
		} else {
			return v.Values[i]
		}
	}
}

func (v NullableNumeric) Truthy(i int) bool {
	return v.IsNullArr[i] == false && v.Values[i] != 0
}

func (v NullableNumeric) TruthyArr() []bool {
	arr := make([]bool, len(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false && v.Values[i] != 0
	}
	return arr
}

type NullableText struct {
	NullableVector
	Values []string
}

func (v NullableText) Set(i int, val string, isNull bool) {
	v.Values[i] = val
	v.IsNullArr[i] = isNull
}

func (v NullableText) Seti(i int, val interface{}) {
	if vval, ok := val.(string); ok {
		v.Set(i, vval, false)
	} else {
		v.SetNull(i, true)
	}
}

func (v *NullableText) Init(length int) {
	v.IsNullArr = make([]bool, length)
	v.Values = make([]string, length)
}

func (v NullableText) Type() BaseType {
	return Text
}

func (v NullableText) Index(i int) interface{} {
	if v.IsScalaV {
		if v.IsNullArr[0] {
			return nil
		} else {
			return v.Values[0]
		}
	} else {
		if v.IsNullArr[i] {
			return nil
		} else {
			return v.Values[i]
		}
	}
}
func (v NullableText) Truthy(i int) bool {
	return v.IsNullArr[i] == false && v.Values[i] != ""
}

func (v NullableText) TruthyArr() []bool {
	arr := make([]bool, len(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false && v.Values[i] != ""
	}
	return arr
}

type NullableTimestamp struct {
	NullableVector
	Values []int64
}

func (v NullableTimestamp) Type() BaseType {
	return Timestamp
}

func (v NullableTimestamp) Set(i int, val int64, isNull bool) {
	v.Values[i] = val
	v.IsNullArr[i] = isNull
}

func (v NullableTimestamp) Seti(i int, val interface{}) {
	if vval, ok := val.(int64); ok {
		v.Set(i, vval, false)
	} else {
		v.SetNull(i, true)
	}
}

func (v *NullableTimestamp) Init(length int) {
	v.IsNullArr = make([]bool, length)
	v.Values = make([]int64, length)
}

func (v NullableTimestamp) Index(i int) interface{} {
	if v.IsScalaV {
		if v.IsNullArr[0] {
			return nil
		} else {
			return v.Values[0]
		}
	} else {
		if v.IsNullArr[i] {
			return nil
		} else {
			return v.Values[i]
		}
	}
}

func (v NullableTimestamp) Truthy(i int) bool {
	return !v.IsNullArr[i]
}

func (v NullableTimestamp) TruthyArr() []bool {
	arr := make([]bool, len(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false
	}
	return arr
}

func BuildValue(valueType BaseType, values ...interface{}) INullableVector {
	l := len(values)
	switch valueType {
	case Int:
		v := &NullableInt{
			NullableVector: NullableVector{
				IsNullArr: make([]bool, l),
				IsScalaV:  false,
			},
			Values: make([]int64, l),
		}
		for i := 0; i < l; i++ {
			if values[i] == nil {
				v.IsNullArr[i] = true
			} else {
				switch values[i].(type) {
				case int64:
					v.Values[i] = values[i].(int64)
				case int:
					v.Values[i] = int64(values[i].(int))
				}

			}
		}
		return v
	case Float:
		v := &NullableFloat{
			NullableVector: NullableVector{
				IsNullArr: make([]bool, l),
				IsScalaV:  false,
			},
			Values: make([]float64, l),
		}
		for i := 0; i < l; i++ {
			if values[i] == nil {
				v.IsNullArr[i] = true
			} else {
				v.Values[i] = values[i].(float64)
			}
		}
		return v
	case Text:
		v := &NullableText{
			NullableVector: NullableVector{
				IsNullArr: make([]bool, l),
				IsScalaV:  false,
			},
			Values: make([]string, l),
		}
		for i := 0; i < l; i++ {
			if values[i] == nil {
				v.IsNullArr[i] = true
			} else {
				v.Values[i] = values[i].(string)
			}
		}
		return v
	case Bool:
		v := &NullableBool{
			NullableVector: NullableVector{
				IsNullArr: make([]bool, l),
				IsScalaV:  false,
			},
			Values: make([]bool, l),
		}
		for i := 0; i < l; i++ {
			if values[i] == nil {
				v.IsNullArr[i] = true
			} else {
				v.Values[i] = values[i].(bool)
			}
		}
		return v
	}
	return nil
}
