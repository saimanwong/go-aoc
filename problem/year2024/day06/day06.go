package day06

import (
	"fmt"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	matrix := toolbox.ToByteMatrix(p.input)

	fmt.Println("Part 1:", p1(matrix))
	fmt.Println("Part 2:", p2(p.input, matrix))
}

func p1(m toolbox.ByteMatrix) int {
	walk(m.Find('^'), m, nil)
	return m.Count('X')
}

func p2(input []string, afterMap toolbox.ByteMatrix) int {
	paths := []toolbox.Coord{}
	for r, rr := range afterMap {
		for c, cc := range rr {
			if cc != 'X' {
				continue
			}
			paths = append(paths, toolbox.Coord{R: r, C: c})
		}
	}

	originalMap := toolbox.ToByteMatrix(input)
	start := originalMap.Find('^')
	count := 0
	for _, p := range paths {
		if p.R == start.R && p.C == start.C {
			continue
		}
		tmp := toolbox.ToByteMatrix(input)
		tmp[p.R][p.C] = 'O'
		tmpStart := &toolbox.Coord{R: start.R, C: start.C}
		if walk(tmpStart, tmp, &p) {
			count++
		}
	}
	return count
}

func walk(
	start *toolbox.Coord,
	m toolbox.ByteMatrix,
	obstacle *toolbox.Coord,
) bool {
	rotate := map[byte]byte{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}
	visited := map[string]int{}
	for m.Inside(start.R, start.C) {
		prevR, prevC := start.R, start.C
		var dir byte
		switch m[start.R][start.C] {
		case '^':
			start.Move('U')
			dir = '^'
		case '>':
			start.Move('R')
			dir = '>'
		case 'v':
			start.Move('D')
			dir = 'v'
		case '<':
			start.Move('L')
			dir = '<'
		}
		if m.Inside(start.R, start.C) && (m[start.R][start.C] == '#' || m[start.R][start.C] == 'O') {
			tmp := "%d,%d,%s"
			if _, ok := visited[fmt.Sprintf(tmp, start.R, start.C, string(dir))]; !ok {
				visited[fmt.Sprintf(tmp, start.R, start.C, string(dir))] = 0
			}
			visited[fmt.Sprintf(tmp, start.R, start.C, string(dir))]++
			if visited[fmt.Sprintf(tmp, start.R, start.C, string(dir))] > 1 {
				return true // loop identified
			}
			m[prevR][prevC] = rotate[m[prevR][prevC]]
			start.R = prevR
			start.C = prevC
			continue
		}
		if m.Inside(start.R, start.C) {
			m[start.R][start.C] = dir
		}
		m[prevR][prevC] = 'X'
	}
	return false
}
