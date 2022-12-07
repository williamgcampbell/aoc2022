package _6

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "1723", day.SolvePart1())
}

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "3708", day.SolvePart2())
}

func TestSolve(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
		part1 bool
	}{
		"Example": {
			input: example,
			want:  "7",
			part1: true,
		},
		"Example 2": {
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  "5",
			part1: true,
		},
		"Example 3": {
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  "6",
			part1: true,
		},
		"Example 4": {
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  "10",
			part1: true,
		},
		"Example 5": {
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  "11",
			part1: true,
		},
		"Example part 2": {
			input: example,
			want:  "19",
		},
		"Example part 2 2": {
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  "23",
		},
		"Example part 2 3": {
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  "23",
		},
		"Example part 2 4": {
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  "29",
		},
		"Example part 2 5": {
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  "26",
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
