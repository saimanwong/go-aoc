package day12

import (
	"fmt"
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
	garden := toolbox.ToByteMatrix(p.input)
	regions := getRegions(garden)
	fmt.Println("Part 1:", p1(garden, regions))
	fmt.Println("Part 2:", p2(garden, regions))
}

type Plant struct {
	Coord *toolbox.Coord
	Sides int
}

func p1(garden toolbox.ByteMatrix, regions [][]*toolbox.Coord) int {
	sum := 0
	for _, region := range regions {
		sides := 0
		for _, plant := range region {
			for _, dir := range toolbox.DirectionSlice { // perimeter
				newPos := &toolbox.Coord{
					R: plant.R - dir.R,
					C: plant.C - dir.C,
				}
				if !garden.Inside(newPos.R, newPos.C) {
					sides++
					continue
				}
				if garden.GetVal(newPos) != garden.GetVal(plant) {
					sides++
				}
			}
		}
		sum += len(region) * sides
	}
	return sum
}

// not the prettiest...
func p2(garden toolbox.ByteMatrix, regions [][]*toolbox.Coord) int {
	masks := []toolbox.ByteMatrix{
		toolbox.ToByteMatrix([]string{
			"-#-",
			"#x-",
			"---",
		}),
		toolbox.ToByteMatrix([]string{
			"-#-",
			"-x#",
			"---",
		}),
		toolbox.ToByteMatrix([]string{
			"---",
			"-x#",
			"-#-",
		}),
		toolbox.ToByteMatrix([]string{
			"---",
			"#x-",
			"-#-",
		}),
		toolbox.ToByteMatrix([]string{
			"-x#",
			"-xx",
			"---",
		}),
		toolbox.ToByteMatrix([]string{
			"---",
			"-xx",
			"-x#",
		}),
		toolbox.ToByteMatrix([]string{
			"---",
			"xx-",
			"#x-",
		}),
		toolbox.ToByteMatrix([]string{
			"#x-",
			"xx-",
			"---",
		}),
	}
	sum := 0
	for _, plants := range regions {
		corner := 0
		for _, plant := range plants {
			for _, mask := range masks {
				str := []string{}
				builder := strings.Builder{}
				for r := plant.R - 1; r < plant.R+2; r++ {
					for c := plant.C - 1; c < plant.C+2; c++ {
						val := garden.GetVal(plant)
						if garden.Inside(r, c) && garden[r][c] == val {
							builder.WriteRune('x')
							continue
						}
						builder.WriteRune('#')
					}
					str = append(str, builder.String())
					builder.Reset()
				}
				subset := toolbox.ToByteMatrix(str)
				for r := 0; r < 3; r++ {
					for c := 0; c < 3; c++ {
						if mask[r][c] == '-' {
							subset[r][c] = '-'
						}
					}
				}
				if mask.Equal(subset) {
					corner++
				}
			}
		}
		sum += corner * len(plants)
	}
	return sum
}

func getRegions(garden toolbox.ByteMatrix) [][]*toolbox.Coord {
	visited := map[string]interface{}{}
	q := []*toolbox.Coord{}
	regions := [][]*toolbox.Coord{}
	garden.Loop(func(r, c int, val rune) {
		if _, ok := visited[fmt.Sprintf("%d,%d", r, c)]; ok {
			return
		}
		q = append(q, &toolbox.Coord{R: r, C: c})
		currRegion := []*toolbox.Coord{}
		for len(q) > 0 {
			curr := q[len(q)-1]
			q = q[:len(q)-1]
			if _, ok := visited[fmt.Sprintf("%s", curr)]; ok {
				continue
			}
			visited[fmt.Sprintf("%s", curr)] = nil
			currRegion = append(currRegion, curr)

			for _, dir := range toolbox.Direction {
				newPos := &toolbox.Coord{R: curr.R + dir.R, C: curr.C + dir.C}
				if !garden.Inside(newPos.R, newPos.C) {
					continue
				}
				if garden.GetVal(newPos) != garden.GetVal(curr) {
					continue
				}
				if _, ok := visited[fmt.Sprintf("%s", newPos)]; ok {
					continue
				}
				q = append(q, newPos)
			}
		}
		if len(currRegion) > 0 {
			regions = append(regions, currRegion)
		}
	})
	return regions
}
