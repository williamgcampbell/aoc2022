package days

type Solver interface {
	SolvePart1() string
	SolvePart2() string
}

type SolverRegistry map[int]Solver

func NewDayRegistry() SolverRegistry {
	return make(SolverRegistry)
}

func (r SolverRegistry) Register(dayNumber int, day Solver) {
	r[dayNumber] = day
}
