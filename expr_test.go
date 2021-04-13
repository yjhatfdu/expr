package expr

import (
	"github.com/yjhatfdu/expr/types"
	"testing"
	"time"
)

func TestExpr_Minus(t *testing.T) {
	code := `-1`
	p, err := Compile(code, nil, nil)
	if err != nil {
		t.Error(err)
		return
	}
	ret, err := p.Run(nil, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(types.ToString(ret))
}

func TestExpr(t *testing.T) {
	code := `1-1`
	p, err := Compile(code, nil, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run(nil, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}
func TestExpr2(t *testing.T) {
	code := `length($1)<10 and length($1)>2`
	p, err := Compile(code, []types.BaseType{types.Text}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run([]types.INullableVector{types.BuildValue(types.Text, nil, "1", "12", "123", "1234")}, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func TestExpr3(t *testing.T) {
	code := "($1+$1*0.1):Float:Text"
	p, err := Compile(code, []types.BaseType{types.Int}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run([]types.INullableVector{types.BuildValue(types.Int, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10)}, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func TestExprCoalesce(t *testing.T) {
	code := "coalesce($1,$2,1)"
	p, err := Compile(code, []types.BaseType{types.Int, types.Int}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run([]types.INullableVector{types.BuildValue(types.Int, nil, 2, 3, 4, 5, 6, nil, 8, 9, 10), types.BuildValue(types.Int, 10, 2, 3, 4, 5, 6, nil, 8, 9, 10)}, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func TestExprMultiIf(t *testing.T) {
	code := `multiIf($1==2,"is2",$1==3,"is3","err")`
	p, err := Compile(code, []types.BaseType{types.Int, types.Int}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run([]types.INullableVector{types.BuildValue(types.Int, nil, 2, 3, 4, 5, 6, nil, 8, 9, 10), types.BuildValue(types.Int, 10, 2, 3, 4, 5, 6, nil, 8, 9, 10)}, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}
func TestExprCast(t *testing.T) {
	code := "`12`"
	p, err := Compile(code, []types.BaseType{types.Int, types.Int}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run([]types.INullableVector{types.BuildValue(types.Int, nil, 2, 3, 4, 5, 6, nil, 8, 9, 10), types.BuildValue(types.Int, 10, 2, 3, 4, 5, 6, nil, 8, 9, 10)}, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func TestExprNow(t *testing.T) {
	code := `now()`
	p, err := Compile(code, []types.BaseType{}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run(nil, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func TestExprNowDate(t *testing.T) {
	code := `now():Text`
	p, err := Compile(code, []types.BaseType{}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run(nil, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func TestExprNowTime(t *testing.T) {
	code := `now():Time`
	p, err := Compile(code, []types.BaseType{}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run(nil, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}
func TestExprNowYear(t *testing.T) {
	code := `now|getYear|add(10)`
	p, err := Compile(code, []types.BaseType{}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run(nil, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func TestExprNowYearWithoutP(t *testing.T) {
	code := `" nihao   "|trim|length`
	p, err := Compile(code, []types.BaseType{}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run(nil, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func BenchmarkExpr(b *testing.B) {
	code := "!$1"
	p, err := Compile(code, []types.BaseType{types.Int}, nil)
	if err != nil {
		panic(err)
	}
	input := []types.INullableVector{types.BuildValue(types.Int, 0, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10)}
	b.ReportAllocs()
	b.ResetTimer()
	var ret types.INullableVector
	for i := 0; i < b.N; i++ {
		ret, _ = p.Run(input, nil)

	}
	b.Log(types.ToString(ret))
}

func BenchmarkExprShort(b *testing.B) {
	code := "1 + $1:Float"
	p, err := Compile(code, []types.BaseType{types.Int}, nil)
	if err != nil {
		panic(err)
	}
	input := []types.INullableVector{types.BuildValue(types.Int, 1, 2, 3)}
	b.ReportAllocs()
	b.ResetTimer()
	var ret types.INullableVector
	for i := 0; i < b.N; i++ {
		ret, _ = p.Run(input, nil)

	}
	b.Log(types.ToString(ret))
}

func TestExprSIMD(t *testing.T) {
	code := "$1 + $1"
	p, err := Compile(code, []types.BaseType{types.Int}, nil)
	if err != nil {
		panic(err)
	}
	input := []types.INullableVector{types.BuildValue(types.Int, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10, 1, 2, 3, 4, 5, 6, nil, 8, 9, 10)}
	for i := 0; i < 1000000; i++ {
		_, _ = p.Run(input, nil)
	}

}

func TestExprTRUE(t *testing.T) {
	code := `true`
	p, err := Compile(code, []types.BaseType{}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run(nil, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func BenchmarkParse(b *testing.B) {
	code := `(1+1):Text+"text"`
	for i := 0; i < b.N; i++ {
		_, _ = Compile(code, []types.BaseType{types.Int}, nil)
	}
}

func TestExprRegexp(t *testing.T) {
	code := "$1|similar(`[a-zA-Z]+`)"
	p, err := Compile(code, []types.BaseType{types.Text}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run([]types.INullableVector{types.BuildValue(types.Text, "nil", "1", ".", "text")}, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func Test2Time(t *testing.T) {
	code := "toTime(toTimestamp($1))"
	p, err := Compile(code, []types.BaseType{types.Text}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run([]types.INullableVector{types.BuildValue(types.Text, time.Now().Format(time.RFC3339))}, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}

func TestRegexpReplace(t *testing.T) {
	code := `regexpReplace($1,"\\d","")`
	p, err := Compile(code, []types.BaseType{types.Text}, nil)
	if err != nil {
		panic(err)
	}
	ret, err := p.Run([]types.INullableVector{types.BuildValue(types.Text, "2016-11-10 09:41:51+08")}, nil)
	if err != nil {
		panic(err)
	}
	t.Log(types.ToString(ret))
}
