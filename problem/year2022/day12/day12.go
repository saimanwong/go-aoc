package day12

import (
	"fmt"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	Mat   [][]rune
	Start *toolbox.Coord
	End   *toolbox.Coord
}

const (
	Start = 'a' - 1
	End   = 'z' + 1
)

func (p *Problem) SetInput(input []string) {
	p.Mat = make([][]rune, len(input))
	for r := 0; r < len(input); r++ {
		p.Mat[r] = []rune{}
		for c, letter := range input[r] {
			tmp := letter
			if letter == 'S' {
				p.Start = &toolbox.Coord{R: r, C: c}
				tmp = 'a' - 1
			}
			if letter == 'E' {
				p.End = &toolbox.Coord{R: r, C: c}
				tmp = 'z' + 1
			}
			p.Mat[r] = append(p.Mat[r], tmp)
		}
	}
}

func (p *Problem) Run() {
	fmt.Println("Part 1:", p.search(p.Start)-1)

	a := []*toolbox.Coord{}
	for r := 0; r < len(p.Mat); r++ {
		for c := 0; c < len(p.Mat[r]); c++ {
			if p.Mat[r][c] == 'a' {
				a = append(a, &toolbox.Coord{R: r, C: c})
			}
		}
	}
	min := int(^uint(0) >> 1)
	for _, start := range a {
		steps := p.search(start)
		if steps == 1 {
			continue
		}
		if min > steps {
			min = steps
		}
	}
	fmt.Println("Part 2:", min-1)
}

func (p *Problem) search(start *toolbox.Coord) int {
	visited := map[string]struct{}{}
	q := []*toolbox.Coord{start}
	prev := make([][]*toolbox.Coord, len(p.Mat))
	for r := 0; r < len(p.Mat); r++ {
		prev[r] = make([]*toolbox.Coord, len(p.Mat[r]))
	}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, adj := range p.neighs(curr) {
			if _, ok := visited[adj.String()]; ok {
				continue
			}
			q = append(q, adj)
			prev[adj.R][adj.C] = curr
			visited[adj.String()] = struct{}{}
		}
	}
	sum := 0
	prev[start.R][start.C] = nil
	curr := p.End
	for curr != nil {
		sum++
		curr = prev[curr.R][curr.C]
	}
	return sum
}

func (p *Problem) neighs(currCord *toolbox.Coord) []*toolbox.Coord {
	neighs := []*toolbox.Coord{}
	for _, d := range []rune{'U', 'L', 'D', 'R'} {
		r, c := currCord.R+toolbox.Direction[d].R, currCord.C+toolbox.Direction[d].C
		if r < 0 || c < 0 || r >= len(p.Mat) || c >= len(p.Mat[0]) { // oob
			continue
		}
		if p.Mat[r][c] > p.Mat[currCord.R][currCord.C]+1 { // invalid
			continue
		}
		neighs = append(neighs, &toolbox.Coord{R: r, C: c})
	}
	return neighs
}

func (p *Problem) Debug() {
	fmt.Println("start:", p.Start)
	fmt.Println("end:", p.End)
	for _, r := range p.Mat {
		for _, c := range r {
			fmt.Printf("%s ", string(c))
		}
		fmt.Println()
	}
}
