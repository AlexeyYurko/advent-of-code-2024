package day17

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day17", "input.txt"))
	return &Solver{input: string(input)}
}

func executeProgram(registers map[string]int, program []int) string {
	var outputs []string
	instr := 0

	for instr < len(program)-1 {
		opcode := program[instr]
		operand := program[instr+1]

		getValue := func(n int) int {
			if n <= 3 {
				return n
			}
			registerMap := map[int]string{
				4: "A",
				5: "B",
				6: "C",
			}
			if reg, ok := registerMap[n]; ok {
				return registers[reg]
			}
			return 0
		}

		switch opcode {
		case 0: // adv
			registers["A"] /= 1 << getValue(operand)
		case 1: // bxl
			registers["B"] ^= operand
		case 2: // bst
			registers["B"] = getValue(operand) % 8
		case 3: // jnz
			if registers["A"] == 0 {
				instr += 2
				continue
			}
			instr = operand
			continue
		case 4: // bxc
			registers["B"] ^= registers["C"]
		case 5: // out
			outputs = append(outputs, strconv.Itoa(getValue(operand)%8))
		case 6: // bdv
			registers["B"] = registers["A"] / (1 << getValue(operand))
		case 7: // cdv
			registers["C"] = registers["A"] / (1 << getValue(operand))
		}
		instr += 2
	}
	return strings.Join(outputs, ",")
}

func parseInput(s string) (registers map[string]int, program []int) {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	registers = make(map[string]int, 3)
	for i, reg := range []string{"A", "B", "C"} {
		registers[reg], _ = strconv.Atoi(lines[i][12:])
	}

	programStr := strings.Split(lines[4][9:], ",")
	program = make([]int, len(programStr))
	for i, str := range programStr {
		program[i], _ = strconv.Atoi(str)
	}
	return registers, program
}

func (s *Solver) Part1() (interface{}, error) {
	registers, program := parseInput(s.input)
	return executeProgram(registers, program), nil
}

func (s *Solver) Part2() (interface{}, error) {
	registers, program := parseInput(s.input)

	valid := []int{0}
	for length := 1; length < len(program)+1; length++ {
		var newValid []int
		target := program[len(program)-length:]

		for _, num := range valid {
			for offset := 0; offset < 8; offset++ {
				candidate := 8*num + offset

				testRegisters := make(map[string]int)
				for k, v := range registers {
					testRegisters[k] = v
				}
				testRegisters["A"] = candidate

				output := strings.Split(executeProgram(testRegisters, program), ",")
				if len(output) == len(target) {
					match := true
					for i := range output {
						outVal, _ := strconv.Atoi(output[i])
						if outVal != target[i] {
							match = false
							break
						}
					}
					if match {
						newValid = append(newValid, candidate)
					}
				}
			}
		}
		valid = newValid
	}

	if len(valid) == 0 {
		return 0, nil
	}

	minValue := valid[0]
	for _, v := range valid[1:] {
		if v < minValue {
			minValue = v
		}
	}

	return minValue, nil
}
