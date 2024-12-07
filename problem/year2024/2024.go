package year2024

import (
	"github.com/saimanwong/go-aoc/problem"
	"github.com/saimanwong/go-aoc/problem/year2024/day01"
	"github.com/saimanwong/go-aoc/problem/year2024/day02"
	"github.com/saimanwong/go-aoc/problem/year2024/day03"
	"github.com/saimanwong/go-aoc/problem/year2024/day04"
	"github.com/saimanwong/go-aoc/problem/year2024/day05"
	"github.com/saimanwong/go-aoc/problem/year2024/day06"
	"github.com/saimanwong/go-aoc/problem/year2024/day07"
)

func GetAllProblems() problem.Problems {
	return problem.Problems{
		"01": &day01.Problem{},
		"02": &day02.Problem{},
		"03": &day03.Problem{},
		"04": &day04.Problem{},
		"05": &day05.Problem{},
		"06": &day06.Problem{},
		"07": &day07.Problem{},
	}
}
