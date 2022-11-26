package day17

import (
	"github.com/saimanwong/go-aoc/problem/year2020/day17/p1"
	"github.com/saimanwong/go-aoc/problem/year2020/day17/p2"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	p1.Run(p.input)
	p2.Run(p.input)
}
