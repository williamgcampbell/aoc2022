package _2

import (
	_ "embed"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"io"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Outcome int64

const (
	OutcomeUndefined Outcome = iota
	Win
	Lose
	Tie
)

var outcomeMap = map[string]Outcome{
	"X": Lose,
	"Y": Tie,
	"Z": Win,
}

func ParseOutcomeString(str string) Outcome {
	c, ok := outcomeMap[str]
	if ok {
		return c
	}
	return OutcomeUndefined
}

type Move int64

const (
	MoveUndefined Move = iota - 1
	Rock
	Paper
	Scissors
)

var moveMap = map[string]Move{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var compare = []Outcome{
	0: Tie,
	1: Lose,
	2: Win,
}

func (m Move) LosesTo() Move {
	return (m + 1) % 3
}

func (m Move) Beats() Move {
	return (m + 2) % 3
}

func (m Move) Result(o Move) Outcome {
	return compare[((o-m)%3+3)%3]
}

func ParseMoveString(str string) Move {
	c, ok := moveMap[str]
	if ok {
		return c
	}
	return MoveUndefined
}

type Solver struct{}

func (d *Solver) SolvePart1() string {
	return solve(strings.NewReader(input), true)
}

func (d *Solver) SolvePart2() string {
	return solve(strings.NewReader(input), false)
}

func solve(reader io.Reader, part1 bool) string {
	lines := scanner.ScanLines(reader)
	score := totalScore(lines, part1)
	return strconv.Itoa(score)
}

func totalScore(rounds []string, part1 bool) int {
	score := 0
	for _, round := range rounds {
		code := strings.Split(round, " ")
		if part1 {
			// the two inputs represent the opponent's play and your play respectively
			score += roundScore(ParseMoveString(code[0]), ParseMoveString(code[1]))
		} else {
			// the two inputs represent the opponent's play and what the outcome of the game should be respectively
			opponent := ParseMoveString(code[0])
			me := getMyMove(opponent, ParseOutcomeString(code[1]))
			score += roundScore(opponent, me)
		}
	}
	return score
}

// getMyMove returns the Move that represent what your play will be.
// This is determined by taking the opponents move and the desired outcome.
//
// Opponents play will be A for Rock, B for Paper, and Z for Scissors
func getMyMove(opponentMove Move, desiredOutcome Outcome) Move {
	if desiredOutcome == Tie {
		return opponentMove
	}

	if desiredOutcome == Lose {
		return opponentMove.Beats()
	}

	return opponentMove.LosesTo()
}

func roundScore(opponent, me Move) int {
	score := outcomeScore(opponent, me)
	if me == Rock {
		score += 1
	} else if me == Paper {
		score += 2
	} else if me == Scissors {
		score += 3
	}
	return score
}

func outcomeScore(opponent, me Move) int {
	if me.Result(opponent) == Win {
		return 6
	}

	if me.Result(opponent) == Tie {
		return 3
	}

	return 0
}
