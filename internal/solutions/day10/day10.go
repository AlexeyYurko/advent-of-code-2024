package day10

import (
	"os"
	"path/filepath"
	"strings"
)

type Point struct {
	x, y int
	val  int
	done bool
}

var directions = []struct{ dx, dy int }{
	{0, -1}, // up
	{0, 1},  // down
	{-1, 0}, // left
	{1, 0},  // right
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day10", "input.txt"))
	return &Solver{input: string(input)}
}

func (p Point) getValidMoves(grid [][]int) []Point {
	var moves []Point
	height, width := len(grid), len(grid[0])
	nextVal := p.val + 1

	for _, dir := range directions {
		newX, newY := p.x+dir.dx, p.y+dir.dy

		if newX >= 0 && newX < width &&
			newY >= 0 && newY < height &&
			grid[newY][newX] == nextVal {

			next := Point{
				x:    newX,
				y:    newY,
				val:  nextVal,
				done: nextVal == 9,
			}
			moves = append(moves, next)
		}
	}
	return moves
}

func (p Point) findPaths(grid [][]int, peaks map[Point]int) {
	moves := p.getValidMoves(grid)

	for _, next := range moves {
		if next.done {
			peaks[next]++
		} else {
			next.findPaths(grid, peaks)
		}
	}
}

func parseGrid(s *Solver) (grid [][]int, starts []Point) {
	lines := strings.Split(strings.TrimSpace(s.input), "\n")

	grid = make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, char := range line {
			val := int(char - '0')
			grid[y][x] = val

			if val == 0 {
				starts = append(starts, Point{x: x, y: y})
			}
		}
	}
	return
}

func (s *Solver) Part1() (interface{}, error) {
	grid, starts := parseGrid(s)
	total := 0

	peaks := make(map[Point]int)
	for _, start := range starts {
		start.findPaths(grid, peaks)
		total += len(peaks)
		clear(peaks)
	}

	return total, nil
}

func (s *Solver) Part2() (interface{}, error) {
	grid, starts := parseGrid(s)
	total := 0

	peaks := make(map[Point]int)
	for _, start := range starts {
		start.findPaths(grid, peaks)
		for _, count := range peaks {
			total += count
		}
		clear(peaks)
	}

	return total, nil
}
