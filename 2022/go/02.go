package main

import (
	"fmt"
	"strings"
)

type Day2 struct {
}

type roundValues struct {
	opponent string
	me       string
}

type handShape int
type strategy int

const (
	drawHand     handShape = 0
	rockHand     handShape = 1
	paperHand    handShape = 2
	scissorsHand handShape = 3
)

const (
	winStrategy  strategy = 0
	drawStrategy strategy = 1
	loseStrategy strategy = 2
)

const (
	loseScore = 0
	drawScore = 3
	winScore  = 6
)

func (*Day2) Day() int {
	return 2
}

func (*Day2) IsTest() bool {
	return false
}

func (puzzle *Day2) Run(input string) {
	roundValues := puzzle.createRoundValuesList(input)

	puzzle.part1(roundValues)
	puzzle.part2(roundValues)
}

func (*Day2) getWinner(opponent, me handShape) handShape {
	if opponent == me {
		return drawHand
	}

	if opponent == rockHand && me == paperHand ||
		opponent == paperHand && me == scissorsHand ||
		opponent == scissorsHand && me == rockHand {
		return me
	}

	return opponent
}

func (*Day2) getWinHandShape(opponent handShape) handShape {
	if opponent == rockHand {
		return paperHand
	} else if opponent == paperHand {
		return scissorsHand
	}
	return rockHand
}

func (*Day2) getLoseHandShape(opponent handShape) handShape {
	if opponent == rockHand {
		return scissorsHand
	} else if opponent == paperHand {
		return rockHand
	}
	return paperHand
}

func (*Day2) convertHandShape(character string) handShape {
	if character == "A" || character == "X" {
		return rockHand
	} else if character == "B" || character == "Y" {
		return paperHand
	} else if character == "C" || character == "Z" {
		return scissorsHand
	}

	panic(fmt.Sprintf("Invalid round character %q", character))
}

func (*Day2) convertStrategy(character string) strategy {
	if character == "X" {
		return loseStrategy
	} else if character == "Y" {
		return drawStrategy
	} else if character == "Z" {
		return winStrategy
	}

	panic(fmt.Sprintf("Invalid round character %q", character))
}

func (puzzle *Day2) getHandShapeFromStrategy(opponent handShape, strategy strategy) handShape {
	if strategy == winStrategy {
		return puzzle.getWinHandShape(opponent)
	} else if strategy == drawStrategy {
		return opponent
	} else if strategy == loseStrategy {
		return puzzle.getLoseHandShape(opponent)
	}

	panic(fmt.Sprintf("Unknown strategy %v", strategy))
}

func (puzzle *Day2) getRoundScore(opponent, me handShape) int {
	winner := puzzle.getWinner(opponent, me)

	if opponent == me {
		return drawScore
	} else if winner == me {
		return winScore
	} else {
		return loseScore
	}
}

func (puzzle *Day2) part1(roundValues []roundValues) {
	score := 0

	for _, roundValue := range roundValues {
		opponent := puzzle.convertHandShape(roundValue.opponent)
		me := puzzle.convertHandShape(roundValue.me)

		score += puzzle.getRoundScore(opponent, me) + (int)(me)
	}

	fmt.Printf("Part1: score is %d\n", score)
}

func (puzzle *Day2) part2(roundValues []roundValues) {
	score := 0

	for _, roundValue := range roundValues {
		opponent := puzzle.convertHandShape(roundValue.opponent)
		strategy := puzzle.convertStrategy(roundValue.me)
		me := puzzle.getHandShapeFromStrategy(opponent, strategy)

		score += puzzle.getRoundScore(opponent, me) + (int)(me)
	}

	fmt.Printf("Part2: score is %d\n", score)
}

func (*Day2) createRoundValuesList(input string) []roundValues {
	result := []roundValues{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(strings.TrimSpace(line), " ")
		result = append(result, roundValues{opponent: parts[0], me: parts[1]})
	}

	return result
}
