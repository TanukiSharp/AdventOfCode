package day6

import "fmt"

type Puzzle struct{}

func (*Puzzle) Day() int     { return 6 }
func (*Puzzle) IsTest() bool { return false }

func (*Puzzle) Run(input string) {
	fmt.Printf("Part1: %d\n", findStartOfPacket(input, 4))
	fmt.Printf("Part2: %d\n", findStartOfPacket(input, 14))
}

func findStartOfPacket(input string, packetSize int) int {
	length := len(input) - packetSize

	for i := 0; i < length; i++ {
		slice := input[i : i+packetSize]
		if isStartOfPacket(slice) {
			return i + packetSize
		}
	}

	return -1
}

func isStartOfPacket(packet string) bool {
	tester := map[rune]bool{}

	for _, c := range packet {
		if tester[c] {
			return false
		}
		tester[c] = true
	}

	return true
}
