package day01

import (
	"fmt"
	"os"
	"strconv"
)

type Problem struct {
	input []string
}

func (p *Problem) Run() {
	fmt.Println("Part 1: " + problemOne(p.input))
	fmt.Println("Part 2: " + problemTwo(p.input))
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func problemOne(input []string) string {
	m := make(map[int]bool)
	for _, i := range input {
		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		m[num] = true
	}
	for _, i := range input {
		curr, _ := strconv.Atoi(i)
		num := 2020 - curr
		_, ok := m[num]
		if !ok {
			continue
		}
		ans := strconv.Itoa(curr * num)
		return ans
	}
	return "something went horribly wrong"
}

func problemTwo(input []string) string {
	m := make(map[int]bool)
	for _, i := range input {
		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		m[num] = true
	}
	for _, i := range input {
		for _, j := range input[1:] {
			curr, _ := strconv.Atoi(i)
			next, _ := strconv.Atoi(j)
			num := 2020 - curr - next
			_, ok := m[num]
			if !ok {
				continue
			}
			ans := strconv.Itoa(curr * next * num)
			return ans
		}
	}
	return "something went horribly wrong"
}
