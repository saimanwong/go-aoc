package day03

import (
	"fmt"
	"math"
	"regexp"
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

func p1(lines []string) int {
	ans := 0
	for _, line := range lines {
		matches := regexp.MustCompile(`mul\((\d+,\d+)\)`).FindAllStringSubmatch(line, int(math.MaxUint>>1))
		for _, m := range matches {
			nums := strings.Split(m[1], ",")
			ans += toolbox.ToInt(nums[0]) * toolbox.ToInt(nums[1])
		}
	}
	return ans
}

const (
	do   = "do()"
	dont = `don't()`
)

func p2(lines []string) int {
	ans := 0
	// ofc it's one single line...
	enabled := true
	for _, line := range lines {
		matches := regexp.MustCompile(`mul\((\d+,\d+)\)|don't\(\)|do\(\)`).FindAllStringSubmatch(line, int(math.MaxUint>>1))
		for _, m := range matches {
			if m[0] == dont {
				enabled = false
				continue
			}
			if m[0] == do {
				enabled = true
				continue
			}
			if !enabled {
				continue
			}
			nums := strings.Split(m[1], ",")
			ans += toolbox.ToInt(nums[0]) * toolbox.ToInt(nums[1])
		}
	}
	return ans
}
