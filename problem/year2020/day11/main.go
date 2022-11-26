package day11

import (
	"fmt"
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

func parseInput(input []string) [][]rune {
	ret := [][]rune{}
	for r, line := range input {
		ret = append(ret, []rune{})
		ret[r] = append(ret[r], '.')
		for _, char := range line {
			ret[r] = append(ret[r], char)
		}
		ret[r] = append(ret[r], '.')
	}

	pad := []rune{}
	for i := 0; i < len(input[0])+2; i++ {
		pad = append(pad, '.')
	}

	ret = append(ret, pad)
	ret = append([][]rune{pad}, ret...)

	return ret
}

func visual(m [][]rune) {
	for _, row := range m {
		for _, char := range row {
			fmt.Printf("%s ", string(char))
		}
		fmt.Printf("\n")
	}
}

func createEmpty(rL int, cL int) [][]rune {
	ret := [][]rune{}
	for r := 0; r < rL; r++ {
		ret = append(ret, []rune{})
		for c := 0; c < cL; c++ {
			ret[r] = append(ret[r], '.')
		}
	}
	return ret
}

func checkOccu1(m [][]rune, r int, c int) int {
	ret := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if m[r+i][c+j] == '#' {
				ret++
			}
		}
	}
	return ret
}

func checkOccu2(m [][]rune, r int, c int) int {
	ret := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			a := 1
			rN := r + i
			cN := c + j
			for rN >= 0 && cN >= 0 && rN < len(m) && cN < len(m[0]) {
				if m[rN][cN] == 'L' {
					break
				}
				if m[rN][cN] == '#' {
					ret++
					break
				}
				a++
				rN = r + (i * a)
				cN = c + (j * a)
			}
		}
	}
	return ret
}

func applyRules(m [][]rune, p int) [][]rune {
	ret := createEmpty(len(m), len(m[0]))
	for r := 1; r < len(m)-1; r++ {
		occu := 0
		for c := 1; c < len(m[r])-1; c++ {
			if p == 1 {
				occu = checkOccu1(m, r, c)
			}

			if p == 2 {
				occu = checkOccu2(m, r, c)
			}

			if m[r][c] == 'L' && occu == 0 {
				ret[r][c] = '#'
				continue
			}

			if m[r][c] == '#' && ((occu >= 4 && p == 1) || (occu >= 5 && p == 2)) {
				ret[r][c] = 'L'
				continue
			}

			ret[r][c] = m[r][c]
		}
	}
	return ret
}

func checkTotOccu(m [][]rune) int {
	ret := 0
	for _, row := range m {
		for _, char := range row {
			if char == '#' {
				ret++
			}
		}
	}
	return ret
}

func problemOne(m [][]rune) string {
	occu1 := 0
	occu2 := -1
	newM := applyRules(m, 1)
	for occu1 != occu2 {
		occu2 = occu1
		occu1 = checkTotOccu(newM)
		newM = applyRules(newM, 1)
	}
	return fmt.Sprintf("%d", occu1)
}

func problemTwo(m [][]rune) string {
	occu1 := 0
	occu2 := -1
	newM := applyRules(m, 2)
	for occu1 != occu2 {
		occu2 = occu1
		occu1 = checkTotOccu(newM)
		newM = applyRules(newM, 2)
	}
	return fmt.Sprintf("%d", occu1)
}
