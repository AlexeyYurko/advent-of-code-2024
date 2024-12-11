package day11

import (
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Stone struct {
	value int
	count int
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day11", "input.txt"))
	return &Solver{input: string(input)}
}

func parseInput(input string) map[int]int {
	stones := make(map[int]int)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for _, s := range strings.Fields(line) {
			if num, err := strconv.Atoi(s); err == nil {
				stones[num]++
			}
		}
	}
	return stones
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(float64(n)))) + 1
}

func splitNumber(n, digits int) (first, second int) {
	divisor := intPow(10, digits/2)
	return n / divisor, n % divisor
}

func intPow(base, exp int) int {
	if exp == 0 {
		return 1
	}

	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

func transformStone(stone Stone) []Stone {
	digits := countDigits(stone.value)

	switch {
	case stone.value == 0:
		return []Stone{{value: 1, count: stone.count}}
	case digits%2 == 0:
		first, second := splitNumber(stone.value, digits)
		return []Stone{
			{value: first, count: stone.count},
			{value: second, count: stone.count},
		}
	default:
		return []Stone{{value: stone.value * 2024, count: stone.count}}
	}
}

func solve(stones map[int]int, iterations int) uint64 {
	current := stones

	for i := 0; i < iterations; i++ {
		next := make(map[int]int)
		for value, count := range current {
			for _, transformed := range transformStone(Stone{value: value, count: count}) {
				next[transformed.value] += transformed.count
			}
		}
		current = next
	}

	var sum uint64
	for _, count := range current {
		sum += uint64(count)
	}
	return sum
}

func (s *Solver) Part1() (interface{}, error) {
	return solve(parseInput(s.input), 25), nil
}

func (s *Solver) Part2() (interface{}, error) {
	return solve(parseInput(s.input), 75), nil
}
