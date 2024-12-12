package day12

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Boundary struct {
	row, col int
	counted  bool
}

type Region struct {
	area       int
	perimeter  int
	positions  map[string]bool
	boundaries map[Direction][]Boundary
	symbol     rune
}

type Solver struct {
	input string
}

type RegionExplorer interface {
	ProcessRegion(grid [][]rune, x, y int, region *Region)
}

type Part1Explorer struct{}
type Part2Explorer struct{}

func (p Part1Explorer) ProcessRegion(grid [][]rune, x, y int, region *Region) {
	region.perimeter += 4

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dirs {
		nx, ny := x+d[0], y+d[1]
		if isInBounds(grid, nx, ny) && grid[ny][nx] == grid[y][x] {
			region.perimeter--
		}
	}
}

func (p Part2Explorer) ProcessRegion(grid [][]rune, x, y int, region *Region) {
	region.symbol = grid[y][x]

	if y == 0 || grid[y-1][x] != region.symbol {
		region.boundaries[North] = append(region.boundaries[North], Boundary{y, x, true})
	}
	if y == len(grid)-1 || grid[y+1][x] != region.symbol {
		region.boundaries[South] = append(region.boundaries[South], Boundary{y, x, true})
	}
	if x == 0 || grid[y][x-1] != region.symbol {
		region.boundaries[West] = append(region.boundaries[West], Boundary{y, x, true})
	}
	if x == len(grid[0])-1 || grid[y][x+1] != region.symbol {
		region.boundaries[East] = append(region.boundaries[East], Boundary{y, x, true})
	}
}

func parseGrid(s *Solver) [][]rune {
	lines := strings.Split(strings.TrimSpace(s.input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func isInBounds(grid [][]rune, x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func exploreRegion(grid [][]rune, x, y int, visited map[string]bool, region *Region, explorer RegionExplorer) {
	pos := fmt.Sprintf("%d,%d", x, y)
	if visited[pos] {
		return
	}

	region.area++
	region.positions[pos] = true
	visited[pos] = true

	explorer.ProcessRegion(grid, x, y, region)

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dirs {
		nx, ny := x+d[0], y+d[1]
		if isInBounds(grid, nx, ny) && grid[ny][nx] == grid[y][x] {
			exploreRegion(grid, nx, ny, visited, region, explorer)
		}
	}
}

func findRegions(grid [][]rune, explorer RegionExplorer) []Region {
	var regions []Region
	visited := make(map[string]bool)

	for y := range grid {
		for x := range grid[y] {
			pos := fmt.Sprintf("%d,%d", x, y)
			if !visited[pos] {
				region := Region{
					positions:  make(map[string]bool),
					boundaries: make(map[Direction][]Boundary),
				}
				exploreRegion(grid, x, y, visited, &region, explorer)
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func pruneRedundantBoundaries(region *Region) {
	for dir := range region.boundaries {
		bounds := region.boundaries[dir]
		if dir == North || dir == South {
			sort.Slice(bounds, func(i, j int) bool { return bounds[i].col < bounds[j].col })
		} else {
			sort.Slice(bounds, func(i, j int) bool { return bounds[i].row < bounds[j].row })
		}

		for i := range bounds {
			var next int
			if dir == North || dir == South {
				next = bounds[i].col
			} else {
				next = bounds[i].row
			}

			for {
				next++
				found := false
				for j := range bounds {
					if (dir == North || dir == South) &&
						bounds[j].col == next && bounds[j].row == bounds[i].row {
						bounds[j].counted = false
						found = true
						break
					} else if (dir == East || dir == West) &&
						bounds[j].row == next && bounds[j].col == bounds[i].col {
						bounds[j].counted = false
						found = true
						break
					}
				}
				if !found {
					break
				}
			}
		}
	}
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day12", "input.txt"))
	return &Solver{input: string(input)}
}

func (s *Solver) Part1() (interface{}, error) {
	grid := parseGrid(s)
	regions := findRegions(grid, Part1Explorer{})

	sum := 0
	for _, region := range regions {
		sum += region.area * region.perimeter
	}
	return sum, nil
}

func (s *Solver) Part2() (interface{}, error) {
	grid := parseGrid(s)
	regions := findRegions(grid, Part2Explorer{})

	sum := 0
	for _, region := range regions {
		pruneRedundantBoundaries(&region)
		boundaryCount := 0
		for _, bounds := range region.boundaries {
			for _, b := range bounds {
				if b.counted {
					boundaryCount++
				}
			}
		}
		sum += region.area * boundaryCount
	}
	return sum, nil
}
