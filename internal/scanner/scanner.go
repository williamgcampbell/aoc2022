package scanner

import (
	"bufio"
	"encoding/csv"
	"io"
	"strconv"
)

// ReadFromCsvInt returns a two-dimensional array of int.
// The outer array is split into lines and the inner arrays are split by commas.
//
// Note:
// - Leading spaces are trimmed
// - Equal number of fields per line are NOT required
//
// Example.
// A file containing two lines:
// 1, 2, 3, 4
// 5, 6, 7, 8
//
// Will produce an array:
// [[1, 2, 3, 4], [5, 6, 7, 8]]
func ReadFromCsvInt(handle io.Reader) ([][]int, error) {
	csvR := csv.NewReader(handle)
	csvR.TrimLeadingSpace = true

	// don't require equal number of fields per line
	csvR.FieldsPerRecord = -1

	lines, err := csvR.ReadAll()
	if err != nil {
		return nil, err
	}

	var r [][]int
	for _, line := range lines {
		var c []int
		for _, val := range line {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			c = append(c, valInt)
		}
		r = append(r, c)
	}

	return r, nil
}

// ScanIntLines returns an array of int, split by new line
func ScanIntLines(handle io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(handle)
	scanner.Split(bufio.ScanLines)
	var lines []int

	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, v)
	}
	return lines, nil
}

// ScanLines returns an array of strings, split by new lines
func ScanLines(handle io.Reader) []string {
	scanner := bufio.NewScanner(handle)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ScanUntilEmptyLine returns an array of strings, split by empty lines in the reader.
// When continuing from one line to the next, newlineDelimiter will be put between the values.
func ScanUntilEmptyLine(handle io.Reader, newlineDelimiter string) []string {
	scanner := bufio.NewScanner(handle)
	scanner.Split(bufio.ScanLines)
	var lines []string
	var line string
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 {
			lines = append(lines, line)
			line = ""
			continue
		}

		if len(line) == 0 {
			line = t
		} else {
			line += newlineDelimiter + t
		}
	}

	if len(line) != 0 {
		lines = append(lines, line)
	}

	return lines
}
