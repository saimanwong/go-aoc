package day03

import (
	"fmt"
	"strconv"
	"strings"
)

type Problem struct {
	input []string
}

func (p *Problem) Run() {
	slopes := []slope{
		{
			right: 3,
			down:  1,
		},
		{
			right: 1,
			down:  1,
		},
		{
			right: 5,
			down:  1,
		},
		{
			right: 7,
			down:  1,
		},
		{
			right: 1,
			down:  2,
		},
	}

	max_right := 0
	for _, s := range slopes {
		if s.right > max_right {
			max_right = s.right
		}
	}

	inputParsed := parseInput(p.input, max_right)

	fmt.Println("Answer 1: ", problemOne(inputParsed, slopes[0]))
	fmt.Println("Answer 2: ", problemTwo(inputParsed, slopes))
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

type slope struct {
	right int
	down  int
}

func parseInput(input []string, max_right int) []string {
	var ret []string
	for _, s := range input {
		ret = append(ret, strings.Repeat(s, len(input)*max_right))
	}
	return ret
}

func problemOne(input []string, s slope) string {
	ret := 0
	down, right := 0, 0
	for down < len(input) {
		if input[down][right] == '#' {
			ret++
		}

		down += s.down
		right += s.right
	}
	return strconv.Itoa(ret)
}

func problemTwo(input []string, s []slope) string {
	ret := 1
	for _, val := range s {
		n, _ := strconv.Atoi(problemOne(input, val))
		ret *= n
	}
	return strconv.Itoa(ret)
}
