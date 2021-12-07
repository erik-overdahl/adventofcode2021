package main

import (
	"aoc2021/day01"
	"aoc2021/day02"
	"aoc2021/day03"
	"aoc2021/day04"
	"aoc2021/day05"
	"aoc2021/day06"
	"aoc2021/day07"
	"aoc2021/utils"
	"embed"
	"strings"
	"sync"
	"time"

	"fmt"
)

var wg sync.WaitGroup

//go:embed inputs/*
var inputs embed.FS

func main() {
	solutions := []utils.AOCDay{
		&day01.Solution{},
		&day02.Solution{},
		&day03.Solution{},
		&day04.Solution{},
		&day05.Solution{},
		&day06.Solution{},
		&day07.Solution{},
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
	filename := fmt.Sprintf("inputs/0%02d.txt", day)
	inputBytes, err := inputs.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputBlob := strings.Trim(string(inputBytes), "\n")
	s.Init(inputBlob)
	start := time.Now()
	p1 := s.Part1()
	elapsed := time.Since(start)
	s1 := fmt.Sprintf("Day %d Part 1 solution: %s  (took %s)\n", day, p1, elapsed)

	start = time.Now()
	p2 := s.Part2()
	elapsed = time.Since(start)
	s2 := fmt.Sprintf("Day %d Part 2 solution: %s  (took %s)\n", day, p2, elapsed)

	fmt.Printf("%s%s", s1, s2)
}
