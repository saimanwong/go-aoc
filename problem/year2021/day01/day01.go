package day01

import (
	"fmt"

	tb "github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	nums := tb.ToIntSlice(p.input...)

	// p1
	p1 := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			p1++
		}
	}

	// p2
	p2 := 0
	a, b, c := 1, 2, 3
	for c != len(nums) {
		if nums[a]+nums[b]+nums[c] > nums[a-1]+nums[b-1]+nums[c-1] {
			p2++
		}
		a++
		b++
		c++
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
