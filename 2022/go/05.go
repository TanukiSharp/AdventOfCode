package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Day5 struct {
	stacks    []*stack
	nStacks   []*nStack
	movements []*movement
}

func (*Day5) Day() int {
	return 5
}

func (puzzle *Day5) Run(input string) {
	puzzle.parseInput(input)
	puzzle.part1()
	puzzle.part2()
}

func (puzzle *Day5) part1() {
	for _, movement := range puzzle.movements {
		for i := 0; i < movement.quantity; i++ {
			c := puzzle.stacks[movement.fromIndex].pop()
			puzzle.stacks[movement.toIndex].push(c)
		}
	}

	fmt.Print("Part1: ")

	for _, stack := range puzzle.stacks {
		fmt.Printf("%c", stack.peek())
	}

	fmt.Println()
}

func (puzzle *Day5) part2() {
	for _, movement := range puzzle.movements {
		c := puzzle.nStacks[movement.fromIndex].pop(movement.quantity)
		puzzle.nStacks[movement.toIndex].push(c)
	}

	fmt.Print("Part2: ")

	for _, stack := range puzzle.nStacks {
		fmt.Printf("%c", stack.peek(1)[0])
	}

	fmt.Println()
}

func (puzzle *Day5) parseInput(input string) {
	isParsingStacks := true

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimRightFunc(line, unicode.IsSpace)

		if line == "" {
			isParsingStacks = false
			continue
		}

		if isParsingStacks {
			puzzle.parseCrateLine(line)
		} else {
			puzzle.parseMovement(line)
		}
	}

	for _, stack := range puzzle.stacks {
		stack.isReversed = false
		puzzle.nStacks = append(puzzle.nStacks, puzzle.cloneStackToNStack(stack))
	}
}

func (puzzle *Day5) parseCrateLine(line string) {
	if strings.HasPrefix(line, " 1 ") {
		return
	}

	stackIndex := 0

	runes := ([]rune)(line)

	for len(runes) > 0 {
		if len(puzzle.stacks) <= stackIndex {
			puzzle.stacks = append(puzzle.stacks, &stack{isReversed: true})
		}

		if runes[0] == '[' && runes[2] == ']' {
			c := runes[1]
			puzzle.stacks[stackIndex].push(c)
		}

		runes = runes[Min(4, len(runes)):]
		stackIndex++
	}
}

type movement struct {
	quantity  int
	fromIndex int
	toIndex   int
}

func (puzzle *Day5) parseMovement(line string) {
	regex, _ := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
	matches := regex.FindSubmatch(([]byte)(line))

	quantity, _ := strconv.Atoi((string)(matches[1]))
	from, _ := strconv.Atoi((string)(matches[2]))
	to, _ := strconv.Atoi((string)(matches[3]))

	puzzle.movements = append(puzzle.movements, &movement{
		quantity:  quantity,
		fromIndex: from - 1,
		toIndex:   to - 1,
	})
}

type stack struct {
	items      []rune
	isReversed bool
}

func (stack *stack) push(c rune) {
	if stack.isReversed {
		stack.items = append([]rune{c}, stack.items...)
	} else {
		stack.items = append(stack.items, c)
	}
}

func (stack *stack) peek() rune {
	if stack.isReversed {
		return stack.items[0]
	}
	return stack.items[len(stack.items)-1]
}

func (stack *stack) pop() rune {
	c := stack.peek()
	if stack.isReversed {
		stack.items = stack.items[1:]
	} else {
		stack.items = stack.items[:len(stack.items)-1]
	}
	return c
}

type nStack struct {
	items []rune
}

func (*Day5) cloneStackToNStack(stack *stack) *nStack {
	newArray := make([]rune, len(stack.items))

	for i := 0; i < len(newArray); i++ {
		newArray[i] = stack.items[i]
	}

	return &nStack{items: newArray}
}

func (stack *nStack) push(c []rune) {
	stack.items = append(stack.items, c...)
}

func (stack *nStack) peek(n int) []rune {
	return stack.items[len(stack.items)-n:]
}

func (stack *nStack) pop(n int) []rune {
	result := stack.peek(n)
	stack.items = stack.items[:len(stack.items)-n]
	return result
}
