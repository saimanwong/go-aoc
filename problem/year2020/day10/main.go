package day10

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
	inputParsed := parseInput(p.input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}

func parseInput(input []string) []int {
	ret := []int{}
	for _, v := range input {
		toInt, _ := strconv.Atoi(v)
		ret = append(ret, toInt)
	}
	sort.Ints(ret)
	return ret
}

func problemOne(input []int) string {
	input = append([]int{0}, input...)
	input = append(input, input[len(input)-1]+3)
	count := map[int]int{
		1: 0,
		3: 0,
	}

	prev := 0
	for _, v := range input {
		count[v-prev]++
		prev = v
	}
	return strconv.Itoa(count[1] * count[3])
}

func problemTwo(input []int) string {
	rem := map[int]int{0: 1}

	for _, n := range input {
		rem[n] = 0

		_, ok := rem[n-1]
		if ok {
			rem[n] += rem[n-1]
		}

		_, ok = rem[n-2]
		if ok {
			rem[n] += rem[n-2]
		}

		_, ok = rem[n-3]
		if ok {
			rem[n] += rem[n-3]
		}
	}

	ret := 0
	for _, n := range rem {
		if n > ret {
			ret = n
		}
	}

	return fmt.Sprintf("%d", ret)
}
