// toolbox helper package, most of the funcs panics by design.
package toolbox

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

type ByteMatrix [][]byte

func ToByteMatrix(lines []string) ByteMatrix {
	ret := make([][]byte, len(lines))
	for r := 0; r < len(lines); r++ {
		ret[r] = make([]byte, len(lines[r]))
		for c := 0; c < len(lines[r]); c++ {
			ret[r][c] = lines[r][c]
		}
	}
	return ret
}

func (b ByteMatrix) Copy() ByteMatrix {
	ret := make([][]byte, len(b))
	for r := 0; r < len(b); r++ {
		ret[r] = make([]byte, len(b[r]))
		for c := 0; c < len(b[r]); c++ {
			ret[r][c] = b[r][c]
		}
	}
	return ret
}

func (b ByteMatrix) Inside(r int, c int) bool {
	return r >= 0 && r < len(b) && c >= 0 && c < len(b[0])
}

func (b ByteMatrix) Count(n byte) int {
	ret := 0
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[r]); c++ {
			if b[r][c] != n {
				continue
			}
			ret++
		}
	}
	return ret
}

func (b ByteMatrix) CountExcept(n byte) int {
	ret := 0
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[r]); c++ {
			if b[r][c] == n {
				continue
			}
			ret++
		}
	}
	return ret
}

func (b ByteMatrix) Find(n byte) *Coord {
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[r]); c++ {
			if b[r][c] != n {
				continue
			}
			return &Coord{R: r, C: c}
		}
	}
	return nil
}

func (b ByteMatrix) String() string {
	builder := strings.Builder{}
	for _, r := range b {
		for _, c := range r {
			builder.WriteString(fmt.Sprintf("%s ", string(c)))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (b ByteMatrix) Loop(
	fn func(r int, c int, curr rune),
) {
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[r]); c++ {
			fn(r, c, rune(b[r][c]))
		}
	}
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

func (c *Coord) Add(y *Coord) {
	c.R = c.R + y.R
	c.C = c.C + y.C
}

func (c *Coord) Subtract(y *Coord) {
	c.R = c.R - y.R
	c.C = c.C - y.C
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
