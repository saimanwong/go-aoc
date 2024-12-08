package day08

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

type Antennas map[rune][]*toolbox.Coord

func (p *Problem) Run() {
	matrix := toolbox.ToByteMatrix(p.input)
	antennas := Antennas{}
	matrix.Loop(func(r, c int, curr rune) {
		if (curr >= '0' && curr <= '9') ||
			(curr >= 'a' && curr <= 'z') ||
			(curr >= 'A' && curr <= 'Z') {
			if _, ok := antennas[curr]; !ok {
				antennas[curr] = []*toolbox.Coord{}
			}
			antennas[curr] = append(antennas[curr], &toolbox.Coord{
				R: r,
				C: c,
			})
		}
	})

	fmt.Println("Part 1:", p1(matrix.Copy(), antennas))
	fmt.Println("Part 2:", p2(matrix.Copy(), antennas))
}

func p1(matrix toolbox.ByteMatrix, antennas Antennas) int {
	antiNodes := []*toolbox.Coord{}
	for _, v := range antennas {
		for i := 0; i < len(v)-1; i++ {
			curr := v[i]
			for j := i + 1; j < len(v); j++ {
				next := v[j]
				diff := &toolbox.Coord{R: curr.R - next.R, C: curr.C - next.C}
				n1, n2 := *curr, *next
				n1.Add(diff)
				n2.Subtract(diff)
				if matrix.Inside(n1.R, n1.C) {
					antiNodes = append(antiNodes, &n1)
				}
				if matrix.Inside(n2.R, n2.C) {
					antiNodes = append(antiNodes, &n2)
				}
			}
		}
	}
	for _, n := range antiNodes {
		matrix[n.R][n.C] = '#'
	}
	return matrix.Count('#')
}

func p2(matrix toolbox.ByteMatrix, antennas Antennas) int {
	antiNodes := []toolbox.Coord{}
	for _, v := range antennas {
		for i := 0; i < len(v)-1; i++ {
			curr := v[i]
			for j := i + 1; j < len(v); j++ {
				next := v[j]
				diff := &toolbox.Coord{R: curr.R - next.R, C: curr.C - next.C}
				n1, n2 := *curr, *next
				n1.Add(diff)
				n2.Subtract(diff)
				for matrix.Inside(n1.R, n1.C) {
					antiNodes = append(antiNodes, n1)
					n1.Add(diff)
				}
				for matrix.Inside(n2.R, n2.C) {
					antiNodes = append(antiNodes, n2)
					n2.Subtract(diff)
				}
			}
		}
	}
	for _, n := range antiNodes {
		matrix[n.R][n.C] = '#'
	}
	return matrix.CountExcept('.')
}
