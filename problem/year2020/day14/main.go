package day14

import (
	"fmt"
	"math"
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
	inputParsed := parseInput(p.input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}

type memAddr map[uint64]uint64

type instruct struct {
	mask   [36]rune
	writes []memAddr
}

func newValue(m [36]rune, v uint64) uint64 {
	vBin := strconv.FormatUint(v, 2)
	src := []rune(fmt.Sprintf("%036s", vBin))
	newBin := [36]rune{}
	for i, c := range src {
		mC := m[i]
		if mC != 'X' {
			newBin[i] = mC
			continue
		}
		newBin[i] = c
	}

	var newStr strings.Builder
	for _, c := range newBin {
		newStr.WriteString(string(c))
	}

	ret, err := strconv.ParseUint(newStr.String(), 2, 64)
	if err != nil {
		panic(err)
	}
	return ret
}

func parseInput(input []string) []instruct {
	ret := []instruct{}

	toAdd := instruct{}
	reMask := regexp.MustCompile("mask = ([01X]{36})")
	reMemAddr := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	for _, line := range input {
		if reMask.MatchString(line) {
			if len(toAdd.writes) > 0 {
				ret = append(ret, toAdd)
			}
			toAdd = instruct{
				mask:   [36]rune{},
				writes: []memAddr{},
			}

			s := reMask.FindStringSubmatch(line)
			for n, r := range s[1] {
				toAdd.mask[n] = r
			}
			continue
		}

		s := reMemAddr.FindStringSubmatch(line)
		addr, _ := strconv.ParseUint(s[1], 10, 64)
		val, _ := strconv.ParseUint(s[2], 10, 64)

		toAdd.writes = append(toAdd.writes, memAddr{
			addr: val,
		})
	}
	ret = append(ret, toAdd)

	return ret
}

func problemOne(input []instruct) string {
	mem := memAddr{}
	for _, ins := range input {
		mask := ins.mask
		for _, w := range ins.writes {
			for addr, val := range w {
				_, ok := mem[addr]
				if !ok {
					mem[addr] = 0
				}

				newVal := newValue(mask, val)
				mem[addr] = newVal
			}
		}
	}

	ret := uint64(0)
	for _, v := range mem {
		ret += v
	}

	return fmt.Sprintf("%d", ret)
}

func getAllCombo(n int) []string {
	ret := []string{}
	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		bin := strconv.FormatUint(uint64(i), 2)
		binPad := fmt.Sprintf("%0"+fmt.Sprintf("%d", n)+"s", bin)
		ret = append(ret, binPad)
	}
	return ret
}

func newAddr(m [36]rune, addr uint64) []uint64 {
	nX := 0
	for _, c := range m {
		if c == 'X' {
			nX++
		}
	}
	combos := getAllCombo(nX)
	addrBin := fmt.Sprintf("%036s", strconv.FormatUint(addr, 2))
	coll := [][36]rune{}
	for _, c := range combos {
		newBin := [36]rune{}
		count := 0
		for i, mC := range m {
			if mC == 'X' {
				newBin[i] = rune(c[count])
				count++
				continue
			}

			if mC == '1' {
				newBin[i] = mC
				continue
			}

			newBin[i] = rune(addrBin[i])
		}
		coll = append(coll, newBin)
	}

	ret := []uint64{}
	for _, a := range coll {
		var s strings.Builder
		for _, r := range a {
			s.WriteString(string(r))
		}
		val, _ := strconv.ParseUint(s.String(), 2, 64)
		ret = append(ret, val)
	}
	return ret
}

func problemTwo(input []instruct) string {
	mem := memAddr{}
	for _, ins := range input {
		mask := ins.mask
		for _, w := range ins.writes {
			for addr, val := range w {
				allAddr := newAddr(mask, addr)
				for _, a := range allAddr {
					_, ok := mem[a]
					if !ok {
						mem[a] = uint64(0)
					}
					mem[a] = val
				}
			}
		}
	}

	ret := uint64(0)
	for _, v := range mem {
		ret += v
	}
	return fmt.Sprintf("%d", ret)
}
