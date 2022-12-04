// toolbox helper package, most of the funcs panics by design.
package toolbox

import "strconv"

func ToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func ToIntSlice(s ...string) []int {
	ns := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		ns[i] = ToInt(s[i])
	}
	return ns
}

func ToFloat64(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func ToFloat64Slice(s ...string) []float64 {
	ns := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		ns[i] = ToFloat64(s[i])
	}
	return ns
}
