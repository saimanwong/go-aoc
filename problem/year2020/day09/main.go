package day09

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
	fmt.Println("Answer 1: " + problemOne(inputParsed, 25))
	fmt.Println("Answer 2: " + problemTwo(inputParsed, 25))
	// fmt.Println("Answer 1: " + problemOne(inputParsed, 5))
	// fmt.Println("Answer 2: " + problemTwo(inputParsed, 5))
}

func parseInput(input []string) []int {
	ret := []int{}
	for _, i := range input {
		n, _ := strconv.Atoi(i)
		ret = append(ret, n)
	}

	return ret
}

func findGoal(input []int, start int, end int, curr int, prev []int, m map[int]bool) bool {
	for j := start; j >= end; j-- {
		prev := input[j]
		num := curr - prev
		_, ok := m[num]
		if ok {
			return false
		}

	}
	return true
}

func rollMap(curr int, prev []int, m map[int]bool) {
	del := prev[0]
	delete(m, del)
	m[curr] = true
	prev = prev[1:]
	prev = append(prev, curr)
}

func problemOne(input []int, preamble int) string {
	ret := -1

	m := map[int]bool{}
	prev := []int{}

	for i := 0; i < len(input); i++ {
		curr := input[i]
		if i < preamble {
			m[curr] = true
			prev = append(prev, curr)
			continue
		}

		if findGoal(input, i-1, i-preamble, curr, prev, m) {
			ret = curr
			break
		}

		rollMap(curr, prev, m)
	}

	return strconv.Itoa(ret)
}

func bruteForceTwo(input []int, preamble int, goal int) []int {
	curr := preamble
	ret := []int{}
	sum := 0
	for curr <= len(input) {
		for i := curr - 1; i >= curr-preamble; i-- {
			sum += input[i]
			ret = append(ret, input[i])
		}

		if sum == goal {
			return ret
		}

		curr++
		sum = 0
		ret = []int{}
	}

	return ret
}
func problemTwo(input []int, preamble int) string {
	goal, _ := strconv.Atoi(problemOne(input, preamble))
	preamble2 := 2
	ret := []int{}

	for preamble2 < 1000 {
		ret = bruteForceTwo(input, preamble2, goal)
		if len(ret) > 0 {
			break
		}
		preamble2++

	}

	sort.Ints(ret)
	return strconv.Itoa(ret[0] + ret[len(ret)-1])
}
