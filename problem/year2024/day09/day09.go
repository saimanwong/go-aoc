package day09

import (
	"fmt"
	"slices"
	"strings"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	input []string
}

func (p *Problem) SetInput(input []string) {
	p.input = input
}

func (p *Problem) Run() {
	nums := []int{}
	id := 0
	for idx, r := range p.input[0] {
		n := toolbox.ToInt(string(r))
		w := id
		if idx%2 == 1 {
			w = -1
		}
		for i := 0; i < n; i++ {
			nums = append(nums, w)
		}
		if w != -1 {
			id++
		}
	}
	// fmt.Println(len(p.input[0]))

	fmt.Println("Part 1:", calc(p1move(nums)))
	// debug(nums)
	p2num := p2move(nums)
	fmt.Println("Part 2:", calc(p2num))
}

func debug(nums []int) {
	b := strings.Builder{}
	for _, n := range nums {
		if n == -1 {
			b.WriteString(".")
			continue
		}
		b.WriteString(fmt.Sprintf("%d", n))
	}
	fmt.Println(b.String())
}

func p1move(nums []int) []int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	i, j := 0, len(tmp)-1
	for i < j {
		if tmp[i] != -1 {
			i++
			continue
		}
		if tmp[j] == -1 {
			j--
			continue
		}
		tmp = slices.Replace(tmp, i, i+1, tmp[j])
		tmp[j] = -1
		i++
	}
	return tmp[:j]
}

func calc(nums []int) uint64 {
	var ret uint64
	for idx, n := range nums {
		if n == -1 {
			continue
		}
		ret += uint64(idx) * uint64(n)
	}
	return ret
}

type Block struct {
	Start int
	End   int
	Len   int
	Val   int
}

func p2move(nums []int) []int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	blocks := []*Block{}
	curr := &Block{}
	for i, n := range tmp {
		if i == 0 {
			curr.Start = i
			curr.Val = n
			continue
		}
		if curr.Val != n {
			curr.End = i
			curr.Len = curr.End - curr.Start
			blocks = append(blocks, curr)
			curr = &Block{Start: i, Val: n}
		}
		if i == len(tmp)-1 {
			curr.End = i + 1
			curr.Len = curr.End - curr.Start
			blocks = append(blocks, curr)
		}
	}
	frees, drives := []*Block{}, []*Block{}
	for _, b := range blocks {
		if b.Val == -1 {
			frees = append(frees, b)
			continue
		}
		drives = append(drives, b)
	}
	for len(drives) > 0 {
		curr := drives[len(drives)-1]
		drives = drives[:len(drives)-1]
		for _, free := range frees {
			if free.Len < curr.Len {
				continue
			}
			// COME ON! gotta learn to read...
			// "If there is no span of free space to the left of a file that is large enough to fit the file, the file does not move."
			// https://www.reddit.com/r/adventofcode/comments/1ha7bab/comment/m16h5h9
			if free.End > curr.Start {
				break
			}
			for i := free.Start; i < free.Start+curr.Len; i++ {
				tmp[i] = curr.Val
			}
			for i := curr.Start; i < curr.End; i++ {
				tmp[i] = -1
			}
			free.Start = free.Start + curr.Len
			free.Len = free.Len - curr.Len
			break
		}
	}
	return tmp
}
