package day15

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input []string) []int {
	ret := []int{}
	s := strings.Split(input[0], ",")
	for _, n := range s {
		toInt, _ := strconv.Atoi(n)
		ret = append(ret, toInt)
	}
	return ret
}

func problemOne(input []int, goal int) string {
	memo := map[int][]int{}
	for i, n := range input {
		_, ok := memo[n]
		if !ok {
			memo[n] = []int{}
		}
		memo[n] = append(memo[n], i+1)
	}

	for i := len(input); i < goal; i++ {
		last := input[i-1]
		curr := 0

		if len(memo[last]) > 1 {
			lastSpo := memo[last][len(memo[last])-1]
			lastLastSpo := memo[last][len(memo[last])-2]
			curr = lastSpo - lastLastSpo
		}

		input = append(input, curr)

		_, ok := memo[curr]
		if !ok {
			memo[curr] = []int{}
		}
		memo[curr] = append(memo[curr], i+1)
	}

	return fmt.Sprintf("%d", input[len(input)-1])
}

func Run(input []string) {
	inputParsed := parseInput(input)
	fmt.Println("Answer 1: " + problemOne(inputParsed, 2020))
	fmt.Println("Answer 2: " + problemOne(inputParsed, 30000000))
}
