package day02

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
	fmt.Println("Part 1:", p1(p.input))
	fmt.Println("Part 2:", p2(p.input))
}

func p1(input []string) int {
	reports := [][]int{}
	for _, line := range input {
		reports = append(reports, toolbox.ToIntSlice(
			strings.Split(line, " ")...,
		))
	}
	ans, _ := solve(reports)
	return ans
}

func p2(input []string) int {
	reports := [][]int{}
	for _, line := range input {
		reports = append(reports, toolbox.ToIntSlice(
			strings.Split(line, " ")...,
		))
	}
	count, rest := solve(reports)
	for _, r := range rest {
		for i := 0; i < len(r); i++ {
			co := make([]int, len(r))
			copy(co, r)
			tmp := append(co[:i], co[i+1:]...)
			ans, _ := solve([][]int{tmp})
			if ans > 0 {
				count += ans
				break
			}
		}
	}
	return count
}

func solve(reports [][]int) (int, [][]int) {
	count := 0
	rest := [][]int{}
	for _, report := range reports {
		increasing := report[0] < report[1]
		ok := true
		for i := 1; i < len(report); i++ {
			a, b := report[i-1], report[i]
			if a == b {
				ok = false
				break
			}
			dist := a - b
			if dist < 0 {
				dist *= -1
			}
			if dist < 1 || dist > 3 {
				ok = false
				break
			}
			if increasing && a > b {
				ok = false
				break
			}
			if !increasing && a < b {
				ok = false
				break
			}
		}
		if ok {
			count += 1
			continue
		}
		rest = append(rest, report)
	}
	return count, rest
}
