package day18

import (
	"fmt"
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

type assignment interface {
	eval() int
	eval2() int
}

type sequence struct {
	opr   []string
	num   []int
	child map[int]*sequence
	line  string
}

type paren struct {
	opened int
	closed int
}

type homework map[int]*sequence

func parseInput(input []string) []homework {
	ret := []homework{}
	for _, line := range input {
		s := strings.ReplaceAll(line, " ", "")
		spl := strings.Split(s, "")

		opened := []int{}
		pairs := []paren{}
		for i, c := range spl {
			if c == "(" {
				opened = append(opened, i)
			}
			if c == ")" {
				p := paren{
					opened: opened[len(opened)-1],
					closed: i,
				}
				pairs = append(pairs, p)
				opened = opened[:len(opened)-1]
			}
		}

		visited := map[int]int{}
		id := -1
		seqs := homework{}
		for len(pairs) > 0 {
			curr := pairs[0]
			pairs = pairs[1:]

			seq := &sequence{
				opr:   []string{},
				num:   []int{},
				child: map[int]*sequence{},
			}

			for i := curr.opened; i <= curr.closed; i++ {
				visitedID, ok := visited[i]
				if i == curr.opened || i == curr.closed {
					visited[i] = id
					continue
				}
				if !ok {
					visited[i] = id
					if spl[i] == "*" || spl[i] == "+" {
						seq.opr = append(seq.opr, spl[i])
						continue
					}

					n, _ := strconv.Atoi(spl[i])
					seq.num = append(seq.num, n)
					continue
				}

				found := false
				for _, n := range seq.num {
					if n == visitedID {
						found = true
						break
					}
				}
				if !found {
					seq.num = append(seq.num, visitedID)
				}
				seq.child[visitedID] = seqs[visitedID]
			}

			seqs[id] = seq
			id--
		}

		// Create root
		seqs[0] = &sequence{
			opr:   []string{},
			num:   []int{},
			child: map[int]*sequence{},
			line:  line,
		}
		min := 0
		root := seqs[0]
		for i, c := range spl {
			visitedID, ok := visited[i]
			if ok {
				found := false
				for _, n := range root.num {
					if n == visitedID {
						found = true
						break
					}
				}

				if visitedID < min && !found {
					root.num = append(root.num, visitedID)
					root.child[visitedID] = seqs[visitedID]
					min = visitedID
				}
				continue
			}

			min = 0
			if c == "*" || c == "+" {
				root.opr = append(root.opr, c)
				continue
			}

			n, _ := strconv.Atoi(c)
			root.num = append(root.num, n)
		}

		ret = append(ret, seqs)
	}

	return ret
}

func (s sequence) eval() int {
	for len(s.opr) > 0 {
		currOpr := s.opr[0]
		s.opr = s.opr[1:]
		n1 := s.num[0]
		s.num = s.num[1:]
		n2 := s.num[0]
		s.num = s.num[1:]

		if n1 < 0 {
			n1 = s.child[n1].eval()
		}

		if n2 < 0 {
			n2 = s.child[n2].eval()
		}

		if currOpr == "*" {
			s.num = append([]int{n1 * n2}, s.num...)
		}

		if currOpr == "+" {
			s.num = append([]int{n1 + n2}, s.num...)
		}
	}

	return s.num[0]
}

func (s sequence) eval2() int {
	for len(s.opr) > 0 {
		idxAdd := 0
		for i, o := range s.opr {
			if o == "+" {
				idxAdd = i
				break
			}
		}

		currOpr := s.opr[idxAdd]
		leftOpr := s.opr[:idxAdd]
		rightOpr := s.opr[idxAdd+1:]
		s.opr = append(leftOpr, rightOpr...)

		n1 := s.num[idxAdd]
		n2 := s.num[idxAdd+1]

		leftNum := s.num[:idxAdd]
		rightNum := s.num[idxAdd+2:]
		s.num = append(leftNum, rightNum...)

		if n1 < 0 {
			n1 = s.child[n1].eval2()
		}

		if n2 < 0 {
			n2 = s.child[n2].eval2()
		}

		if currOpr == "*" {
			s.num = append([]int{n1 * n2}, s.num...)
		}

		if currOpr == "+" {
			leftNum = s.num[:idxAdd]
			rightNum = s.num[idxAdd:]
			rightNum = append([]int{n1 + n2}, rightNum...)
			s.num = append(leftNum, rightNum...)
		}
	}

	return s.num[0]
}

func problemOne(input []homework) string {
	ret := 0
	for _, c := range input {
		ret += c[0].eval()
	}
	return fmt.Sprintf("%d", ret)
}

func problemTwo(input []homework) string {
	ret := 0
	for _, c := range input {
		res := c[0].eval2()
		// fmt.Printf("%s %d\n", c[0].line, res)
		ret += res
	}
	return fmt.Sprintf("%d", ret)
}
