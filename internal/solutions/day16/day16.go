package day16

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type Maze struct {
	grid  [][]rune
	start Point
	end   Point
}

type QueueItem struct {
	pos   Point
	dir   int
	score int
	path  []Point
}

var directions = []Direction{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

const (
	Wall     = '#'
	Start    = 'S'
	End      = 'E'
	TurnCost = 1000
	MoveCost = 1
	StartDir = 1
)

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day16", "input.txt"))
	return &Solver{input: string(input)}
}

func (p Point) add(d Direction) Point {
	return Point{p.x + d.dx, p.y + d.dy}
}

func (p Point) key(dir int) string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, dir)
}

func (m Maze) isValid(p Point) bool {
	return p.y >= 0 && p.y < len(m.grid) &&
		p.x >= 0 && p.x < len(m.grid[0]) &&
		m.grid[p.y][p.x] != Wall
}

func (m Maze) isEnd(p Point) bool {
	return p == m.end
}

func parseMaze(input string) Maze {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	var start, end Point

	for y, line := range lines {
		grid[y] = []rune(line)
		for x, ch := range grid[y] {
			switch ch {
			case Start:
				start = Point{x, y}
			case End:
				end = Point{x, y}
			}
		}
	}

	return Maze{grid, start, end}
}

func findLowestScore(m Maze) int {
	queue := []QueueItem{{m.start, StartDir, 0, nil}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].score < queue[j].score
		})

		current := queue[0]
		queue = queue[1:]

		if m.isEnd(current.pos) {
			return current.score
		}

		key := current.pos.key(current.dir)
		if visited[key] {
			continue
		}
		visited[key] = true

		nextPos := current.pos.add(directions[current.dir])
		if m.isValid(nextPos) {
			queue = append(queue, QueueItem{
				nextPos,
				current.dir,
				current.score + MoveCost,
				nil,
			})
		}

		queue = append(queue,
			QueueItem{current.pos, (current.dir + 1) % 4, current.score + TurnCost, nil},
			QueueItem{current.pos, (current.dir + 3) % 4, current.score + TurnCost, nil},
		)
	}

	return -1
}

func findAllOptimalPaths(m Maze, targetScore int) [][]Point {
	queue := []QueueItem{{m.start, StartDir, 0, []Point{m.start}}}
	visited := make(map[string]int)
	var paths [][]Point

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.score > targetScore {
			continue
		}

		key := current.pos.key(current.dir)
		if score, exists := visited[key]; exists && score < current.score {
			continue
		}
		visited[key] = current.score

		if m.isEnd(current.pos) && current.score == targetScore {
			paths = append(paths, current.path)
			continue
		}

		nextPos := current.pos.add(directions[current.dir])
		if m.isValid(nextPos) {
			newPath := make([]Point, len(current.path))
			copy(newPath, current.path)
			queue = append(queue, QueueItem{
				nextPos,
				current.dir,
				current.score + MoveCost,
				append(newPath, nextPos),
			})
		}

		for _, newDir := range []int{(current.dir + 1) % 4, (current.dir + 3) % 4} {
			queue = append(queue, QueueItem{
				current.pos,
				newDir,
				current.score + TurnCost,
				current.path,
			})
		}
	}

	return paths
}

func countUniqueTiles(paths [][]Point) int {
	unique := make(map[string]bool)
	for _, path := range paths {
		for _, p := range path {
			unique[p.key(0)] = true
		}
	}
	return len(unique)
}

func (s *Solver) Part1() (interface{}, error) {
	maze := parseMaze(s.input)
	return findLowestScore(maze), nil
}

func (s *Solver) Part2() (interface{}, error) {
	maze := parseMaze(s.input)
	lowestScore := findLowestScore(maze)
	paths := findAllOptimalPaths(maze, lowestScore)
	return countUniqueTiles(paths), nil
}
