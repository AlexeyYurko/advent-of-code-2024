package runner

import (
	"fmt"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day01"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day02"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day03"
)

type Result struct {
	Part1 interface{}
	Part2 interface{}
}

type Solver interface {
	Part1() (interface{}, error)
	Part2() (interface{}, error)
}

func Run(day int) (*Result, error) {
	solver, err := getSolver(day)
	if err != nil {
		return nil, err
	}

	p1, err := solver.Part1()
	if err != nil {
		return nil, fmt.Errorf("part 1: %w", err)
	}

	p2, err := solver.Part2()
	if err != nil {
		return nil, fmt.Errorf("part 2: %w", err)
	}

	return &Result{
		Part1: p1,
		Part2: p2,
	}, nil
}

func getSolver(day int) (Solver, error) {
	switch day {
	case 1:
		return day01.New(), nil
	case 2:
		return day02.New(), nil
	case 3:
		return day03.New(), nil
	default:
		return nil, fmt.Errorf("invalid day: %d", day)
	}
}
