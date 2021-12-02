package main

import (
	"aoc2021/day01"
	"aoc2021/day02"
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
		fmt.Printf("Day 1 Part %d solution: %d\n", i, solution)
	}

	filename = "../inputs/002.txt"
	input2, err := utils.ReadlinesStr(filename)
	if err != nil {
		panic(err)
	}
	for i, f := range [](func([]string) int){day02.Part1, day02.Part2} {
		solution := f(input2)
		fmt.Printf("Day 1 Part %d solution: %d\n", i, solution)
	}
}
