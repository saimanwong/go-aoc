package day10

import (
	"fmt"
	"strings"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []struct {
		ins string
		n   int
		cc  int
	}
}

func (p *Problem) SetInput(input []string) {
	in := make([]struct {
		ins string
		n   int
		cc  int
	}, len(input))
	for i, l := range input {
		spl := strings.Split(l, " ")
		tmp := struct {
			ins string
			n   int
			cc  int
		}{ins: spl[0]}
		if tmp.ins == "noop" {
			tmp.cc = 1
			in[i] = tmp
			continue
		}
		if tmp.ins == "addx" {
			tmp.n = toolbox.ToInt(spl[1])
			tmp.cc = 2
			in[i] = tmp
		}
	}

	p.input = in
}

func (p *Problem) Run() {
	var (
		pc         = 1
		basePC     = 20
		overBasePC = 40
		multi      = 0
		registerX  = 1
		spriteLit  = map[int]struct{}{registerX: {}, registerX + 1: {}, registerX + 2: {}}
		crt        = strings.Builder{}
	)
	p1 := 0
	for _, in := range p.input {
		for i := 0; i < in.cc; i++ { // each cycle
			if pc%(basePC+overBasePC*multi) == 0 {
				p1 += registerX * pc
				multi++
			}
			pixel := '.'
			if _, ok := spriteLit[pc%40]; ok {
				pixel = '#'
			}
			crt.WriteRune(pixel)
			if pc%40 == 0 {
				crt.WriteRune('\n')
			}
			pc++
		}
		registerX += in.n
		spriteLit = map[int]struct{}{registerX: {}, registerX + 1: {}, registerX + 2: {}}
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:")
	fmt.Println(crt.String())
}
