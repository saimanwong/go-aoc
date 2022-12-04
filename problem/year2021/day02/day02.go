package day02

import (
	"fmt"
	"strings"

	tb "github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	// p1
	hz1, dp1 := 0, 0
	// p2
	dp2, aim2 := 0, 0
	for _, s := range p.input {
		spl := strings.Split(s, " ")
		n := tb.ToInt(spl[1])
		switch spl[0] {
		case "forward":
			hz1 += n
			dp2 += aim2 * n
		case "down":
			dp1 += n
			aim2 += n
		case "up":
			dp1 -= n
			aim2 -= n
		default:
			panic("no no no")
		}
	}
	fmt.Println("Part 1:", hz1*dp1)
	fmt.Println("Part 2:", hz1*dp2)
}
