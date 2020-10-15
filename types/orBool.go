package types

func orBool(data1, data2, out []bool) {
	for i := range out {
		out[i] = data1[i] || data2[i]
	}
}
