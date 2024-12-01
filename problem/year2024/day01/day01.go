package day01

import (
	"fmt"
	"sort"
	"strings"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	fmt.Println("Part 1:", p1(p.input))
	fmt.Println("Part 2:", p2(p.input))
}

func p1(input []string) int {
	left, right := []int{}, []int{}
	for _, s := range input {
		spl := strings.Split(s, "   ")
		left = append(left, toolbox.ToInt(spl[0]))
		right = append(right, toolbox.ToInt(spl[1]))
	}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := 0; i < len(left); i++ {
		dist := left[i] - right[i]
		if dist < 0 {
			dist = dist * -1
		}
		sum += dist
	}
	return sum
}

func p2(input []string) int {
	left, right := []int{}, map[int]int{}
	for _, s := range input {
		spl := strings.Split(s, "   ")
		left = append(left, toolbox.ToInt(spl[0]))

		tmp := toolbox.ToInt(spl[1])
		_, ok := right[tmp]
		if !ok {
			right[tmp] = 1
			continue
		}
		right[tmp] += 1
	}
	sum := 0
	for _, n := range left {
		mulitplier := 0
		if count, ok := right[n]; ok {
			mulitplier = count
		}
		sum = sum + n*mulitplier
	}
	return sum
}
