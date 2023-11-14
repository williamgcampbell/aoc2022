package _8

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `30373
25512
65332
33549
35390`

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "", day.SolvePart1())
}

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "", day.SolvePart2())
}

func TestSolve(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
		part1 bool
	}{
		"Example": {
			input: example,
			want:  "21",
			part1: true,
		},
		"Example part 2": {
			input: example,
			want:  "",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := solve(strings.NewReader(test.input), test.part1)
			if actual != test.want {
				t.Errorf("Got: %s, Want: %s.", actual, test.want)
			}
		})
	}
}
