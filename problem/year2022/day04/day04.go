package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

type pair struct {
	a minMax
	b minMax
}

type minMax struct {
	min int
	max int
}

func (p *Problem) Run() {
	pairs := []pair{}
	for _, line := range p.input {
		pairs = append(pairs, parse(line))
	}
	fmt.Println("Part 1:", solve(true, pairs...))
	fmt.Println("Part 2:", solve(false, pairs...))
}

func solve(fully bool, pairs ...pair) int {
	n := 0
	for _, p := range pairs {
		if contained(p, fully) {
			n++
		}
	}
	return n
}

func contained(p pair, fully bool) bool {
	m := map[int]struct{}{}
	maxN := p.a
	minN := p.b
	if (p.b.max - p.b.min) > (p.a.max - p.a.min) {
		maxN = p.b
		minN = p.a
	}
	for i := maxN.min; i <= maxN.max; i++ {
		m[i] = struct{}{}
	}
	for i := minN.min; i <= minN.max; i++ {
		_, ok := m[i]
		// part1
		if !ok && fully {
			return false
		}
		// part2
		if ok && !fully {
			return true
		}
	}
	// part1
	if fully {
		return true
	}
	// part2
	return false
}

func parse(line string) pair {
	spl := strings.Split(line, ",")
	a, b := strings.Split(spl[0], "-"), strings.Split(spl[1], "-")
	return pair{
		a: minMax{
			min: parseInt(a[0]),
			max: parseInt(a[1]),
		},
		b: minMax{
			min: parseInt(b[0]),
			max: parseInt(b[1]),
		},
	}
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
