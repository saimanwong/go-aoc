package day03

import (
	"fmt"
	"sync"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	// part 1
	sum1 := 0
	for _, l := range p.input {
		str1, str2 := l[:len(l)/2], l[len(l)/2:]
		sum1 += toPrio(part1(str1, str2))
	}

	// part2
	groups := [3]string{}
	sum2 := 0
	for i, l := range p.input {
		mod := i % len(groups)
		if mod == 2 {
			groups[mod] = l
			sum2 += toPrio(part2(groups))
			continue
		}
		groups[mod] = l
	}
	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}

func part1(str1, str2 string) rune {
	for _, a := range str1 {
		for _, b := range str2 {
			if a == b {
				return a
			}
		}
	}
	panic("something wrong")
}

func part2(g [3]string) rune {
	var wg sync.WaitGroup
	freq := initFreq(map[rune][]bool{})
	for idx, line := range g {
		wg.Add(1)
		line := line
		idx := idx
		go func() {
			defer wg.Done()
			for _, r := range line {
				freq[r][idx] = true
			}
		}()
	}
	wg.Wait()
	for k, v := range freq {
		if v[0] && v[1] && v[2] {
			return k
		}
	}
	panic("part2")
}

func initFreq(m map[rune][]bool) map[rune][]bool {
	for i := 'a'; i <= 'z'; i++ {
		m[i] = make([]bool, 3)
	}
	for i := 'A'; i <= 'Z'; i++ {
		m[i] = make([]bool, 3)
	}
	return m
}

func toPrio(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}
	return int(r) - int('A') + 27
}
