package main

import (
	"aoc2021/day01"
	"aoc2021/day02"
	"aoc2021/day03"
	"aoc2021/day04"
	"aoc2021/day05"
	"aoc2021/day06"
	"aoc2021/day07"
	"aoc2021/day08"
	"aoc2021/day09"
	"aoc2021/day10"
	"aoc2021/day11"
	"aoc2021/utils"
	"embed"
	"strings"
	"sync"
	"time"

	"fmt"
)

var solutions = []utils.AOCDay{
	&day01.Solution{},
	&day02.Solution{},
	&day03.Solution{},
	&day04.Solution{},
	&day05.Solution{},
	&day06.Solution{},
	&day07.Solution{},
	&day08.Solution{},
	&day09.Solution{},
	&day10.Solution{},
	&day11.Solution{},
}

var wg sync.WaitGroup

//go:embed inputs/*
var inputs embed.FS

func main() {
	outChan := make(chan output)
	res := make(chan []string, len(solutions))
	go func() {
		outputs := make([]string, len(solutions))
		for m := range outChan {
			outputs[m.pos] = m.msg
		}
		res <- outputs
	}()
	for _, s := range solutions {
		wg.Add(1)
		go func(solutions utils.AOCDay) {
			defer wg.Done()
			msg := runDay(solutions)
			outChan <- output{solutions.Day() - 1, msg}
		}(s)
	}
	wg.Wait()
	close(outChan)
	outputs := <-res
	close(res)
	for day := range outputs {
		fmt.Println(outputs[day])
	}
}

func runDay(s utils.AOCDay) string {
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
	s1 := fmt.Sprintf("Part 1: %-19d  (took %s)", p1, elapsed)

	start = time.Now()
	p2 := s.Part2()
	elapsed = time.Since(start)
	s2 := fmt.Sprintf("Part 2: %-19d  (took %s)", p2, elapsed)

	return fmt.Sprintf("Day %d\n%s\n%s\n", day, s1, s2)
}

type output struct {
	pos int
	msg string
}
