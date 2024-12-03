package day03

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Solver struct {
	input string
}

func getNumbers(match string) (num1, num2 int) {
	numsRegExp := regexp.MustCompile("(\\d+),(\\d+)")
	numbers := numsRegExp.FindAllStringSubmatch(match, -1)
	num1, _ = strconv.Atoi(numbers[0][1])
	num2, _ = strconv.Atoi(numbers[0][2])
	return num1, num2
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day03", "input.txt"))
	return &Solver{input: string(input)}
}

func (s *Solver) Part1() (interface{}, error) {
	mulRegExp := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := mulRegExp.FindAllString(s.input, -1)
	summary := 0

	for _, instruction := range matches {
		num1, num2 := getNumbers(instruction)
		summary += num1 * num2
	}
	return summary, nil
}

func (s *Solver) Part2() (interface{}, error) {
	instructionRegExp := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	matches := instructionRegExp.FindAllString(s.input, -1)
	summary := 0
	mulEnabled := true

	for _, instruction := range matches {
		switch instruction {
		case "do()":
			mulEnabled = true
		case "don't()":
			mulEnabled = false
		default:
			if mulEnabled && strings.HasPrefix(instruction, "mul(") {
				num1, num2 := getNumbers(instruction)
				summary += num1 * num2
			}
		}
	}
	return summary, nil
}
