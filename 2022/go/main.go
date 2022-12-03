package main

import (
	"fmt"
	"os"
)

type Puzzle interface {
	Day() int
	Run(input string)
}

func main() {
	var puzzle Puzzle = &Day2{}

	filename := fmt.Sprintf("../%02d.txt", puzzle.Day())

	bytes, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	input := (string)(bytes)

	puzzle.Run(input)
}
