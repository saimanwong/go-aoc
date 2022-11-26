package day13

import (
	"fmt"
	"math"
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
	inputParsed := parseInput(p.input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}

type day13Input struct {
	goal int
	ids  []int
}

type bus struct {
	id  int
	idx int
}

func parseInput(input []string) day13Input {
	ret := day13Input{}
	ret.goal, _ = strconv.Atoi(input[0])
	ret.ids = []int{}
	s := strings.Split(input[1], ",")
	for _, id := range s {
		if id == "x" {
			ret.ids = append(ret.ids, -1)
			continue
		}

		idInt, _ := strconv.Atoi(id)
		ret.ids = append(ret.ids, idInt)
	}
	return ret
}

func problemOne(input day13Input) string {
	currDiff := 1.0
	precision := 5
	precisionString := "%." + strconv.Itoa(precision) + "f"
	ans := map[string]int{
		"id":    -1,
		"multi": -1,
	}

	for _, id := range input.ids {
		if id == -1 {
			continue
		}

		curr := float64(input.goal) / float64(id)
		splitCurr := strings.Split(fmt.Sprintf(precisionString, curr), ".")
		n, _ := strconv.Atoi(splitCurr[0])
		d, _ := strconv.ParseFloat(splitCurr[1], 16)
		d = d / math.Pow(10, float64(precision))

		if d > 0.5 {
			n++
			d = 1 - d
		}

		if d < currDiff {
			currDiff = d
			ans["id"] = id
			ans["multi"] = n
		}

	}

	earliest := ans["id"] * ans["multi"]
	diffMin := int(math.Abs(float64(earliest - input.goal)))
	ret := strconv.Itoa(ans["id"] * diffMin)
	return ret
}

func problemTwo(input day13Input) string {
	busInput := []bus{}
	for idx, id := range input.ids {
		if id == -1 {
			continue
		}

		busInput = append(busInput, bus{
			id:  id,
			idx: idx,
		})
	}

	t := 0
	factor := 1
	for _, b := range busInput {
		for (t+b.idx)%b.id != 0 {
			t += factor
		}

		factor *= b.id
	}

	return strconv.Itoa(t)
}
