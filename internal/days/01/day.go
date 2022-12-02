package _1

import (
	_ "embed"
	"fmt"
	"github.com/williamgcampbell/aoc2022/internal"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Solver struct{}

func (d *Solver) SolvePart1() string {
	r := strings.NewReader(input)
	lines := scanner.ScanLines(r)
	return strconv.Itoa(mostCalories(lines))
}

func mostCalories(lines []string) int {
	r := 0
	currentCalories := 0
	for _, line := range lines {
		if len(line) == 0 {
			if currentCalories > r {
				r = currentCalories
			}
			currentCalories = 0
		} else {
			lineInt, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("converting to int: %w", err)
				return 0
			}
			currentCalories += lineInt
		}
	}
	return r
}

func (d *Solver) SolvePart2() string {
	r := strings.NewReader(input)
	lines := scanner.ScanLines(r)
	return strconv.Itoa(topThreeCalories(lines))
}

func topThreeCalories(lines []string) int {
	first, second, third := 0, 0, 0

	currentCalories := 0
	for _, line := range lines {
		if len(line) == 0 {
			first, second, third = maxThree(first, second, third, currentCalories)
			currentCalories = 0
		} else {
			lineInt := internal.MustAtoI(line)
			currentCalories += lineInt
		}
	}
	first, second, third = maxThree(first, second, third, currentCalories)
	return first + second + third
}

func maxThree(f, s, t, d int) (int, int, int) {
	if d > f {
		temp := f
		f = d
		d = temp
	}

	if d > s {
		temp := s
		s = d
		d = temp
	}

	if d > t {
		temp := s
		t = d
		d = temp
	}
	return f, s, t
}
