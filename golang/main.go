package main

import (
	"aoc2021/day01"
	"aoc2021/day02"
	"aoc2021/day03"
	"aoc2021/day04"
	"aoc2021/utils"
	"fmt"
)

func main() {
	filename := "../inputs/001.txt"
	input, err := utils.ReadlinesInt(filename)
	if err != nil {
		panic(err)
	}
	for i, f := range [](func([]int) int){day01.Part1, day01.Part2} {
		solution := f(input)
		fmt.Printf("Day 1 Part %d solution: %d\n", i+1, solution)
	}

	filename = "../inputs/002.txt"
	input2, err := utils.ReadlinesStr(filename)
	if err != nil {
		panic(err)
	}
	for i, f := range [](func([]string) int){day02.Part1, day02.Part2} {
		solution := f(input2)
		fmt.Printf("Day 2 Part %d solution: %d\n", i+1, solution)
	}

	filename = "../inputs/003.txt"
	input3, err := utils.ReadlinesStr(filename)
	if err != nil {
		panic(err)
	}
	for i, f := range [](func([]string) int){day03.Part1, day03.Part2} {
		solution := f(input3)
		fmt.Printf("Day 3 Part %d solution: %d\n", i+1, solution)
	}

	filename = "../inputs/004.txt"
	input4, err := utils.ReadlinesStr(filename)
	if err != nil {
		panic(err)
	}
	for i, f := range [](func([]string) int){day04.Part1} {
		solution := f(input4)
		fmt.Printf("Day 4 Part %d solution: %d\n", i+1, solution)
	}
}
