package day05

import (
	"fmt"
	"sort"
	"strconv"
)

func findSeat(seat string, min int, max int) int {
	for _, c := range seat {
		if c == 'F' || c == 'L' {
			max = (max-min)/2 + min
		}
		if c == 'B' || c == 'R' {
			min = (max-min)/2 + 1 + min
		}
	}
	return min
}

func getId(input string) int {
	r, c := input[:7], input[7:]
	row, column := findSeat(r, 0, 127), findSeat(c, 0, 7)
	return row*8 + column
}

func problemOne(input []string) string {
	ret := 0
	for _, s := range input {
		id := getId(s)
		if id > ret {
			ret = id
		}
	}
	return strconv.Itoa(ret)
}

func problemTwo(input []string) string {
	var ids []int
	for _, i := range input {
		ids = append(ids, getId(i))
	}
	sort.Ints(ids)
	for i := 0; i < len(ids)-1; i++ {
		curr := ids[i]
		next := ids[i+1]
		if next-curr > 1 {
			return strconv.Itoa(curr + 1)
		}
	}

	return "something went wrong..."
}

func Run(input []string) {
	fmt.Println("Answer 1: " + problemOne(input))
	fmt.Println("Answer 2: " + problemTwo(input))
}
