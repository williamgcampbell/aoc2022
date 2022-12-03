package _3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "7446", day.SolvePart1())
}

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "2646", day.SolvePart2())
}

func TestSolvePart1_samples(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
		part1 bool
	}{
		"Example": {
			input: example,
			want:  "157",
			part1: true,
		},
		"Example part 2": {
			input: example,
			want:  "70",
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
