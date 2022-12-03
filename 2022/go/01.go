package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day1 struct{}

func (*Day1) Day() int {
	return 1
}

func (*Day1) Run(input string) {
	calories := createCaloriesList(input)

	if len(calories) < 3 {
		panic(fmt.Sprintf("Not enough entries, expected at least 3, got only %d.", len(calories)))
	}

	sort.Slice(calories, func(i1, i2 int) bool { return calories[i1] > calories[i2] })

	fmt.Printf("The Elf with the most calories has %d calories.\n", calories[0])
	fmt.Printf("The total calories of the top three Elves is %d calories.\n", calories[0]+calories[1]+calories[2])
}

func createCaloriesList(input string) []int {
	lineNumber := -1
	calories := []int{}
	totalCurrentCalories := 0

	for _, line := range strings.Split(input, "\n") {
		lineNumber++

		line = strings.TrimSpace(line)

		if line == "" {
			if totalCurrentCalories > 0 {
				calories = append(calories, totalCurrentCalories)
				totalCurrentCalories = 0
			}

			continue
		}

		currentCalories, err := strconv.Atoi(line)

		if err != nil {
			panic(fmt.Sprintf("Invalid calories entry at line %d: '%q'.", lineNumber, line))
		}

		totalCurrentCalories += currentCalories
	}

	if totalCurrentCalories > 0 {
		calories = append(calories, totalCurrentCalories)
	}

	return calories
}
