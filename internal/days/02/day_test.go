package _2

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var example = `A Y
B X
C Z`

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "12586", day.SolvePart1())
}

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "13193", day.SolvePart2())
}

func TestSolve(t *testing.T) {
	tests := map[string]struct {
		lines string
		part1 bool
		want  string
	}{
		"Example": {
			lines: example,
			part1: true,
			want:  "15",
		},
		"Example Part 2": {
			lines: example,
			want:  "12",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := solve(strings.NewReader(test.lines), test.part1)
			if actual != test.want {
				t.Errorf("Got: %s, Want: %s.", actual, test.want)
			}
		})
	}
}
