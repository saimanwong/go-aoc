package day05

import (
	"fmt"
	"sort"
	"strings"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	c1, moves := parse(p.input)
	c2, _ := parse(p.input)
	fmt.Println("Part 1:", solve(c1, moves, false))
	fmt.Println("Part 2:", solve(c2, moves, true))
}

func solve(crates Crates, moves Moves, multiple bool) string {
	for _, m := range moves {
		// move stacks
		if multiple {
			tmp := make([]rune, m.Num)
			copy(tmp, crates[m.From][:m.Num])
			crates[m.From] = crates[m.From][m.Num:]
			crates[m.To] = append(tmp, crates[m.To]...)
			continue
		}
		// move one-by-one
		for i := 0; i < m.Num; i++ {
			tmp := crates[m.From][0]
			crates[m.From] = crates[m.From][1:]
			crates[m.To] = append([]rune{tmp}, crates[m.To]...)
		}
	}
	keys := make([]int, 0, len(crates))
	for k := range crates {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	res := ""
	for _, v := range keys {
		res += string(crates[v][0])
	}
	return res
}

type Crates map[int][]rune

type Move struct {
	Num  int
	From int
	To   int
}

type Moves []Move

func parse(lines []string) (Crates, Moves) {
	c := Crates{}
	movesIdx := 0
	for idx, l := range lines {
		if l == "" {
			movesIdx = idx + 1
			break
		}
		crateIdx := 1
		for i := 1; i < len(l); i += 4 {
			r := rune(l[i])
			if r < 'A' || r > 'Z' {
				crateIdx++
				continue
			}
			if _, ok := c[crateIdx]; !ok {
				c[crateIdx] = []rune{}
			}
			c[crateIdx] = append(c[crateIdx], r)
			crateIdx++
		}
	}

	m := Moves{}
	for i := movesIdx; i < len(lines); i++ {
		spl := strings.Split(lines[i], " ")
		m = append(m, Move{
			Num:  toolbox.ToInt(spl[1]),
			From: toolbox.ToInt(spl[3]),
			To:   toolbox.ToInt(spl[5]),
		})
	}
	return c, m
}
