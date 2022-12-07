package _6

import (
	_ "embed"
	"fmt"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"io"
	"strings"
)

//go:embed input.txt
var input string

type Solver struct{}

func (d *Solver) SolvePart1() string {
	return solve(strings.NewReader(input), true)
}

func (d *Solver) SolvePart2() string {
	return solve(strings.NewReader(input), false)
}

func solve(reader io.Reader, part1 bool) string {
	lines := scanner.ScanLines(reader)
	var sos int
	if part1 {
		sos = findStartOfSomething(lines[0], 3)
	} else {
		sos = findStartOfSomething(lines[0], 13)
	}
	return fmt.Sprintf("%d", sos)
}

func findStartOfSomething(packet string, distinctCharacters int) int {
	lastFew := ""
	for i := 0; i < len(packet); i++ {
		chString := string(packet[i])
		lastIndex := strings.LastIndex(lastFew, chString)
		if lastIndex < 0 {
			if len(lastFew) == distinctCharacters {
				return i + 1
			} else {
				lastFew += chString
			}
		} else {
			lastFew = lastFew[lastIndex+1:] + chString
		}
	}
	return -1
}
