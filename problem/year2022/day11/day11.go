package day11

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	Monkeys map[int]*Monkey
	input   []string
}

type Monkey struct {
	Items     []int
	Operation Operation
	Test      struct {
		Mod   int
		True  int
		False int
	}
}

type Operation struct { // -1 == old
	Left     int
	Operator rune
	Right    int
}

func (p *Problem) SetInput(input []string) {
	p.input = input
	var currMonk *Monkey
	p.Monkeys = map[int]*Monkey{}
	for _, l := range input {
		if l == "" {
			continue
		}
		spl := strings.Split(l, ": ")
		if strings.Contains(spl[0], "Monkey") {
			n := toolbox.ToInt(strings.Trim(strings.Split(spl[0], " ")[1], ":"))
			p.Monkeys[n] = &Monkey{}
			currMonk = p.Monkeys[n]
		}
		if strings.Contains(spl[0], "Starting items") {
			spl := strings.Split(spl[1], ", ")
			currMonk.Items = append(currMonk.Items, toolbox.ToIntSlice(spl...)...)
			continue
		}
		if strings.Contains(spl[0], "Operation") {
			spl := strings.Split(spl[1], " ") // new = old + 8
			currMonk.Operation.Operator = rune(spl[3][0])
			left, right := -1, -1
			if spl[2] != "old" {
				left = toolbox.ToInt(spl[4])
			}
			if spl[4] != "old" {
				right = toolbox.ToInt(spl[4])
			}
			currMonk.Operation.Left = left
			currMonk.Operation.Right = right
			continue
		}
		if strings.Contains(spl[0], "Test") {
			currMonk.Test.Mod = toolbox.ToInt(strings.Split(spl[1], "divisible by ")[1])
			continue
		}
		if strings.Contains(spl[0], "If true") {
			currMonk.Test.True = toolbox.ToInt(strings.Split(spl[1], "throw to monkey ")[1])
			continue
		}
		if strings.Contains(spl[0], "If false") {
			currMonk.Test.False = toolbox.ToInt(strings.Split(spl[1], "throw to monkey ")[1])
			continue
		}
	}
}

func (p *Problem) Debug(m int) {
	b, err := json.MarshalIndent(p.Monkeys[m], "", " ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(b))
}

func (p *Problem) Reset() {
	p.SetInput(p.input)
}

func (p *Problem) Run() {
	monkFreq1 := map[int]int{}
	for i := 0; i < len(p.Monkeys); i++ { // init
		monkFreq1[i] = 0
	}
	for r := 0; r < 20; r++ { // p1
		p.solve(false, monkFreq1, 3)
	}
	// p2
	p.Reset()
	monkFreq2 := map[int]int{}
	newMod := 1
	for _, m := range p.Monkeys {
		newMod *= m.Test.Mod
	}
	for r := 0; r < 10_000; r++ { // p2
		p.solve(true, monkFreq2, newMod)
	}
	sFreq1, sFreq2 := []int{}, []int{}
	for i := 0; i < len(p.Monkeys); i++ {
		sFreq1 = append(sFreq1, monkFreq1[i])
		sFreq2 = append(sFreq2, monkFreq2[i])
	}
	sort.Ints(sFreq1)
	sort.Ints(sFreq2)
	fmt.Println("Part 1:", sFreq1[len(sFreq1)-1]*sFreq1[len(sFreq1)-2])
	fmt.Println("Part 2:", sFreq2[len(sFreq2)-1]*sFreq2[len(sFreq2)-2])
}

func (p *Problem) solve(p2 bool, monkFreq map[int]int, mod int) {
	for i := 0; i < len(p.Monkeys); i++ {
		currMonk := p.Monkeys[i]
		for len(currMonk.Items) > 0 { // inpect item
			monkFreq[i]++
			currItem := currMonk.Items[0]
			currMonk.Items = currMonk.Items[1:]
			wl := worryLevel(currItem, currMonk.Operation) / mod
			if p2 {
				wl = worryLevel(currItem, currMonk.Operation) % mod
			}
			if wl%currMonk.Test.Mod == 0 { // true
				p.Monkeys[currMonk.Test.True].Items = append(p.Monkeys[currMonk.Test.True].Items, wl)
				continue
			}
			p.Monkeys[currMonk.Test.False].Items = append(p.Monkeys[currMonk.Test.False].Items, wl)
		}
	}
}

func worryLevel(n int, op Operation) int {
	left, right := op.Left, op.Right
	if op.Left == -1 {
		left = n
	}
	if op.Right == -1 {
		right = n
	}
	if op.Operator == '*' {
		return left * right
	}
	if op.Operator == '+' {
		return left + right
	}
	panic("oh no")
}
