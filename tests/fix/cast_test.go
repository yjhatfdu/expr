package fix

import (
	"github.com/yjhatfdu/expr"
	"github.com/yjhatfdu/expr/types"
	"testing"
)

func TestTimestamp2TimeWithoutTimezone(t *testing.T) {
	code := `toTime(toTimestamp($1,"yyyy-MM-ddTHH:mm:ss+hh:mm"))`
	p, err := expr.Compile(code, []types.BaseType{types.Text}, nil)
	if err != nil {
		t.Error(err)
		return
	}
	ret, err := p.Run([]types.INullableVector{types.BuildValue(types.Text, "2020-01-01T12:00:00+08:00")}, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if ret.Index(0) != 43200 {
		t.Error("2020-01-01T12:00:00+08:00 can not transform to 20:00.")
	}
}
