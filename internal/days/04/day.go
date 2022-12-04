package _4

import (
	_ "embed"
	"github.com/williamgcampbell/aoc2022/internal"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"io"
	"strconv"
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
	var fullyContained int
	for _, line := range lines {
		pair := strings.Split(line, ",")
		a := ParseAssignment(pair[0])
		b := ParseAssignment(pair[1])
		if part1 {
			if FullyOverlap(a, b) {
				fullyContained += 1
			}
		} else {
			if PartiallyOverlap(a, b) {
				fullyContained += 1
			}
		}
	}
	return strconv.Itoa(fullyContained)
}

type Assignment struct {
	beginningID int
	endID       int
}

func PartiallyOverlap(a *Assignment, b *Assignment) bool {
	return (a.beginningID <= b.beginningID && a.endID >= b.beginningID) ||
		(b.beginningID <= a.beginningID && b.endID >= a.beginningID)
}

func FullyOverlap(a *Assignment, b *Assignment) bool {
	return (a.beginningID <= b.beginningID && a.endID >= b.endID) ||
		(a.beginningID >= b.beginningID && a.endID <= b.endID)
}

func ParseAssignment(str string) *Assignment {
	a := strings.Split(str, "-")
	return &Assignment{
		beginningID: internal.MustAtoI(a[0]),
		endID:       internal.MustAtoI(a[1]),
	}
}
