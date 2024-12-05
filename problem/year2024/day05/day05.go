package day05

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

type Rules map[int]map[int]interface{}

func (p *Problem) Run() {
	newSection := false
	rules, updates := Rules{}, [][]int{}
	for _, line := range p.input {
		if line == "" {
			newSection = true
			continue
		}
		if !newSection {
			spl := strings.Split(line, "|")
			before, after := toolbox.ToInt(spl[0]), toolbox.ToInt(spl[1])
			if _, ok := rules[before]; !ok {
				rules[before] = map[int]interface{}{}
			}
			rules[before][after] = nil
			continue
		}
		updates = append(updates, toolbox.ToIntSlice(strings.Split(line, ",")...))
	}
	fmt.Println("Part 1:", p1(rules, updates))
	fmt.Println("Part 2:", p2(rules, updates))
}

func p1(beforeRules Rules, updates [][]int) int {
	okUpdates, _ := okUpdates(beforeRules, updates)
	ans := 0
	for _, u := range okUpdates {
		ans += u[len(u)/2]
	}
	return ans
}

func p2(rules Rules, updates [][]int) int {
	_, notOKUpdates := okUpdates(rules, updates)
	newUpdates := [][]int{}
	type Number struct {
		N   int
		Len int
	}
	for _, update := range notOKUpdates {
		lengths := []Number{}
		for _, n := range update {
			count := 0
			for _, nn := range update {
				if _, ok := rules[n][nn]; ok {
					count++
				}
			}
			lengths = append(lengths, Number{N: n, Len: count})
		}

		sort.Slice(lengths, func(i, j int) bool {
			return lengths[i].Len > lengths[j].Len
		})
		u := []int{}
		for _, n := range lengths {
			u = append(u, n.N)
		}
		newUpdates = append(newUpdates, u)
	}
	ans := 0
	for _, u := range newUpdates {
		ans += u[len(u)/2]
	}
	return ans
}

func okUpdates(
	rules Rules,
	updates [][]int,
) ([][]int, [][]int) {
	okUpdates := [][]int{}
	notOKUpdates := [][]int{}
	for _, u := range updates {
		okUpdate := true
		for i := len(u) - 1; i >= 0; i-- {
			curr := u[i]
			for j := i - 1; j >= 0; j-- {
				if i == j {
					continue
				}
				next := u[j]
				if _, ok := rules[curr]; !ok {
					continue
				}
				if _, ok := rules[curr][next]; !ok {
					continue
				}
				okUpdate = false
				break
			}
		}
		if okUpdate {
			okUpdates = append(okUpdates, u)
			continue
		}
		notOKUpdates = append(notOKUpdates, u)
	}
	return okUpdates, notOKUpdates
}

// nope
// func p2(rules Rules, updates [][]int) int {
// 	_, notOKUpdates := okUpdates(rules, updates)
// 	newUpdates := [][]int{}
// 	for x, u := range notOKUpdates {
// 		result := []int{}
// 		for i := 0; i < len(u); i++ {
// 			curr := u[i]
// 			if i == 0 {
// 				result = append(result, curr)
// 				continue
// 			}
// 			j := len(result) - 1
// 			foundPos := false
// 			for !foundPos {
// 				if j < 0 {
// 					result = append(result, curr)
// 					break
// 				}
// 				prev := result[j]
// 				if _, ok := rules[curr]; ok { // has rules
// 					if _, ok := rules[curr][prev]; ok { // curr must be before prev
// 						j--
// 						if x == 0 {
// 							fmt.Println("curr before prev", curr, prev)
// 						}
// 						continue
// 					}
// 				}
//
// 				if _, ok := rules[prev]; ok { // prev has rules
// 					if _, ok := rules[prev][curr]; ok { // must be after prev
// 						front := append([]int{}, result[j+1:]...)
// 						tmp := append(result[0:j+1], curr)
// 						result = append(tmp, front...)
// 						foundPos = true
// 						if x == 0 {
// 							fmt.Println("curr after prev", curr, prev)
// 						}
// 					}
// 				}
// 			}
// 		}
// 		newUpdates = append(newUpdates, result)
// 		if x == 0 {
// 			fmt.Println(updates[x], newUpdates[x])
// 		}
// 	}
// 	ans := 0
// 	for _, u := range newUpdates {
// 		ans += u[len(u)/2]
// 	}
// 	return ans
// }
