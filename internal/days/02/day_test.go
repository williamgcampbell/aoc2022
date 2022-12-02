package _2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

var example = []string{"A Y", "B X", "C Z"}

func TestTotalScore(t *testing.T) {
	tests := map[string]struct {
		lines   []string
		partTwo bool
		want    int
	}{
		"Example": {
			lines: example,
			want:  15,
		},
		"Example Part 2": {
			lines:   example,
			partTwo: true,
			want:    12,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := totalScore(test.lines, test.partTwo)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
