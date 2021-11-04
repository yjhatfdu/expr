package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type BaseType int

var _, LocalOffsetSec = time.Now().Local().Zone()
var LocalOffsetNano = int64(LocalOffsetSec) * int64(time.Second)

const (
	ScalaOffset          = 1000
	Any         BaseType = iota + 1000
	Null
	Int
	Float
	Numeric
	Text
	Bool
	Timestamp
	Date
	Time
	Interval
	Blob
	IntA
	TextA
	ScalaTypes = Any + ScalaOffset
	IntS       = Int + ScalaOffset
	FloatS     = Float + ScalaOffset
	TextS      = Text + ScalaOffset
	BoolS      = Bool + ScalaOffset
	NumericS   = Numeric + ScalaOffset
	TimestampS = Timestamp + ScalaOffset
	DateS      = Date + ScalaOffset
	TimeS      = Time + ScalaOffset
	IntervalS  = Interval + ScalaOffset
	BlobS      = Blob + ScalaOffset
	IntAS      = IntA + ScalaOffset
	TextAS     = TextA + ScalaOffset
)

var typeNames = map[BaseType]string{
	Any:        "Any",
	Int:        "Int",
	Float:      "Float",
	Numeric:    "Numeric",
	Text:       "Text",
	Bool:       "Bool",
	Timestamp:  "Timestamp",
	Date:       "Date",
	Time:       "Time",
	Interval:   "Interval",
	Blob:       "Blob",
	IntA:       "IntA",
	TextA:      "TextA",
	BlobS:      "BlobS",
	IntS:       "IntS",
	FloatS:     "FloatS",
	NumericS:   "NumericS",
	TextS:      "TextS",
	BoolS:      "BoolS",
	TimestampS: "TimestampS",
	DateS:      "DateS",
	TimeS:      "TimeS",
	IntervalS:  "IntervalS",
	IntAS:      "IntAS",
	TextAS:     "TextAS",
}

var typeMapping = map[string]BaseType{
	"Any":        Any,
	"Int":        Int,
	"IntA":       IntA,
	"TextA":      TextA,
	"Float":      Float,
	"Numeric":    Numeric,
	"Text":       Text,
	"Bool":       Bool,
	"Timestamp":  Timestamp,
	"Date":       Date,
	"Time":       Time,
	"Interval":   Interval,
	"Blob":       Blob,
	"BlobS":      BlobS,
	"IntS":       IntS,
	"FloatS":     FloatS,
	"NumericS":   NumericS,
	"TextS":      TextS,
	"BoolS":      BoolS,
	"TimestampS": TimestampS,
	"DateS":      DateS,
	"TimeS":      TimeS,
	"IntervalS":  IntervalS,
	"IntAS":      IntAS,
	"TextAS":     TextAS,
}

func (t *BaseType) MarshalJSON() ([]byte, error) {
	if *t < ScalaTypes {
		if name, ok := typeNames[*t]; ok {
			return json.Marshal(name + "[S]")
		} else {
			return json.Marshal("undefined type")
		}
	} else {
		if name, ok := typeNames[*t]; ok {
			return json.Marshal(name)
		} else {
			return json.Marshal("undefined type")
		}
	}
}

func GetTypeName(t BaseType) string {
	if t < ScalaTypes {
		if name, ok := typeNames[t]; ok {
			return name + "[S]"
		} else {
			return "undefined type"
		}
	} else {
		if name, ok := typeNames[t]; ok {
			return name
		} else {
			return "undefined type"
		}
	}

}

func GetTypeByName(n string) (BaseType, bool) {
	t, ok := typeMapping[n]
	return t, ok
}

type VectorError struct {
	Index int
	Error error
}

type INullableVector interface {
	IsNull(i int) bool
	GetIsNullArr() []bool
	IsScala() bool
	Length() int
	Index(i int) interface{}
	Truthy(i int) bool
	TruthyArr() []bool
	FalseArr() []bool
	Type() BaseType
	SetNull(i int, isNull bool)
	Init(length int)
	SetScala(isScala bool)
	Seti(i int, v interface{})
	AddError(err *VectorError)
	GetErrors() []*VectorError
	GetFilterArr() []bool
	SetFilterArr(arr []bool)
	InitFilterArr() []bool
	Copy() INullableVector
	Concat(vector INullableVector) (INullableVector, error)
	InterfaceArr() []interface{}
}

type NullableVector struct {
	IsNullArr []bool
	IsScalaV  bool
	errors    []*VectorError
	FilterArr []bool
}

func (v *NullableVector) AddError(err *VectorError) {
	v.errors = append(v.errors, err)
}

func (v *NullableVector) GetErrors() []*VectorError {
	return v.errors
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

func (v NullableVector) GetFilterArr() []bool {
	return v.FilterArr
}
func (v *NullableVector) SetFilterArr(arr []bool) {
	v.FilterArr = arr
}

func (v *NullableVector) InitFilterArr() []bool {
	v.FilterArr = make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	return v.FilterArr
}

func (v NullableVector) copy() NullableVector {
	v2 := NullableVector{
		IsNullArr: make([]bool, len(v.IsNullArr), cap(v.IsNullArr)),
		IsScalaV:  v.IsScalaV,
		errors:    nil,
		FilterArr: nil,
	}
	copy(v2.IsNullArr, v.IsNullArr)
	if v.errors != nil {
		v2.errors = make([]*VectorError, len(v.errors))
		copy(v2.errors, v.errors)
	}
	if v.FilterArr != nil {
		v2.FilterArr = make([]bool, len(v.FilterArr), cap(v.FilterArr))
		copy(v2.FilterArr, v.FilterArr)
	}
	return v2
}

// for debug
// todo optimize performance
func ToString(v INullableVector) string {
	if v == nil {
		return "null"
	}
	var retSegs []string
	for i := 0; i < v.Length(); i++ {
		if v.IsNull(i) == true {
			retSegs = append(retSegs, "NULL")
		} else {
			time.Now().Month()
			val := v.Index(i)
			switch v.Type() {
			case Int, Float, Bool, IntS, FloatS, BoolS, IntA, TextA:
				retSegs = append(retSegs, fmt.Sprintf("%v", val))
			case Text, TextS:
				retSegs = append(retSegs, strconv.Quote(val.(string)))
			case Numeric, NumericS:
				retSegs = append(retSegs, Numeric2Text(val.(int64), v.(*NullableNumeric).Scale))
			case Timestamp, TimestampS:
				t := time.Unix(0, val.(int64)).In(time.Local)
				retSegs = append(retSegs, t.Format(time.RFC3339))
			case Time, TimeS:
				t := time.Unix(0, val.(int64)).In(time.UTC)
				retSegs = append(retSegs, t.Format("15:04:05"))
			case Date, DateS:
				t := time.Unix(0, val.(int64)).In(time.UTC)
				retSegs = append(retSegs, t.Format("2006-01-02"))
			}
		}
	}
	tname := GetTypeName(v.Type())
	tname = tname[:len(tname)-3]
	if v.IsScala() {
		return tname + "S(" + strings.Join(retSegs, ",") + ")"
	} else {
		return tname + "V[" + strings.Join(retSegs, ",") + "]"
	}
}

type NullableInt struct {
	NullableVector
	Values []int64
}

func (v *NullableInt) InterfaceArr() []interface{} {
	out := make([]interface{}, v.Length())
	for i := 0; i < len(out); i++ {
		if v.IsNullArr[i] {
			out[i] = nil
		} else {
			out[i] = v.Values[i]
		}
	}
	return out
}

func (v *NullableInt) Concat(vector INullableVector) (INullableVector, error) {
	other, ok := vector.(*NullableInt)
	if !ok {
		return nil, fmt.Errorf("NullableInt must concat NullableInt")
	}

	var r = v.Copy().(*NullableInt)
	if !v.IsScalaV {
		r.Values = append(r.Values, other.Values...)
		r.errors = append(r.errors, other.errors...)
		r.FilterArr = append(r.FilterArr, other.FilterArr...)
		r.IsNullArr = append(r.IsNullArr, other.IsNullArr...)
		return r, nil
	} else {
		return v, nil
	}
}

func (v *NullableInt) Init(length int) {
	v.IsNullArr = make([]bool, length, 32*(length/32+1))
	v.Values = make([]int64, length, 8*(length/8+1))
}
func (v NullableInt) Set(i int, val int64, isNull bool) {
	if i >= len(v.Values) {
		return
	}
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
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = v.Values[i] != 0 && (v.IsNullArr[i] == false)
	}
	return arr
}

func (v NullableInt) FalseArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = v.Values[i] == 0 || v.IsNullArr[i]
	}
	return arr
}

func (v NullableInt) Copy() INullableVector {
	return &v
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

func (v *NullableFloat) InterfaceArr() []interface{} {
	out := make([]interface{}, v.Length())
	for i := 0; i < len(out); i++ {
		if v.IsNullArr[i] {
			out[i] = nil
		} else {
			out[i] = v.Values[i]
		}
	}
	return out
}

func (v *NullableFloat) Concat(vector INullableVector) (INullableVector, error) {
	other, ok := vector.(*NullableFloat)
	if !ok {
		return nil, fmt.Errorf("NullableFloat must concat NullableFloat")
	}

	var r = v.Copy().(*NullableFloat)
	if !v.IsScalaV {
		r.Values = append(r.Values, other.Values...)
		r.errors = append(r.errors, other.errors...)
		r.FilterArr = append(r.FilterArr, other.FilterArr...)
		r.IsNullArr = append(r.IsNullArr, other.IsNullArr...)
		return r, nil
	} else {
		return v, nil
	}
}

func (v *NullableFloat) Init(length int) {
	v.IsNullArr = make([]bool, length, 32*(length/32+1))
	v.Values = make([]float64, length, 4*(length/4+1))
}

func (v NullableFloat) Set(i int, val float64, isNull bool) {
	if i >= len(v.Values) {
		return
	}
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
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	l := len(v.IsNullArr)
	for i := 0; i < l; i++ {
		arr[i] = v.IsNullArr[i] == false && v.Values[i] != 0
	}
	return arr
}

func (v NullableFloat) FalseArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = v.Values[i] == 0 || v.IsNullArr[i]
	}
	return arr
}

func (v NullableFloat) Copy() INullableVector {
	return &v
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

func (v *NullableBool) InterfaceArr() []interface{} {
	out := make([]interface{}, v.Length())
	for i := 0; i < len(out); i++ {
		if v.IsNullArr[i] {
			out[i] = nil
		} else {
			out[i] = v.Values[i]
		}
	}
	return out
}

func (v *NullableBool) Concat(vector INullableVector) (INullableVector, error) {
	other, ok := vector.(*NullableBool)
	if !ok {
		return nil, fmt.Errorf("NullableBool must concat NullableBool")
	}

	var r = v.Copy().(*NullableBool)
	if !v.IsScalaV {
		r.Values = append(r.Values, other.Values...)
		r.errors = append(r.errors, other.errors...)
		r.FilterArr = append(r.FilterArr, other.FilterArr...)
		r.IsNullArr = append(r.IsNullArr, other.IsNullArr...)
		return r, nil
	} else {
		return v, nil
	}
}

func (v NullableBool) Set(i int, val bool, isNull bool) {
	if i >= len(v.Values) {
		return
	}
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
	v.IsNullArr = make([]bool, length, 32*(length/32+1))
	v.Values = make([]bool, length, 32*(length/32+1))
}

func (v NullableBool) Type() BaseType {
	return Bool
}

func (v NullableBool) Truthy(i int) bool {
	return v.IsNullArr[i] == false && v.Values[i]
}
func (v NullableBool) TruthyArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false && v.Values[i]
	}
	return arr
}

func (v NullableBool) FalseArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = v.Values[i] == false || v.IsNullArr[i]
	}
	return arr
}
func (v NullableBool) Copy() INullableVector {
	return &v
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
	Scale  int
}

func (v *NullableNumeric) InterfaceArr() []interface{} {
	out := make([]interface{}, v.Length())
	for i := 0; i < len(out); i++ {
		if v.IsNullArr[i] {
			out[i] = nil
		} else {
			out[i] = v.Values[i]
		}
	}
	return out
}

func (v *NullableNumeric) Concat(vector INullableVector) (INullableVector, error) {
	other, ok := vector.(*NullableNumeric)
	if !ok {
		return nil, fmt.Errorf("NullableNumeric must concat NullableNumeric")
	}

	var r = v.Copy().(*NullableNumeric)
	if !v.IsScalaV {
		r.Values = append(r.Values, other.Values...)
		r.errors = append(r.errors, other.errors...)
		r.FilterArr = append(r.FilterArr, other.FilterArr...)
		r.IsNullArr = append(r.IsNullArr, other.IsNullArr...)
		return r, nil
	} else {
		return v, nil
	}
}

func (v NullableNumeric) Set(i int, val int64, isNull bool) {
	if i >= len(v.Values) {
		return
	}
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
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false && v.Values[i] != 0
	}
	return arr
}

func (v NullableNumeric) Copy() INullableVector {
	return &v
}

func (v NullableNumeric) FalseArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = v.Values[i] == 0 || v.IsNullArr[i]
	}
	return arr
}

type NullableText struct {
	NullableVector
	Values []string
}

func (v *NullableText) InterfaceArr() []interface{} {
	out := make([]interface{}, v.Length())
	for i := 0; i < len(out); i++ {
		if v.IsNullArr[i] {
			out[i] = nil
		} else {
			out[i] = v.Values[i]
		}
	}
	return out
}

func (v *NullableText) Concat(vector INullableVector) (INullableVector, error) {
	other, ok := vector.(*NullableText)
	if !ok {
		return nil, fmt.Errorf("NullableText must concat NullableText")
	}

	var r = v.Copy().(*NullableText)
	if !v.IsScalaV {
		r.Values = append(r.Values, other.Values...)
		r.errors = append(r.errors, other.errors...)
		r.FilterArr = append(r.FilterArr, other.FilterArr...)
		r.IsNullArr = append(r.IsNullArr, other.IsNullArr...)
		return r, nil
	} else {
		return v, nil
	}
}

func (v NullableText) Set(i int, val string, isNull bool) {
	if i >= len(v.Values) {
		return
	}
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
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false && v.Values[i] != ""
	}
	return arr
}

func (v NullableText) FalseArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = v.Values[i] == "" || v.IsNullArr[i]
	}
	return arr
}

func (v NullableText) Copy() INullableVector {
	return &v
}

type NullableTimestamp struct {
	NullableVector
	Values []int64
	TsType BaseType //one of Timestamp,Date,Time, Interval
}

func (v *NullableTimestamp) InterfaceArr() []interface{} {
	out := make([]interface{}, v.Length())
	for i := 0; i < len(out); i++ {
		if v.IsNullArr[i] {
			out[i] = nil
		} else {
			out[i] = v.Values[i]
		}
	}
	return out
}

func (v *NullableTimestamp) Concat(vector INullableVector) (INullableVector, error) {
	other, ok := vector.(*NullableTimestamp)
	if !ok {
		return nil, fmt.Errorf("NullableTimestamp must concat NullableTimestamp")
	}

	var r = v.Copy().(*NullableTimestamp)
	if !v.IsScalaV {
		r.Values = append(r.Values, other.Values...)
		r.errors = append(r.errors, other.errors...)
		r.FilterArr = append(r.FilterArr, other.FilterArr...)
		r.IsNullArr = append(r.IsNullArr, other.IsNullArr...)
		return r, nil
	} else {
		return v, nil
	}
}

func (v NullableTimestamp) Type() BaseType {
	return v.TsType
}

func (v NullableTimestamp) Set(i int, val int64, isNull bool) {
	if i >= len(v.Values) {
		return
	}
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
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := 0; i < len(v.IsNullArr); i++ {
		arr[i] = v.IsNullArr[i] == false
	}
	return arr
}

func (v NullableTimestamp) FalseArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = v.IsNullArr[i]
	}
	return arr
}

func (v NullableTimestamp) Copy() INullableVector {
	return &v
}

func BuildValue(valueType BaseType, values ...interface{}) INullableVector {
	l := len(values)
	switch valueType {
	case Int:
		v := &NullableInt{}
		v.Init(l)
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
		v := &NullableFloat{}
		v.Init(l)
		for i := 0; i < l; i++ {
			if values[i] == nil {
				v.IsNullArr[i] = true
			} else {
				v.Values[i] = values[i].(float64)
			}
		}
		return v
	case Text:
		v := &NullableText{}
		v.Init(l)
		for i := 0; i < l; i++ {
			if values[i] == nil {
				v.IsNullArr[i] = true
			} else {
				v.Values[i] = values[i].(string)
			}
		}
		return v
	case Bool:
		v := &NullableBool{}
		v.Init(l)
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

func GetFilteredMaskOfVectors(vec []INullableVector) []bool {
	v0 := vec[0]
	l := v0.Length()
	out := make([]bool, l, 32*(l/32+1))
	for _, v := range vec {
		farr := v.GetFilterArr()
		if farr != nil {
			orBool(out, farr, out)
		}
	}
	return out
}

type NullableIntArray struct {
	NullableVector
	Values [][]int64
}

func (v *NullableIntArray) InterfaceArr() []interface{} {
	out := make([]interface{}, v.Length())
	for i := 0; i < len(out); i++ {
		if v.IsNullArr[i] {
			out[i] = nil
		} else {
			out[i] = v.Values[i]
		}
	}
	return out
}

func (v *NullableIntArray) Concat(vector INullableVector) (INullableVector, error) {
	other, ok := vector.(*NullableIntArray)
	if !ok {
		return nil, fmt.Errorf("NullableIntArray must concat NullableIntArray")
	}

	var r = v.Copy().(*NullableIntArray)
	if !v.IsScalaV {
		r.Values = append(r.Values, other.Values...)
		r.errors = append(r.errors, other.errors...)
		r.FilterArr = append(r.FilterArr, other.FilterArr...)
		r.IsNullArr = append(r.IsNullArr, other.IsNullArr...)
		return r, nil
	} else {
		return v, nil
	}
}

func (v *NullableIntArray) Init(length int) {
	v.IsNullArr = make([]bool, length, 32*(length/32+1))
	v.Values = make([][]int64, length, 8*(length/8+1))
}

func (v NullableIntArray) Set(i int, val []int64, isNull bool) {
	if i >= len(v.Values) {
		return
	}
	v.Values[i] = val
	v.IsNullArr[i] = isNull
}

func (v NullableIntArray) Seti(i int, val interface{}) {
	if vval, ok := val.([]int64); ok {
		v.Set(i, vval, false)
	} else {
		v.SetNull(i, true)
	}
}

func (v NullableIntArray) Type() BaseType {
	return IntA
}

func (v NullableIntArray) Truthy(i int) bool {
	return v.IsNullArr[i] == false && len(v.Values[i]) > 0
}

func (v NullableIntArray) TruthyArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = len(v.Values[i]) > 0 && (v.IsNullArr[i] == false)
	}
	return arr
}

func (v NullableIntArray) FalseArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = len(v.Values[i]) == 0 || v.IsNullArr[i]
	}
	return arr
}

func (v NullableIntArray) Copy() INullableVector {
	return &v
}

func (v NullableIntArray) Index(i int) interface{} {

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

type NullableTextArray struct {
	NullableVector
	Values [][]string
}

func (v *NullableTextArray) InterfaceArr() []interface{} {
	out := make([]interface{}, v.Length())
	for i := 0; i < len(out); i++ {
		if v.IsNullArr[i] {
			out[i] = nil
		} else {
			out[i] = v.Values[i]
		}
	}
	return out
}

func (v *NullableTextArray) Concat(vector INullableVector) (INullableVector, error) {
	other, ok := vector.(*NullableIntArray)
	if !ok {
		return nil, fmt.Errorf("NullableIntArray must concat NullableIntArray")
	}

	var r = v.Copy().(*NullableIntArray)
	if !v.IsScalaV {
		r.Values = append(r.Values, other.Values...)
		r.errors = append(r.errors, other.errors...)
		r.FilterArr = append(r.FilterArr, other.FilterArr...)
		r.IsNullArr = append(r.IsNullArr, other.IsNullArr...)
		return r, nil
	} else {
		return v, nil
	}
}

func (v *NullableTextArray) Init(length int) {
	v.IsNullArr = make([]bool, length, 32*(length/32+1))
	v.Values = make([][]string, length, 8*(length/8+1))
}

func (v *NullableTextArray) Set(i int, val []string, isNull bool) {
	if i >= len(v.Values) {
		return
	}
	v.Values[i] = val
	v.IsNullArr[i] = isNull
}

func (v *NullableTextArray) Seti(i int, val interface{}) {
	if vval, ok := val.([]string); ok {
		v.Set(i, vval, false)
	} else {
		v.SetNull(i, true)
	}
}

func (v *NullableTextArray) Type() BaseType {
	return TextA
}

func (v *NullableTextArray) Truthy(i int) bool {
	return v.IsNullArr[i] == false && len(v.Values[i]) > 0
}

func (v *NullableTextArray) TruthyArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = len(v.Values[i]) > 0 && (v.IsNullArr[i] == false)
	}
	return arr
}

func (v *NullableTextArray) FalseArr() []bool {
	arr := make([]bool, len(v.IsNullArr), cap(v.IsNullArr))
	for i := range v.IsNullArr {
		arr[i] = len(v.Values[i]) == 0 || v.IsNullArr[i]
	}
	return arr
}

func (v *NullableTextArray) Copy() INullableVector {
	return v
}

func (v *NullableTextArray) Index(i int) interface{} {

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

type NullVector struct {
	IsNullArr []bool
	FilterArr []bool
	IsScalaV  bool
	Values    []string
}

func (v NullVector) InterfaceArr() []interface{} {
	out := make([]interface{}, v.Length())
	for i := 0; i < len(out); i++ {
			out[i] = nil
	}
	return out
}

func (n NullVector) IsNull(i int) bool {
	return true
}

func (n NullVector) GetIsNullArr() []bool {
	return n.IsNullArr
}

func (n NullVector) IsScala() bool {
	return true
}

func (n NullVector) Length() int {
	return 1
}

func (n NullVector) Index(i int) interface{} {
	return nil
}

func (n NullVector) Truthy(i int) bool {
	return false
}

func (n NullVector) TruthyArr() []bool {
	return []bool{false}
}

func (n NullVector) FalseArr() []bool {
	return []bool{true}
}

func (n NullVector) Type() BaseType {
	return Any
}

func (n NullVector) SetNull(i int, isNull bool) {
	return
}

func (n NullVector) Init(length int) {
	return
}

func (n NullVector) SetScala(isScala bool) {
	return
}

func (n NullVector) Seti(i int, v interface{}) {
	return
}

func (n NullVector) AddError(err *VectorError) {
	return
}

func (n NullVector) GetErrors() []*VectorError {
	return nil
}

func (n NullVector) GetFilterArr() []bool {
	return nil
}

func (n NullVector) SetFilterArr(arr []bool) {
	return
}

func (n NullVector) InitFilterArr() []bool {
	return n.FilterArr
}

func (n NullVector) Copy() INullableVector {
	return n
}

func (n NullVector) Concat(vector INullableVector) (INullableVector, error) {
	return n, nil
}
