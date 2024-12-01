package day01

import (
	"os"
	"path/filepath"
)

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day01", "input.txt"))
	return &Solver{input: string(input)}
}

func (s *Solver) Part1() (interface{}, error) {
	return nil, nil
}

func (s *Solver) Part2() (interface{}, error) {
	return nil, nil
}
