package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day11 struct {
	monkeysMap  map[int]*monkey
	monkeysList []*monkey
}

func (*Day11) Day() int     { return 11 }
func (*Day11) IsTest() bool { return true }

func (puzzle *Day11) Run(input string) {
	puzzle.parseInput(input)

	puzzle.part1()
	puzzle.part2()
}

func (puzzle *Day11) part1() {
	reliefFunc := func(worryLevel, _ int) int {
		return (int)(math.Floor((float64)(worryLevel) / 3.0))
	}

	for i := 0; i < 20; i++ {
		puzzle.playOneRound(reliefFunc)
	}

	fmt.Printf("Part1: %d\n", puzzle.computeMonkeyBusiness())
}

func (puzzle *Day11) part2() {
	reliefFunc := func(worryLevel, divisibleBy int) int {
		panic("Not implemented yet.")
	}

	for _, monkey := range puzzle.monkeysList {
		monkey.reset()
	}

	for i := 0; i < 10_000; i++ {
		if i == 20 || i%1000 == 0 {
			fmt.Print("") // For breakpoint.
		}
		puzzle.playOneRound(reliefFunc)
	}

	fmt.Printf("Part2: %d\n", puzzle.computeMonkeyBusiness())
}

func (puzzle *Day11) computeMonkeyBusiness() int {
	inspectionCounts := []int{}

	for _, monkey := range puzzle.monkeysList {
		inspectionCounts = append(inspectionCounts, monkey.inspectionsCount)
	}

	sort.Slice(inspectionCounts, func(i, j int) bool {
		return inspectionCounts[i] > inspectionCounts[j]
	})

	return inspectionCounts[0] * inspectionCounts[1]
}

func (puzzle *Day11) playOneRound(reliefFunc reliefFunc) {
	for _, monkey := range puzzle.monkeysList {
		monkey.playTurn(reliefFunc)
	}
}

func (puzzle *Day11) parseInput(input string) {
	var currentMonkey *monkey

	puzzle.monkeysMap = map[int]*monkey{}
	puzzle.monkeysList = []*monkey{}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			currentMonkey = nil
			continue
		}

		if strings.HasPrefix(line, "Monkey ") {
			currentMonkey = puzzle.parseMonkey(line)
			puzzle.monkeysMap[currentMonkey.id] = currentMonkey
			puzzle.monkeysList = append(puzzle.monkeysList, currentMonkey)
		} else if strings.HasPrefix(line, "Starting items: ") {
			puzzle.parseStartingItems(currentMonkey, line)
		} else if strings.HasPrefix(line, "Operation: ") {
			puzzle.parseOperation(currentMonkey, line)
		} else if strings.HasPrefix(line, "Test: ") {
			puzzle.parseTest(currentMonkey, line)
		} else if strings.HasPrefix(line, "If ") {
			puzzle.parseTestAction(currentMonkey, line)
		}
	}
}

func (puzzle *Day11) parseMonkey(line string) *monkey {
	regex, _ := regexp.Compile(`Monkey (\d+):`)

	matches := regex.FindSubmatch(([]byte)(line))

	id, _ := strconv.Atoi((string)(matches[1]))

	return &monkey{
		id:                    id,
		originalStartingItems: []int{},
	}
}

func (puzzle *Day11) parseStartingItems(currentMonkey *monkey, line string) {
	listStr := strings.Split(line[16:], ", ")

	for _, listItemStr := range listStr {
		worryLevel, _ := strconv.Atoi(listItemStr)
		currentMonkey.originalStartingItems = append(currentMonkey.originalStartingItems, worryLevel)
		currentMonkey.reset()
	}
}

func (puzzle *Day11) parseOperation(currentMonkey *monkey, line string) {
	line = line[17:]

	parts := strings.Split(line, " ")

	lhs := parts[0]
	op := parts[1]
	rhs := parts[2]

	currentMonkey.operation = &operationInfo{
		lhs:      lhs,
		operator: operators[op],
		rhs:      rhs,
	}
}

func (puzzle *Day11) parseTest(currentMonkey *monkey, line string) {
	line = line[6:]

	regex, _ := regexp.Compile(`divisible by (\d+)`)
	matches := regex.FindSubmatch(([]byte)(line))

	divisibleBy, err := strconv.Atoi((string)(matches[1]))

	if err != nil {
		panic(err)
	}

	currentMonkey.test = &test{
		divisibleBy: divisibleBy,
		condition: func(input int) bool {
			return input%divisibleBy == 0
		},
	}
}

func (puzzle *Day11) parseTestAction(currentMonkey *monkey, line string) {
	regex, _ := regexp.Compile(`If (true|false): throw to monkey (\d+)`)

	matches := regex.FindSubmatch(([]byte)(line))

	targetMonkeyId, _ := strconv.Atoi((string)(matches[2]))

	action := func(itemWorryLevel int) {
		currentMonkey.startingItems = currentMonkey.startingItems[1:]
		puzzle.monkeysMap[targetMonkeyId].startingItems = append(puzzle.monkeysMap[targetMonkeyId].startingItems, itemWorryLevel)
	}

	if (string)(matches[1]) == "true" {
		currentMonkey.test.onTrueAction = action
	} else if (string)(matches[1]) == "false" {
		currentMonkey.test.onFalseAction = action
	}
}

var operators = map[string]operatorFunc{
	"+": addOperator,
	"*": mulOperator,
}

func addOperator(lhs, rhs int) int {
	return lhs + rhs
}

func mulOperator(lhs, rhs int) int {
	return lhs * rhs
}

type operatorFunc func(lhs, rhs int) int
type reliefFunc func(worryLevel, disibleBy int) int

type operationInfo struct {
	operator operatorFunc
	lhs      string
	rhs      string
}

type test struct {
	divisibleBy   int
	condition     func(input int) bool
	onTrueAction  func(itemWorryLevel int)
	onFalseAction func(itemWorryLevel int)
}

type monkey struct {
	id                    int
	originalStartingItems []int
	startingItems         []int
	operation             *operationInfo
	test                  *test
	inspectionsCount      int
}

func (m *monkey) inspectItem() int {
	var lhs int
	var rhs int

	itemWorryLevel := m.startingItems[0]
	m.inspectionsCount += 1

	if m.operation.lhs == "old" {
		lhs = itemWorryLevel
	} else {
		lhs, _ = strconv.Atoi(m.operation.lhs)
	}

	if m.operation.rhs == "old" {
		rhs = itemWorryLevel
	} else {
		rhs, _ = strconv.Atoi(m.operation.rhs)
	}

	return m.operation.operator(lhs, rhs)
}

func (m *monkey) throwItem(itemWorryLevel int) {
	if m.test.condition(itemWorryLevel) {
		m.test.onTrueAction(itemWorryLevel)
	} else {
		m.test.onFalseAction(itemWorryLevel)
	}
}

func (m *monkey) playTurn(reliefFunc reliefFunc) {
	for len(m.startingItems) > 0 {
		worryLevel := m.inspectItem()

		worryLevel = reliefFunc(worryLevel, m.test.divisibleBy)

		m.throwItem(worryLevel)
	}
}

func (m *monkey) reset() {
	m.inspectionsCount = 0
	m.startingItems = append([]int{}, m.originalStartingItems...)
}
