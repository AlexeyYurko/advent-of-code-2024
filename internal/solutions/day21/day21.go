package day21

import (
	"github.com/AlexeyYurko/advent-of-code-2024/internal/aoc"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

type padMappings struct {
	numeric, directional map[string]aoc.Point
}

const (
	startingPoint = "A"
)

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day21", "input.txt"))
	return &Solver{input: string(input)}
}

func extractFirstNumberFromString(line string) int {
	var build strings.Builder
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		return int(localNum)
	}
	return 0
}

func initializePadMappings() padMappings {
	numeric := map[string]aoc.Point{
		"A": {2, 0}, "0": {1, 0}, "1": {0, 1},
		"2": {1, 1}, "3": {2, 1}, "4": {0, 2},
		"5": {1, 2}, "6": {2, 2}, "7": {0, 3},
		"8": {1, 3}, "9": {2, 3},
	}

	directional := map[string]aoc.Point{
		"A": {2, 1}, "^": {1, 1}, "<": {0, 0},
		"v": {1, 0}, ">": {2, 0},
	}

	return padMappings{numeric, directional}
}

func parseInput(s *Solver) []string {
	return strings.Split(s.input, "\n")
}

func (s *Solver) Part1() (interface{}, error) {
	input := parseInput(s)

	padMaps := initializePadMappings()

	result := calculateTotalPressCount(input, padMaps, 2)
	return result, nil
}

func (s *Solver) Part2() (interface{}, error) {
	input := parseInput(s)

	padMaps := initializePadMappings()

	result := calculateTotalPressCount(input, padMaps, 25)
	return result, nil
}

func calculateButtonPresses(input []string, start string, padMap map[string]aoc.Point, isNumeric bool) []string {
	current := padMap[start]
	var output []string

	for _, char := range input {
		dest := padMap[char]
		diffX, diffY := dest.X-current.X, dest.Y-current.Y

		var horizontal []string
		var vertical []string

		for i := 0; i < aoc.Abs(diffX); i++ {
			if diffX >= 0 {
				horizontal = append(horizontal, ">")
			} else {
				horizontal = append(horizontal, "<")
			}
		}

		for i := 0; i < aoc.Abs(diffY); i++ {
			if diffY >= 0 {
				vertical = append(vertical, "^")
			} else {
				vertical = append(vertical, "v")
			}
		}

		if isNumeric {
			if current.Y == 0 && dest.X == 0 {
				output = append(output, vertical...)
				output = append(output, horizontal...)
			} else if current.X == 0 && dest.Y == 0 {
				output = append(output, horizontal...)
				output = append(output, vertical...)
			} else if diffX < 0 {
				output = append(output, horizontal...)
				output = append(output, vertical...)
			} else {
				output = append(output, vertical...)
				output = append(output, horizontal...)
			}
		} else {
			if current.X == 0 && dest.Y == 1 {
				output = append(output, horizontal...)
				output = append(output, vertical...)
			} else if current.Y == 1 && dest.X == 0 {
				output = append(output, vertical...)
				output = append(output, horizontal...)
			} else if diffX < 0 {
				output = append(output, horizontal...)
				output = append(output, vertical...)
			} else {
				output = append(output, vertical...)
				output = append(output, horizontal...)
			}
		}

		current = dest
		output = append(output, startingPoint)
	}
	return output
}

func calculateTotalPressCount(input []string, padMaps padMappings, robots int) int {
	count := 0
	cache := make(map[string][]int)
	for _, line := range input {
		row := strings.Split(line, "")
		sequence := calculateButtonPresses(row, startingPoint, padMaps.numeric, true)
		num := calculatePressCountWithRobots(sequence, robots, 1, cache, padMaps.directional)
		count += extractFirstNumberFromString(line) * num
	}
	return count
}

func calculatePressCountWithRobots(input []string, maxRobots, robot int, cache map[string][]int, directionalMap map[string]aoc.Point) int {
	seqKey := strings.Join(input, "")

	if val, ok := cache[seqKey]; ok {
		if robot-1 < len(val) && val[robot-1] != 0 {
			return val[robot-1]
		}
	} else {
		cache[seqKey] = make([]int, maxRobots)
	}

	seq := calculateButtonPresses(input, startingPoint, directionalMap, false)

	cache[seqKey][0] = len(seq)

	if robot == maxRobots {
		return len(seq)
	}

	splitSeq := splitSequenceIntoSteps(seq)
	count := 0
	for _, s := range splitSeq {
		c := calculatePressCountWithRobots(s, maxRobots, robot+1, cache, directionalMap)
		cacheKey := strings.Join(s, "")
		if _, ok := cache[cacheKey]; !ok {
			cache[cacheKey] = make([]int, maxRobots)
		}
		cache[cacheKey][0] = c
		count += c
	}

	if robot-1 < len(cache[seqKey]) {
		cache[seqKey][robot-1] = count
	}

	return count
}

func splitSequenceIntoSteps(input []string) [][]string {
	var output [][]string
	var current []string
	for _, char := range input {
		current = append(current, char)

		if char == startingPoint {
			output = append(output, current)
			current = []string{}
		}
	}
	return output
}
