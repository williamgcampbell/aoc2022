package intersect

import (
	"strings"
	"testing"
)

func TestUnique(t *testing.T) {
	tests := map[string]struct {
		a    string
		b    string
		want string
	}{
		"Example": {
			a:    "vJrwpWtwJgWrhcsFMMfFFhFp",
			b:    "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			want: "rsFMf",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := Unique(strings.Split(test.a, ""), strings.Split(test.b, ""))
			actualStr := strings.Join(actual, "")
			if actualStr != test.want {
				t.Errorf("Got: %s, Want: %s.", actualStr, test.want)
			}
		})
	}
}
