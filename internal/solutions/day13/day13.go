package day13

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

type MachineMap map[string]Coordinate

func parseMachines(input string) ([]MachineMap, error) {
	var machines []MachineMap
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i := 0; i < len(lines); i += 4 {
		locations := make(MachineMap)
		for j := 0; j < 3; j++ {
			parts := strings.Split(lines[i+j], ":")

			key := "C"
			if strings.HasPrefix(parts[0], "Button") {
				key = strings.Fields(parts[0])[1]
			}

			x, y, _ := parseCoordinates(parts[1])

			locations[key] = Coordinate{X: x, Y: y}
		}
		machines = append(machines, locations)
	}
	return machines, nil
}

func parseCoordinates(input string) (x, y int, err error) {
	coords := strings.Split(strings.TrimSpace(input), ",")

	xStr := strings.TrimSpace(coords[0])
	yStr := strings.TrimSpace(coords[1])

	if strings.Contains(xStr, "+") && strings.Contains(yStr, "+") {
		x, _ = strconv.Atoi(strings.Replace(xStr, "X", "", -1))
		y, _ = strconv.Atoi(strings.Replace(yStr, "Y", "", -1))
	} else {
		x, _ = strconv.Atoi(strings.Replace(xStr, "X=", "", -1))
		y, _ = strconv.Atoi(strings.Replace(yStr, "Y=", "", -1))
	}
	return x, y, nil
}

func countTokens(moves [][2]int) int {
	tokens := 0
	for _, move := range moves {
		tokens += move[0]*3 + move[1]
	}
	return tokens
}

func findWinningMoves(machines []MachineMap, offset int64) (int, error) {
	moves := make([][2]int, 0, len(machines))

	for i, machine := range machines {
		a1 := int64(machine["A"].X)
		a2 := int64(machine["A"].Y)
		b1 := int64(machine["B"].X)
		b2 := int64(machine["B"].Y)
		c1 := int64(machine["C"].X) + offset
		c2 := int64(machine["C"].Y) + offset

		denominator := b2*a1 - b1*a2
		if denominator == 0 {
			return 0, fmt.Errorf("division by zero in machine %d", i)
		}

		A := float64(b2*c1-b1*c2) / float64(denominator)
		B := (float64(c1) - float64(a1)*A) / float64(b1)

		if A == float64(int64(A)) && B == float64(int64(B)) {
			moves = append(moves, [2]int{int(A), int(B)})
		}
	}

	return countTokens(moves), nil
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day13", "input.txt"))
	return &Solver{input: string(input)}
}
func (s *Solver) Part1() (interface{}, error) {
	machines, _ := parseMachines(s.input)
	return findWinningMoves(machines, 0)
}

func (s *Solver) Part2() (interface{}, error) {
	machines, _ := parseMachines(s.input)
	return findWinningMoves(machines, 10000000000000)
}
