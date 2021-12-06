package main

import (
	"aoc2021/day01"
	"aoc2021/day02"
	"aoc2021/day03"
	"aoc2021/day04"
	"aoc2021/day05"
	"aoc2021/day06"
	"aoc2021/utils"
	"path/filepath"
	"sync"

	"fmt"
)

var wg sync.WaitGroup

func main() {
	solutions := []utils.AOCDay{
		&day01.Solution{},
		&day02.Solution{},
		&day03.Solution{},
		&day04.Solution{},
		&day05.Solution{},
		&day06.Solution{},
	}
	for _, s := range solutions {
		wg.Add(1)
		go func(solutions utils.AOCDay) {
			defer wg.Done()
			runDay(solutions)
		}(s)
	}
	wg.Wait()
}

func runDay(s utils.AOCDay) {
	day := s.Day()
	filename := fmt.Sprintf("../inputs/0%02d.txt", day)
	path, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	s.Init(path)
	p1 := fmt.Sprintf("Day %d Part 1 solution: %s\n", day, s.Part1())
	p2 := fmt.Sprintf("Day %d Part 2 solution: %s\n", day, s.Part2())
	fmt.Printf("%s%s", p1, p2)
}
