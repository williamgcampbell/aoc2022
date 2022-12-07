package _5

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const example = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "QGTHFZBHV", day.SolvePart1())
}

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Solver{}
	require.Equal(t, "MGDMPSZTM", day.SolvePart2())
}

func TestSolve(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
		part1 bool
	}{
		"Example": {
			input: example,
			want:  "CMZ",
			part1: true,
		},
		"Example part 2": {
			input: example,
			want:  "MCD",
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

func TestMove(t *testing.T) {
	tests := map[string]struct {
		stack [][]string
		count int
		from  int
		to    int
		want  [][]string
	}{
		"Example": {
			stack: [][]string{
				{"A", "B"},
				{},
				{"D"},
			},
			count: 2,
			from:  1,
			to:    2,
			want: [][]string{
				{},
				{"B", "A"},
				{"D"},
			},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := crateMover9000(test.stack, test.count, test.from, test.to)
			if !assertEq(actual, test.want) {
				t.Errorf("Got: %s, Want: %s.", actual, test.want)
			}
		})
	}
}

func assertEq(test [][]string, ans [][]string) bool {
	return reflect.DeepEqual(test, ans)
}
