package internal

import (
	"fmt"
	"strconv"
)

func MustAtoI(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("converting to int: %w", err)
		panic(err)
	}

	return result
}
