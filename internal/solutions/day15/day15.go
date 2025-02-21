package day15

import (
	"github.com/AlexeyYurko/advent-of-code-2024/internal/aoc"
	"os"
	"path/filepath"
	"strings"
)

var (
	deltaPoints = map[rune]aoc.Point{
		'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0},
		'[': {1, 0}, ']': {-1, 0},
	}
)

func parseGrid(gridInput string) (map[aoc.Point]rune, aoc.Point) {
	grid := make(map[aoc.Point]rune, len(gridInput))
	robot := aoc.Point{}
	lines := strings.Fields(gridInput)
	for y, line := range lines {
		for x, r := range line {
			if r == '@' {
				robot = aoc.Point{x, y}
				r = '.'
			}
			grid[aoc.Point{x, y}] = r
		}
	}
	return grid, robot
}

func processMove(grid map[aoc.Point]rune, robot aoc.Point, move rune) (map[aoc.Point]rune, aoc.Point, bool) {
	queue := []aoc.Point{robot}
	boxes := make(map[aoc.Point]rune, len(grid)/4)

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if _, ok := boxes[p]; ok {
			continue
		}
		boxes[p] = grid[p]

		delta := deltaPoints[move]
		n := p.Add(delta)

		if grid[n] == '#' {
			return grid, robot, false
		}

		if grid[n] == '[' || grid[n] == ']' {
			queue = append(queue, n.Add(deltaPoints[grid[n]]))
		}
		if grid[n] == '[' || grid[n] == ']' || grid[n] == 'O' {
			queue = append(queue, n)
		}
	}

	updates := make(map[aoc.Point]rune, len(boxes))
	for b := range boxes {
		grid[b] = '.'
		updates[b.Add(deltaPoints[move])] = boxes[b]
	}
	for p, r := range updates {
		grid[p] = r
	}
	return grid, robot.Add(deltaPoints[move]), true
}

func calculateGPS(grid map[aoc.Point]rune) int {
	gps := 0
	for p, r := range grid {
		if r == 'O' || r == '[' {
			gps += 100*p.Y + p.X
		}
	}
	return gps
}

type Solver struct {
	input string
}

func (s *Solver) solve(useDoubleWidth bool) (interface{}, error) {
	parts := strings.Split(strings.TrimSpace(s.input), "\n\n")
	gridInput := parts[0]
	moves := strings.ReplaceAll(parts[1], "\n", "")

	if useDoubleWidth {
		r := strings.NewReplacer("#", "##", "O", "[]", ".", "..", "@", "@.")
		gridInput = r.Replace(gridInput)
	}

	grid, robot := parseGrid(gridInput)

	for _, move := range moves {
		var ok bool
		grid, robot, ok = processMove(grid, robot, move)
		if !ok {
			continue
		}
	}

	return calculateGPS(grid), nil
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day15", "input.txt"))
	return &Solver{input: string(input)}
}

func (s *Solver) Part1() (interface{}, error) {
	return s.solve(false)
}

func (s *Solver) Part2() (interface{}, error) {
	return s.solve(true)
}
