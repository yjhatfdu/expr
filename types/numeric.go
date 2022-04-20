package types

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

//var pow10 = [16]int64{
//	1, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, 1e10, 1e11, 1e12, 1e13, 1e14, 1e15,
//}

type Decimal struct {
	i     *big.Int
	scale int
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.String())), nil
}

func Int2Decimal(i int64, scale int) Decimal {
	if scale == 0 {
		scale = 4
	}
	return Decimal{
		i:     new(big.Int).Mul(big.NewInt(i), GenPow(scale)),
		scale: scale,
	}
}

func Float2Decimal(f float64, scale int) Decimal {
	d, _ := Text2Decimal(strconv.FormatFloat(f, 'f', -1, 64))
	if d.scale < scale {
		_scale := scale - d.scale
		d.i = big.NewInt(0).Mul(d.i, GenPow(_scale))
		d.scale = scale
	}
	return d
}

func GenPow(scale int) *big.Int {
	src := "1"
	for i := 0; i < scale; i++ {
		src += "0"
	}
	num := &big.Int{}
	num.SetString(src, 10)
	return num
}

//func Float2numeric(f float64, scale int) int64 {
//	return int64(f * float64(pow10[scale]))
//}
//
//func Numeric2Float(n int64, scale int) float64 {
//	return float64(n) / float64(pow10[scale])
//}
//
//func Numeric2Int(n int64, scale int) int64 {
//	return n / pow10[scale]
//}

func (d Decimal) ToInt() int64 {
	if d.scale == 0 {
		return d.i.Int64()
	}
	return (&big.Int{}).Div(d.i, GenPow(d.scale)).Int64()
}

func (d Decimal) ToFloat() float64 {
	str := d.String()
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func (d Decimal) abs() Decimal {
	return Decimal{i: (&big.Int{}).Abs(d.i), scale: d.scale}
}

func (d Decimal) IsZero() bool {
	return d.i.Cmp(&big.Int{}) == 0
}

//
//func abs(n int64) int64 {
//	return int64(math.Abs(float64(n)))
//}

func (d Decimal) String() string {
	if d.scale == 0 {
		return d.i.String()
	}
	//absN := d.abs()
	isNeg := d.i.Sign() < 0

	neg := ""
	if isNeg {
		neg = "-"
	}

	x := d.i.String()

	if len(x) > d.scale {
		return fmt.Sprintf("%s%s.%s", neg, x[0:len(x)-d.scale], x[len(x)-d.scale:])
	} else {
		return fmt.Sprintf("%s0.%s", neg, strings.Repeat("0", d.scale-len(x))+x)
	}

	//num, frac := (&big.Int{}).DivMod(absN.i, GenPow(d.scale), &big.Int{})
	//if frac.Cmp(&big.Int{}) == 0 {
	//	return num.String()
	//}
	//if isNeg {
	//	return "-" + num.String() + "." + frac.String()
	//} else {
	//	return num.String() + "." + frac.String()
	//}
}

func (d Decimal) StringScale(n int) string {
	if d.scale == 0 {
		return d.i.String()
	}
	absN := d.abs()
	isNeg := d.i.Sign() < 0

	num, frac := (&big.Int{}).DivMod(absN.i, GenPow(d.scale), &big.Int{})
	if frac.Cmp(&big.Int{}) == 0 {
		return num.String()
	}
	fracString := frac.String()
	if n == 0 {
		return d.String()
	}
	if len(fracString) < n {
		fracString = fracString + strings.Repeat("0", n-len(fracString))
	} else {
		fracString = fracString[:n]
	}
	if isNeg {
		return "-" + num.String() + "." + fracString
	} else {
		return num.String() + "." + fracString
	}
}

//func Numeric2Text(n int64, scale int) string {
//	if scale == 0 {
//		return strconv.FormatInt(n/pow10[scale], 10)
//	}
//
//	absN := abs(n)
//	frac := strconv.FormatInt(absN%pow10[scale], 10)
//	if len(frac) < scale {
//		frac = strings.Repeat("0", scale-len(frac)) + frac
//	}
//	return strconv.FormatInt(n/pow10[scale], 10) + "." + frac
//}

func CompareDecimal(d1, d2 Decimal) int {
	if d1.scale == d2.scale {
		return d1.i.Cmp(d2.i)
	}
	if d1.scale > d2.scale {
		i2 := (&big.Int{}).Mul(d2.i, GenPow(d1.scale-d2.scale))
		return d1.i.Cmp(i2)
	}
	i1 := (&big.Int{}).Mul(d1.i, GenPow(d2.scale-d1.scale))
	return i1.Cmp(d2.i)
}

func DivideDecimal(d1, d2 Decimal) Decimal {
	if d2.IsZero() {
		panic("divide by zero")
	}
	//if d1.scale == d2.scale {
	//	return Decimal{i: (&big.Int{}).Div(d1.i, d2.i), scale: d1.scale}
	//}
	//if d1.scale > d2.scale {
	//	i2 := (&big.Int{}).Mul(d2.i, big.NewInt(pow10[d1.scale-d2.scale]))
	//	return Decimal{i: (&big.Int{}).Div(d1.i, i2), scale: d1.scale}
	//}
	//i1 := (&big.Int{}).Mul(d1.i, big.NewInt(pow10[d2.scale-d1.scale]))
	return Decimal{i: (&big.Int{}).Div((&big.Int{}).Mul(d1.i, GenPow(d2.scale)), d2.i), scale: d2.scale}
}

func MulDecimal(d1, d2 Decimal) Decimal {
	return Decimal{
		i:     (&big.Int{}).Mul(d1.i, d2.i),
		scale: d1.scale + d2.scale,
	}
}

func AddDecimal(d1, d2 Decimal) Decimal {
	if d1.scale == d2.scale {
		return Decimal{
			i:     (&big.Int{}).Add(d1.i, d2.i),
			scale: d1.scale,
		}
	}
	if d1.scale > d2.scale {
		i2 := (&big.Int{}).Mul(d2.i, GenPow(d1.scale-d2.scale))
		return Decimal{
			i:     (&big.Int{}).Add(d1.i, i2),
			scale: d1.scale,
		}
	}
	i1 := (&big.Int{}).Mul(d1.i, GenPow(d2.scale-d1.scale))
	return Decimal{
		i:     (&big.Int{}).Add(i1, d2.i),
		scale: d2.scale,
	}
}

func MinusDecimal(d1, d2 Decimal) Decimal {
	if d1.scale == d2.scale {
		return Decimal{
			i:     (&big.Int{}).Sub(d1.i, d2.i),
			scale: d1.scale,
		}
	}
	if d1.scale > d2.scale {
		i2 := (&big.Int{}).Mul(d2.i, GenPow(d1.scale-d2.scale))
		return Decimal{
			i:     (&big.Int{}).Sub(d1.i, i2),
			scale: d1.scale,
		}
	}
	i1 := (&big.Int{}).Mul(d1.i, GenPow(d2.scale-d1.scale))
	return Decimal{
		i:     (&big.Int{}).Sub(i1, d2.i),
		scale: d2.scale,
	}
}

//func CompareNumeric(n1 int64, scale1 int, n2 int64, scale2 int) int64 {
//	if scale1 != scale2 {
//		if scale1 > scale2 {
//			n2 *= pow10[scale1-scale2]
//		} else {
//			n1 *= pow10[scale2-scale1]
//		}
//	}
//	return n2 - n1
//}

//func NumericScale(s1, s2 int) int {
//	if s1 > s2 {
//		return s1
//	} else {
//		return s2
//	}
//}
//
//func NormalizeNumeric(n int64, fromScale int, toScale int) int64 {
//	if toScale >= fromScale {
//		return n * pow10[toScale-fromScale]
//	} else {
//		return n / pow10[fromScale-toScale]
//	}
//}

//func CompareNumericInt(n int64, scale int, intV int64) int64 {
//	return CompareNumeric(n, scale, Int2numeric(intV, scale), scale)
//}

//func CompareNumericFloat(n int64, scale int, floatV float64) int64 {
//	fn := Numeric2Float(n, scale)
//	if math.Abs(fn-floatV) <= math.SmallestNonzeroFloat64 {
//		return 0
//	}
//	if fn > floatV {
//		return 1
//	} else {
//		return 0
//	}
//	//return CompareNumeric(n, scale, Float2numeric(floatV, scale), scale)
//}

//func Text2Numeric(s string, scale int) (int64, error) {
//	segs := strings.SplitN(s, ".", 2)
//	if len(segs) == 1 {
//		i, err := strconv.ParseInt(segs[0], 10, 64)
//		if err != nil {
//			return 0, err
//		} else {
//			return pow10[scale] * i, nil
//		}
//	} else if len(segs) == 2 {
//		i, err := strconv.ParseInt(segs[0], 10, 64)
//		if err != nil {
//			return 0, err
//		}
//		f, err := strconv.ParseInt(segs[1], 10, 64)
//		if err != nil {
//			return 0, err
//		}
//		rs := len(segs[1])
//		n := i*pow10[rs] + f
//		if rs > scale {
//			n = n / (pow10[rs-scale])
//		} else if rs < scale {
//			n = n * (pow10[scale-rs])
//		}
//
//		return n, nil
//	} else {
//		return 0, fmt.Errorf("invalid numeric '%s'", s)
//	}
//}

func Text2Decimal(s string) (n Decimal, err error) {
	segs := strings.SplitN(s, ".", 2)
	if len(segs) == 1 {
		i := &big.Int{}
		_, ok := i.SetString(s, 10)
		i.Mul(i, big.NewInt(10000))
		if !ok {
			return Decimal{}, fmt.Errorf("invalid numeric '%s'", s)
		}
		return Decimal{i: i, scale: 4}, nil
	} else if len(segs) == 2 {
		num := &big.Int{}
		if _, ok := num.SetString(segs[0], 10); !ok {
			return Decimal{}, fmt.Errorf("invalid numeric '%s'", s)
		}
		frac := &big.Int{}
		if _, ok := frac.SetString(segs[1], 10); !ok {
			return Decimal{}, fmt.Errorf("invalid numeric '%s'", s)
		}
		scale := len(segs[1])
		i := &big.Int{}
		i = i.Add(num.Mul(num, GenPow(scale)), frac)
		return Decimal{i: i, scale: scale}, nil
	} else {
		return Decimal{}, fmt.Errorf("invalid numeric '%s'", s)
	}
}
