package day20

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var ADJMAP = map[int]map[int]string{
	-1: {
		0: "T",
	},
	0: {
		-1: "L",
		1:  "R",
	},
	1: {
		0: "B",
	},
}

type helper interface {
	debug()
	rotate()
	flip()
	compare(*tile, string) bool
}

type pair struct {
	r int
	c int
}

type tile struct {
	id  int
	mtx [][]rune
}

type image [][]*tile

func (t *tile) debug() {
	fmt.Println("Tile", t.id)
	for _, r := range t.mtx {
		for _, c := range r {
			fmt.Printf("%s ", string(c))
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (t *tile) rotate() {
	for r := 0; r < len(t.mtx)/2; r++ {
		for c := r; c < len(t.mtx[r])-r-1; c++ {
			// bl -> tl
			tmp := t.mtx[r][c]
			t.mtx[r][c] = t.mtx[len(t.mtx)-c-1][r]

			// tl -> tr
			tmp2 := t.mtx[c][len(t.mtx)-r-1]
			t.mtx[c][len(t.mtx)-r-1] = tmp

			// tr -> br
			tmp = t.mtx[len(t.mtx)-r-1][len(t.mtx)-c-1]
			t.mtx[len(t.mtx)-r-1][len(t.mtx)-c-1] = tmp2

			// br -> bl
			t.mtx[len(t.mtx)-c-1][r] = tmp
		}
	}
}

func (t *tile) flip() {
	for r := 0; r < len(t.mtx)/2; r++ {
		tmp := t.mtx[r]
		t.mtx[r] = t.mtx[len(t.mtx)-r-1]
		t.mtx[len(t.mtx)-r-1] = tmp
	}
}

func (t1 *tile) compare(t2 *tile, e string) (*pair, bool) {
	p := &pair{
		r: 0,
		c: 0,
	}
	// T -> B
	if e == "T" {
		for i := 0; i < len(t1.mtx); i++ {
			if t1.mtx[0][i] != t2.mtx[len(t2.mtx)-1][i] {
				return nil, false
			}
		}
		p.r = -1
		p.c = 0
	}

	// R -> L
	if e == "R" {
		for i := 0; i < len(t1.mtx); i++ {
			if t1.mtx[i][len(t1.mtx[0])-1] != t2.mtx[i][0] {
				return nil, false
			}
		}
		p.r = 0
		p.c = 1
	}

	// B -> T
	if e == "B" {
		for i := 0; i < len(t1.mtx); i++ {
			if t1.mtx[len(t1.mtx)-1][i] != t2.mtx[0][i] {
				return nil, false
			}
		}
		p.r = 1
		p.c = 0
	}

	// L -> R
	if e == "L" {
		for i := 0; i < len(t1.mtx); i++ {
			if t1.mtx[i][0] != t2.mtx[i][len(t2.mtx[0])-1] {
				return nil, false
			}
		}
		p.r = 0
		p.c = -1
	}

	return p, true
}

func parseInput(input []string) map[int]*tile {
	ret := map[int]*tile{}

	reTile := regexp.MustCompile(`Tile (\d+):`)
	t := &tile{mtx: [][]rune{}}
	for _, line := range input {
		if line == "" {
			ret[t.id] = t
			t = &tile{mtx: [][]rune{}}
			continue
		}

		if reTile.MatchString(line) {
			s := reTile.FindStringSubmatch(line)
			id, _ := strconv.Atoi(s[1])
			t.id = id
			continue
		}

		// Lines
		t.mtx = append(t.mtx, []rune{})
		for _, r := range line {
			t.mtx[len(t.mtx)-1] = append(t.mtx[len(t.mtx)-1], r)
		}
	}
	ret[t.id] = t
	return ret
}

func genAdjMap(dim int) map[int]map[int][]string {
	ret := map[int]map[int][]string{}

	// Map direction of each position
	for r := 0; r < dim; r++ {
		ret[r] = map[int][]string{}
		for c := 0; c < dim; c++ {
			ret[r][c] = []string{}
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					adjR := r + i
					adjC := c + j

					// Do not get same position
					if i == 0 && j == 0 {
						continue
					}

					// Only inside
					if adjR >= 0 && adjR < dim && adjC >= 0 && adjC < dim {
						ret[r][c] = append(ret[r][c], ADJMAP[i][j])
					}

				}
			}
		}
	}

	return ret
}

func checkAdj(t1 *tile, t2 *tile, dir string) (*pair, bool) {
	// T2
	// Rotate Rotate Rotate Rotate
	for i := 0; i < 4; i++ {
		t2.rotate()
		p, ok := t1.compare(t2, dir)
		if ok {
			return p, true
		}
	}
	// fmt.Println("t2 rotate 4 times")

	// Flip
	// Rotate Rotate Rotate Rotate
	t2.flip()
	p, ok := t1.compare(t2, dir)
	if ok {
		return p, true
	}
	// fmt.Println("t2 flipped")
	for i := 0; i < 4; i++ {
		t2.rotate()
		p, ok := t1.compare(t2, dir)
		if ok {
			return p, true
		}
	}
	// fmt.Println("t2 rotate 4 times")

	// T1
	// Rotate T2 Rotate T2 Rotate T2 Rotate T2
	// Flip T2
	// Rotate T2 Rotate T2 Rotate T2 Rotate T2
	for i := 0; i < 4; i++ {
		t1.rotate()
		for j := 0; j < 4; j++ {
			t2.rotate()
			p, ok := t1.compare(t2, dir)
			if ok {
				return p, true
			}
		}
	}
	// fmt.Println("t1 t2 rotate 8 times")
	t2.flip()
	p, ok = t1.compare(t2, dir)
	if ok {
		return p, true
	}
	// fmt.Println("t1 flipped")
	for i := 0; i < 4; i++ {
		t1.rotate()
		for j := 0; j < 4; j++ {
			t2.rotate()
			p, ok := t1.compare(t2, dir)
			if ok {
				return p, true
			}
		}
	}
	// fmt.Println("t1 rotate 4 times")

	return nil, false
}
func problemOne(input map[int]*tile) string {
	dim := int(math.Sqrt(float64(len(input))))
	adjMap := genAdjMap(dim) // For example, [0][0] = [R B BR]

	memo := map[int][]int{}
	for _, t1 := range input {
		for _, t2 := range input {
			if t1.id == t2.id {
				continue
			}

			_, ok := memo[t1.id]
			if !ok {
				memo[t1.id] = []int{}
			}

			for _, dir := range adjMap[dim/2][dim/2] {
				_, ok := checkAdj(t1, t2, dir)
				if ok {
					memo[t1.id] = append(memo[t1.id], t2.id)
				}
			}
		}
	}

	memoN := map[int][]int{}
	for id, oks := range memo {
		n := len(oks)
		_, ok := memoN[n]
		if !ok {
			memoN[n] = []int{}
		}
		memoN[n] = append(memoN[n], id)
	}

	min := 9999999
	for n, ids := range memoN {
		if len(ids) == 4 && n < min {
			min = n
		}
	}

	ret := 1
	for _, corner := range memoN[min] {
		ret *= corner
	}

	return fmt.Sprintf("%d", ret)
}

func problemTwo(input map[int]*tile) string {
	ret := "To be implemented..."
	return ret
}

func Run(input []string) {
	inputParsed := parseInput(input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}
