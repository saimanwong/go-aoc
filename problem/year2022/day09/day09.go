package day09

import (
	"fmt"
	"strings"

	"github.com/saimanwong/go-aoc/internal/toolbox"
)

type Problem struct {
	Room1 *room1
	Room2 *room2
	moves []move
	input []string
}

type move struct {
	dir rune // direction
	n   int  // steps
}

type room1 struct {
	Head *toolbox.Coord
	Tail *toolbox.Coord
}

type room2 struct {
	Head  *toolbox.Coord
	Tails [9]*toolbox.Coord
}

func (p *Problem) SetInput(input []string) {
	for _, l := range input {
		spl := strings.Split(l, " ")
		p.moves = append(p.moves, move{
			dir: rune(spl[0][0]),
			n:   toolbox.ToInt(spl[1]),
		})
	}
	p.Room1 = &room1{
		Head: &toolbox.Coord{},
		Tail: &toolbox.Coord{},
	}
	p.Room2 = &room2{
		Head: &toolbox.Coord{},
	}
	for i := range p.Room2.Tails {
		p.Room2.Tails[i] = &toolbox.Coord{}
	}
}

func (p *Problem) Run() {
	fmt.Println("Part 1:", p.p1())
	fmt.Println("Part 2:", p.p2())
}

func (p *Problem) p2() int {
	visitedCount := 1
	visited := map[string]struct{}{"0,0": {}}
	for _, m := range p.moves { // moves
		for i := 0; i < m.n; i++ { // steps
			p.Room2.Head.Move(m.dir)
			head := &toolbox.Coord{R: p.Room2.Head.R, C: p.Room2.Head.C}
			for tailIdx, tail := range p.Room2.Tails { // move accordingly
				newTail := follow(tail, head)
				if newTail == nil { // still connected
					break
				}
				p.Room2.Tails[tailIdx] = newTail
				head = p.Room2.Tails[tailIdx]
				if tailIdx == 8 {
					key := fmt.Sprintf("%d,%d", head.R, head.C)
					if _, ok := visited[key]; !ok {
						visitedCount++
						visited[key] = struct{}{}
					}
				}
			}
		}
	}
	return visitedCount
}

func follow(t, h *toolbox.Coord) *toolbox.Coord {
	dist := t.Distance(h)
	if dist < 1.5 {
		return nil
	}
	minDst := 100.0
	var retDst *toolbox.Coord
	for _, m := range []string{"U", "UR", "R", "DR", "D", "DL", "L", "UL"} {
		tmpDst := &toolbox.Coord{R: t.R, C: t.C}
		movement := []rune{}
		for _, r := range m {
			movement = append(movement, r)
		}
		tmpDst.Move(movement...)
		dist := tmpDst.Distance(h)
		if dist < minDst {
			minDst = dist
			retDst = tmpDst
		}
		if dist == 1 {
			return tmpDst
		}
	}
	return retDst
}

func (p *Problem) debug() {
	offset := 15
	n := 30
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if p.Room2.Head.R+offset == r && p.Room2.Head.C+offset == c {
				fmt.Printf("H ")
				continue
			}
			found := false
			for idx, t := range p.Room2.Tails {
				if t.R+offset == r && t.C+offset == c {
					fmt.Printf("%d ", idx+1)
					found = true
					break
				}
			}
			if !found {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("= ", n))
}

func (p *Problem) p1() int {
	visitedCount := 1
	visited := map[string]struct{}{"0,0": {}}
	for _, m := range p.moves {
		for i := 0; i < m.n; i++ {
			prevR, prevC := p.Room1.Head.R, p.Room1.Head.C
			p.Room1.Head.Move(m.dir)
			dist := p.Room1.Tail.Distance(p.Room1.Head)
			if !(dist < 1.5) {
				p.Room1.Tail = &toolbox.Coord{R: prevR, C: prevC}
				key := fmt.Sprintf("%d,%d", prevR, prevC)
				if _, ok := visited[key]; !ok {
					visitedCount++
					visited[key] = struct{}{}
				}
			}
		}
	}
	return visitedCount
}
