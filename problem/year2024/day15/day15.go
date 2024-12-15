package day15

import (
	"fmt"
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
	warehouseStr := []string{}
	moves := []rune{}
	for _, line := range p.input {
		if strings.HasPrefix(line, "#") {
			warehouseStr = append(warehouseStr, line)
			continue
		}
		if line == "" {
			continue
		}
		for _, r := range line {
			switch r {
			case '^':
				moves = append(moves, 'U')
			case '>':
				moves = append(moves, 'R')
			case 'v':
				moves = append(moves, 'D')
			case '<':
				moves = append(moves, 'L')
			default:
				panic("not a move")
			}
		}
	}
	warehouse := toolbox.ToByteMatrix(warehouseStr)
	fmt.Println("Part 1:", p1(warehouse, moves))
	fmt.Println("Part 2:")
}

func p1(warehouse toolbox.ByteMatrix, moves []rune) int {
	curr := warehouse.Find('@')
	next := curr.Copy()
	for _, move := range moves {
		next.Move(move)
		if warehouse.GetVal(next) == '#' {
			next = curr.Copy()
			continue
		}

		if warehouse.GetVal(next) == '.' {
			warehouse.SetVal(next, '@')
			warehouse.SetVal(curr, '.')
			curr = next
			next = curr.Copy()
			continue
		}

		prevs := []*toolbox.Coord{}
		hitWall := false
		for {
			prevs = append(prevs, next)
			next = next.Copy()
			next.Move(move)
			if warehouse.GetVal(next) == '.' {
				break
			}
			if warehouse.GetVal(next) == '#' {
				hitWall = true
				break
			}
		}

		for len(prevs) > 0 && !hitWall {
			prev := prevs[len(prevs)-1]
			prevs = prevs[:len(prevs)-1]
			warehouse.SetVal(next, warehouse.GetVal(prev))
			next = prev
		}
		if !hitWall {
			warehouse.SetVal(curr, '.')
			curr.Move(move)
			warehouse.SetVal(curr, '@')
		}
		next = curr.Copy()
	}
	sum := 0
	for _, b := range warehouse.FindMany('O') {
		sum += b.R*100 + b.C
	}
	return sum
}
