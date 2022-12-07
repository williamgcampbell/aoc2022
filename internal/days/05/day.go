package _5

import (
	_ "embed"
	"github.com/williamgcampbell/aoc2022/internal"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"io"
	"regexp"
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

	drawing, index := getDrawing(lines)
	instructions := getInstructions(lines, index)
	for _, instruction := range instructions {
		if part1 {
			drawing = crateMover9000(drawing, instruction[0], instruction[1], instruction[2])
		} else {
			drawing = crateMover9001(drawing, instruction[0], instruction[1], instruction[2])
		}
	}

	r := ""
	for _, arr := range drawing {
		n := len(arr) - 1
		r += arr[n]
	}

	return r
}

var instructionRegex = regexp.MustCompile(`move (?P<count>\d+) from (?P<from>\d+) to (?P<to>\d+)`)

func getInstructions(lines []string, index int) [][]int {
	var r [][]int
	for i := index; i < len(lines); i++ {
		line := lines[i]
		matches := instructionRegex.FindStringSubmatch(line)
		r = append(r, []int{internal.MustAtoI(matches[1]), internal.MustAtoI(matches[2]), internal.MustAtoI(matches[3])})
	}
	return r
}

func getDrawing(lines []string) (drawing [][]string, inx int) {
	index := 0
	var stacks [][]string
	for i := index; i < len(lines); i++ {
		line := lines[i]
		index++
		if len(line) == 0 {
			// end of crate drawing
			break
		}

		row := getRow(line)
		for j, column := range row {

			if len(stacks) <= j {
				stacks = append(stacks, make([]string, 0))
			}

			crateStr := ""
			if strings.HasPrefix(column, "[") && strings.HasSuffix(column, "]") {
				crateStr = column[1:2]
			}

			if len(crateStr) == 0 {
				continue
			}

			// Push
			stacks[j] = append([]string{crateStr}, stacks[j]...)
		}
	}
	return stacks, index
}

// crateMover9000 will move [count] elements from [from] to [to ]of the internal arrays of [stack].
// stack should be arranged in order of left to right and from bottom to top.
//
// For example, the following array would represent a stack of elements with B stacked
// on top of A on the leftmost space, no elements in the middle space, and a single element
// D on the rightmost space.
// e.g. [[A, B], [], [D]]
//
// Moving two elements from space 1 to space 2 would result in.
// [[], [B, A], [D]]
func crateMover9000(stack [][]string, count, from, to int) [][]string {
	for i := 0; i < count; i++ {
		n := len(stack[from-1]) - 1
		el := stack[from-1][n]
		stack[from-1][n] = "" // Erase element (write zero value)
		stack[from-1] = stack[from-1][:n]
		stack[to-1] = append(stack[to-1], el)
	}
	return stack
}

func crateMover9001(stack [][]string, count, from, to int) [][]string {
	n := len(stack[from-1]) - 1
	crates := stack[from-1][n-count+1:]

	stack[from-1] = stack[from-1][:n-count+1]
	stack[to-1] = append(stack[to-1], crates...)
	return stack
}

func getRow(line string) []string {
	b, e := 0, 3
	var s []string
	for e <= len(line) {

		crate := line[b:e]
		s = append(s, crate)
		b += 4
		e += 4
	}
	return s
}
