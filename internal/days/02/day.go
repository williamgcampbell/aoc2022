package _2

import (
	_ "embed"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	OpponentRock     = "A"
	OpponentPaper    = "B"
	OpponentScissors = "C"

	MeRock     = "X"
	MePaper    = "Y"
	MeScissors = "Z"

	ShouldLose = "X"
	ShouldDraw = "Y"
	ShouldWin  = "Z"
)

type Solver struct{}

func (d *Solver) SolvePart1() string {
	r := strings.NewReader(input)
	lines := scanner.ScanLines(r)
	return strconv.Itoa(totalScore(lines, false))
}

func totalScore(rounds []string, isPartTwo bool) int {
	score := 0
	for _, round := range rounds {
		code := strings.Split(round, " ")
		if isPartTwo {
			// the two inputs represent the opponent's play and what the outcome of the game should be respectively
			opponent := code[0]
			me := myPlay(opponent, code[1])
			score += roundScore(opponent, me)
		} else {
			// the two inputs represent the opponent's play and your play respectively
			score += roundScore(code[0], code[1])
		}
	}
	return score
}

// myPlay returns a string letter X for Rock, Y for Paper, and Z for Scissors to represent what your play will be
// based on the opponents play and the desired outcome.
//
// Opponents play will be A for Rock, B for Paper, and Z for Scissors
func myPlay(opponentPlay string, desiredOutcome string) string {
	if opponentPlay == OpponentRock {
		if desiredOutcome == ShouldLose {
			return MeScissors
		}
		if desiredOutcome == ShouldWin {
			return MePaper
		}
		return MeRock
	}

	if opponentPlay == OpponentPaper {
		if desiredOutcome == ShouldLose {
			return MeRock
		}
		if desiredOutcome == ShouldWin {
			return MeScissors
		}
		return MePaper
	}

	if desiredOutcome == ShouldLose {
		return MePaper
	}
	if desiredOutcome == ShouldWin {
		return MeRock
	}
	return MeScissors
}

func roundScore(opponent, me string) int {
	score := outcomeScore(opponent, me)
	if me == MeRock {
		score += 1
	} else if me == MePaper {
		score += 2
	} else if me == MeScissors {
		score += 3
	}
	return score
}

func outcomeScore(opponent, me string) int {
	if isWin(opponent, me) {
		return 6
	}

	if isTie(opponent, me) {
		return 3
	}

	return 0
}

func isWin(opponent, me string) bool {
	return opponent == OpponentRock && me == MePaper || opponent == OpponentPaper && me == MeScissors || opponent == OpponentScissors && me == MeRock
}

func isTie(opponent, me string) bool {
	return opponent == OpponentRock && me == MeRock || opponent == OpponentPaper && me == MePaper || opponent == OpponentScissors && me == MeScissors
}

func (d *Solver) SolvePart2() string {
	r := strings.NewReader(input)
	lines := scanner.ScanLines(r)
	return strconv.Itoa(totalScore(lines, true))
}
