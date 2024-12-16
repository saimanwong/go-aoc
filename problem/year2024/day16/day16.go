package day16

import (
	"fmt"
	"math"
	"slices"
	"sort"
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
	maze := toolbox.ToByteMatrix(p.input)
	fmt.Println("Part 1:", p1(maze))
	fmt.Println("Part 2:")
}

type Direction rune

const (
	UnknownDir Direction = 'X'
	UpDir                = 'U'
	RightDir             = 'R'
	DownDir              = 'D'
	LeftDir              = 'L'
)

type Step struct {
	Coord toolbox.Coord
	Dir   Direction
}

type Visited map[toolbox.Coord]any

type Distance map[toolbox.Coord]int

type Queue []Step

func p1(maze toolbox.ByteMatrix) int {
	start, end := maze.Find('S'), maze.Find('E')
	dist := Distance{}
	visit := Visited{}
	queue := Queue{}
	maze.LoopCoord(func(c *toolbox.Coord, curr rune) {
		dist[*c] = math.MaxInt
		dir := UnknownDir
		if *c == *start { // ofc, facing east, at the top of the page...
			dir = RightDir
		}
		queue = append(queue, Step{Coord: *c, Dir: dir})
	})
	dist[*start] = 0

	for len(queue) > 0 {
		curr := queue.GetShortestDistance(dist)
		if curr.Coord == *end {
			break
		}
		if visit.Visited(curr.Coord) {
			continue
		}
		visit[curr.Coord] = nil
		neighs := []Step{}
		for dir, c := range toolbox.Direction {
			neigh := Step{
				Coord: toolbox.Coord{R: curr.Coord.R + c.R, C: curr.Coord.C + c.C},
				Dir:   Direction(dir),
			}
			if !maze.InsideCoord(&neigh.Coord) || maze.GetVal(&neigh.Coord) == '#' || visit.Visited(neigh.Coord) {
				continue
			}
			queue.SetDirection(neigh, Direction(dir))
			neighs = append(neighs, neigh)
		}
		for _, n := range neighs {
			if curr.Dir == n.Dir || curr.Dir == UnknownDir { // not rotated
				dist[n.Coord] = min(dist[curr.Coord]+1, dist[n.Coord])
				continue
			}
			dist[n.Coord] = min(dist[curr.Coord]+1000+1, dist[n.Coord])
		}
	}
	return dist[*end]
}

func (q Queue) SetDirection(c Step, dir Direction) {
	for idx, s := range q {
		if s.Coord == c.Coord {
			q[idx].Dir = dir
			return
		}
	}
}

func (q Queue) GetShortestDistance(dist Distance) Step {
	shortestDist, shortestDistIdx := math.MaxInt, -1
	for idx, s := range q {
		if dist[s.Coord] < shortestDist {
			shortestDist = dist[s.Coord]
			shortestDistIdx = idx
		}
	}
	if shortestDistIdx == -1 {
		panic("no shortest")
	}
	step := q[shortestDistIdx]
	slices.Delete(q, shortestDistIdx, shortestDistIdx+1)
	return step
}

func (v Visited) Visited(c toolbox.Coord) bool {
	_, ok := v[c]
	return ok
}

func (d Distance) String() string {
	builder := strings.Builder{}
	keys := []toolbox.Coord{}
	for k := range d {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return d[keys[i]] < d[keys[j]]
	})
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf("%s: %d\n", k.String(), d[k]))
	}
	return builder.String()
}
