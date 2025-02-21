package day09

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type File struct {
	ID   int
	Size int
}

type Solver struct {
	input string
}

func New() *Solver {
	input, _ := os.ReadFile(filepath.Join("internal", "solutions", "day09", "input.txt"))
	return &Solver{input: string(input)}
}

func (s *Solver) defragment(files []File) []File {
	fs := make([]File, len(files))
	copy(fs, files)

	for file := len(fs) - 1; file >= 0; file-- {
		for free := 0; free < file; free++ {
			if fs[file].ID != -1 && fs[free].ID == -1 && fs[free].Size >= fs[file].Size {
				fs = slices.Insert(fs, free, fs[file])
				fs[file+1].ID, fs[free+1].ID = -1, -1
				fs[free+1].Size -= fs[file+1].Size
			}
		}
	}
	return fs
}

func (s *Solver) calculateChecksum(files []File) int {
	var disk []int
	for _, f := range files {
		disk = append(disk, slices.Repeat([]int{f.ID}, f.Size)...)
	}

	checksum := 0
	for i, v := range disk {
		if v != -1 {
			checksum += i * v
		}
	}
	return checksum
}

func (s *Solver) Part1() (interface{}, error) {
	diskMap := strings.TrimSpace(s.input) + "0"
	var files []File

	for id := 0; id*2 < len(diskMap); id++ {
		size := int(diskMap[id*2] - '0')
		free := int(diskMap[id*2+1] - '0')
		files = append(files, slices.Repeat([]File{{id, 1}}, size)...)
		files = append(files, slices.Repeat([]File{{-1, 1}}, free)...)
	}
	deFragmentResult := s.defragment(files)
	return s.calculateChecksum(deFragmentResult), nil
}

func (s *Solver) Part2() (interface{}, error) {
	diskMap := strings.TrimSpace(s.input) + "0"
	var files []File

	for id := 0; id*2 < len(diskMap); id++ {
		size := int(diskMap[id*2] - '0')
		free := int(diskMap[id*2+1] - '0')
		files = append(files, File{id, size}, File{-1, free})
	}
	deFragmentResult := s.defragment(files)
	return s.calculateChecksum(deFragmentResult), nil
}
