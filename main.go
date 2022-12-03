package main

import (
	"fmt"
	"log"

	"github.com/williamgcampbell/aoc2022/internal/days"
)

const (
	chorus  = "On the %s day of Christmas the part %s solution is: %s\n"
	numDays = 25
)

func main() {
	r := days.NewDayRegistry()
	days.RegisterAll(r)

	for i := 1; i <= numDays; i++ {
		if day, ok := r[i]; ok {
			singDay(i, day)
		} else {
			log.Fatalf("Could not find day %d", i)
		}
	}
}

func singDay(num int, day days.Solver) {
	singPart("one", num, day.SolvePart1())
	singPart("two", num, day.SolvePart2())
}

func singPart(part string, day int, solution string) {
	if day == 5 && part == "one" {
		fmt.Println("FIIIIIIIIIIVE GOOOOOLDEN R....just kidding. ")
	}
	fmt.Printf(chorus, ordinalString(day), part, normalizeSolution(solution))
}

func normalizeSolution(r string) string {
	if r == "" {
		return "I don't know"
	}
	return r
}

func ordinalString(i int) string {
	j := i % 10
	k := i % 100
	if j == 1 && k != 11 {
		return fmt.Sprintf("%dst", i)
	}
	if j == 2 && k != 12 {
		return fmt.Sprintf("%dnd", i)
	}
	if j == 3 && k != 13 {
		return fmt.Sprintf("%drd", i)
	}
	return fmt.Sprintf("%dth", i)
}
