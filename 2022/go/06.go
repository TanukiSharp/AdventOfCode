package main

import "fmt"

type Day6 struct{}

func (*Day6) Day() int {
	return 6
}

func (*Day6) IsTest() bool {
	return false
}

func (puzzle *Day6) Run(input string) {
	fmt.Printf("Part1: %d\n", puzzle.findStartOfPacket(input, 4))
	fmt.Printf("Part2: %d\n", puzzle.findStartOfPacket(input, 14))
}

func (puzzle *Day6) findStartOfPacket(input string, packetSize int) int {
	length := len(input) - packetSize

	for i := 0; i < length; i++ {
		slice := input[i : i+packetSize]
		if puzzle.isStartOfPacket(slice) {
			return i + packetSize
		}
	}

	return -1
}

func (puzzle *Day6) isStartOfPacket(packet string) bool {
	tester := map[rune]bool{}

	for _, c := range packet {
		if tester[c] {
			return false
		}
		tester[c] = true
	}

	return true
}
