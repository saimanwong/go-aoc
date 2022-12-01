package day01

import (
	"fmt"
	"sort"
	"strconv"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	elves := []int{}
	tot := 0
	for _, line := range p.input {
		if line == "" {
			elves = append(elves, tot)
			tot = 0
			continue
		}
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		tot += n
	}
	sort.Ints(elves)
	fmt.Println("Part 1:", elves[len(elves)-1])
	fmt.Println("Part 2:", elves[len(elves)-1]+elves[len(elves)-2]+elves[len(elves)-3])
}
