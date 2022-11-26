package day08

import (
	"fmt"
	"regexp"
	"strconv"
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

type instruction struct {
	opr string
	arg int
}

func parseInput(input []string) []instruction {
	ins := []instruction{}
	re := regexp.MustCompile("^([a-z]{3}) ([+-])([0-9]+)$")
	for _, line := range input {
		s := re.FindStringSubmatch(line)
		opr := s[1]           // operation
		argStr := s[2] + s[3] // + -
		if rune(argStr[0]) == '+' {
			argStr = argStr[1:]
		}
		arg, err := strconv.Atoi(argStr)
		if err != nil {
			fmt.Println(err)
		}
		ins = append(ins, instruction{
			opr: opr,
			arg: arg,
		})
	}
	return ins
}

func executeIns(ins instruction, curr *int, acc *int) {
	if ins.opr == "nop" {
		*curr++
	}

	if ins.opr == "acc" {
		*acc += ins.arg
		*curr++
	}

	if ins.opr == "jmp" {
		*curr += ins.arg
	}
}

func problemOne(input []instruction) string {
	acc, curr := 0, 0
	seen := map[int]bool{}

	for {
		ins := input[curr]
		_, ok := seen[curr]
		if !ok {
			seen[curr] = false
		}

		if !seen[curr] {
			seen[curr] = true
			executeIns(ins, &curr, &acc)
			continue
		}

		break
	}

	return strconv.Itoa(acc)
}

func problemTwo(input []instruction) string {
	input = append(input, instruction{opr: "fin"})
	inputModified := input
	updatedIdx := map[int]bool{}
	acc, curr := 0, 0
	seen := map[int]bool{}

	for {
		ins := inputModified[curr]
		if ins.opr == "fin" {
			break
		}

		_, ok := seen[curr]
		if !ok {
			seen[curr] = false
		}

		if !seen[curr] {
			seen[curr] = true
			executeIns(ins, &curr, &acc)
			continue
		}

		// Reset
		if seen[curr] {
			acc, curr = 0, 0
			seen = map[int]bool{}
			inputModified = make([]instruction, len(input))
			_ = copy(inputModified, input)

			for i, ins := range inputModified {
				_, ok := updatedIdx[i]
				if ok {
					continue
				}

				if ins.opr == "nop" && ins.arg != 0 {
					inputModified[i].opr = "jmp"
					updatedIdx[i] = true
					break
				}

				if ins.opr == "jmp" {
					inputModified[i].opr = "nop"
					updatedIdx[i] = true
					break
				}
			}
		}
	}

	return strconv.Itoa(acc)
}
