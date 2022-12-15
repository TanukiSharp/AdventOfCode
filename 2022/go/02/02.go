package day2

import (
	"fmt"
	"strings"
)

type Puzzle struct{}

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

func (*Puzzle) Day() int     { return 2 }
func (*Puzzle) IsTest() bool { return false }

func (*Puzzle) Run(input string) {
	roundValues := createRoundValuesList(input)

	part1(roundValues)
	part2(roundValues)
}

func getWinner(opponent, me handShape) handShape {
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

func getWinHandShape(opponent handShape) handShape {
	if opponent == rockHand {
		return paperHand
	} else if opponent == paperHand {
		return scissorsHand
	}
	return rockHand
}

func getLoseHandShape(opponent handShape) handShape {
	if opponent == rockHand {
		return scissorsHand
	} else if opponent == paperHand {
		return rockHand
	}
	return paperHand
}

func convertHandShape(character string) handShape {
	if character == "A" || character == "X" {
		return rockHand
	} else if character == "B" || character == "Y" {
		return paperHand
	} else if character == "C" || character == "Z" {
		return scissorsHand
	}

	panic(fmt.Sprintf("Invalid round character %q", character))
}

func convertStrategy(character string) strategy {
	if character == "X" {
		return loseStrategy
	} else if character == "Y" {
		return drawStrategy
	} else if character == "Z" {
		return winStrategy
	}

	panic(fmt.Sprintf("Invalid round character %q", character))
}

func getHandShapeFromStrategy(opponent handShape, strategy strategy) handShape {
	if strategy == winStrategy {
		return getWinHandShape(opponent)
	} else if strategy == drawStrategy {
		return opponent
	} else if strategy == loseStrategy {
		return getLoseHandShape(opponent)
	}

	panic(fmt.Sprintf("Unknown strategy %v", strategy))
}

func getRoundScore(opponent, me handShape) int {
	winner := getWinner(opponent, me)

	if opponent == me {
		return drawScore
	} else if winner == me {
		return winScore
	} else {
		return loseScore
	}
}

func part1(roundValues []roundValues) {
	score := 0

	for _, roundValue := range roundValues {
		opponent := convertHandShape(roundValue.opponent)
		me := convertHandShape(roundValue.me)

		score += getRoundScore(opponent, me) + (int)(me)
	}

	fmt.Printf("Part1: score is %d\n", score)
}

func part2(roundValues []roundValues) {
	score := 0

	for _, roundValue := range roundValues {
		opponent := convertHandShape(roundValue.opponent)
		strategy := convertStrategy(roundValue.me)
		me := getHandShapeFromStrategy(opponent, strategy)

		score += getRoundScore(opponent, me) + (int)(me)
	}

	fmt.Printf("Part2: score is %d\n", score)
}

func createRoundValuesList(input string) []roundValues {
	result := []roundValues{}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")
		result = append(result, roundValues{opponent: parts[0], me: parts[1]})
	}

	return result
}
