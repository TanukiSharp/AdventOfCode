package day14

import (
	"aoc/2022/shared"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Puzzle struct{}

func (*Puzzle) Day() int     { return 14 }
func (*Puzzle) IsTest() bool { return false }

func (*Puzzle) Run(input string) {
	rockTraces := parseInput(input)

	part1(rockTraces)
	part2(rockTraces)
}

func part1(rocks rockTraces) {
	fmt.Printf("Part1: %d\n", runPart1(false, rocks))
}

func part2(rocks rockTraces) {
	_, _, _, mapMaxY := findMapBoundingBox(rocks)

	mapMaxY += 2

	rockTrace := &rockTrace{}
	rockTrace.addCoord(coord{x: math.MinInt, y: mapMaxY})
	rockTrace.addCoord(coord{x: math.MaxInt, y: mapMaxY})
	rocks = append(rocks, rockTrace)

	fmt.Printf("Part2: %d\n", runPart2(rocks))
}

func runPart1(show bool, rocks rockTraces) int {
	sands := &sands{}

	count := 0

	for sands.emit(rocks) {
		for sands.update(rocks) {
			for _, grain := range sands.grains {
				if rocks.isBelow(grain.pos.y) {
					return count
				}
			}
			if show {
				print(rocks, sands)
				fmt.Printf("count: %d\n", count)
				fmt.Print("") // For breakpoint.
			}
		}

		count++

		if show {
			print(rocks, sands)
			fmt.Printf("count: %d\n", count)
			fmt.Print("") // For breakpoint.
		}
	}

	return -1
}

func runPart2(rocks rockTraces) int {
	sands := &sands{}

	count := 0

	for sands.emit(rocks) {
		for sands.update(rocks) {
		}

		count++
	}

	return count
}

type coord struct {
	x int
	y int
}

func parseInput(input string) rockTraces {
	rockTraces := rockTraces{}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			break
		}

		rockTraces = append(rockTraces, parseLine(line))
	}

	return rockTraces
}

func parseLine(line string) *rockTrace {
	coordinates := strings.Split(line, " -> ")

	rockTrace := &rockTrace{}

	for _, coordStr := range coordinates {
		coordParts := strings.Split(coordStr, ",")

		x, _ := strconv.Atoi(coordParts[0])
		y, _ := strconv.Atoi(coordParts[1])

		rockTrace.addCoord(coord{x: x, y: y})
	}

	return rockTrace
}

func print(traces rockTraces, sands *sands) {
	sb := strings.Builder{}

	minX, minY, maxX, maxY := findMapBoundingBox(traces)

	width := maxX - minX + 1
	height := maxY - minY + 1

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rx := x + minX
			ry := y + minY

			if rx == 500 && ry == 0 {
				sb.WriteRune('+')
			} else if traces.isOn(rx, ry) {
				sb.WriteRune('#')
			} else if sands.isOnAny(rx, ry) {
				sb.WriteRune('O')
			} else {
				sb.WriteRune('.')
			}
		}
		fmt.Println(sb.String())
		sb.Reset()
	}
}

func findMapBoundingBox(traces rockTraces) (minX, minY, maxX, maxY int) {
	minX = math.MaxInt
	minY = math.MaxInt
	maxX = math.MinInt
	maxY = math.MinInt

	for _, trace := range traces {
		for _, coord := range trace.coords {
			minX = shared.Min(minX, coord.x)
			minY = shared.Min(minY, coord.y)
			maxX = shared.Max(maxX, coord.x)
			maxY = shared.Max(maxY, coord.y)
		}
	}

	minX = shared.Min(minX, 500)
	minY = shared.Min(minY, 0)
	maxX = shared.Max(maxX, 500)
	maxY = shared.Max(maxY, 0)

	return
}
