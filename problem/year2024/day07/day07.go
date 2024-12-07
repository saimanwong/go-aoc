package day07

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

type Equation struct {
	Target int
	Nums   []int
}

func (p *Problem) Run() {
	equations := make([]Equation, len(p.input))
	for i, line := range p.input {
		e := Equation{}
		spl := strings.Split(line, ": ")
		e.Target = toolbox.ToInt(spl[0])
		spl = strings.Split(spl[1], " ")
		e.Nums = toolbox.ToIntSlice(spl...)
		equations[i] = e
	}

	ans1 := 0
	leftovers := []Equation{}
	for _, e := range equations {
		combos := combos(len(e.Nums), 2)
		if solveEquation(e, combos) {
			ans1 += e.Target
			continue
		}
		leftovers = append(leftovers, e)
	}

	ans2 := 0
	for _, e := range leftovers {
		combos := combos(len(e.Nums), 3)
		if solveEquation(e, combos) {
			ans2 += e.Target
		}
	}
	fmt.Println("Part 1:", ans1)
	fmt.Println("Part 2:", ans1+ans2)
}

func solveEquation(e Equation, combos [][]int) bool {
	for _, c := range combos {
		sum := e.Nums[0]
		for i := 1; i < len(e.Nums); i++ {
			if sum > e.Target {
				break
			}
			if c[i-1] == 0 {
				sum += e.Nums[i]
				continue
			}
			if c[i-1] == 1 {
				sum *= e.Nums[i]
				continue
			}
			// ||
			sum = toolbox.ToInt(fmt.Sprintf("%d%d", sum, e.Nums[i]))
		}
		if sum == e.Target {
			return true
		}
	}
	return false
}

func combos(n int, base float64) [][]int {
	nCombos := int(math.Pow(base, float64(n-1)))
	combos := [][]int{}
	for i := 0; i < nCombos; i++ {
		combo := make([]int, n-1)
		b := fmt.Sprintf("%b", i)
		if base == 3 {
			b = strconv.FormatInt(int64(i), 3)
		}
		for len(b) < len(combo) {
			b = "0" + b
		}
		for j := 0; j < len(b); j++ {
			combo[j] = toolbox.ToInt(string(b[j]))
		}
		combos = append(combos, combo)
	}
	return combos
}
