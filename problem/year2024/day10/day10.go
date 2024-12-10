package day10

import (
	"fmt"
	"strconv"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	m := toolbox.ToByteMatrix(p.input)
	fmt.Println("Part 1:", solve(m, false))
	fmt.Println("Part 2:", solve(m, true))
}

func solve(m toolbox.ByteMatrix, rating bool) int {
	q := m.FindMany('0')
	count := 0
	counts := []int{}
	visited := map[string]interface{}{}
	for len(q) > 0 {
		curr := q[len(q)-1]
		q = q[:len(q)-1]
		if m.GetVal(curr) == '9' {
			str := fmt.Sprintf("%s", curr)
			if _, ok := visited[str]; !rating && !ok {
				visited[str] = nil
				count++
				continue
			}
			if rating {
				count++
			}
			continue
		}
		if m.GetVal(curr) == '0' {
			if count > 0 {
				counts = append(counts, count)
			}
			visited = map[string]interface{}{}
			count = 0
		}
		for _, dir := range toolbox.Direction {
			neigh := &toolbox.Coord{R: curr.R + dir.R, C: curr.C + dir.C}
			if !m.Inside(neigh.R, neigh.C) {
				continue
			}
			if _, err := strconv.Atoi(string(m.GetVal(neigh))); err != nil {
				continue
			}
			if toolbox.ToInt(string(m.GetVal(neigh))) == toolbox.ToInt(string(m.GetVal(curr)))+1 {
				q = append(q, neigh)
			}
		}
	}
	sum := count
	for _, n := range counts {
		sum += n
	}
	return sum
}
