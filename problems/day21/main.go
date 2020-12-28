package day21

import (
	"fmt"
	"sort"
	"strings"
)

type food struct {
	ingr []string
	alr  []string
}

type ingrMap map[string]int

type helper interface {
	max() []string
}

func parseInput(input []string) []*food {
	ret := []*food{}
	del := []string{"(", ")", ","}
	for _, line := range input {
		s := line
		for _, c := range del {
			s = strings.ReplaceAll(s, c, "")
		}

		spl := strings.Split(s, " contains ")
		ingr := strings.Split(spl[0], " ")
		alr := strings.Split(spl[1], " ")

		ret = append(ret, &food{
			ingr: ingr,
			alr:  alr,
		})
	}

	return ret
}

func (i ingrMap) max() []string {
	max := -1
	for _, n := range i {
		if n > max {
			max = n
		}
	}

	ret := []string{}
	for ingr, n := range i {
		if max == n {
			ret = append(ret, ingr)
		}
	}

	return ret
}

func problemOne(input []*food) (map[string][]string, string) {
	m := map[string]ingrMap{}
	for _, f := range input {
		for _, a := range f.alr {
			_, ok := m[a]
			if !ok {
				m[a] = map[string]int{}
			}
			for _, i := range f.ingr {
				_, ok := m[a][i]
				if !ok {
					m[a][i] = 0
				}
				m[a][i]++
			}
		}
	}

	m1 := map[string][]string{}
	for alr, ingr := range m {
		m1[alr] = ingr.max()
	}

	taken := map[string]bool{}
	for len(taken) != len(m1) {
		for alr, ingr := range m1 {
			for idx, i := range ingr {
				if len(ingr) == 1 {
					taken[i] = true
					break
				}

				_, ok := taken[i]
				if ok {
					left := m1[alr][:idx]
					right := m1[alr][idx+1:]
					m1[alr] = append(left, right...)
				}

			}
		}
	}

	ret := 0
	for _, f := range input {
		for _, ingr := range f.ingr {
			_, ok := taken[ingr]
			if ok {
				continue
			}
			ret++
		}
	}

	return m1, fmt.Sprintf("%d", ret)
}

func problemTwo(input map[string][]string) string {
	ret := strings.Builder{}
	keyOrder := []string{}
	for k, _ := range input {
		keyOrder = append(keyOrder, k)
	}
	sort.Strings(keyOrder)
	for _, k := range keyOrder {
		ret.WriteString(input[k][0] + ",")
	}

	return strings.TrimRight(ret.String(), ",")
}

func Run(input []string) {
	inputParsed := parseInput(input)
	p2input, ans1 := problemOne(inputParsed)
	fmt.Println("Answer 1: " + ans1)
	fmt.Println("Answer 2: " + problemTwo(p2input))
}
