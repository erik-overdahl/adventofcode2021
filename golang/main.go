package main

import (
	"aoc2021/day01"
	"aoc2021/day02"
	"aoc2021/day03"
	"aoc2021/day04"
	"aoc2021/day06"
	"aoc2021/utils"
	"path/filepath"

	"fmt"
)

func main() {
	solutions := []utils.AOCDay{
		&day01.Solution{},
		&day02.Solution{},
		&day03.Solution{},
		&day04.Solution{},
		&day06.Solution{},
	}
	for _, s := range solutions {
		day := s.Day()
		filename := fmt.Sprintf("../inputs/0%02d.txt", day)
		path, err := filepath.Abs(filename)
		if err != nil {
			panic(err)
		}
		s.Init(path)
		fmt.Printf("Day %d Part 1 solution: %s\n", day, s.Part1())
		fmt.Printf("Day %d Part 2 solution: %s\n", day, s.Part2())
	}
}
