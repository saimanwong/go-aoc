package day14

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

type Robot struct {
	P *toolbox.Coord
	V *toolbox.Coord
}

func (r *Robot) Move() {
	r.P.R = r.P.R + r.V.R
	r.P.C = r.P.C + r.V.C
}

func (p *Problem) Run() {
	robots := []*Robot{}
	for _, line := range p.input {
		spl := strings.Split(line, " ")
		r := &Robot{}
		p := strings.Split(strings.Split(spl[0], "=")[1], ",")
		r.P = &toolbox.Coord{R: toolbox.ToInt(p[1]), C: toolbox.ToInt(p[0])}
		v := strings.Split(strings.Split(spl[1], "=")[1], ",")
		r.V = &toolbox.Coord{R: toolbox.ToInt(v[1]), C: toolbox.ToInt(v[0])}
		robots = append(robots, r)
	}
	// bathroom := toolbox.NewByteMatrix(7, 11, '-') // example
	bathroom := toolbox.NewByteMatrix(103, 101, '-')

	newBathroom := bathroom.Copy()
	newRobots := []*Robot{}
	for _, r := range robots {
		newRobots = append(newRobots, &Robot{
			P: &toolbox.Coord{
				C: r.P.C,
				R: r.P.R,
			},
			V: &toolbox.Coord{
				C: r.V.C,
				R: r.V.R,
			},
		})
	}

	fmt.Println("Part 1:", p1(bathroom, robots))
	fmt.Println("Part 2:", p2(10_000, newBathroom, newRobots))
}

func p1(bathroom toolbox.ByteMatrix, robots []*Robot) int {
	const iter = 100
	for i := 0; i < iter; i++ {
		runSecond(bathroom, robots)
	}
	for _, robot := range robots {
		if bathroom[robot.P.R][robot.P.C] == '-' {
			bathroom[robot.P.R][robot.P.C] = '1'
			continue
		}
		bathroom[robot.P.R][robot.P.C]++
	}
	ret := 1
	for _, q := range [][]int{
		{0, bathroom.Height() / 2, 0, bathroom.Width() / 2},
		{0, bathroom.Height() / 2, bathroom.Width()/2 + 1, bathroom.Width()},
		{bathroom.Height()/2 + 1, bathroom.Height(), 0, bathroom.Width() / 2},
		{bathroom.Height()/2 + 1, bathroom.Height(), bathroom.Width()/2 + 1, bathroom.Width()},
	} {
		sum := 0
		for r := q[0]; r < q[1]; r++ {
			for c := q[2]; c < q[3]; c++ {
				if bathroom[r][c] == '-' {
					continue
				}
				sum += toolbox.ToInt(string(bathroom[r][c]))
			}
		}
		ret *= sum
	}
	return ret
}

func p2(iter int, bathroom toolbox.ByteMatrix, robots []*Robot) int {
	for i := 0; i < iter; i++ {
		tmp := bathroom.Copy()
		runSecond(tmp, robots)
		for _, r := range robots {
			if tmp[r.P.R][r.P.C] == '-' {
				tmp[r.P.R][r.P.C] = '1'
				continue
			}
			tmp[r.P.R][r.P.C]++
		}

		// catch the frame
		if strings.Contains(fmt.Sprintf("%s", tmp), "1 1 1 1 1 1 1 1 1 1 1") {
			return i + 1
		}
	}
	return 0
}

func runSecond(bathroom toolbox.ByteMatrix, robots []*Robot) {
	for _, robot := range robots {
		robot.Move()
		if bathroom.InsideCoord(robot.P) {
			continue
		}
		if robot.P.R < 0 {
			robot.P.R = bathroom.Height() + robot.P.R
		}
		if robot.P.R >= bathroom.Height() {
			robot.P.R = robot.P.R - bathroom.Height()
		}
		if robot.P.C < 0 {
			robot.P.C = bathroom.Width() + robot.P.C
		}
		if robot.P.C >= bathroom.Width() {
			robot.P.C = robot.P.C - bathroom.Width()
		}
	}
}
