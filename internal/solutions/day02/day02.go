package day02

import (
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day02", "input_.txt"))
	return &Solver{input: string(input)}
}

func isValidReport(numbers []int) bool {
	isValid := true
	isIncreasing := false

	prev := numbers[0]
	second := numbers[1]
	diff := second - prev

	if math.Abs(float64(diff)) > 3 || math.Abs(float64(diff)) < 1 {
		isValid = false
	}
	isIncreasing = diff > 0
	prev = second

	for i := 2; i < len(numbers); i++ {
		curr := numbers[i]
		diff := curr - prev

		if math.Abs(float64(diff)) > 3 || math.Abs(float64(diff)) < 1 {
			isValid = false
		}

		if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) {
			isValid = false
		}

		prev = curr
	}
	return isValid
}

func (s *Solver) Part1() (interface{}, error) {
	lines := strings.Split(s.input, "\n")
	safeReports := 0

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		intNumbers := make([]int, len(numbers))
		for i, str := range numbers {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			intNumbers[i] = num
		}

		isValid := isValidReport(intNumbers)

		if isValid {
			safeReports++
		}
	}

	return safeReports, nil
}

func (s *Solver) Part2() (interface{}, error) {
	lines := strings.Split(s.input, "\n")
	safeReports := 0

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		intNumbers := make([]int, len(numbers))
		for i, str := range numbers {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			intNumbers[i] = num
		}

		if isValidReport(intNumbers) {
			safeReports++
			continue
		}

		for k := 0; k < len(intNumbers); k++ {
			var cleanNumbers []int
			for i := 0; i < len(intNumbers); i++ {
				if i != k {
					cleanNumbers = append(cleanNumbers, intNumbers[i])
				}
			}
			if isValidReport(cleanNumbers) {
				safeReports++
				break
			}

		}
	}
	return safeReports, nil
}
