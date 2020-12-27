package day07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	name string
	qty  int
}

type bags struct {
	bags map[string][]bag
}

func parseInput(input []string) bags {
	ret := bags{bags: map[string][]bag{}}

	for _, line := range input {
		s := strings.Split(line, "contain") // 0 outer, 1 inner

		outer := strings.TrimRight(strings.Join(strings.Split(s[0], " "), ""), "s")
		_, ok := ret.bags[outer]
		if !ok {
			ret.bags[outer] = []bag{}
		}

		s = strings.Split(s[1], ",")
		if len(s) == 0 {
			continue
		}

		reInner := regexp.MustCompile("([0-9)]) ([a-z]+) ([a-z]+) (bag).*")
		for _, inner := range s {
			if !reInner.MatchString(inner) {
				continue
			}

			r := reInner.FindStringSubmatch(inner)
			qty, _ := strconv.Atoi(r[1])
			name := r[2] + r[3] + r[4]
			b := bag{
				name: name,
				qty:  qty,
			}
			ret.bags[outer] = append(ret.bags[outer], b)
		}

	}

	return ret
}

func problemOne(input bags) string {
	// Reverse bags... can be contained by
	b := bags{bags: map[string][]bag{}}
	for outer, innerBags := range input.bags {
		for _, inner := range innerBags {
			_, ok := b.bags[inner.name]
			if !ok {
				b.bags[inner.name] = []bag{}
			}
			b.bags[inner.name] = append(b.bags[inner.name], bag{name: outer})
		}
	}

	memo := map[string]bool{}
	seen := map[string]bool{}
	q := []string{"shinygoldbag"}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		_, ok := seen[curr]
		if !ok {
			seen[curr] = false
		}

		if !seen[curr] {
			seen[curr] = true
			memo[curr] = true
			for _, v := range b.bags[curr] {
				q = append(q, v.name)
			}
		}
	}

	return strconv.Itoa(len(memo) - 1)
}

func problemTwo(input bags) string {
	tot := []string{"shinygoldbag"}
	i := 0
	for {
		if i >= len(tot) {
			break
		}

		curr := tot[i]
		for _, innerBag := range input.bags[curr] {
			for j := 0; j < innerBag.qty; j++ {
				tot = append(tot, innerBag.name)
			}
		}

		i++
	}

	return strconv.Itoa(len(tot) - 1)
}

func Run(input []string) {
	parsedInput := parseInput(input)
	fmt.Println("Answer 1: " + problemOne(parsedInput))
	fmt.Println("Answer 2: " + problemTwo(parsedInput))
}
