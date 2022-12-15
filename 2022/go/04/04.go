package day4

import (
	"aoc/2022/shared"
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct{}

func (*Puzzle) Day() int     { return 4 }
func (*Puzzle) IsTest() bool { return false }

func (*Puzzle) Run(input string) {
	assignmentPairList := createAssignmentPairList(input)

	part1(assignmentPairList)
	part2(assignmentPairList)
}

func part1(assignmentPairList []pair) {
	result := 0

	for _, pair := range assignmentPairList {
		if pair.lhs.isFullyContaining(pair.rhs) || pair.rhs.isFullyContaining(pair.lhs) {
			result++
		}
	}

	fmt.Printf("Part1: %d\n", result)
}

func part2(assignmentPairList []pair) {
	result := 0

	for _, pair := range assignmentPairList {
		if pair.lhs.isFullyContaining(pair.rhs) ||
			pair.rhs.isFullyContaining(pair.lhs) ||
			pair.lhs.isOverlapping(pair.rhs) {
			result++
		}
	}

	fmt.Printf("Part2: %d\n", result)
}

type assignment struct {
	start int
	end   int
}

type pair struct {
	lhs *assignment
	rhs *assignment
}

func newAssignment(pair string) *assignment {
	parts := strings.Split(strings.TrimSpace(pair), "-")

	start, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	end, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

	return &assignment{
		start: start,
		end:   end,
	}
}

func (a *assignment) isFullyContaining(b *assignment) bool {
	return b.start >= a.start && b.end <= a.end
}

func (a *assignment) isOverlapping(b *assignment) bool {
	start := shared.Max(a.start, b.start)
	end := shared.Min(a.end, b.end)

	return start <= end
}

func createAssignmentPairList(input string) []pair {
	pairs := []pair{}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")

		pairs = append(pairs, pair{
			lhs: newAssignment(parts[0]),
			rhs: newAssignment(parts[1]),
		})
	}

	return pairs
}
