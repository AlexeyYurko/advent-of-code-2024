package day22

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Pair struct {
	Delta   int
	Bananas int
}

type Window struct {
	A, B, C, D int
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day22", "input.txt"))
	return &Solver{input: string(input)}
}

func calculateSecretNumber(currentNumber, iterations int) int {
	for i := 0; i < iterations; i++ {
		currentNumber = ((currentNumber << 6) ^ currentNumber) & 0xFFFFFF
		currentNumber = ((currentNumber >> 5) ^ currentNumber) & 0xFFFFFF
		currentNumber = ((currentNumber << 11) ^ currentNumber) & 0xFFFFFF
	}
	return currentNumber
}

func populateSecretNumberCache(currentNumber, iterations int, windowCache map[Window]int) []Pair {
	visited := make(map[Window]bool)
	deltaBananaPairs := make([]Pair, 0, iterations-1)
	currentDeltaWindow := make([]int, 0, 4)
	previousNumber := currentNumber % 10

	for i := 1; i < iterations; i++ {
		currentNumber = calculateSecretNumber(currentNumber, 1)
		delta := currentNumber%10 - previousNumber
		deltaBananaPairs = append(deltaBananaPairs, Pair{Delta: delta, Bananas: currentNumber % 10})
		currentDeltaWindow = append(currentDeltaWindow, delta)

		if len(currentDeltaWindow) == 4 {
			key := Window{A: currentDeltaWindow[0], B: currentDeltaWindow[1], C: currentDeltaWindow[2], D: currentDeltaWindow[3]}
			if !visited[key] {
				windowCache[key] += currentNumber % 10
				visited[key] = true
			}
			currentDeltaWindow = currentDeltaWindow[1:]
		}
		previousNumber = currentNumber % 10
	}
	return deltaBananaPairs
}

func parseInput(s *Solver) []int {
	lines := strings.Split(s.input, "\n")
	nums := make([]int, 0, len(lines))
	for _, number := range lines {
		if num, err := strconv.Atoi(number); err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func (s *Solver) Part1() (interface{}, error) {
	input := parseInput(s)

	result := 0
	for _, num := range input {
		result += calculateSecretNumber(num, 2000)
	}
	return result, nil
}

func (s *Solver) Part2() (interface{}, error) {
	input := parseInput(s)

	result := 0
	cache := make(map[Window]int)
	for _, num := range input {
		populateSecretNumberCache(num, 2000, cache)
	}

	for _, value := range cache {
		result = max(result, value)
	}
	return result, nil
}
