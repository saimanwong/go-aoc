package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	r, m := parseInput(p.input)
	fmt.Println("Answer 1: " + problemOne(r, m))
	fmt.Println("Answer 2: " + problemTwo(r, m))
}

type rules struct {
	val map[int][]string
}

type msg struct {
	val []string
}

func parseInput(input []string) (rules, msg) {
	newLine := 0
	re := regexp.MustCompile("([ab]{1})")

	r := rules{val: map[int][]string{}}
	m := msg{val: []string{}}
	for _, line := range input {
		if line == "" {
			newLine++
		}

		// msg
		if newLine > 0 {
			m.val = append(m.val, line)
			continue
		}

		// rule
		s := strings.Split(strings.ReplaceAll(line, ":", ""), " ")
		rInt, _ := strconv.Atoi(s[0])
		r.val[rInt] = []string{}
		for _, c := range s[1:] {
			toApp := c
			if sub := re.FindStringSubmatch(c); len(sub) > 0 {
				toApp = sub[0]
			}
			r.val[rInt] = append(r.val[rInt], toApp)
		}
	}

	return r, m
}

func debug(r rules, from int, end int) {
	for i := from; i < end; i++ {
		fmt.Println(i, r.val[i])
	}

}

func findPipes(line []string) []int {
	ret := []int{}
	for i, c := range line {
		if c == "|" {
			ret = append(ret, i)
		}
	}
	return ret
}

func getRegex(r rules) string {
	reFin := regexp.MustCompile("^[ab\\+\\|\\(\\)]+$")
	reAlp := regexp.MustCompile("^[ab]{1}$")
	reNum := regexp.MustCompile("^[0-9\\+\\?]+$")

	for !reFin.MatchString(strings.Join(r.val[0], "")) {
		tmpFinal := []string{}
		for _, rule := range r.val[0] {
			if reNum.MatchString(rule) {
				n, _ := strconv.Atoi(rule)
				tmpFinal = append(tmpFinal, r.val[n]...)
				continue
			}

			if reAlp.MatchString(rule) {
				tmpFinal = append(tmpFinal, rule)
				continue
			}

			tmpFinal = append(tmpFinal, rule)
		}

		// fmt.Println(tmpFinal)
		r.val[0] = tmpFinal
	}

	return fmt.Sprintf("^%s$", strings.Join(r.val[0], ""))

}

func parsedRules(r rules) rules {
	ret := rules{val: map[int][]string{}}

	for k, v := range r.val {
		ret.val[k] = []string{}
		containsPipe := false
		for _, c := range v {
			ret.val[k] = append(ret.val[k], c)
			if c == "|" {
				containsPipe = true
			}
		}

		if containsPipe {
			ret.val[k] = append([]string{"("}, ret.val[k]...)
			ret.val[k] = append(ret.val[k], ")")
		}
	}
	return ret
}

func problemOne(r rules, m msg) string {
	rn := parsedRules(r)
	regex := getRegex(rn)
	reNew := regexp.MustCompile(regex)
	ret := 0
	for _, m := range m.val {
		if reNew.MatchString(m) {
			ret++
		}
	}

	return fmt.Sprintf("%d", ret)
}

func modInputTwo(r rules, nr int, start int, end int) rules {
	newR := []string{}
	newR = append(newR, r.val[nr]...)
	newR = append(newR, "|")
	for i, c := range r.val[nr] {
		newR = append(newR, c)
		if i%2 == 0 {
			newR = append(newR, fmt.Sprintf("%d", start))
		}
	}

	for i := start; i < end; i++ {
		r.val[i] = []string{"|"}

		for j, c := range r.val[nr] {
			r.val[i] = append(r.val[i], c)
			if j%2 == 0 {
				r.val[i] = append(r.val[i], fmt.Sprintf("%d", i+1))
			}
		}

		r.val[i] = append(r.val[nr], r.val[i]...)
	}

	r.val[nr] = newR
	return r
}

func problemTwo(r rules, m msg) string {
	rn := rules{val: map[int][]string{}}
	for k, v := range r.val {
		rn.val[k] = []string{}
		for _, c := range v {
			rn.val[k] = append(rn.val[k], c)
		}
	}

	modInputTwo(rn, 8, 200, 210)
	modInputTwo(rn, 11, 210, 220)

	return problemOne(rn, m)
}
