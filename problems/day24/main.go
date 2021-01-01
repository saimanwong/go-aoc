package day24

import (
	"fmt"
)

type tile struct {
	moves []string
}

type hexGrid map[int]map[int]map[int]bool

type hexCoord struct {
	x, y, z int
}

var hexCoordMap map[string]hexCoord = map[string]hexCoord{
	"e": hexCoord{
		x: 1,
		y: -1,
		z: 0,
	},
	"se": hexCoord{
		x: 0,
		y: -1,
		z: 1,
	},
	"sw": hexCoord{
		x: -1,
		y: 0,
		z: 1,
	},
	"w": hexCoord{
		x: -1,
		y: 1,
		z: 0,
	},
	"nw": hexCoord{
		x: 0,
		y: 1,
		z: -1,
	},
	"ne": hexCoord{
		x: 1,
		y: 0,
		z: -1,
	},
}

func (g *hexGrid) insert(x, y, z int, val bool) {
	_, ok := (*g)[x]
	if !ok {
		(*g)[x] = map[int]map[int]bool{}
	}

	_, ok = (*g)[x][y]
	if !ok {
		(*g)[x][y] = map[int]bool{}
	}

	_, ok = (*g)[x][y][z]
	if !ok {
		(*g)[x][y][z] = false
	}

	(*g)[x][y][z] = val
}

func (g *hexGrid) exists(x, y, z int) bool {
	_, ok := (*g)[x]
	if !ok {
		return false
	}

	_, ok = (*g)[x][y]
	if !ok {
		return false
	}

	_, ok = (*g)[x][y][z]
	if !ok {
		return false
	}

	return true
}

func parseInput(input []string) []*tile {
	ret := []*tile{}
	for _, line := range input {
		t := &tile{
			moves: []string{},
		}

		for i := 0; i < len(line); i++ {
			curr := string(line[i])
			if i+1 >= len(line) {
				t.moves = append(t.moves, curr)
				continue
			}
			next := string(line[i+1])
			if (curr == "n" || curr == "s") && (next == "e" || next == "w") {
				t.moves = append(t.moves, curr+next)
				i++
				continue
			}

			t.moves = append(t.moves, curr)
		}

		ret = append(ret, t)
	}

	return ret
}

func (g *hexGrid) placeAll(input []*tile) {
	for _, t := range input {
		var x, y, z int
		for _, m := range t.moves {
			h := hexCoordMap[m]
			x += h.x
			y += h.y
			z += h.z
		}

		ok := g.exists(x, y, z)
		if !ok {
			g.insert(x, y, z, false)
		}

		if (*g)[x][y][z] {
			(*g)[x][y][z] = false
			continue
		}

		(*g)[x][y][z] = true
	}
}

func problemOne(input []*tile) string {
	g := &hexGrid{}
	g.placeAll(input)
	count := g.getTotBlack()
	return fmt.Sprintf("%d", count)
}

func (g *hexGrid) checkAdj(x, y, z int) ([]hexCoord, int) {
	potentials := []hexCoord{}
	allDir := []string{"e", "se", "sw", "w", "nw", "ne"}
	count := 0

	for _, dir := range allDir {
		h := hexCoordMap[dir]
		ok := (*g)[x+h.x][y+h.y][z+h.z]
		if ok {
			if (*g)[x+h.x][y+h.y][z+h.z] {
				count++
			}
			continue
		}

		potentials = append(potentials, hexCoord{
			x: x + h.x,
			y: y + h.y,
			z: z + h.z,
		})
		// g.insert(x+h.x, y+h.y, z+h.z, false)
	}

	return potentials, count
}

func (g *hexGrid) getTotBlack() int {
	count := 0
	for x, _ := range *g {
		for y, _ := range (*g)[x] {
			for z, _ := range (*g)[x][y] {
				if (*g)[x][y][z] {
					count++
				}
			}
		}
	}
	return count
}

func problemTwo(input []*tile, d int) string {
	g := &hexGrid{}
	g.placeAll(input)

	// Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
	// Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
	toWhite := []hexCoord{}
	toBlack := []hexCoord{}
	potentials := []hexCoord{}
	for i := 0; i < d; i++ {
		for x, _ := range *g {
			for y, _ := range (*g)[x] {
				for z, _ := range (*g)[x][y] {
					pot, nb := g.checkAdj(x, y, z)
					potentials = append(potentials, pot...)

					if (nb == 0 || nb > 2) && (*g)[x][y][z] {
						toWhite = append(toWhite, hexCoord{
							x: x,
							y: y,
							z: z,
						})
					}

					if nb == 2 && !(*g)[x][y][z] {
						toBlack = append(toBlack, hexCoord{
							x: x,
							y: y,
							z: z,
						})
					}
				}
			}
		}

		for _, pot := range potentials {
			x, y, z := pot.x, pot.y, pot.z
			ok := g.exists(x, y, z)
			if !ok {
				g.insert(x, y, z, false)
			}
			_, nb := g.checkAdj(x, y, z)
			if (nb == 0 || nb > 2) && (*g)[x][y][z] {
				toWhite = append(toWhite, hexCoord{
					x: x,
					y: y,
					z: z,
				})
			}

			if nb == 2 && !(*g)[x][y][z] {
				toBlack = append(toBlack, hexCoord{
					x: x,
					y: y,
					z: z,
				})
			}
		}

		for _, c := range toWhite {
			(*g)[c.x][c.y][c.z] = false
		}
		for _, c := range toBlack {
			(*g)[c.x][c.y][c.z] = true
		}

		toWhite = []hexCoord{}
		toBlack = []hexCoord{}
		potentials = []hexCoord{}
	}

	return fmt.Sprintf("%d", g.getTotBlack())
}

func Run(input []string) {
	inputParsed := parseInput(input)
	fmt.Println("Answer 1: " + problemOne(inputParsed))
	fmt.Println("Answer 2: " + problemTwo(inputParsed, 100))
}
