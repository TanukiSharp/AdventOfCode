package main

import (
	"fmt"
	"os"
	"path/filepath"
)

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
	var puzzle Puzzle = &Day8{}

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

	puzzle.Run(input)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
