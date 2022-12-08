package main

import (
	"fmt"
	"strings"
)

type Day8 struct {
	treeMap [][]int8
}

func (*Day8) Day() int {
	return 8
}

func (*Day8) IsTest() bool {
	return false
}

func (puzzle *Day8) Run(input string) {
	puzzle.parseInput(input)

	puzzle.part1()
	puzzle.part2()
}

func (puzzle *Day8) part1() {
	height := len(puzzle.treeMap)
	width := len(puzzle.treeMap[0])

	visibilityMap := puzzle.newSet()

	for x := 0; x < width; x++ {
		puzzle.scan(x, 0, height, 0, +1, visibilityMap)
		puzzle.scan(x, height-1, height, 0, -1, visibilityMap)
	}

	for y := 0; y < height; y++ {
		puzzle.scan(0, y, width, +1, 0, visibilityMap)
		puzzle.scan(width-1, y, width, -1, 0, visibilityMap)
	}

	fmt.Printf("Part1: %d\n", visibilityMap.size())
}

func (puzzle *Day8) part2() {
	bestViewingDistance := 0

	height := len(puzzle.treeMap)
	width := len(puzzle.treeMap[0])

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			viewingDistance := 1
			viewingDistance *= puzzle.computeViewingDistance(x, y, width, height, -1, 0)
			viewingDistance *= puzzle.computeViewingDistance(x, y, width, height, +1, 0)
			viewingDistance *= puzzle.computeViewingDistance(x, y, width, height, 0, -1)
			viewingDistance *= puzzle.computeViewingDistance(x, y, width, height, 0, +1)

			if viewingDistance > bestViewingDistance {
				bestViewingDistance = viewingDistance
			}
		}
	}

	fmt.Printf("Part2: %d\n", bestViewingDistance)
}

func (puzzle *Day8) computeViewingDistance(x, y, width, height, dx, dy int) int {
	viewingDistance := 0
	consideringTreeHeight := puzzle.treeMap[y][x]

	for {
		x += dx
		y += dy

		if x < 0 || y < 0 || x >= width || y >= height {
			break
		}

		scanningTreeHeight := puzzle.treeMap[y][x]
		viewingDistance++

		if scanningTreeHeight >= consideringTreeHeight {
			break
		}
	}

	return viewingDistance
}

func (puzzle *Day8) scan(startX, startY, count, dx, dy int, visibilityMap set) {
	var largetTreeHeight int8 = -1

	x := startX
	y := startY

	for count > 0 {
		treeHeight := puzzle.treeMap[y][x]
		if treeHeight > largetTreeHeight {
			visibilityMap.add(x, y)
			largetTreeHeight = treeHeight
		}
		if treeHeight == 9 {
			break
		}
		count--
		x += dx
		y += dy
	}
}

func (puzzle *Day8) parseInput(input string) {
	treeMap := [][]int8{}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		treeMapLine := []int8{}

		for _, c := range line {
			treeMapLine = append(treeMapLine, (int8)(c-'0'))
		}

		treeMap = append(treeMap, treeMapLine)
	}

	puzzle.treeMap = treeMap
}

type set struct {
	data map[string]struct{}
}

func (*Day8) newSet() set {
	return set{
		data: map[string]struct{}{},
	}
}

func (s *set) add(x, y int) {
	key := fmt.Sprintf("%d-%d", x, y)
	s.data[key] = struct{}{}
}

func (s *set) size() int {
	return len(s.data)
}
