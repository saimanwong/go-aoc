package day13

import (
	"fmt"
	"math"
	"regexp"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

type ClawMachine struct {
	A     Location
	B     Location
	Prize Location
}

type Location struct {
	X float64
	Y float64
}

func (p *Problem) Run() {
	machine := &ClawMachine{}
	machines := []*ClawMachine{}
	for i, line := range p.input {
		matches := regexp.MustCompile(`X[=\+](\d+), Y[=\+](\d+)`).FindAllStringSubmatch(line, -1)
		switch i % 4 {
		case 0:
			machine.A = Location{X: toolbox.ToFloat64(matches[0][1]), Y: toolbox.ToFloat64(matches[0][2])}
		case 1:
			machine.B = Location{X: toolbox.ToFloat64(matches[0][1]), Y: toolbox.ToFloat64(matches[0][2])}
		case 2:
			machine.Prize = Location{X: toolbox.ToFloat64(matches[0][1]), Y: toolbox.ToFloat64(matches[0][2])}
		default:
			machines = append(machines, machine)
			machine = &ClawMachine{}
		}
	}
	machines = append(machines, machine)
	fmt.Println("Part 1:", p1(machines))
	fmt.Println("Part 2:", p2(machines))
}

func p1(machines []*ClawMachine) int {
	// cost A = 3 tokens
	// cost B = 1 tokens
	const maxPress = 200
	sum := 0.0
	for _, machine := range machines {
		maxToken := 0.0
		for a := float64(0); a <= maxPress; a++ {
			for b := float64(1); b <= maxPress-a; b++ {
				xb, yb := machine.B.X*b, machine.B.Y*b
				xa, ya := machine.A.X*a, machine.A.Y*a
				if xb+xa == machine.Prize.X && yb+ya == machine.Prize.Y {
					token := b*1 + a*3
					maxToken = max(maxToken, token)
				}
			}
		}
		sum += maxToken
	}
	return int(sum)
}

func p2(machines []*ClawMachine) int {
	sum := 0
	const add = 10000000000000

	for _, m := range machines {
		m.Prize.X += add
		m.Prize.Y += add
		b := (m.Prize.X*m.A.Y/m.A.X - m.Prize.Y) / (m.B.X*m.A.Y/m.A.X - m.B.Y)
		a := (m.Prize.X - b*m.B.X) / m.A.X
		aDiff, bDiff := math.Abs(a-float64(int(a))), math.Abs(b-float64(int(b)))
		if aDiff > 0.99 {
			a = math.Ceil(a)
		}
		if aDiff < 0.01 {
			a = math.Floor(a)
		}
		if bDiff > 0.99 {
			b = math.Ceil(b)
		}
		if bDiff < 0.01 {
			b = math.Floor(b)
		}
		if a == float64(int(a)) && b == float64(int(b)) {
			sum += int(1*b) + int(3*a)
		}
	}
	return int(sum)
}
