package day08

import (
	"github.com/AlexeyYurko/advent-of-code-2024/internal/aoc"
	"os"
	"path/filepath"
	"strings"
)

func isWithinBounds(p aoc.Point, width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
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
	antennas := make(map[rune][]aoc.Point)
	antinodeSet := make(map[aoc.Point]bool)
	grid := parseGrid(s)

	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], aoc.Point{x, y})
			}
		}
	}

	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1, p2 := points[i], points[j]
				dx, dy := p2.X-p1.X, p2.Y-p1.Y
				antinode1 := aoc.Point{p1.X - dx, p1.Y - dy}
				antinode2 := aoc.Point{p2.X + dx, p2.Y + dy}

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
	antennas := make(map[rune][]aoc.Point)
	antinodeSet := make(map[aoc.Point]bool)
	grid := parseGrid(s)

	for y, row := range grid {
		for x, cell := range row {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], aoc.Point{x, y})
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

				dx := p2.X - p1.X
				dy := p2.Y - p1.Y

				gcd := GCD(aoc.Abs(dx), aoc.Abs(dy))
				if gcd != 0 {
					dx /= gcd
					dy /= gcd
				}

				curr := p1
				for isWithinBounds(curr, len(grid[0]), len(grid)) {
					antinodeSet[curr] = true
					curr.X -= dx
					curr.Y -= dy
				}

				curr = p1
				for isWithinBounds(curr, len(grid[0]), len(grid)) {
					antinodeSet[curr] = true
					curr.X += dx
					curr.Y += dy
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
