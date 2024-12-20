package runner

import (
	"fmt"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day01"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day02"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day03"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day04"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day05"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day06"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day07"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day08"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day09"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day10"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day11"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day12"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day13"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day14"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day15"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day16"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day17"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day18"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day19"
	"github.com/AlexeyYurko/advent-of-code-2024/internal/solutions/day20"
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
	case 4:
		return day04.New(), nil
	case 5:
		return day05.New(), nil
	case 6:
		return day06.New(), nil
	case 7:
		return day07.New(), nil
	case 8:
		return day08.New(), nil
	case 9:
		return day09.New(), nil
	case 10:
		return day10.New(), nil
	case 11:
		return day11.New(), nil
	case 12:
		return day12.New(), nil
	case 13:
		return day13.New(), nil
	case 14:
		return day14.New(), nil
	case 15:
		return day15.New(), nil
	case 16:
		return day16.New(), nil
	case 17:
		return day17.New(), nil
	case 18:
		return day18.New(), nil
	case 19:
		return day19.New(), nil
	case 20:
		return day20.New(), nil
	default:
		return nil, fmt.Errorf("invalid day: %d", day)
	}
}
