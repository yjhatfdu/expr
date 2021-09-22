package tests

import (
	"github.com/yjhatfdu/expr"
	"github.com/yjhatfdu/expr/functions"
	"github.com/yjhatfdu/expr/types"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCase(t *testing.T) {
	var err error
	types.LocalOffsetNano = 8 * 3600 * 1e9
	time.Local, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	//_, err = functions.NewFunction("add")
	//if err == nil {
	//	panic("should error")
	//}
	f, _ := functions.NewFunction("test")
	f.Generic(func(types []types.BaseType) (types.BaseType, error) {
		return 0, nil
	}, func(vectors []types.INullableVector, env map[string]string) (types.INullableVector, error) {
		return nil, nil
	})
	f.Print()
	p, err := expr.Compile("now", nil, nil)
	if err != nil {
		panic(err)
	}
	_, err = p.Run(nil, nil)
	if err != nil {
		panic(err)
	}
	err = filepath.Walk("./case", func(p string, info os.FileInfo, err error) error {
		if os.Getenv("NO_TIME_TEST") != "" {
			if !info.IsDir() {
				if strings.Contains(info.Name(), "time") {
				} else {
					loader(p)
				}
			}
			return nil
		}
		if !info.IsDir() {
			loader(p)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func TestPrintFunctions(t *testing.T) {
	t.Log(functions.PrintAllFunctions())
}

func TestTrue(t *testing.T) {
	p, err := expr.Compile("false", nil, nil)
	if err != nil {
		panic(err)
	}
	t.Log(p.Run(nil, nil))
}

func TestAutoCast(t *testing.T) {
	p, err := expr.Compile("1111 like \"11\"", nil, nil)
	if err != nil {
		panic(err)
	}
	t.Log(p.Run(nil, nil))

	p, err = expr.Compile("toNumeric($1) like \"12\"", []types.BaseType{types.Int}, nil)
	if err != nil {
		panic(err)
	}

	t.Log(p.Run([]types.INullableVector{types.BuildValue(types.Int, 123)}, nil))
}

// test *
func TestSimilarToCase1(t *testing.T) {
	p, err := expr.Compile("`"+`abbc`+"`"+` similar to `+"`"+`a?b*c+`+"`", nil, nil)
	if err != nil {
		panic(err)
	}

	v, _ := p.Run(nil, nil)
	if !v.Index(0).(bool) {
		panic(types.ToString(v))
	}
}

// test |
func TestSimilarToCase2(t *testing.T) {
	p, err := expr.Compile("`"+`abc`+"`"+` similar to `+"`"+`%(b|d)%`+"`", nil, nil)
	if err != nil {
		panic(err)
	}

	v, _ := p.Run(nil, nil)
	if !v.Index(0).(bool) {
		panic(types.ToString(v))
	}
}

func TestNotSimilarToCase1(t *testing.T) {
	p, err := expr.Compile("`"+`abc`+"`"+`not similar to `+"`"+`ab`+"`", nil, nil)
	if err != nil {
		panic(err)
	}

	v, _ := p.Run(nil, nil)
	if !v.Index(0).(bool) {
		panic(types.ToString(v))
	}
}
