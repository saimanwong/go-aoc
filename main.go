package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/saimanwong/go-aoc/problem"
	"github.com/saimanwong/go-aoc/problem/year2020"
)

var problems map[problem.Year]problem.Problems = map[problem.Year]problem.Problems{
	"2020": year2020.GetAllProblems(),
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("exactly 3 arguments required")
		os.Exit(1)
	}

	year := os.Args[1]
	re := regexp.MustCompile("^[0-9]{4}$")
	if !re.MatchString(year) {
		fmt.Println("first argument must be a valid year")
		os.Exit(2)
	}

	day := os.Args[2]
	re = regexp.MustCompile("^[0-9]{2}$")
	if !re.MatchString(day) {
		fmt.Println("second argument must be a valid day")
		os.Exit(3)
	}

	inputFile := os.Args[3]
	re = regexp.MustCompile("^[a-z0-9]+$")
	if !re.MatchString(inputFile) {
		fmt.Println("third argument must only contain alphanumeric characters")
		os.Exit(4)
	}

	probY, ok := problems[problem.Year(year)]
	if !ok {
		fmt.Println("year", year, "does not exist")
		os.Exit(1)
	}
	prob, ok := probY[problem.Day(day)]
	if !ok {
		fmt.Println("day", day, "does not exist")
		os.Exit(1)
	}
	path := filepath.Join(
		"problem",
		fmt.Sprintf("year%s", year),
		fmt.Sprintf("day%s", day),
		"inputs",
		inputFile,
	)
	if err := parse(path, prob); err != nil {
		fmt.Printf("parse fail: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("===== %s-12-%s %s =====\n", year, day, inputFile)
	prob.Run()
}

func parse(path string, prob problem.Problemer) error {
	var lines []string
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	prob.SetInput(lines)
	return nil
}
