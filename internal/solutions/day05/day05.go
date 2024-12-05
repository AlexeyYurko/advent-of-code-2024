package day05

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Solver struct {
	input string
}

func parseRules(lines []string) map[int]map[int]bool {
	rules := make(map[int]map[int]bool)
	for _, line := range lines {
		parts := strings.Split(line, "|")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		if rules[a] == nil {
			rules[a] = make(map[int]bool)
		}
		rules[a][b] = true
	}
	return rules
}

func parseUpdates(lines []string) [][]int {
	var updates [][]int
	for _, line := range lines {
		var update []int
		for _, num := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(num)
			update = append(update, n)
		}
		updates = append(updates, update)
	}
	return updates
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day05", "input.txt"))
	return &Solver{input: string(input)}
}

func (s *Solver) Part1() (interface{}, error) {
	parts := strings.Split(strings.TrimSpace(s.input), "\n\n")
	rules := parseRules(strings.Split(parts[0], "\n"))
	updates := parseUpdates(strings.Split(parts[1], "\n"))

	sum := 0
	for _, update := range updates {
		if isValidOrder(update, rules) {
			sum += update[len(update)/2]
		}
	}

	return sum, nil
}

func (s *Solver) Part2() (interface{}, error) {
	parts := strings.Split(strings.TrimSpace(s.input), "\n\n")
	rules := parseRules(strings.Split(parts[0], "\n"))
	updates := parseUpdates(strings.Split(parts[1], "\n"))

	sum := 0
	for _, update := range updates {
		if !isValidOrder(update, rules) {
			sorted := topologicalSort(update, rules)
			sum += sorted[len(sorted)/2]
		}
	}

	return sum, nil
}

func isValidOrder(update []int, rules map[int]map[int]bool) bool {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if rules[update[j]] != nil && rules[update[j]][update[i]] {
				return false
			}
		}
	}
	return true
}

func topologicalSort(update []int, rules map[int]map[int]bool) []int {
	graph := make(map[int]map[int]bool)
	nodes := make(map[int]bool)

	for _, n := range update {
		nodes[n] = true
		graph[n] = make(map[int]bool)
	}

	for i := range nodes {
		for j := range nodes {
			if i != j && rules[i] != nil && rules[i][j] {
				graph[i][j] = true
			}
		}
	}

	inDegree := make(map[int]int)
	for n := range nodes {
		inDegree[n] = 0
	}
	for _, edges := range graph {
		for dest := range edges {
			inDegree[dest]++
		}
	}

	var result []int
	var queue []int

	for n := range nodes {
		if inDegree[n] == 0 {
			queue = append(queue, n)
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		result = append(result, curr)

		for next := range graph[curr] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return result
}
