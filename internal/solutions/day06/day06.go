package day06

import (
	"os"
	"path/filepath"
	"strings"
)

type Point struct {
	y, x int
}

type directionType struct {
	dX int
	dY int
}

var up = directionType{0, -1}
var right = directionType{1, 0}
var down = directionType{0, 1}
var left = directionType{-1, 0}

var directions = []directionType{up, right, down, left}

type Solver struct {
	input string
}

type VisitedState struct {
	point     Point
	direction int
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day06", "input.txt"))
	return &Solver{input: string(input)}
}

func parseGrid(s *Solver) (map[Point]rune, int, int) {
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	maxY := len(lines)
	maxX := len(lines[0])

	grid := make(map[Point]rune)
	for y, line := range lines {
		for x, cell := range line {
			grid[Point{y, x}] = cell
		}
	}
	return grid, maxX, maxY
}

func findStart(grid map[Point]rune) Point {
	for k, v := range grid {
		if v == '^' {
			return k
		}
	}
	return Point{}
}

func walk(grid map[Point]rune, start Point, maxX, maxY int) []Point {
	var positions []Point

	directionOrder := 0
	grid[start] = 'X'
	positions = append(positions, start)
	currentPosition := start

	for {
		nextX := currentPosition.x + directions[directionOrder].dX
		nextY := currentPosition.y + directions[directionOrder].dY

		if nextX < 0 || nextX >= maxX || nextY < 0 || nextY >= maxY {
			break
		}

		nextCell := grid[Point{nextY, nextX}]
		if nextCell == '#' {
			directionOrder = (directionOrder + 1) % 4
			continue
		}

		if nextCell != 'X' {
			grid[Point{nextY, nextX}] = 'X'
			positions = append(positions, Point{nextY, nextX})
		}

		currentPosition = Point{nextY, nextX}
	}
	return positions
}

func hasLoop(grid map[Point]rune, start Point, maxX, maxY int) bool {
	visited := make(map[VisitedState]bool)
	pos := start
	directionOrder := 0

	for {
		state := VisitedState{pos, directionOrder}

		if visited[state] {
			return true
		}

		visited[state] = true

		nextPos := Point{
			x: pos.x + directions[directionOrder].dX,
			y: pos.y + directions[directionOrder].dY,
		}

		if nextPos.x < 0 || nextPos.x >= maxX || nextPos.y < 0 || nextPos.y >= maxY {
			return false
		}

		if grid[Point{nextPos.y, nextPos.x}] == '#' {
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
