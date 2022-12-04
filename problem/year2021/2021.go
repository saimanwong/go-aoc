package year2021

import (
	"github.com/saimanwong/go-aoc/problem"
	"github.com/saimanwong/go-aoc/problem/year2021/day01"
	"github.com/saimanwong/go-aoc/problem/year2021/day02"
)

func GetAllProblems() problem.Problems {
	return problem.Problems{
		"01": &day01.Problem{},
		"02": &day02.Problem{},
	}
}
