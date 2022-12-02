package _1

import (
	_ "embed"
	"github.com/williamgcampbell/aoc2022/internal"
	"github.com/williamgcampbell/aoc2022/internal/scanner"
	"io"
	"sort"
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
	calories := getCalories(lines)
	if part1 {
		return strconv.Itoa(calories[len(calories)-1])
	} else {
		return strconv.Itoa(calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3])
	}
}

// getCalories calculates the number of calories carried by each elf
func getCalories(lines []string) []int {
	var calories []int
	currentCalories := 0
	for _, line := range lines {
		if len(line) == 0 {
			calories = insertSorted(calories, currentCalories)
			currentCalories = 0
		} else {
			lineInt := internal.MustAtoI(line)
			currentCalories += lineInt
		}
	}
	calories = insertSorted(calories, currentCalories)
	return calories
}

// insertSorted inserts an int v into the array ss in ascending order
func insertSorted(ss []int, v int) []int {
	i := sort.SearchInts(ss, v)
	ss = append(ss, 0)
	copy(ss[i+1:], ss[i:])
	ss[i] = v
	return ss
}
