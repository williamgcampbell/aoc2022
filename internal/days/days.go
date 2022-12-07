package days

import (
	_1 "github.com/williamgcampbell/aoc2022/internal/days/01"
	_2 "github.com/williamgcampbell/aoc2022/internal/days/02"
	_3 "github.com/williamgcampbell/aoc2022/internal/days/03"
	_4 "github.com/williamgcampbell/aoc2022/internal/days/04"
	_5 "github.com/williamgcampbell/aoc2022/internal/days/05"
	_6 "github.com/williamgcampbell/aoc2022/internal/days/06"
	_7 "github.com/williamgcampbell/aoc2022/internal/days/07"
	_8 "github.com/williamgcampbell/aoc2022/internal/days/08"
	_9 "github.com/williamgcampbell/aoc2022/internal/days/09"
)

func RegisterAll(registry SolverRegistry) {
	registry.Register(1, &_1.Solver{})
	registry.Register(2, &_2.Solver{})
	registry.Register(3, &_3.Solver{})
	registry.Register(4, &_4.Solver{})
	registry.Register(5, &_5.Solver{})
	registry.Register(6, &_6.Solver{})
	registry.Register(7, &_7.Solver{})
	registry.Register(8, &_8.Solver{})
	registry.Register(9, &_9.Solver{})
}
