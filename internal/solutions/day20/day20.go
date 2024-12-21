package day20

import (
	"github.com/AlexeyYurko/advent-of-code-2024/internal/aoc"
	"os"
	"path/filepath"
	"strings"
)

var directions = [4]aoc.Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day20", "input.txt"))
	return &Solver{input: string(input)}
}

const (
	StartPoint       = 'S'
	WallPoint        = '#'
	DefaultQueueSize = 1000
)

func buildDist(grid map[aoc.Point]rune, start aoc.Point) map[aoc.Point]int {
	dist := make(map[aoc.Point]int, len(grid))
	queue := make([]aoc.Point, 0, DefaultQueueSize)
	visited := make(map[aoc.Point]struct{}, len(grid))

	queue = append(queue, start)
	dist[start] = 0
	visited[start] = struct{}{}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			n := p.Add(d)
			if _, seen := visited[n]; !seen && grid[n] != WallPoint {
				visited[n] = struct{}{}
				dist[n] = dist[p] + 1
				queue = append(queue, n)
			}
		}
	}
	return dist
}

func parseGrid(s *Solver) (map[aoc.Point]rune, aoc.Point) {
	lines := strings.Fields(s.input)

	grid := make(map[aoc.Point]rune, len(lines)*len(lines[0]))
	var start aoc.Point

	for y, line := range lines {
		for x, r := range line {
			p := aoc.Point{x, y}
			if r == StartPoint {
				start = p
			}
			grid[p] = r
		}
	}

	return grid, start
}

func solve(dist map[aoc.Point]int, maxDist int) int {
	count := 0
	points := make([]aoc.Point, 0, len(dist))

	for p := range dist {
		points = append(points, p)
	}

	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			d := p1.Manhattan(p2)
			if d <= maxDist {
				if dist[p2] >= dist[p1]+d+100 {
					count++
				}
				if dist[p1] >= dist[p2]+d+100 {
					count++
				}
			}
		}
	}
	return count
}

func (s *Solver) Part1() (interface{}, error) {
	grid, start := parseGrid(s)
	dist := buildDist(grid, start)
	return solve(dist, 2), nil
}

func (s *Solver) Part2() (interface{}, error) {
	grid, start := parseGrid(s)
	dist := buildDist(grid, start)
	return solve(dist, 20), nil
}
