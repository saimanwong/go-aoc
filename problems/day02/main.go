package day02

import (
	"fmt"
	"regexp"
	"strconv"
)

type password struct {
	min  int
	max  int
	char rune
	pass string
}

func parseInput(input []string) []password {
	var ret []password
	re := regexp.MustCompile(`\b([0-9].*)-([0-9].*) ([a-z]{1}): ([a-z].*)\b`)
	for _, i := range input {
		match := re.FindStringSubmatch(i)
		min, _ := strconv.Atoi(match[1])
		max, _ := strconv.Atoi(match[2])
		char := []rune(match[3])[0]
		pass := match[4]
		ret = append(ret, password{
			min:  min,
			max:  max,
			char: char,
			pass: pass,
		})

	}
	return ret
}

func checkCharsOne(p password) bool {
	count := 0
	for _, char := range p.pass {
		if char == p.char {
			count++
		}

		if count > p.max {
			return false
		}

	}

	if count < p.min {
		return false
	}
	return true
}

func problemOne(input []password) string {
	ret := 0
	for _, p := range input {
		if checkCharsOne(p) {
			ret++
		}
	}
	return strconv.Itoa(ret)
}

func checkCharsTwo(p password) bool {
	first := false
	second := false
	if p.char == rune(p.pass[p.min-1]) {
		first = true
	}
	if p.char == rune(p.pass[p.max-1]) {
		second = true
	}
	if first && !second {
		return true
	}
	if !first && second {
		return true
	}

	return false
}

func problemTwo(input []password) string {
	ret := 0
	for _, p := range input {
		if checkCharsTwo(p) {
			ret++
		}
	}
	return strconv.Itoa(ret)
}

func Run(input []string) {
	inputParsed := parseInput(input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}
