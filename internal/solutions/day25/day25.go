package day25

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/AlexeyYurko/advent-of-code-2024/internal/aoc"
)

const (
	gridWidth  = 5
	heightBits = 4
	maskValue  = 0x88888
	offsetKey  = 0x22222
)

var (
	directions = struct {
		down, up, origin aoc.Point
	}{
		down:   aoc.Point{Y: 1},
		up:     aoc.Point{Y: -1},
		origin: aoc.Point{},
	}
)

type Block struct {
	grid    [][]rune
	heights int
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day25", "input.txt"))
	return &Solver{input: string(input)}
}

func (s *Solver) parseInput() []string {
	return strings.Split(s.input, "\n\n")
}

func parseBlock(block string) [][]rune {
	lines := strings.Split(block, "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func calculateBlockHeights(grid [][]rune, isLock bool) int {
	heights := 0
	for x := 0; x < gridWidth; x++ {
		var position aoc.Point
		if isLock {
			position = aoc.Point{X: x, Y: 1}
			for position.Y < len(grid) && grid[position.Y][position.X] == '#' {
				position = position.Add(directions.down)
			}
			heights = (heights << heightBits) + (position.Y - 1)
		} else {
			position = aoc.Point{X: x, Y: gridWidth}
			for position.Y >= 0 && grid[position.Y][position.X] == '#' {
				position = position.Add(directions.up)
			}
			heights = (heights << heightBits) + (gridWidth - position.Y)
		}
	}
	return heights
}

func processBlocks(blocks []string) (locks, keys []int) {
	for _, block := range blocks {
		grid := parseBlock(block)
		if grid[directions.origin.Y][directions.origin.X] == '#' {
			locks = append(locks, calculateBlockHeights(grid, true))
		} else {
			keys = append(keys, calculateBlockHeights(grid, false))
		}
	}
	return locks, keys
}

func countMatches(locks, keys []int) int {
	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			if (lock+key+offsetKey)&maskValue == 0 {
				count++
			}
		}
	}
	return count
}

func (s *Solver) Part1() (interface{}, error) {
	blocks := s.parseInput()
	locks, keys := processBlocks(blocks)
	return countMatches(locks, keys), nil
}

func (s *Solver) Part2() (interface{}, error) {
	return "Merry Christmas and Happy New Year", nil
}
