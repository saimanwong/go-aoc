package day04

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
	fmt.Println("Part 1:", p1(p.input))
	fmt.Println("Part 2:", p2(p.input))
}

func p1(lines []string) int {
	matrix := toolbox.ToByteMatrix(lines)
	count := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] != 'X' {
				continue
			}
			for _, move := range [][]toolbox.Coord{
				// up
				{{R: -1, C: 0}, {R: -2, C: 0}, {R: -3, C: 0}},
				// up right
				{{R: -1, C: 1}, {R: -2, C: 2}, {R: -3, C: 3}},
				// right
				{{R: 0, C: 1}, {R: 0, C: 2}, {R: 0, C: 3}},
				// down right
				{{R: 1, C: 1}, {R: 2, C: 2}, {R: 3, C: 3}},
				// down
				{{R: 1, C: 0}, {R: 2, C: 0}, {R: 3, C: 0}},
				// down left
				{{R: 1, C: -1}, {R: 2, C: -2}, {R: 3, C: -3}},
				// left
				{{R: 0, C: -1}, {R: 0, C: -2}, {R: 0, C: -3}},
				// up left
				{{R: -1, C: -1}, {R: -2, C: -2}, {R: -3, C: -3}},
			} {
				if !matrix.Inside(r+move[len(move)-1].R, c+move[len(move)-1].C) {
					continue
				}
				if matrix[r+move[0].R][c+move[0].C] != 'M' ||
					matrix[r+move[1].R][c+move[1].C] != 'A' ||
					matrix[r+move[2].R][c+move[2].C] != 'S' {
					continue
				}
				count += 1
			}
		}
	}
	return count
}

func p2(lines []string) int {
	matrix := toolbox.ToByteMatrix(lines)
	count := 0
	x := []toolbox.Coord{
		{R: -1, C: 1},  // up right
		{R: 1, C: -1},  // down left
		{R: -1, C: -1}, // up left
		{R: 1, C: 1},   // down right
	}
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] != 'A' {
				continue
			}
			ok := true
			for _, m := range x {
				if !matrix.Inside(r+m.R, c+m.C) {
					ok = false
					break
				}
				if !(matrix[r+m.R][c+m.C] == 'M' || matrix[r+m.R][c+m.C] == 'S') {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			if (matrix[r+x[0].R][c+x[0].C] == 'M' && matrix[r+x[1].R][c+x[1].C] != 'S') ||
				(matrix[r+x[0].R][c+x[0].C] == 'S' && matrix[r+x[1].R][c+x[1].C] != 'M') {
				continue
			}
			if (matrix[r+x[2].R][c+x[2].C] == 'M' && matrix[r+x[3].R][c+x[3].C] != 'S') ||
				(matrix[r+x[2].R][c+x[2].C] == 'S' && matrix[r+x[3].R][c+x[3].C] != 'M') {
				continue
			}
			count += 1
		}
	}
	return count
}
