package types

func int2bool(in []int64, out []bool) {
	for i := range in {
		out[i] = in[i] != 0
	}
}

func float2bool(in []float64, out []bool) {
	for i := range in {
		out[i] = in[i] != 0
	}
}
