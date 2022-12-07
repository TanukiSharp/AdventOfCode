package main

import (
	"fmt"
	"strings"
)

type Day3 struct{}

func (*Day3) Day() int {
	return 3
}

func (*Day3) IsTest() bool {
	return false
}

func (puzzle *Day3) Run(input string) {
	rucksacks := puzzle.createRucksackList(input)

	puzzle.part1(rucksacks)
	puzzle.part2(rucksacks)
}

type rucksack struct {
	compartment1 string
	compartment2 string

	_ struct{}
}

type runeSet = map[rune]int

func (puzzle *Day3) part1(rucksacks []rucksack) {
	rucksackNumber := -1
	total := 0

	for _, rucksack := range rucksacks {
		rucksackNumber++

		duplicates := puzzle.findDuplicates(rucksack.compartment1, rucksack.compartment2)

		duplicateCount := len(duplicates)

		if duplicateCount > 1 {
			panic(fmt.Sprintf("Rucksack number %d has %d duplicates.", rucksackNumber, duplicateCount))
		}

		if duplicateCount == 0 {
			continue
		}

		total += puzzle.getItemPriority(puzzle.getDuplicateRune(duplicates))
	}

	fmt.Printf("Part1: %d\n", total)
}

func (puzzle *Day3) part2(rucksacks []rucksack) {
	priorities := 0

	for i := 0; i < len(rucksacks); i += 3 {
		group := rucksacks[i : i+3]
		badge := puzzle.findBadgeOfGroup(group)
		priorities += puzzle.getItemPriority(badge)
	}

	fmt.Printf("Part2: %d\n", priorities)
}

func (*Day3) getDuplicateRune(runeSet runeSet) rune {
	for c := range runeSet {
		return c
	}
	panic("Unreachable.")
}

func (*Day3) getItemPriority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return (int)(c-'a') + 1
	}
	if c >= 'A' && c <= 'Z' {
		return (int)(c-'A') + 27
	}
	panic(fmt.Sprintf("Unknown character %v.", c))
}

func (*Day3) findDuplicates(a, b string) runeSet {
	duplicates := runeSet{}

	for _, c1 := range a {
		for _, c2 := range b {
			if c1 == c2 {
				duplicates[c1] = 0
			}
		}
	}

	return duplicates
}

func (puzzle *Day3) findBadgeOfGroup(rucksacks []rucksack) rune {
	runeSets := make([]runeSet, 3)

	for i := 0; i < len(rucksacks); i++ {
		j := (i + 1) % 3
		concat1 := rucksacks[i].compartment1 + rucksacks[i].compartment2
		concat2 := rucksacks[j].compartment1 + rucksacks[j].compartment2

		runeSets[i] = puzzle.findDuplicates(concat1, concat2)
	}

	check := map[rune]int{}

	for i := 0; i < len(runeSets); i++ {
		for r := range runeSets[i] {
			check[r]++
		}
	}

	for r, count := range check {
		if count == len(rucksacks) {
			return r
		}
	}

	panic("Could not find common item.")
}

func (*Day3) createRucksackList(input string) []rucksack {
	lineNumber := -1
	result := []rucksack{}

	for _, line := range strings.Split(input, "\n") {
		lineNumber++

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if len(line)%2 != 0 {
			panic(fmt.Sprintf("Line %d has an odd number of characters (%d).", lineNumber, len(line)))
		}

		result = append(result, rucksack{
			compartment1: line[:len(line)/2],
			compartment2: line[len(line)/2:],
		})
	}

	return result
}
