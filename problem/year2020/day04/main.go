package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	input []string
}

func (p *Problem) Run() {
	inputParsed := parseInput(p.input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func parseInput(input []string) []map[string]string {
	var passports []map[string]string
	curr := map[string]string{}
	input = append(input, "")
	for _, line := range input {
		if line == "" {
			passports = append(passports, curr)
			curr = map[string]string{}
			continue
		}
		for _, field := range strings.Split(line, " ") {
			s := strings.Split(field, ":")
			k, v := s[0], s[1]
			curr[k] = v
		}
	}
	return passports
}

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID) - not required
func checkValidBasic(input map[string]string) bool {
	req := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, r := range req {
		_, ok := input[r]
		if !ok {
			return false
		}
	}
	return true
}

func problemOne(input []map[string]string) string {
	ret := 0
	for _, p := range input {
		if checkValidBasic(p) {
			ret++
		}
	}
	return strconv.Itoa(ret)
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
func checkFourDigits(val string, min int, max int) bool {
	rFourDigits := regexp.MustCompile("^[0-9]{4}$")
	if !rFourDigits.MatchString(val) {
		return false
	}

	n, _ := strconv.Atoi(rFourDigits.FindString(val))
	if n < min || n > max {
		return false
	}

	return true
}

func checkHeight(val string) bool {
	rHeight := regexp.MustCompile("^([0-9]+)(in|cm)")
	if !rHeight.MatchString(val) {
		return false
	}
	s := rHeight.FindStringSubmatch(val)
	hgt, _ := strconv.Atoi(s[1])
	metric := s[2]

	if metric == "cm" && hgt >= 150 && hgt <= 193 {
		return true
	}

	if metric == "in" && hgt >= 59 && hgt <= 76 {
		return true
	}

	return false
}

func checkValidStrict(input map[string]string) bool {
	rHair := regexp.MustCompile("^#[0-9a-f]{6}$")
	rEye := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	rPid := regexp.MustCompile("^[0-9]{9}$")

	if !checkFourDigits(input["byr"], 1920, 2002) {
		return false
	}

	if !checkFourDigits(input["iyr"], 2010, 2020) {
		return false
	}

	if !checkFourDigits(input["eyr"], 2020, 2030) {
		return false
	}

	if !checkHeight(input["hgt"]) {
		return false
	}

	if !rHair.MatchString(input["hcl"]) {
		return false
	}

	if !rEye.MatchString(input["ecl"]) {
		return false
	}

	if !rPid.MatchString(input["pid"]) {
		return false
	}

	return true
}

func problemTwo(input []map[string]string) string {
	ret := 0
	for _, p := range input {
		if !checkValidBasic(p) {
			continue
		}

		if !checkValidStrict(p) {
			continue
		}
		ret++
	}
	return strconv.Itoa(ret)
}
