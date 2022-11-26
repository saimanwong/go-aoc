package day06

import (
	"fmt"
	"strconv"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	fmt.Println("Answer 1: " + problemOne(p.input))
	fmt.Println("Answer 2: " + problemTwo(p.input))
}

func problemOne(input []string) string {
	input = append(input, "")
	ret := 0
	set := make(map[rune]bool)
	for _, line := range input {
		if line == "" {
			ret += len(set)
			set = map[rune]bool{}
		}
		for _, c := range line {
			set[c] = true
		}
	}

	return strconv.Itoa(ret)
}

func problemTwo(input []string) string {
	input = append(input, "")
	ret, persons := 0, 0
	set := make(map[rune]int)
	for _, line := range input {
		if line == "" {
			for _, yes := range set {
				if persons == yes {
					ret++
				}
			}
			set = map[rune]int{}
			persons = 0
			continue
		}
		for _, c := range line {
			_, ok := set[c]
			if !ok {
				set[c] = 0
			}
			set[c]++
		}
		persons++
	}

	return strconv.Itoa(ret)
}
