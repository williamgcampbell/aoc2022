package scanner

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestReadFromCsvInt(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    [][]int
		wantErr bool
	}{
		"Empty file is nil": {
			input: ``,
			want:  nil,
		},
		"Empty line is ignored": {
			input: "1\n\n2",
			want:  [][]int{{1}, {2}},
		},
		"Uneven fields per record ok": {
			input: "1,2,3,4\n5,6,7,8",
			want:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
		},
		"Empty field is error": {
			input:   "1,2,3,4\n5,6,,8",
			wantErr: true,
		},
		"Non integer value returns error": {
			input:   `1,2,a`,
			wantErr: true,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual, err := ReadFromCsvInt(strings.NewReader(test.input))

			if test.wantErr {
				require.Error(t, err)
			} else {
				if diff := cmp.Diff(actual, test.want); diff != "" {
					t.Error(diff)
				}
			}
		})
	}
}

func TestScanIntLines(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    []int
		wantErr bool
	}{
		"Empty file is nil": {
			input:   ``,
			want:    nil,
			wantErr: false,
		},
		"Empty line is ignored": {
			input: `1

2`,
			want:    []int{1, 2},
			wantErr: true,
		},
		"Integer values are separated by line": {
			input:   "1\n2\n3\n4",
			want:    []int{1, 2, 3, 4},
			wantErr: false,
		},
		"Non integer value returns error": {
			input:   `1test`,
			want:    nil,
			wantErr: true,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual, err := ScanIntLines(strings.NewReader(test.input))

			if test.wantErr {
				require.Error(t, err)
			} else {
				if diff := cmp.Diff(actual, test.want); diff != "" {
					t.Error(diff)
				}
			}
		})
	}
}

func TestScanLines(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"Empty file is nil": {
			input: ``,
			want:  nil,
		},
		"Empty line is empty": {
			input: `l1

l3`,
			want: []string{"l1", "", "l3"},
		},
		"Values are separated by line": {
			input: "l1\n2\nl3\nl4",
			want:  []string{"l1", "2", "l3", "l4"},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := ScanLines(strings.NewReader(test.input))

			if diff := cmp.Diff(actual, test.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestScanUntilEmptyLine(t *testing.T) {
	tests := map[string]struct {
		input     string
		want      []string
		delimiter string
	}{
		"Empty file is nil": {
			input:     ``,
			want:      nil,
			delimiter: "",
		},
		"Empty lines are ignored": {
			input: `l1

l3`,
			want:      []string{"l1", "l3"},
			delimiter: "",
		},
		"Lines are grouped by empty lines": {
			input:     "l1\n2\n\nl3\nl4",
			want:      []string{"l12", "l3l4"},
			delimiter: "",
		},
		"Lines are separated by delimiter": {
			input:     "l1\n2\n\nl3\nl4",
			want:      []string{"l1,2", "l3,l4"},
			delimiter: ",",
		},
		"Consecutive empty lines create empty values": {
			input:     "\n\n\n\n\n",
			want:      []string{"", "", "", "", ""},
			delimiter: "",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := ScanUntilEmptyLine(strings.NewReader(test.input), test.delimiter)

			if diff := cmp.Diff(actual, test.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
