package day13

import (
	"aoc/2022/shared"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Puzzle struct{}

func (*Puzzle) Day() int     { return 13 }
func (*Puzzle) IsTest() bool { return false }

func (puzzle *Puzzle) Run(input string) {
	packets := parseInput(input)

	part1(packets)
	part2(packets)
}

func part1(packets []*packetElement) {
	total := 0

	var left *packetElement
	var right *packetElement

	for i, packet := range packets {

		if left == nil {
			left = packet
			continue
		} else {
			right = packet
		}

		cmp := left.compareTo(right)

		if cmp < 0 {
			total += (i / 2) + 1
		}

		left = nil
		right = nil
	}

	fmt.Printf("Part1: %d\n", total)
}

func part2(packets []*packetElement) {
	packets = append(
		packets,
		createDividerPacket(2),
		createDividerPacket(6),
	)

	sort.SliceStable(packets, func(i, j int) bool {
		return packets[i].compareTo(packets[j]) < 0
	})

	total := 1

	for i, packet := range packets {
		if packet.isDivider {
			total *= i + 1
		}
	}

	fmt.Printf("Part2: %d\n", total)
}

func createDividerPacket(num int) *packetElement {
	return &packetElement{
		isDivider: true,
		list: []*packetElement{
			{
				list: []*packetElement{
					{
						isPrimitive: true,
						primitive:   num,
					},
				},
			},
		},
	}
}

func parseInput(input string) []*packetElement {
	result := []*packetElement{}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		result = append(result, parseLine(&line))
	}

	return result
}

func parseLine(line *string) *packetElement {
	if strings.HasPrefix(*line, "[") {
		*line = (*line)[1:]
		list := []*packetElement{}
		for {
			if strings.HasPrefix(*line, "]") {
				*line = (*line)[1:]
				return &packetElement{isPrimitive: false, list: list}
			} else if strings.HasPrefix(*line, ",") {
				*line = (*line)[1:]
			} else {
				list = append(list, parseLine(line))
			}
		}
	} else {
		index := strings.IndexFunc(*line, func(c rune) bool { return !unicode.IsNumber(c) })
		num, _ := strconv.Atoi((*line)[:index])
		*line = (*line)[index:]
		return &packetElement{isPrimitive: true, primitive: num}
	}
}

type packetElement struct {
	isPrimitive bool
	isDivider   bool
	primitive   int
	list        []*packetElement
}

func (p *packetElement) compareTo(other *packetElement) int {
	if p.isPrimitive && other.isPrimitive {
		return p.primitive - other.primitive
	} else {
		return comparePacketElementLists(p.asList(), other.asList())
	}
}

func (p *packetElement) asList() []*packetElement {
	if p.isPrimitive {
		return []*packetElement{p}
	}
	return p.list
}

func (p *packetElement) String() string {
	if p.isPrimitive {
		return fmt.Sprint(p.primitive)
	}

	items := []string{}

	for _, item := range p.list {
		items = append(items, item.String())
	}

	return fmt.Sprintf("[%s]", strings.Join(items, ","))
}

func comparePacketElementLists(lhs []*packetElement, rhs []*packetElement) int {
	len1 := len(lhs)
	len2 := len(rhs)

	len := shared.Min(len1, len2)

	for i := 0; i < len; i++ {
		cmp := lhs[i].compareTo(rhs[i])
		if cmp != 0 {
			return cmp
		}
	}

	return len1 - len2
}
