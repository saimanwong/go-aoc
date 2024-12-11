package day11

import (
	"fmt"
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
	stones := toolbox.ToIntSlice(strings.Split(p.input[0], " ")...)
	fmt.Println("Part 1:", solve1(stones, 25))
	fmt.Println("Part 2:", solve2(stones, 75))
}

func solve1(stones []int, n int) int {
	currStones := make([]int, len(stones))
	copy(currStones, stones)
	for i := 0; i < n; i++ {
		newStones := rules(currStones)
		currStones = make([]int, len(newStones))
		copy(currStones, newStones)
	}
	return len(currStones)
}

func solve2(stones []int, n int) int {
	freq := map[int]int{}
	for _, s := range stones {
		if _, ok := freq[s]; !ok {
			freq[s] = 0
		}
		freq[s]++
	}

	for i := 0; i < n; i++ {
		newFreq := map[int]int{}
		for s, n := range freq {
			newStones := rules([]int{s})
			tmp := map[int]int{}
			for _, ss := range newStones {
				if _, ok := tmp[ss]; !ok {
					tmp[ss] = 0
				}
				tmp[ss]++
			}
			for k := range tmp {
				tmp[k] *= n
				if _, ok := newFreq[k]; !ok {
					newFreq[k] = 0
				}
				newFreq[k] += tmp[k]
			}
		}
		freq = newFreq
	}
	sum := 0
	for _, n := range freq {
		sum += n
	}
	return sum
}

func rules(stones []int) []int {
	newStones := []int{}
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}
		str := fmt.Sprintf("%d", stone)
		if len(str)%2 == 0 {
			a, b := str[:len(str)/2], str[len(str)/2:]
			newStones = append(newStones, toolbox.ToInt(a), toolbox.ToInt(b))
			continue
		}
		newStones = append(newStones, stone*2024)
	}
	return newStones
}
