package days

import (
	_1 "github.com/williamgcampbell/aoc2022/internal/days/01"
	_2 "github.com/williamgcampbell/aoc2022/internal/days/02"
)

func RegisterAll(registry SolverRegistry) {
	registry.Register(1, &_1.Solver{})
	registry.Register(2, &_2.Solver{})
}
