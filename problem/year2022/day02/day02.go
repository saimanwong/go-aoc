package day02

import (
	"fmt"
	"strings"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	sum1 := 0
	sum2 := 0
	for _, line := range p.input {
		spl := strings.Split(line, " ")
		opp, me := spl[0], spl[1]
		sum1 += shapeNum(me) + outcomeNum(me, opp)
		sum2 += problem2[me][opp]
	}
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}

func shapeNum(m string) int {
	switch m {
	case "A", "X": // rock
		return 1
	case "B", "Y": // paper
		return 2
	case "C", "Z": // scissor
		return 3
	}
	panic("not supported")
}

func toMove2(m string) int {
	switch m {
	case "X": // lose
		return 0
	case "Y": // draw
		return 3
	case "Z": // win
		return 6
	}
	panic("not supported")
}

func outcomeNum(me, opp string) int {
	m, o := shapeNum(me), shapeNum(opp)
	if m == o {
		return 3
	}
	if (m == 1 && o == 3) ||
		(m == 2 && o == 1) ||
		(m == 3 && o == 2) {
		return 6
	}
	return 0
}

var problem2 = map[string]map[string]int{ // mat[expected][opponent] = movement + expected
	"X": { // lose
		"A": shapeNum("C"),
		"B": shapeNum("A"),
		"C": shapeNum("B"),
	},
	"Y": { // draw
		"A": shapeNum("A") + 3,
		"B": shapeNum("B") + 3,
		"C": shapeNum("C") + 3,
	},
	"Z": { // win
		"A": shapeNum("B") + 6,
		"B": shapeNum("C") + 6,
		"C": shapeNum("A") + 6,
	},
}
