package day08

import (
	"fmt"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	m [][]int
}

func (p *Problem) SetInput(input []string) {
	p.m = make([][]int, 0, len(input))
	for _, r := range input {
		m := make([]int, 0, len(input[0]))
		for _, c := range r {
			m = append(m, toolbox.ToInt(string(c)))
		}
		p.m = append(p.m, m)
	}
}

func (p *Problem) Run() {
	p1 := 0
	for r := range p.m {
		for c := range p.m {
			visible := p.visible1(r, c)
			if visible {
				p1++
			}
		}
	}
	fmt.Println("Part 1:", p1)

	p2 := 0
	for r := range p.m {
		for c := range p.m {
			score := 1
			for _, m := range []string{"up", "left", "down", "right"} {
				n, _ := p.check(m, r, c)
				score *= n
			}
			if p2 < score {
				p2 = score
			}
		}
	}
	fmt.Println("Part 2:", p2)
}

func (p *Problem) visible1(r, c int) bool {
	if r == 0 || r == len(p.m)-1 || c == 0 || c == len(p.m[0])-1 {
		return true
	}
	for _, m := range []string{"up", "right", "down", "left"} {
		if _, ok := p.check(m, r, c); ok {
			return true
		}
	}
	return false
}

var direction map[string]struct {
	r int
	c int
} = map[string]struct {
	r int
	c int
}{
	"up": {
		r: -1,
		c: 0,
	},
	"right": {
		r: 0,
		c: 1,
	},
	"down": {
		r: 1,
		c: 0,
	},
	"left": {
		r: 0,
		c: -1,
	},
}

func (p *Problem) check(dir string, r, c int) (int, bool) {
	curr := p.m[r][c]
	r, c = r+direction[dir].r, c+direction[dir].c
	count := 0
	for r >= 0 && c >= 0 && r < len(p.m) && c < len(p.m) {
		count++
		if p.m[r][c] >= curr {
			return count, false
		}
		r, c = r+direction[dir].r, c+direction[dir].c
	}
	return count, true
}
