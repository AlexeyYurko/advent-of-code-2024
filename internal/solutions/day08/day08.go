package day08

import (
	"os"
	"path/filepath"
	"strings"
)

type Point struct {
	x, y int
}

func isWithinBounds(p Point, width, height int) bool {
	return p.x >= 0 && p.x < width && p.y >= 0 && p.y < height
}

func parseGrid(s *Solver) [][]rune {
	lines := strings.Split(strings.TrimSpace(s.input), "\n")

	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = []rune(line)
	}
	return grid
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day08", "input.txt"))
	return &Solver{input: string(input)}
}

func (s *Solver) Part1() (interface{}, error) {
	antennas := make(map[rune][]Point)
	antinodeSet := make(map[Point]bool)
	grid := parseGrid(s)

	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], Point{x, y})
			}
		}
	}

	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1, p2 := points[i], points[j]
				dx, dy := p2.x-p1.x, p2.y-p1.y
				antinode1 := Point{p1.x - dx, p1.y - dy}
				antinode2 := Point{p2.x + dx, p2.y + dy}

				if isWithinBounds(antinode1, len(grid[0]), len(grid)) {
					antinodeSet[antinode1] = true
				}
				if isWithinBounds(antinode2, len(grid[0]), len(grid)) {
					antinodeSet[antinode2] = true
				}
			}
		}
	}

	return len(antinodeSet), nil
}

func (s *Solver) Part2() (interface{}, error) {
	antennas := make(map[rune][]Point)
	antinodeSet := make(map[Point]bool)
	grid := parseGrid(s)

	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], Point{x, y})
			}
		}
	}

	for _, points := range antennas {
		if len(points) < 2 {
			continue
		}

		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1, p2 := points[i], points[j]

				dx := p2.x - p1.x
				dy := p2.y - p1.y

				gcd := GCD(abs(dx), abs(dy))
				if gcd != 0 {
					dx /= gcd
					dy /= gcd
				}

				curr := p1
				for isWithinBounds(curr, len(grid[0]), len(grid)) {
					antinodeSet[curr] = true
					curr.x -= dx
					curr.y -= dy
				}

				curr = p1
				for isWithinBounds(curr, len(grid[0]), len(grid)) {
					antinodeSet[curr] = true
					curr.x += dx
					curr.y += dy
				}
			}
		}
	}

	return len(antinodeSet), nil
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
