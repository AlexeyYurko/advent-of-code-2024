package day24

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var gatePattern = regexp.MustCompile(`^(.*) (AND|OR|XOR) (.*) -> (.*)$`)

type Operation int

const (
	AND Operation = iota
	OR
	XOR
)

type Gate struct {
	inputA    string
	inputB    string
	operation Operation
	output    string
}

type Circuit struct {
	values   map[string]uint16
	gates    []Gate
	resolved map[string]bool
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day24", "input.txt"))
	return &Solver{input: string(input)}
}

func NewCircuit() *Circuit {
	return &Circuit{
		values:   make(map[string]uint16),
		gates:    make([]Gate, 0),
		resolved: make(map[string]bool),
	}
}

func parseInput(input string) *Circuit {
	circuit := NewCircuit()
	sections := strings.Split(input, "\n\n")

	for _, line := range strings.Split(strings.TrimSpace(sections[0]), "\n") {
		parts := strings.SplitN(line, ": ", 2)
		if value, err := strconv.ParseUint(parts[1], 10, 16); err == nil {
			circuit.values[parts[0]] = uint16(value)
			circuit.resolved[parts[0]] = true
		}
	}

	for _, line := range strings.Split(strings.TrimSpace(sections[1]), "\n") {
		if gate := parseGate(line); gate != nil {
			circuit.gates = append(circuit.gates, *gate)
		}
	}

	return circuit
}

func parseGate(line string) *Gate {
	if match := gatePattern.FindStringSubmatch(line); match != nil {
		var op Operation
		switch match[2] {
		case "AND":
			op = AND
		case "OR":
			op = OR
		case "XOR":
			op = XOR
		}
		return &Gate{
			inputA:    match[1],
			inputB:    match[3],
			operation: op,
			output:    match[4],
		}
	}
	return nil
}

func (c *Circuit) evaluate() {
	changed := true
	for changed {
		changed = false
		for _, gate := range c.gates {
			if c.resolved[gate.output] {
				continue
			}

			valA, okA := c.values[gate.inputA]
			valB, okB := c.values[gate.inputB]

			if !okA || !okB {
				continue
			}

			c.values[gate.output] = c.computeGate(gate, valA, valB)
			c.resolved[gate.output] = true
			changed = true
		}
	}
}

func (c *Circuit) computeGate(gate Gate, a, b uint16) uint16 {
	switch gate.operation {
	case AND:
		return a & b
	case OR:
		return a | b
	case XOR:
		return a ^ b
	default:
		return 0
	}
}

func (c *Circuit) getZWires() []string {
	zWires := make([]string, 0)
	for wire := range c.values {
		if strings.HasPrefix(wire, "z") {
			zWires = append(zWires, wire)
		}
	}
	sort.Slice(zWires, func(i, j int) bool {
		return zWires[i] > zWires[j]
	})
	return zWires
}

func (c *Circuit) computeDecimalValue() int {
	zWires := c.getZWires()
	binary := strings.Builder{}
	binary.Grow(len(zWires))

	for _, wire := range zWires {
		binary.WriteString(strconv.Itoa(int(c.values[wire])))
	}

	if result, err := strconv.ParseInt(binary.String(), 2, 64); err == nil {
		return int(result)
	}
	return 0
}

func (c *Circuit) findWireSwaps() []string {
	swapped := make([]string, 0, 90)
	gateMap := make(map[string]Gate)

	for _, gate := range c.gates {
		gateMap[gate.output] = gate
	}

	var c0 string
	for i := 0; i < 45; i++ {
		n := fmt.Sprintf("%02d", i)
		xn, yn := "x"+n, "y"+n

		var m1, n1, r1, z1, c1 string

		for out, gate := range gateMap {
			if (gate.inputA == xn && gate.inputB == yn) || (gate.inputA == yn && gate.inputB == xn) {
				if gate.operation == XOR {
					m1 = out
				} else if gate.operation == AND {
					n1 = out
				}
			}
		}

		if c0 != "" {
			for out, gate := range gateMap {
				if (gate.inputA == c0 && gate.inputB == m1) || (gate.inputA == m1 && gate.inputB == c0) {
					if gate.operation == AND {
						r1 = out
					} else if gate.operation == XOR {
						z1 = out
					}
				}
			}

			if r1 == "" {
				m1, n1 = n1, m1
				swapped = append(swapped, m1, n1)
				for out, gate := range gateMap {
					if (gate.inputA == c0 && gate.inputB == m1) || (gate.inputA == m1 && gate.inputB == c0) {
						if gate.operation == AND {
							r1 = out
						}
					}
				}
			}

			if strings.HasPrefix(m1, "z") {
				m1, z1 = z1, m1
				swapped = append(swapped, m1, z1)
			}
			if strings.HasPrefix(n1, "z") {
				n1, z1 = z1, n1
				swapped = append(swapped, n1, z1)
			}
			if strings.HasPrefix(r1, "z") {
				r1, z1 = z1, r1
				swapped = append(swapped, r1, z1)
			}

			for out, gate := range gateMap {
				if (gate.inputA == r1 && gate.inputB == n1) || (gate.inputA == n1 && gate.inputB == r1) {
					if gate.operation == OR {
						c1 = out
					}
				}
			}
		}

		if strings.HasPrefix(c1, "z") && c1 != "z45" {
			c1, z1 = z1, c1
			swapped = append(swapped, c1, z1)
		}

		if c0 == "" {
			c0 = n1
		} else {
			c0 = c1
		}
	}

	sort.Strings(swapped)
	return swapped
}

func (s *Solver) Part1() (interface{}, error) {
	circuit := parseInput(s.input)
	circuit.evaluate()
	return circuit.computeDecimalValue(), nil
}

func (s *Solver) Part2() (interface{}, error) {
	circuit := parseInput(s.input)
	swapped := circuit.findWireSwaps()
	return strings.Join(swapped, ","), nil
}
