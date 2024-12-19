package day19

import (
	"os"
	"path/filepath"
	"strings"
)

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day19", "input.txt"))
	return &Solver{input: string(input)}
}

func splitTowelsAndDesigns(s *Solver) (towels, designs []string) {
	lines := strings.Split(strings.TrimSpace(s.input), "\n\n")
	towels = strings.Split(lines[0], ", ")
	designs = strings.Split(lines[1], "\n")
	return towels, designs
}

type TrieNode struct {
	children map[string]*TrieNode
	isEnd    bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[string]*TrieNode),
	}
}

func createTowelPatternTrie(towels []string) *TrieNode {
	root := newTrieNode()
	for _, towel := range towels {
		current := root
		current.children[towel] = newTrieNode()
		current.children[towel].isEnd = true
	}
	return root
}

func countValidTowelPatterns(pattern string, root *TrieNode, cache map[string]int) int {
	if val, ok := cache[pattern]; ok {
		return val
	}

	if pattern == "" {
		return 1
	}

	isPossibleCount := 0
	for prefix, node := range root.children {
		if strings.HasPrefix(pattern, prefix) {
			if len(prefix) == len(pattern) && node.isEnd {
				isPossibleCount++
				continue
			}
			if len(prefix) < len(pattern) {
				isPossibleCount += countValidTowelPatterns(pattern[len(prefix):], root, cache)
			}
		}
	}

	cache[pattern] = isPossibleCount
	return isPossibleCount
}

func countDesignsWithValidPatterns(towels []string, designs []string) (result int) {
	cache := make(map[string]int)
	root := createTowelPatternTrie(towels)

	for _, d := range designs {
		if countValidTowelPatterns(d, root, cache) > 0 {
			result++
		}
	}
	return result
}

func sumAllValidPatternCombinations(towels []string, designs []string) (result int) {
	cache := make(map[string]int)
	root := createTowelPatternTrie(towels)

	for _, d := range designs {
		result += countValidTowelPatterns(d, root, cache)
	}
	return result
}

func (s *Solver) Part1() (interface{}, error) {
	towels, designs := splitTowelsAndDesigns(s)
	res := countDesignsWithValidPatterns(towels, designs)
	return res, nil
}

func (s *Solver) Part2() (interface{}, error) {
	towels, designs := splitTowelsAndDesigns(s)
	res := sumAllValidPatternCombinations(towels, designs)
	return res, nil
}
