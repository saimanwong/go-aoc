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

func NewByteMatrix(height, width int, fill byte) ByteMatrix {
	ret := make([][]byte, height)
	for r := 0; r < height; r++ {
		ret[r] = make([]byte, width)
		for c := 0; c < width; c++ {
			ret[r][c] = fill
		}
	}
	return ret
}

func (b ByteMatrix) Equal(bb ByteMatrix) bool {
	ret := true
	b.Loop(func(r, c int, curr rune) {
		if curr != rune(bb[r][c]) {
			ret = false
			return
		}
	})
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

func (b ByteMatrix) InsideCoord(coord *Coord) bool {
	return coord.R >= 0 && coord.R < len(b) && coord.C >= 0 && coord.C < len(b[0])
}

func (b ByteMatrix) Width() int {
	return len(b[0])
}

func (b ByteMatrix) Height() int {
	return len(b)
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

func (b ByteMatrix) FindMany(n byte) []*Coord {
	ret := []*Coord{}
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[r]); c++ {
			if b[r][c] != n {
				continue
			}
			ret = append(ret, &Coord{R: r, C: c})
		}
	}
	return ret
}

func (b ByteMatrix) GetVal(c *Coord) byte {
	return b[c.R][c.C]
}

func (b ByteMatrix) SetVal(c *Coord, val byte) {
	b[c.R][c.C] = val
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

func (b ByteMatrix) LoopCoord(
	fn func(c *Coord, curr rune),
) {
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[r]); c++ {
			coord := &Coord{R: r, C: c}
			fn(coord, rune(b[r][c]))
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

func (c *Coord) Copy() *Coord {
	return &Coord{R: c.R, C: c.C}
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

var DirectionSlice []Coord = []Coord{
	{
		R: -1,
		C: 0,
	},
	{
		R: 0,
		C: 1,
	},
	{
		R: 1,
		C: 0,
	},
	{
		R: 0,
		C: -1,
	},
}
