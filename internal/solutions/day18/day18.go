package day18

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

const (
	gridSize = 70
	maxLines = 1024
)

var (
	startPoint = Point{0, 0}
	goalPoint  = Point{70, 70}
)

var directions = []Point{
	{1, 0},  // right
	{-1, 0}, // left
	{0, 1},  // down
	{0, -1}, // up
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day18", "input.txt"))
	return &Solver{input: string(input)}
}

func (p Point) add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func parseInput(s *Solver) (map[Point]bool, error) {
	walls := make(map[Point]bool)
	lines := strings.Split(strings.TrimSpace(s.input), "\n")

	for i, line := range lines {
		if i >= maxLines {
			break
		}

		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		walls[Point{x, y}] = true
	}
	return walls, nil
}

func findShortestPath(walls map[Point]bool) (int, error) {
	steps := 0
	front := make([]Point, 0, gridSize*4)
	front = append(front, startPoint)
	seen := make(map[Point]bool, gridSize*gridSize)
	seen[startPoint] = true

	for len(front) > 0 {
		newFront := make([]Point, 0, len(front)*4)
		steps++

		for _, pos := range front {
			for _, dir := range directions {
				next := pos.add(dir)

				if next == goalPoint {
					return steps, nil
				}

				if !isValidPosition(next) || walls[next] || seen[next] {
					continue
				}

				seen[next] = true
				newFront = append(newFront, next)
			}
		}
		front = newFront
	}
	return 0, fmt.Errorf("no path found")
}

func isValidPosition(p Point) bool {
	return p.x >= 0 && p.x <= gridSize && p.y >= 0 && p.y <= gridSize
}

func (s *Solver) Part1() (interface{}, error) {
	grid, err := parseInput(s)
	if err != nil {
		return nil, err
	}
	return findShortestPath(grid)
}

func (s *Solver) Part2() (interface{}, error) {
	var coordinates []Point
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		coordinates = append(coordinates, Point{x, y})
	}

	pathCache := make(map[int]bool)

	left, right := 0, len(coordinates)-1
	var lastBlocking Point

	for left <= right {
		mid := left + (right-left)/2

		if result, exists := pathCache[mid]; exists {
			if result {
				right = mid - 1
				lastBlocking = coordinates[mid]
			} else {
				left = mid + 1
			}
			continue
		}

		testGrid := make(map[Point]bool, mid+1)
		for i := 0; i <= mid; i++ {
			testGrid[coordinates[i]] = true
		}

		_, err := findShortestPath(testGrid)
		pathCache[mid] = err != nil

		if err != nil {
			right = mid - 1
			lastBlocking = coordinates[mid]
		} else {
			left = mid + 1
		}
	}

	return fmt.Sprintf("%d,%d", lastBlocking.x, lastBlocking.y), nil
}
