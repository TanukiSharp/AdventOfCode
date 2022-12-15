package main

import (
	"fmt"
	"os"
	"path/filepath"

	// day1 "aoc/2022/01"
	// day2 "aoc/2022/02"
	// day3 "aoc/2022/03"
	// day4 "aoc/2022/04"
	// day5 "aoc/2022/05"
	// day6 "aoc/2022/06"
	// day7 "aoc/2022/07"
	// day8 "aoc/2022/08"
	// day11 "aoc/2022/11"
	// day13 "aoc/2022/13"
	day14 "aoc/2022/14"
)

var puzzle Puzzle = &day14.Puzzle{}

// -------------------------------------------

type Puzzle interface {
	IsTest() bool
	Day() int
	Run(input string)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir() == false
}

func findFile(basePath string, filename string) string {
	for {
		absoluteFilename := filepath.Join(basePath, filename)

		if fileExists(absoluteFilename) {
			return absoluteFilename
		}

		prevBasePath := basePath
		basePath = filepath.Dir(basePath)

		if basePath == prevBasePath {
			panic(fmt.Sprintf("Could not find file %q anywhere.", filename))
		}
	}
}

func constructFilename(puzzle Puzzle) string {
	test := ""

	if puzzle.IsTest() {
		test = ".test"
	}

	return fmt.Sprintf("%02d%s.txt", puzzle.Day(), test)
}

func main() {
	pwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	filename := findFile(pwd, constructFilename(puzzle))

	bytes, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	input := (string)(bytes)

	fmt.Printf("--- Day %d ---\n", puzzle.Day())
	puzzle.Run(input)
}
