package day20

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	inputParsed := parseInput(p.input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}

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

type tileInterface interface {
	debug()
	rotate()
	flip()
	compare(*tile, string) bool
	totHashtag() int
}

type imageInterface interface {
	debug()
	fill(int, int, []tile, map[int]bool, *bool)
	toTile() *tile
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

func checkAdj(t1 *tile, t2 *tile) (*pair, bool) {
	dir := "R"
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

			_, ok = checkAdj(t1, t2)
			if ok {
				memo[t1.id] = append(memo[t1.id], t2.id)
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

func (t *tile) copyTile() tile {
	curr := tile{}
	curr.id = t.id
	curr.mtx = make([][]rune, len(t.mtx))
	for i, _ := range curr.mtx {
		curr.mtx[i] = make([]rune, len(t.mtx[i]))
		copy(curr.mtx[i], t.mtx[i])
	}

	return curr
}

func genAll(tiles map[int]*tile) []tile {
	ret := []tile{}

	for _, t := range tiles {
		for i := 0; i < 2; i++ {
			for j := 0; j < 4; j++ {
				curr := t.copyTile()
				ret = append(ret, curr)
				t.rotate()
			}
			t.flip()
		}

	}

	return ret
}

func (img *image) debug() {
	dim := len(*img)
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			if (*img)[r][c] != nil {
				fmt.Printf("%d ", (*img)[r][c].id)
			} else {
				fmt.Printf("nil ")
			}
		}
		fmt.Println()
	}
}

func (img *image) fill(r, c int, tiles []tile, visited map[int]bool, found *bool) {
	// Stop when last
	if r == len((*img)) {
		*found = true
		return
	}

	for _, t := range tiles {
		if !visited[t.id] {
			// Check top
			if r > 0 {
				topT := (*img)[r-1][c]
				if topT == nil {
					panic("nil top...")
				}
				_, ok := t.compare(topT, "T")
				if !ok {
					continue
				}
			}

			// Check Left
			if c > 0 {
				leftT := (*img)[r][c-1]
				if leftT == nil {
					panic("nil left...")
				}
				_, ok := t.compare(leftT, "L")
				if !ok {
					continue
				}
			}

			// OK tile
			(*img)[r][c] = &t
			visited[t.id] = true

			// Go to next
			if c < len(*img)-1 {
				img.fill(r, c+1, tiles, visited, found)
			} else {
				img.fill(r+1, 0, tiles, visited, found)
			}

			if *found {
				break
			}

			visited[t.id] = false
		}
	}
}

func (img *image) toTile() *tile {
	ret := &tile{
		mtx: [][]rune{},
		id:  -1,
	}

	for r := 0; r < len(*img)*len((*img)[0][0].mtx)-(2*len(*img)); r++ {
		ret.mtx = append(ret.mtx, []rune{})

	}

	for imgR, imgRow := range *img {
		for _, t := range imgRow {
			for tR, tRow := range t.mtx {
				if tR == 0 || tR == len(t.mtx)-1 {
					continue
				}
				for tC, char := range tRow {
					if tC == 0 || tC == len(tRow)-1 {
						continue
					}
					rowNr := tR - 1 + (imgR * (len(tRow) - 2))
					ret.mtx[rowNr] = append(ret.mtx[rowNr], char)
				}
			}
		}
	}

	return ret
}

func (t *tile) findSeaMonsters(seaMonster []string) int {
	count := 0

	// Start row1
	for r := 0; r < len(t.mtx)-2; r++ {
		for c := 0; c < len(t.mtx[r]); c++ {
			// First line
			firstIndex := strings.Index(seaMonster[0], "#")
			if t.mtx[r][c] == '#' && c >= firstIndex {
				nextIndex := c - firstIndex
				if t.mtx[r+1][nextIndex] == '#' && nextIndex+len(seaMonster[1]) < len(t.mtx[r+1]) {
					line1 := true
					for idx1, char1 := range seaMonster[1] {
						if char1 == '#' && t.mtx[r+1][nextIndex+idx1] != '#' {
							line1 = false
							break
						}
					}

					line2 := true
					if line1 == true && nextIndex+strings.LastIndex(seaMonster[2], "#") < len(t.mtx[r+2]) {
						for idx2, char2 := range seaMonster[2] {
							if char2 == '#' && t.mtx[r+2][nextIndex+idx2] != '#' {
								line1 = false
								break
							}
						}
					}

					if line1 && line2 {
						count++
					}
				}
			}
		}
	}

	return count
}

func (t *tile) totHashtag() int {
	ret := 0
	for _, row := range t.mtx {
		for _, char := range row {
			if char == '#' {
				ret++
			}
		}
	}
	return ret
}

func problemTwo(input map[int]*tile) string {
	dim := int(math.Sqrt(float64(len(input))))

	// Init img
	img := &image{}
	for r := 0; r < dim; r++ {
		*img = append(*img, nil)
		for c := 0; c < dim; c++ {
			(*img)[r] = append((*img)[r], nil)
		}
	}

	// Generate all possibilities...
	tiles := genAll(input)
	visited := map[int]bool{}
	found := false
	for id, _ := range input {
		visited[id] = false
	}

	img.fill(0, 0, tiles, visited, &found)
	bigT := img.toTile()

	seaMonsterTot := -1
	seaMonster := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	for i := 0; i < 4; i++ {
		seaMonsters := bigT.findSeaMonsters(seaMonster)
		if seaMonsters > seaMonsterTot {
			seaMonsterTot = seaMonsters
		}
		bigT.rotate()
	}
	bigT.flip()
	for i := 0; i < 4; i++ {
		seaMonsters := bigT.findSeaMonsters(seaMonster)
		if seaMonsters > seaMonsterTot {
			seaMonsterTot = seaMonsters
		}
		bigT.rotate()
	}

	totHashtag := bigT.totHashtag()
	seaMonsterHashtag := 0
	for _, line := range seaMonster {
		for _, char := range line {
			if char == '#' {
				seaMonsterHashtag++
			}
		}
	}

	ret := totHashtag - seaMonsterHashtag*seaMonsterTot

	return fmt.Sprintf("%d", ret)
}
