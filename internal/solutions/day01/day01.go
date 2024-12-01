package day01

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day01", "input.txt"))
	return &Solver{input: string(input)}
}

func prepareSlices(s string) (slc1, slc2 []int, err error) {
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		sliceLine := strings.Split(line, "   ")
		num1, err1 := strconv.Atoi(sliceLine[0])
		num2, err2 := strconv.Atoi(sliceLine[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("error converting line: %s", line)
		}
		slc1 = append(slc1, num1)
		slc2 = append(slc2, num2)
	}
	return slc1, slc2, nil
}
func (s *Solver) Part1() (interface{}, error) {
	slc1, slc2, err := prepareSlices(s.input)
	if err != nil {
		return nil, err
	}
	sort.Ints(slc1)
	sort.Ints(slc2)
	var diff int
	for i := 0; i < len(slc1); i++ {
		localDiff := slc1[i] - slc2[i]
		if localDiff < 0 {
			localDiff = -localDiff
		}
		diff += localDiff
	}
	return diff, nil
}

func (s *Solver) Part2() (interface{}, error) {
	slc1, slc2, err := prepareSlices(s.input)
	if err != nil {
		return nil, err
	}

	var simScoreMap map[int]int
	simScoreMap = make(map[int]int)

	for i := 0; i < len(slc2); i++ {
		simScoreMap[slc2[i]] += 1
	}

	var simScoreSum int
	for _, leftListValue := range slc1 {
		appears, exists := simScoreMap[leftListValue]
		if exists {
			simScoreSum += leftListValue * appears
		}
	}

	return simScoreSum, nil
}
