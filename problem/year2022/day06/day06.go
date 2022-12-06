package day06

import (
	"fmt"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	// p1
	for _, line := range p.input {
		a, b, c, d := 0, 1, 2, 3
		for d < len(line) {
			if uniq(
				rune(line[a]),
				rune(line[b]),
				rune(line[c]),
				rune(line[d]),
			) {
				fmt.Println("Part 1:", d+1)
				break
			}
			a++
			b++
			c++
			d++
		}
	}

	// p2
	for _, line := range p.input {
		i := 0
		for {
			if i+13 > len(line) {
				fmt.Println(i+13, len(line))
				break
			}
			if uniq2(line[i : i+14]) {
				fmt.Println("Part 2:", i+14)
				break
			}
			i++
		}
	}
}

func uniq(a, b, c, d rune) bool {
	if a == b || a == c || a == d || b == c || b == d || c == d {
		return false
	}
	return true
}

func uniq2(subline string) bool {
	for i := 0; i < len(subline); i++ {
		for j := i + 1; j < len(subline); j++ {
			if rune(subline[i]) == rune(subline[j]) {
				return false
			}
		}
	}
	return true
}
