package p2

import (
	"fmt"
)

const MATRIXSIZE = 12

type cube struct {
	x int
	y int
	z int
	w int
}

type matrix struct {
	values map[int]map[int]map[int]map[int]bool
	active []cube
	from   int
	to     int
}

func initMatrix(from int, to int) matrix {
	ret := matrix{
		values: map[int]map[int]map[int]map[int]bool{},
		from:   from,
		to:     to,
	}

	for x := from; x <= to; x++ {
		ret.values[x] = map[int]map[int]map[int]bool{}
		for y := from; y <= to; y++ {
			ret.values[x][y] = map[int]map[int]bool{}
			for z := from; z <= to; z++ {
				ret.values[x][y][z] = map[int]bool{}
				for w := from; w <= to; w++ {
					ret.values[x][y][z][w] = false
				}
			}
		}

	}

	return ret
}

func parseInput(input []string) matrix {
	m := initMatrix(-MATRIXSIZE, MATRIXSIZE)
	for y, l := range input {
		for x, c := range l {
			if c == '#' {
				m.values[x][y][0][0] = true
				m.active = append(m.active, cube{
					x: x,
					y: y,
					z: 0,
					w: 0,
				})
			}
		}
	}

	return m
}

func visual(m matrix, zFrom int, zTo int) {
	for w := zFrom; w <= zTo; w++ {
		for z := zFrom; z <= zTo; z++ {
			fmt.Println("z", z, "w", w)
			for y := m.from; y <= m.to; y++ {
				for x := m.from; x <= m.to; x++ {
					c := '.'
					if m.values[x][y][z][w] {
						c = '#'
					}
					fmt.Printf("%s ", string(c))
				}
				fmt.Printf("\n")
			}

		}
	}
}

func calcAdj(c cube) []cube {
	ret := []cube{}
	for w := -1; w <= 1; w++ {
		w0 := c.w + w
		for z := -1; z <= 1; z++ {
			z0 := c.z + z
			for y := -1; y <= 1; y++ {
				y0 := c.y + y
				for x := -1; x <= 1; x++ {
					x0 := c.x + x
					if x0 == c.x && y0 == c.y && z0 == c.z && w0 == c.w {
						continue
					}
					ret = append(ret, cube{
						x: x0,
						y: y0,
						z: z0,
						w: w0,
					})
				}
			}

		}
	}
	return ret
}

func checkActive(m matrix, act []cube, from int, to int) ([]cube, []cube) {
	keepActive := []cube{}
	potAdj := map[string]cube{}
	for _, c := range act {
		adj := calcAdj(c)
		n := 0
		for _, a := range adj {
			if m.values[a.x][a.y][a.z][a.w] {
				n++
				continue
			}

			potStr := fmt.Sprintf("%d %d %d", a.x, a.y, a.z, a.w)
			_, ok := potAdj[potStr]
			if !ok {
				potAdj[potStr] = a
			}
		}

		if n >= from && n <= to {
			keepActive = append(keepActive, c)
		}
	}

	retAdj := []cube{}
	for _, v := range potAdj {
		retAdj = append(retAdj, v)
	}
	return keepActive, retAdj
}

func problemTwo(m matrix) string {
	for cycle := 1; cycle <= 6; cycle++ {
		keepActive, potAdj := checkActive(m, m.active, 2, 3)
		keepActive2, _ := checkActive(m, potAdj, 3, 3)
		keepActive = append(keepActive, keepActive2...)

		for _, k := range m.active {
			m.values[k.x][k.y][k.z][k.w] = false
		}

		m.active = []cube{}

		for _, ka := range keepActive {
			m.values[ka.x][ka.y][ka.z][ka.w] = true
			m.active = append(m.active, cube{
				x: ka.x,
				y: ka.y,
				z: ka.z,
				w: ka.w,
			})
		}

		// visual(m, -cycle, cycle)
	}

	return fmt.Sprintf("%d", len(m.active))
}

func Run(input []string) {
	inputParsed := parseInput(input)
	fmt.Println("Answer 2: " + problemTwo(inputParsed))
}
