package functions

import (
	"expr/types"
	"time"
)

func init() {
	now, _ := NewFunction("now")
	now.Overload([]types.BaseType{}, types.Timestamp, func(vectors []types.INullableVector) (types.INullableVector, error) {
		out := &types.NullableTimestamp{
			TsType: types.Timestamp,
		}
		out.Init(1)
		out.SetScala(true)
		out.Set(0, time.Now().UnixNano(), false)
		return out, nil
	})
}
