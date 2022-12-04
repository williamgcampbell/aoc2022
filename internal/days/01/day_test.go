package _1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "69206", day.SolvePart1())
}

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "197400", day.SolvePart2())
}

func TestSolve_Examples(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
		part1 bool
	}{
		"Example": {
			input: example,
			want:  "24000",
			part1: true,
		},
		"Example part 2": {
			input: example,
			want:  "45000",
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
