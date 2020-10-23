package tests

import (
	"github.com/yjhatfdu/expr"
	"github.com/yjhatfdu/expr/functions"
	"github.com/yjhatfdu/expr/types"
	"os"
	"path/filepath"
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
	_, err = functions.NewFunction("add")
	if err == nil {
		panic("should error")
	}
	f, _ := functions.NewFunction("test")
	f.Generic(func(types []types.BaseType) (types.BaseType, error) {
		return 0, nil
	}, func(vectors []types.INullableVector) (types.INullableVector, error) {
		return nil, nil
	})
	f.Print()
	p, err := expr.Compile("now", nil)
	if err != nil {
		panic(err)
	}
	_, err = p.Run(nil)
	if err != nil {
		panic(err)
	}
	err = filepath.Walk("./case", func(p string, info os.FileInfo, err error) error {
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
	p, err := expr.Compile("false", nil)
	if err != nil {
		panic(err)
	}
	t.Log(p.Run(nil))
}

