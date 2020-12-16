package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type validRange struct {
	min int
	max int
}

type field struct {
	name        string
	validRanges []validRange
}

type ticket struct {
	values []int
}

func parseInput(input []string) ([]field, []ticket, ticket) {
	f := []string{}
	n := []string{}
	retMy := ticket{values: []int{}}

	count := 0
	for _, line := range input {
		if line == "" {
			count++
		}

		// fields
		if count == 0 {
			f = append(f, line)
		}

		if count == 1 {
			count++
			continue
		}

		if count == 2 {
			count++
			continue
		}

		// My ticket
		if count == 3 {
			s := strings.Split(line, ",")
			for _, v := range s {
				toInt, _ := strconv.Atoi(v)
				retMy.values = append(retMy.values, toInt)
			}
			count++
			continue
		}

		if count == 4 {
			count++
			continue
		}

		if count == 5 && line != "" {
			n = append(n, line)
		}
	}

	re := regexp.MustCompile("^([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$")
	retField := []field{}
	for _, fi := range f {
		s := re.FindStringSubmatch(fi) // 1 name, 2 min, 3, max, 4 min, 5 max

		min1, _ := strconv.Atoi(s[2])
		max1, _ := strconv.Atoi(s[3])
		v1 := validRange{
			min: min1,
			max: max1,
		}

		min2, _ := strconv.Atoi(s[4])
		max2, _ := strconv.Atoi(s[5])
		v2 := validRange{
			min: min2,
			max: max2,
		}

		f1 := field{
			name:        s[1],
			validRanges: []validRange{v1, v2},
		}

		retField = append(retField, f1)
	}
	// fmt.Println("Field", retField)

	retNearby := []ticket{}
	for i, line := range n {
		if i == 0 {
			continue
		}
		s := strings.Split(line, ",")

		tmp := ticket{values: []int{}}
		for _, i := range s {
			toInt, _ := strconv.Atoi(i)
			tmp.values = append(tmp.values, toInt)
		}
		retNearby = append(retNearby, tmp)
	}

	return retField, retNearby, retMy
}

func isValid(f []field, n int) bool {
	for _, v := range f { // fields
		for _, vr := range v.validRanges { // vr.min, vr.max
			if n >= vr.min && n <= vr.max {
				return true
			}
		}
	}
	return false
}

func isValidWithFields(f []field, n int) []string {
	validFields := []string{}
	for _, v := range f { // fields
		for _, vr := range v.validRanges { // vr.min, vr.max
			if n >= vr.min && n <= vr.max {
				// return true
				validFields = append(validFields, v.name)
			}
		}
	}
	return validFields
}

func problemOne(f []field, n []ticket) string {
	invalid := []int{}
	num := []int{}

	for _, i := range n {
		for _, j := range i.values {
			num = append(num, j)
		}
	}

	for _, n := range num {
		if !isValid(f, n) {
			invalid = append(invalid, n)
		}
	}

	ret := 0
	for _, s := range invalid {
		ret += s
	}
	return strconv.Itoa(ret)
}

func problemTwo(f []field, n []ticket, my ticket) string {
	validTickets := []ticket{}
	// Check valid tickets []ticket[values: []int{}]
	for _, t := range n {
		valid := true
		for _, num := range t.values {
			if !isValid(f, num) {
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, t)
		}
	}

	matrix := map[string][]int{} // class: [[..][..][..]
	for _, fie := range f {
		matrix[fie.name] = []int{}
		for range my.values {
			matrix[fie.name] = append(matrix[fie.name], 0)
		}
	}

	// fmt.Println("Before", matrix)
	for _, tick := range validTickets {
		for idx, val := range tick.values {
			validFields := isValidWithFields(f, val)
			for _, vf := range validFields {
				matrix[vf][idx]++
			}
		}
	}
	// fmt.Println("After", matrix)

	maxLen := len(my.values)
	final := map[string]int{} // field, idx
	sumMap := map[string]int{}

	for len(final) < maxLen {
		// Figure out matrix...
		for name, val := range matrix {
			sumMap[name] = 0
			for _, num := range val {
				sumMap[name] += num
			}
		}

		// Get min tot
		min := 99999999999999999
		minField := ""
		for name, val := range sumMap {
			if val < min {
				min = val
				minField = name
			}
		}
		// fmt.Println("minField", minField)

		// Get max inner
		max := 0
		maxIdx := 0
		for idx, val := range matrix[minField] {
			if val > max {
				max = val
				maxIdx = idx
			}
		}

		// Remove minField from matrix and sumMap, add to final
		final[minField] = maxIdx
		delete(matrix, minField)
		delete(sumMap, minField)

		// Decrement everyone on maxIdx by its value
		for name, _ := range matrix {
			matrix[name][maxIdx] -= max
			sumMap[name] -= max
		}
	}

	departure := []string{
		"departure location",
		"departure station",
		"departure platform",
		"departure track",
		"departure date",
		"departure time",
	}

	ret := 1
	for _, key := range departure {
		ret *= my.values[final[key]]
	}

	// fmt.Println("final", final)
	return strconv.Itoa(ret)
}

func Run(input []string) {
	fields, nearbyTickets, myTicket := parseInput(input)
	fmt.Println("Answer 1: " + problemOne(fields, nearbyTickets))
	fmt.Println("Answer 2: " + problemTwo(fields, nearbyTickets, myTicket))
}
