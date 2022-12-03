package _3

import (
	_ "embed"
	"github.com/williamgcampbell/aoc2022/internal/intersect"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"io"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var alphebet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Solver struct{}

func (d *Solver) SolvePart1() string {
	return solve(strings.NewReader(input), true)
}

func (d *Solver) SolvePart2() string {
	return solve(strings.NewReader(input), false)
}

func solve(reader io.Reader, part1 bool) string {
	lines := scanner.ScanLines(reader)
	sum := 0
	if part1 {
		for _, line := range lines {
			c1 := line[:len(line)/2]
			c2 := line[len(line)/2:]
			d := intersect.StringUnique(c1, c2)[0]
			sum += priority(string(d))
		}
	} else {
		var d string
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			if len(d) == 0 {
				d = line
			}
			d = intersect.StringUnique(d, line)

			if (i+1)%3 == 0 {
				sum += priority(string(d[0]))
				d = ""
			}
		}
	}
	return strconv.Itoa(sum)
}

func dup(c1, c2 string) string {
	var dups string
	for _, c := range c1 {
		if strings.ContainsRune(c2, c) {
			dups += string(c)
		}
	}
	return dups
}

func priority(d string) int {
	return strings.Index(alphebet, d) + 1
}
