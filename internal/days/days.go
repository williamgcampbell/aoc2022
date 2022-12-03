package days

import (
	_1 "github.com/williamgcampbell/aoc2022/internal/days/01"
	_2 "github.com/williamgcampbell/aoc2022/internal/days/02"
	_3 "github.com/williamgcampbell/aoc2022/internal/days/03"
	_4 "github.com/williamgcampbell/aoc2022/internal/days/04"
	_5 "github.com/williamgcampbell/aoc2022/internal/days/05"
)

func RegisterAll(registry SolverRegistry) {
	registry.Register(1, &_1.Solver{})
	registry.Register(2, &_2.Solver{})
	registry.Register(3, &_3.Solver{})
	registry.Register(4, &_4.Solver{})
	registry.Register(5, &_5.Solver{})
}
