package day04

import (
	"os"
	"path/filepath"
	"strings"
)

type Solver struct {
	input string
}

func parseGrid(s *Solver) [][]rune {
	grid := make([][]rune, 0)
	for _, line := range strings.Split(strings.TrimSpace(s.input), "\n") {
		grid = append(grid, []rune(line))
	}
	return grid
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day04", "input.txt"))
	return &Solver{input: string(input)}
}

func (s *Solver) Part1() (interface{}, error) {
	grid := parseGrid(s)

	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	target := "XMAS"
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				found := true

				for k := 0; k < len(target); k++ {
					newX, newY := i+k*dx, j+k*dy
					if newX < 0 || newX >= rows || newY < 0 || newY >= cols || grid[newX][newY] != rune(target[k]) {
						found = false
						break
					}
				}
				if found {
					count++
				}
			}
		}
	}

	return count, nil
}

func (s *Solver) Part2() (interface{}, error) {
	grid := parseGrid(s)

	count := 0
	rows, cols := len(grid), len(grid[0])

	patterns := [][4][2]int{
		// MAS + MAS
		{{-1, -1}, {1, 1}, {-1, 1}, {1, -1}},
		// MAS + SAM
		{{-1, -1}, {1, 1}, {1, -1}, {-1, 1}},
		// SAM + MAS
		{{1, 1}, {-1, -1}, {-1, 1}, {1, -1}},
		// SAM + SAM
		{{1, 1}, {-1, -1}, {1, -1}, {-1, 1}},
	}

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if grid[i][j] != 'A' {
				continue
			}

			for _, pattern := range patterns {
				if isValidMAS(grid, i, j, pattern[0], pattern[1]) && isValidMAS(grid, i, j, pattern[2], pattern[3]) {
					count++
				}
			}
		}
	}

	return count, nil
}

func isValidMAS(grid [][]rune, centerI, centerJ int, start, end [2]int) bool {
	rows, cols := len(grid), len(grid[0])

	mI, mJ := centerI+start[0], centerJ+start[1]
	if mI < 0 || mI >= rows || mJ < 0 || mJ >= cols || grid[mI][mJ] != 'M' {
		return false
	}

	sI, sJ := centerI+end[0], centerJ+end[1]
	if sI < 0 || sI >= rows || sJ < 0 || sJ >= cols || grid[sI][sJ] != 'S' {
		return false
	}

	return true
}
