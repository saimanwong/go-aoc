// toolbox helper package, most of the funcs panics by design.
package toolbox

import (
	"fmt"
	"math"
	"strconv"
)

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

type Coord struct {
	R int
	C int
}

func (c *Coord) Move(m ...rune) {
	for _, d := range m {
		c.R += Direction[d].R
		c.C += Direction[d].C
	}
}

func (c *Coord) Distance(y *Coord) float64 {
	return math.Sqrt(
		math.Pow(float64(c.R-y.R), 2.0) + math.Pow(float64(c.C-y.C), 2.0),
	)
}

func (c *Coord) String() string {
	return fmt.Sprintf("%d,%d", c.R, c.C)
}

var Direction map[rune]Coord = map[rune]Coord{
	'U': {
		R: -1,
		C: 0,
	},
	'R': {
		R: 0,
		C: 1,
	},
	'D': {
		R: 1,
		C: 0,
	},
	'L': {
		R: 0,
		C: -1,
	},
}
