package day14

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	maxX = 101
	maxY = 103
)

type Bot struct {
	x  int
	y  int
	dx int
	dy int
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day14", "input.txt"))
	return &Solver{input: string(input)}
}

func parseBots(input string) []Bot {
	var bots []Bot
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var b Bot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &b.x, &b.y, &b.dx, &b.dy)
		bots = append(bots, b)
	}
	return bots
}

func (s *Solver) Part1() (interface{}, error) {
	bots := parseBots(s.input)

	for i := 0; i < 100; i++ {
		for j := range bots {
			bots[j].x += bots[j].dx
			bots[j].y += bots[j].dy

			if bots[j].x < 0 {
				bots[j].x += maxX
			} else if bots[j].x >= maxX {
				bots[j].x -= maxX
			}

			if bots[j].y < 0 {
				bots[j].y += maxY
			} else if bots[j].y >= maxY {
				bots[j].y -= maxY
			}
		}
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, bot := range bots {
		if bot.x < maxX/2 && bot.y < maxY/2 {
			q1++
		}
		if bot.x > maxX/2 && bot.y < maxY/2 {
			q2++
		}
		if bot.x < maxX/2 && bot.y > maxY/2 {
			q3++
		}
		if bot.x > maxX/2 && bot.y > maxY/2 {
			q4++
		}
	}

	return q1 * q2 * q3 * q4, nil
}

func (s *Solver) Part2() (interface{}, error) {
	bots := parseBots(s.input)

	for i := 0; i < 1e5; i++ {
		grid := make([][]rune, maxY)
		for y := range grid {
			grid[y] = make([]rune, maxX)
			for x := range grid[y] {
				grid[y][x] = '.'
			}
		}

		distinct := true

		for j := range bots {
			bots[j].x += bots[j].dx
			bots[j].y += bots[j].dy

			if bots[j].x < 0 {
				bots[j].x += maxX
			} else if bots[j].x >= maxX {
				bots[j].x -= maxX
			}

			if bots[j].y < 0 {
				bots[j].y += maxY
			} else if bots[j].y >= maxY {
				bots[j].y -= maxY
			}

			if grid[bots[j].y][bots[j].x] == '.' {
				grid[bots[j].y][bots[j].x] = '#'
			} else {
				distinct = false
			}
		}

		if distinct {
			fmt.Printf("\nIteration: %d\n", i+1)
			for _, row := range grid {
				for _, c := range row {
					fmt.Print(string(c))
				}
				fmt.Println()
			}
			return i + 1, nil
		}
	}

	return nil, fmt.Errorf("no solution found within 100,000 iterations")
}
