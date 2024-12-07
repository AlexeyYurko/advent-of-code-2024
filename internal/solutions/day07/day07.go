package day07

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Equation struct {
	result  int
	numbers []int
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day07", "input.txt"))
	return &Solver{input: string(input)}
}

func NewEquation(line string) (Equation, error) {
	parts := strings.Split(line, ": ")
	result, err := strconv.Atoi(parts[0])
	if err != nil {
		return Equation{}, fmt.Errorf("parsing result: %w", err)
	}

	var numbers []int
	for _, num := range strings.Fields(parts[1]) {
		n, err := strconv.Atoi(num)
		if err != nil {
			return Equation{}, fmt.Errorf("parsing number %s: %w", num, err)
		}
		numbers = append(numbers, n)
	}

	return Equation{result: result, numbers: numbers}, nil
}

func parseLines(s string) ([]Equation, error) {
	var equations []Equation

	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		eq, err := NewEquation(line)
		if err != nil {
			return nil, fmt.Errorf("parsing line %q: %w", line, err)
		}
		equations = append(equations, eq)
	}

	return equations, nil
}

func checkEquation(result int, numbers []int, withConcatenation bool) bool {
	var check func(current int, pos int) bool
	check = func(current int, pos int) bool {
		if pos == len(numbers) {
			return current == result
		}

		if check(current+numbers[pos], pos+1) {
			return true
		}

		if check(current*numbers[pos], pos+1) {
			return true
		}

		if withConcatenation {
			currentStr := strconv.Itoa(current)
			nextStr := strconv.Itoa(numbers[pos])
			concatenated, _ := strconv.Atoi(currentStr + nextStr)
			if check(concatenated, pos+1) {
				return true
			}
		}

		return false
	}

	return check(numbers[0], 1)
}

func (e Equation) IsValid(withConcatenation bool) bool {
	return checkEquation(e.result, e.numbers, withConcatenation)
}

func (s *Solver) Part1() (interface{}, error) {
	equations, err := parseLines(s.input)
	if err != nil {
		return nil, fmt.Errorf("parsing input: %w", err)
	}
	var sum int
	for _, equation := range equations {
		equationIsTrue := checkEquation(equation.result, equation.numbers, false)
		if equationIsTrue {
			sum += equation.result
		}
	}
	return sum, nil
}

func (s *Solver) Part2() (interface{}, error) {
	equations, err := parseLines(s.input)
	if err != nil {
		return nil, fmt.Errorf("parsing input: %w", err)
	}
	var sum int
	for _, equation := range equations {
		if equation.IsValid(true) {
			sum += equation.result
		}
	}
	return sum, nil
}
