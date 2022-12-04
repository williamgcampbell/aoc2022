package internal

import (
	"fmt"
	"strconv"
)

// MustAtoI will turn a string into an int.
// If an error occurs a panic is thrown.
func MustAtoI(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("converting to int: %w", err)
		panic(err)
	}

	return result
}
