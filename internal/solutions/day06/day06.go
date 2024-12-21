package day06

import (
	"github.com/AlexeyYurko/advent-of-code-2024/internal/aoc"
	"os"
	"path/filepath"
	"strings"
)

type directionType struct {
	dX int
	dY int
}

var up = directionType{0, -1}
var right = directionType{1, 0}
var down = directionType{0, 1}
var left = directionType{-1, 0}

var directions = []directionType{up, right, down, left}

type VisitedState struct {
	point     aoc.Point
	direction int
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day06", "input.txt"))
	return &Solver{input: string(input)}
}

func parseGrid(s *Solver) (map[aoc.Point]rune, int, int) {
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	maxY := len(lines)
	maxX := len(lines[0])

	grid := make(map[aoc.Point]rune)
	for y, line := range lines {
		for x, cell := range line {
			grid[aoc.Point{y, x}] = cell
		}
	}
	return grid, maxX, maxY
}

func findStart(grid map[aoc.Point]rune) aoc.Point {
	for k, v := range grid {
		if v == '^' {
			return k
		}
	}
	return aoc.Point{}
}

func walk(grid map[aoc.Point]rune, start aoc.Point, maxX, maxY int) []aoc.Point {
	var positions []aoc.Point

	directionOrder := 0
	grid[start] = 'X'
	positions = append(positions, start)
	currentPosition := start

	for {
		nextX := currentPosition.X + directions[directionOrder].dX
		nextY := currentPosition.Y + directions[directionOrder].dY

		if nextX < 0 || nextX >= maxX || nextY < 0 || nextY >= maxY {
			break
		}

		nextCell := grid[aoc.Point{nextY, nextX}]
		if nextCell == '#' {
			directionOrder = (directionOrder + 1) % 4
			continue
		}

		if nextCell != 'X' {
			grid[aoc.Point{nextY, nextX}] = 'X'
			positions = append(positions, aoc.Point{nextY, nextX})
		}

		currentPosition = aoc.Point{nextY, nextX}
	}
	return positions
}

func hasLoop(grid map[aoc.Point]rune, start aoc.Point, maxX, maxY int) bool {
	visited := make(map[VisitedState]bool)
	pos := start
	directionOrder := 0

	for {
		state := VisitedState{pos, directionOrder}

		if visited[state] {
			return true
		}

		visited[state] = true

		nextPos := aoc.Point{
			X: pos.X + directions[directionOrder].dX,
			Y: pos.Y + directions[directionOrder].dY,
		}

		if nextPos.X < 0 || nextPos.X >= maxX || nextPos.Y < 0 || nextPos.Y >= maxY {
			return false
		}

		if grid[aoc.Point{nextPos.Y, nextPos.X}] == '#' {
			directionOrder = (directionOrder + 1) % 4
		} else {
			pos = nextPos
		}
	}
}

func (s *Solver) Part1() (interface{}, error) {
	grid, maxX, maxY := parseGrid(s)

	start := findStart(grid)

	visited := walk(grid, start, maxX, maxY)
	return len(visited), nil
}

func (s *Solver) Part2() (interface{}, error) {
	grid, maxX, maxY := parseGrid(s)
	start := findStart(grid)

	initialPath := walk(grid, start, maxX, maxY)

	count := 0
	for _, pos := range initialPath {
		originalValue := grid[pos]
		grid[pos] = '#'

		if hasLoop(grid, start, maxX, maxY) {
			count++
		}

		grid[pos] = originalValue
	}

	return count, nil
}
