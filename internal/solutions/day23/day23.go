package day23

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Graph map[string]*Node

type Node struct {
	neighbors []string
}

func (g Graph) GetOrCreate(nodeName string) *Node {
	if node, exists := g[nodeName]; exists {
		return node
	}
	g[nodeName] = &Node{
		neighbors: make([]string, 0, 8),
	}
	return g[nodeName]
}

func (n *Node) AddNeighbor(name string) {
	n.neighbors = append(n.neighbors, name)
}

func (n *Node) HasNeighbor(name string) bool {
	for _, neighbor := range n.neighbors {
		if neighbor == name {
			return true
		}
	}
	return false
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day23", "input.txt"))
	return &Solver{input: string(input)}
}

func parseInput(s *Solver) Graph {
	graph := make(Graph)

	for _, line := range strings.Split(s.input, "\n") {
		if line == "" {
			continue
		}

		nodes := strings.Split(line, "-")
		sourceNode := nodes[0]
		targetNode := nodes[1]

		graph.GetOrCreate(sourceNode).AddNeighbor(targetNode)
		graph.GetOrCreate(targetNode).AddNeighbor(sourceNode)
	}
	return graph
}

func (s *Solver) Part1() (interface{}, error) {
	graph := parseInput(s)
	result := findNumberOfInterconnections(graph)
	return result, nil
}

func (s *Solver) Part2() (interface{}, error) {
	graph := parseInput(s)
	result := findPassword(graph)
	return result, nil
}

func findNumberOfInterconnections(graph Graph) int {
	triangles := make(map[string]struct{})

	for nodeName, node := range graph {
		for _, neighbor1 := range node.neighbors {
			nodeA := graph[neighbor1]

			for _, neighbor2 := range node.neighbors {
				if neighbor1 == neighbor2 {
					continue
				}

				nodeB := graph[neighbor2]

				if nodeA.HasNeighbor(neighbor2) && nodeB.HasNeighbor(neighbor1) {
					triangle := make([]string, 3)
					triangle[0] = nodeName
					triangle[1] = neighbor1
					triangle[2] = neighbor2
					sort.Strings(triangle)

					triangles[strings.Join(triangle, "-")] = struct{}{}
				}
			}
		}
	}

	count := 0
	for triangleKey := range triangles {
		if triangleKey[0] == 't' || triangleKey[3] == 't' || triangleKey[6] == 't' {
			count++
		}
	}

	return count
}

func findPassword(graph Graph) string {
	maxClique := make([]string, 0)

	for nodeName, node := range graph {
		currentClique := []string{nodeName}

		for _, neighbor := range node.neighbors {
			neighborNode := graph[neighbor]

			isFullyConnected := true
			for _, existingNode := range currentClique {
				if !neighborNode.HasNeighbor(existingNode) {
					isFullyConnected = false
					break
				}
			}

			if isFullyConnected {
				currentClique = append(currentClique, neighbor)
			}
		}

		sort.Strings(currentClique)

		if len(currentClique) > len(maxClique) {
			maxClique = currentClique
		}
	}

	return strings.Join(maxClique, ",")
}
